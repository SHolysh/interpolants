---
title: Subwatershed delineation of the ORMGP region
author: Oak Ridges Moraine Groundwater Program
date: 2022-02-28
output: html_document
---

* TOC
{:toc}


## Subwatershed Delineation

The 3 million hectare ORMGP jurisdiction is subdivided to a number of 10km² sub-watersheds as a basis for hydrometeorological data analysis. Every sub-watershed has a defined topological order in which headwater sub-watersheds can easily be mapped to subsequent downstream sub-watersheds, and so on until feeding the great lakes. The intent here is to deem these sub-watersheds a "logical unit" for climatological and water budget analyses. Below is a description of the derivation the v.2020 OWRC 10km² sub-watershed map and its derivatives including:

- a catchment area delineation tool
- an interpolated real-time daily meteorological dataset dating back to the year 1900
- an overlay analysis performed to characterize sub-watershed:
 + impervious area
 + canopy coverage
 + water body coverage
 + wetland coverage
 + relative permeability/infiltrability
 + mean slope and dominant aspect
 + mean depth to water table.

Below first describes the processing of a Provincial Digital Elevation Model (DEM) yielding a "hydrologically-correct" model of the ORMGP ground surface. From this information, the ORMGP region is portioned into 2,813 ~10km² sub-watersheds. 

Next, the hydrologically-corrected digital elevation model (HDEM) is further process to derive a number of metrics aggregated at the sub-watershed scale.



