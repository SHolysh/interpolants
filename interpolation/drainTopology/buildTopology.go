package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"github.com/maseology/goHydro/routing"
	"github.com/maseology/mmaths/spatial"
	tp "github.com/maseology/mmaths/topology"
	"github.com/maseology/mmio"
	geojson "github.com/paulmach/go.geojson"
)

const (
	ndthrsh  = 10. //1.5
	strmsFP  = "shp/OHN_WATERCOURSE-export.geojson"
	outletFP = "shp/OHN_WATERCOURSE-export-rootsel.geojson"
)

func main() {

	t := mmio.NewTimer()
	t0 := time.Now()
	defer fmt.Printf("\nTopology complete %v", time.Since(t0))

	fmt.Print("\nloading data..")
	outlets := getOutlets(outletFP)
	var strms [][][3]float64
	var epfs, epls [][3]float64
	var am [][]int
	if _, ok := mmio.FileExists(mmio.RemoveExtension(strmsFP) + "-AM.gob"); ok {
		strms, am, epfs, epls = loadAMgob(mmio.RemoveExtension(strmsFP) + "-AM.gob")
	} else {
		strms = getStreams(strmsFP)
		t.PrintAndReset("segmentizing..")
		strms = spatial.JunctionToJunction(strms, ndthrsh)
		t.PrintAndReset("building adjacency matrix..")
		am, epfs, epls = routing.BuildAdjacencyMatrix(strms, ndthrsh) // VERY SLOW
		saveAMgob(mmio.RemoveExtension(strmsFP)+"-AM.gob", strms, am, epfs, epls, ndthrsh)
	}

	t.PrintAndReset("collecting nodes..\n")
	nodes := routing.CollectNodes(strms, am, epfs, epls) // no direction assigned I[0]: node ID
	t.PrintAndReset("collecting roots..")
	roots := collectRoots(nodes, outlets)
	// func() { // print for testing
	// 	sca := mmio.NewCSVwriter("roots.vertices.csv")
	// 	sca.WriteHead("rid,x,y")
	// 	for i, n := range roots {
	// 		sca.WriteLine(i, n.S[0], n.S[1])
	// 	}
	// 	sca.Close()
	// }()

	t.PrintAndReset("building directional topology..")
	routing.SetDirectionFromRoots(nodes, roots) // I[1]: upstream topologically-safe order
	t.PrintAndReset("collecting junction to junction reaches from roots..")
	reaches := routing.JunctionToJunctionFromRoots(roots) // ordered from roots; I: {dimension, tree ID, tree-segment ID}
	// func() { // print for testing
	// 	sca := mmio.NewCSVwriter("junction.vertices.csv")
	// 	sca.WriteHead("nid,rid,rsid,x,y")
	// 	for i, n := range reaches {
	// 		sca.WriteLine(i, n.I[1], n.I[2], n.S[0], n.S[1])
	// 	}
	// 	sca.Close()
	// }()
	t.PrintAndReset("trimming small reaches..")
	reachesCrop := make([]*tp.Node, 0, len(reaches))
	for _, s := range reaches {
		if len(s.US) == 0 { // leaf
			if len(s.S) <= 2*s.I[0] { // only 2 vertices
				if s.S[0]-s.S[s.I[0]] < ndthrsh && s.S[1]-s.S[s.I[0]+1] < ndthrsh { // too short, remove leaf
					for _, d := range s.DS {
						newdUS := []*tp.Node{}
						for _, u := range d.US {
							if u != s {
								newdUS = append(newdUS, u)
							}
						}
						d.US = newdUS
					}
					continue
				}
			}
		}
		reachesCrop = append(reachesCrop, s)
	}

	t.PrintAndReset("ordering..")
	routing.Strahler(reachesCrop) // I[3]: order

	t.PrintAndReset("saving network..")
	for i, n := range reachesCrop {
		n.I = append(n.I, i) // I[4]: reach ID
		if len(n.I) != 5 {
			log.Fatalln("unequal index count")
		}
	}
	printNetwork(mmio.RemoveExtension(strmsFP)+"-segments.geojson", reachesCrop)
}

func getStreams(fp string) [][][3]float64 {
	fstreams, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	gstreams, err := geojson.UnmarshalFeatureCollection(fstreams)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	streams, nf := [][][3]float64{}, 0
	for _, f := range gstreams.Features {
		if f.Geometry.Type != "MultiLineString" {
			log.Fatalln("todo")
		}
		for _, ln := range f.Geometry.MultiLineString {
			ml := make([][3]float64, len(ln))
			for i, c := range ln {
				ml[i][0] = c[0]
				ml[i][1] = c[1]
				ml[i][2] = -9999.
				nf++
				if len(c) != 2 {
					log.Fatalln("todo")
				}
			}
			streams = append(streams, ml)
		}
	}
	fmt.Printf("  %d stream vertices - ", nf)
	return streams
}

func getOutlets(fp string) [][][3]float64 {
	foutlet, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	goutlet, err := geojson.UnmarshalFeatureCollection(foutlet)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	outlets := [][][3]float64{}
	for _, f := range goutlet.Features {
		if f.Geometry.Type != "MultiPolygon" {
			log.Fatalln("todo")
		}
		for _, pgn := range f.Geometry.MultiPolygon {
			if len(pgn) != 1 {
				log.Fatalln("todo")
			}
			for _, part := range pgn {
				ml := make([][3]float64, len(part))
				for i, c := range part {
					ml[i][0] = c[0]
					ml[i][1] = c[1]
					ml[i][2] = -9999.
					if len(c) != 2 {
						log.Fatalln("todo")
					}
				}
				outlets = append(outlets, ml)
			}
		}
	}
	return outlets
}

// only returns ends within polygons
func collectRoots(nds map[*tp.Node][]*tp.Node, outlets [][][3]float64) []*tp.Node {
	roots := []*tp.Node{}
	pointInPoly := func(plgn [][3]float64, pt []float64) bool {
		// based on PNPoly by W. Randolph Franklin: http://www.ecse.rpi.edu/~wrf/Research/Short_Notes/pnpoly.html
		b := false
		j := len(plgn) - 1
		X, Y := pt[0], pt[1]
		for i, v := range plgn {
			if (v[1] > Y) != (plgn[j][1] > Y) {
				if X < (plgn[j][0]-v[0])*(Y-v[1])/(plgn[j][1]-v[1])+v[0] {
					b = !b
				}
			}
			j = i
		}
		return b
	}

	for _, plgn := range outlets {
		for n, c := range nds {
			if len(c) > 1 { // only considering end points
				continue
			}
			if pointInPoly(plgn, n.S) {
				roots = append(roots, n)
			}
		}
	}

	return roots
}

func saveAMgob(fp string, strms [][][3]float64, am [][]int, epf, epl [][3]float64, radius float64) {
	d := struct {
		S    [][][3]float64
		A    [][]int
		F, L [][3]float64
		R    float64
	}{S: strms, A: am, F: epf, L: epl, R: radius}
	f, err := os.Create(fp)
	if err != nil {
		log.Fatalf("saveAMgob: %v", err)
	}
	enc := gob.NewEncoder(f)
	err = enc.Encode(d)
	if err != nil {
		log.Fatalf("saveAMgob: %v", err)
	}
	err = f.Close()
	if err != nil {
		log.Fatalf("saveAMgob: %v", err)
	}
}

func loadAMgob(fp string) (strms [][][3]float64, am [][]int, epf, epl [][3]float64) {
	d := struct {
		S    [][][3]float64
		A    [][]int
		F, L [][3]float64
		R    float64
	}{}
	f, err := os.Open(fp)
	if err != nil {
		log.Fatalf("loadAMgob: %v", err)
	}
	enc := gob.NewDecoder(f)
	err = enc.Decode(&d)
	if err != nil {
		log.Fatalf("routing/build.go loadAMgob: %v", err)
	}
	err = f.Close()
	if err != nil {
		log.Fatalf("routing/build.go loadAMgob: %v", err)
	}
	return d.S, d.A, d.F, d.L
}

func printNetwork(fp string, nds []*tp.Node) {
	fc := geojson.NewFeatureCollection()
	for _, n := range nds {
		nd := n.I[0] // number of spatial dimensions recorded in set of vertices
		nv := len(n.S) / nd
		vs := make([][]float64, nv)
		for j := 0; j < nv; j++ {
			vs[j] = make([]float64, nd)
			for d := 0; d < nd; d++ {
				vs[j][d] = n.S[j*nd+d]
			}
		}
		ups, dns, dni := []int{}, []int{}, math.MaxInt
		for _, u := range n.US {
			ups = append(ups, u.I[2])
		}
		for _, d := range n.DS {
			dns = append(dns, d.I[2])
			if d.I[4] < dni {
				dni = d.I[4]
			}
		}
		if dni == math.MaxInt {
			dni = -1
		}

		f := geojson.NewLineStringFeature(vs)
		f.SetProperty("segmentID", n.I[4])
		f.SetProperty("downID", dni)
		f.SetProperty("treeID", n.I[1])
		f.SetProperty("treesegID", n.I[2])
		f.SetProperty("topol", fmt.Sprintf("%d %d %d", ups, n.I[2], dns))
		f.SetProperty("order", n.I[3])
		fc.AddFeature(f)
	}
	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		log.Fatalf("routing.PrintSegments: %v\n", err)
	}
	if err := ioutil.WriteFile(fp, rawJSON, 0644); err != nil {
		log.Fatalf("routing.PrintSegments: %v\n", err)
	}
}
