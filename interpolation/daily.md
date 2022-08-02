---
title: Near Real-time Daily Climatology
author: M.Marchildon
output: html_document
---

* TOC
{:toc}


# Introduction
The [ORMGP](https://maps.oakridgeswater.ca/) maintains a current, continuous daily climatology dataset beginning 1901. The data are interpolated to [some 3,000 ~10km² sub-watersheds](https://owrc.github.io/interpolants/interpolation/subwatershed.html) and are made available through our [web portal](https://maps.oakridgeswater.ca/Html5Viewer/index.html?viewer=ORMGPP).

The data serve many purposes from basic overlay with other disparate data sets (e.g., hydrographs) to providing input to hydrological models. 

Clearly, data collection is constrained by public/private funding, open date sharing policies and technology; so the data served are a compilation of what is deemed the best for *our* needs and thus originate from a number of sources. Specifically, we are concerned  climatology *regional scale*. We cover a 3 Million hectare jurisdiction and based on our experience, compiling a continuous interpolated climate dataset required greater emphasis on distribution (quantity) over data quality.  That's not to say that this is a general rule, but for our needs in our humid region with a significant winter presence, attention to how weather distribute is of paramount importance.

Admittedly, greater emphasis is made to "re-packaging" data-products from external agencies; we are not trying to re-invent the wheel. Instead, we are only trying to automate a workflow that until now is constantly repeated by our partner agencies and partner consultants when such data are needed. We don't require any of our partners to use our data, we are only making available the data we use in our internal analyses.

Below is a description of the datasets used in our overall climatology package. Precipitation, for example, comes from a multitude of sources that generally follow technological advances. In order to obtain a century+ continuous dataset, older station-based interpolation is supplanted by more recent *"data assimilation systems"* (DAS), where and when available.

It's also worth noting that the density of [meteorological stations present in our jurisdiction has been in decline since the 1970s](https://owrc.github.io/snapshots/gantt-met.html). Understandably, much of the recent investment in meteorological station operation has been dedicated to Canada's north, a large geographical region that has been grossly overlooked yet is most susceptible to a changing climate.




# Daily climatologies
Below, the data types that are collected, interpolated and delivered through our web portal are described in chronological order. For the most part, the most recent datasets supersede the oldest.

All interpolated (i.e., raster, vector) data are automatically updated and maintained using our [ORMGP-FEWS system](/interpolants/interpolation/fews.html). 


## Precipitation

### Meteorological Service of Canada (*1901—present*)
Meteorological Service of Canada is a division of Environment Canada. Their [online historical data portal](https://climate.weather.gc.ca/index_e.html) provides data collected since the mid 19th century.

These data are collected at (point/local/scalar) weather stations and require spatial interpolation. Here, the nearest neighbour (i.e., Thiessen polygon) method is applied.
