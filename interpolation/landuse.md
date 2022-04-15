---
title: Land use properties of the ORMGP region
author: Oak Ridges Moraine Groundwater Program
date: 2022-02-28
output: html_document
---

* TOC
{:toc}


# Derivative products

An overlay analysis is the process of overlaying 2 or more spatial layers and capturing statistics associated with their relative coverage. In this case, the sub-watershed layer is overlain by Provincial land-use and surficial geology layers to obtain information like percent impervious, relative permeability, etc.

Provincial layers discussed in more detail below have in all cases been re-sampled to the 50x50m² grid associated with the [hydrologically corrected DEM](/interpolants/interpolation/overland.html). It is from these rasters where the aggregation of watershed characteristics is computed.

## Land use

The Ministry of Natural Resources and Forestry (2019) SOLRIS version 3.0 provincial land use layer is employed to aggregate imperviousness and canopy coverage at the sub-watershed scale. In areas to the north, where the SOLRIS coverage discontinues, interpretation was applied by: 

1. Using Provincial mapping of roads, wetlands and water bodies, areas outside of the SOLRIS data bounds *(typically up on the Canadian Shield)* are filled in with the appropriate SOLRIS land use class index (201, 150, 170, respectively--MNRF, 2019); and,
2. All remaining area not covered by SOLRIS is assumed Forest (SOLRIS land use class index of 90), as observed with satellite imagery.

The dominant SOLRIS land use class (by area) is assigned the Land use class index for every 50x50m² grid cell. 

![Final 50x50m SOLRIS mapping.](https://github.com/OWRC/subwatershed/blob/main/jupyter/output/solrisv3_10_infilled_50.png?raw=true)
Final 50x50m SOLRIS mapping. *(For illustrative purposes only [see here](https://github.com/OWRC/subwatershed/blob/main/jupyter/OWRC-SWS.ipynb) to reproduce shown raster.)*


### Land use coverage

For any ~10km² sub-watershed and give a 50x50m² grid , there should be a set of roughly 4,000 SOLRIS land use class indices. Using a look-up system, the set of cells contained within a sub-watershed are assigned a value of imperviousness, water body, wetland and canopy coverage (according to their SOLRIS index) and accumulated to a sub-watershed sum.

Percent impervious and canopy coverage as per SOLRIS v3.0 (MNRF, 2019) land use classification.
<!-- ```{r message=FALSE, warning=FALSE, echo=FALSE}
library(knitr)
library(dplyr)
library(kableExtra)
options(knitr.kable.NA = '')
read.csv('shp/lookup_200731.csv') %>%
  select(1:2,7:8) %>%
  mutate(PerImp=PerImp*100,PerCov=PerCov*100) %>%
  kbl(
    col.names = c("Index", "Name", "Imperviousness (%)", "Canopy cover (%)"),
    align = c("r","l","c","c"),
    digits = 0
  ) %>%
  kable_styling(bootstrap_options = c("striped", "hover", "condensed")) %>%
  add_header_above(c("SOLRIS Land use classification" = 2, " " = 2), align = "l") %>%
  kableExtra::scroll_box(width = "90%", height = "300px")
``` -->


![Final 50x50m impervious mapping.](https://github.com/OWRC/subwatershed/blob/main/jupyter/output/solrisv3_10_infilled_50_perimp.png?raw=true)
*(For illustrative purposes only [see here](https://github.com/OWRC/subwatershed/blob/main/jupyter/OWRC-SWS.ipynb) to reproduce shown raster.)*

![Final 50x50m canopy mapping.](https://github.com/OWRC/subwatershed/blob/main/jupyter/output/solrisv3_10_infilled_50_percov.png?raw=true)
Final 50x50m canopy mapping. *(For illustrative purposes only [see here](https://github.com/OWRC/subwatershed/blob/main/jupyter/OWRC-SWS.ipynb) to reproduce shown raster.)*


## Surficial geology

The Ontario Geological Survey's 2010 Surficial geology of southern Ontario layer also assigns a 50x50m² grid by the dominant class. 

![Final 50x50m permeability mapping.](https://github.com/OWRC/subwatershed/blob/main/jupyter/output/OGSsurfGeo_50.png?raw=true)
Final 50x50m permeability mapping. *(For illustrative purposes only [see here](https://github.com/OWRC/subwatershed/blob/main/jupyter/OWRC-SWS.ipynb) to reproduce shown raster.)*

### Permeability

The OGS classes have been grouped according to the attribute "permeability" using a similar look-up table cross-referencing scheme. OGS (2010) adds: *"Permeability classification is a very generalized one, based purely on characteristics of material types."* 

After assigning an assumed "effective" hydraulic conductivity to every permeability group, sub-watershed "permeability" is then calculated as the geometric mean of 50x50m² grid cells contained within a sub-watershed. Effective hydraulic conductivity value assumed for every permeability group is shown here:

Permeability classifications (after OGS, 2010) and assumed effective hydraulic conductivities.
<!-- ```{r message=FALSE, warning=FALSE, echo=FALSE}
library(knitr)
library(dplyr)
par <- c("Low","Low-medium","Medium","Medium-high","high","unknown/variable","fluvial","organics")
val <- format(c(1e-9,1e-8,1e-7,1e-6,1e-5,1e-8,1e-5,1e-6),digits=3)
data.frame(par,val) %>%
  kbl(
    col.names = c(" ", "K (m/s)"),
    align = c("l","r")
  ) %>%
  kable_styling(full_width = F, bootstrap_options = c("striped", "hover", "condensed"))
``` -->

The resulting effective hydraulic conductivity is then reverted back to the nearest Low--High OGS (2010) classification.


## References

Ministry of Natural Resources and Forestry, 2019. Southern Ontario Land Resource Information System (SOLRIS) Version 3.0: Data Specifications. Science and Research Branch, April 2019

Ontario Geological Survey 2010. Surficial geology of southern Ontario; Ontario Geological Survey, Miscellaneous Release— Data 128 – Revised.