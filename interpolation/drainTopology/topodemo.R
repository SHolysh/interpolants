

library(shiny)
library(leaflet)
library(sf)
library(tidyr)
library(dplyr)
library(stringr)




# segNams <- read.csv("./segList.csv") %>% mutate_if(is.character, str_trim)
ormgp <- st_read("shp/OHN_WATERCOURSE-export-segments-simplWGS.geojson") %>%
  # merge(segNams,by.x="treeID",by.y="treeID",all.x = TRUE) %>%
  # mutate(WatershedName=coalesce(WatershedName, as.character(treeID))) %>%
  filter(order>3) #################################################################################################### for speed, lower order streams disabled
  
  
  
ui <- bootstrapPage(
  title="ORMGPtopology",
  tags$style(type = "text/css", "html, body {width:100%;height:100%}"),
  leafletOutput("map", width = "100%", height = "100%")
)


server <- function(input, output, session) {

  output$map <- renderLeaflet({
    showNotification("Calculation in progress.......... this may take a minute",duration = 60)
    
    pntlab <- paste("Tree ID: ", ormgp$WatershedName, 
                    "</br>Segment ID: ", ormgp$segmentID, 
                    "</br>Strahler order : ", ormgp$order, 
                    "</br>topology : ", ormgp$topol)
    
    leaflet(ormgp) %>%
      addTiles(group='OSM') %>% # OpenStreetMap by default
      addProviderTiles(providers$OpenTopoMap, group='Topo', options = providerTileOptions(attribution=" Map data: © OpenStreetMap contributors, SRTM | Map style: © OpenTopoMap (CC-BY-SA) | Oak Ridges Moraine Groundwater Program")) %>%
      addProviderTiles(providers$Stamen.TonerLite, group = "Toner Lite", options = providerTileOptions(attribution=" Map tiles by Stamen Design, CC BY 3.0 — Map data © OpenStreetMap contributors | Oak Ridges Moraine Groundwater Program")) %>%
      # addProviderTiles(providers$Stamen.TonerLite,options = providerTileOptions(noWrap = TRUE)) %>%
      addPolylines(layerId = ~segmentID,
                   weight = ~order,
                   group = ~paste0("Strahler ", order),
                   color = "blue",
                   highlightOptions = highlightOptions(color = "black",
                                                       weight = 10,
                                                       bringToFront = TRUE),
                   label = lapply(pntlab,htmltools::HTML)) %>%
                   # label = ~paste0("Strahler order ", order)) %>%
                   
      hideGroup(c("Strahler 1","Strahler 2")) %>%
      addLayersControl(baseGroups = c("OSM", "Topo", "Toner Lite"))
  })
  
  observe({
    clk <- input$map_shape_click
    if (!is.null(clk)) {
      i <- clk$id
      lst <- list()
      repeat{
        lst = append(lst,i)
        i <- ormgp[ormgp$segmentID==i,]$downID
        if (i<0) break
      } 
      leafletProxy("map") %>% 
        clearGroup("temp") %>% 
        addPolylines(data=ormgp[ormgp$segmentID %in% lst,],color="yellow",weight=10,group="temp")
    }
  })
  
  observe({
    newzoom <- input$map_zoom
    if (is.null(newzoom)) return()
    isolate({
      if (newzoom >= 12) {
        leafletProxy("map") %>% showGroup(c("Strahler 1","Strahler 2"))
      }
      if (newzoom == 11) {
        leafletProxy("map") %>% hideGroup("Strahler 1") %>% showGroup("Strahler 2")
      }
      if (newzoom <= 10) {
        leafletProxy("map") %>% hideGroup(c("Strahler 1","Strahler 2"))
      }
    })    
  })
  
  session$onSessionEnded(stopApp)
}

shinyApp(ui, server)
