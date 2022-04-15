# Processing steps:

1. Run `drainTopology.exe` (which was build from `drainTopology.go`), requires:
    1. *shp/OHN_WATERCOURSE-export.geojson* - the raw watercourse data.
    1. *shp/OHN_WATERCOURSE-export-rootsel.geojson* - polygons that intersect the outlet line segments of the raw watercourse data.
1. Simplify *shp/OHN_WATERCOURSE-export-segments.geojson* (the output of step 1.) using the `Simplify` (Douglas-Peuker) vector operation of [QGIS](https://www.qgis.org/en/site/) with a tolerance of 10m. Export to *shp/OHN_WATERCOURSE-export-segments_simplWGS.geojson*, re-projected to WGS84. 

More information can be found [here](https://owrc.github.io/interpolants/interpolation/watercourses.html).