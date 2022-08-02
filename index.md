---
title:  ORMGP Interpolants
author: Oak Ridges Moraine Groundwater Program
output: html_document
---

* TOC
{:toc}



# Introduction
The [Oak Ridges Moraine Groundwater Program (ORMGP)](https://www.oakridgeswater.ca/) maintains an authoritative understanding of: i) the geological layering for the area; ii) the groundwater flow system; and iii) its linkage to the region's surface waters and climatologies. 

**_TODO: There's a need for an introduction here, highlighting the need to:_**
* coordinate disparate datasets so that correlation among environmental phenomena can be investigated
* data are sourced primarily from open datasets and are ingested into our database
* by virtue of being in our database, web applications allowing users to analyze the data become available
* this is a system of delivering information needed for water resources and planning.



# Data Sources

The ORMGP data systems described below are active, in that they are [automatically updated at near real-time](/interpolants/sources/sources.html). From hourly and daily intervals, a series of "web-scraper" routines have been written to access open data made available from many sources including international, federal, provincial, municipal and conservation authority agencies. The objective is to "centralize" hydrogeological data for the south-central Ontario (Canada) jurisdiction that currently makes-up the [ORMGP](https://owrc.github.io/snapshots/partners.html).


# Processes, Tools & Technologies
A number of open-source data analysis tools have been employed in our overall system. In contribution to the open-source community, much of our work is also left open [on Github](https://github.com/OWRC).

* [Github pages](https://pages.github.com/) — the hosting of technical information, like what you see now.
* [Python](https://www.python.org/)
   * [Beautiful Soup](https://www.crummy.com/software/BeautifulSoup/bs4/doc/) — web scraping
   * [NumPy](https://numpy.org/), [Pandas](https://pandas.pydata.org/) — data manipulation
   * (and many more)
* [Go](https://go.dev/)
* FORTRAN/C++
* VB.NET/C#
* [R](https://www.r-project.org/) and [RStudio](https://www.rstudio.com/) — a free, open-source "...software environment for statistical computing and graphics."
   * [R-Shiny](https://shiny.rstudio.com/) — online, real-time data analytics
   * [Leaflet for R](https://rstudio.github.io/leaflet/) — web mapping
   * [ggplot2](https://ggplot2.tidyverse.org/) — data visualization
   * [dygraphs for R](https://rstudio.github.io/dygraphs/) — dynamic timeseries visualization

### Data storage systems
* [MSSQL](https://www.microsoft.com/en-us/sql-server/sql-server-2019)
* [Delft-FEWS](https://www.deltares.nl/app/uploads/2015/01/Delft-FEWS_brochure-2017.pdf)


## Servers

### Web scraping
* ORMGP-cron: a linux server with scheduled web-scraping tasks mostly written in Python
* [ORMGP-FEWS](/interpolants/interpolation/fews.html): our operational Delft-FEWS system that scrapes, stores, interpolates and regenerates data served on ORMGP-cron.

### Databases and APIs
* **[Our main MSSQL Server database](https://owrc.github.io/database-manual/Contents/TOC.html)**
* [VertiGIS Studio/GeoCortex](https://www.vertigis.com/vertigis-studio/) — Our principle [web mapping server](https://maps.oakridgeswater.ca/Html5Viewer/index.html?viewer=ORMGPP) used by partners to access our database and products.
* golang server — multi-functional REST API serving a variety of data products and interpolation tools.




# Products

## Near real-time data, spatially interpolated
* [**Daily data**](/interpolants/interpolation/daily.html), updated nightly, including:
   * Min/max daily temperature
   * Atmospheric pressure
   * Precipitation (rainfall and snowfall)
   * Snowmelt
   * Pan evaporation
   * Solar irradiation
* Hourly data, updated every 6-hours, including:
   * Temperature
   * Atmospheric pressure
   * Precipitation
* [Barometry](/interpolants/interpolation/barometry.html)



## Spatial interpolation
_Static 2D fields, updated frequently_
* [Water Table](https://owrc.github.io/watertable/)
* [Land use](/interpolants/interpolation/landuse.html), including:
    * impervious cover
    * canopy cover
    * wetland cover
    * open water
    * relative permeability
* Drainage and Topology
    * [Overland](/interpolants/interpolation/overland.html)
    * [Watercourses](/interpolants/interpolation/watercourses.html)
    * [Sub-watersheds](/interpolants/interpolation/subwatershed.html)



## Modelling
* [Numerical Model Custodianship Program](https://owrc.github.io/snapshots/numerical-model-custodianship-program.html)
* [Hydrograph separation](/interpolants/modelling/hydrographseparation.html)
   * [Hydrograph disaggregation](/interpolants/modelling/hydroparse.html)
* [ORMGP distributed water-budget model](/interpolants/modelling/waterbudgetmodel.html)
   * [Input data](/interpolants/modelling/waterbudget/data.html)
   * [Soil moisture accounting](/interpolants/modelling/waterbudget/sma.html)
   * [Total evaporation](/interpolants/modelling/waterbudget/pet.html)
   * Potential Solar Irradiation
      * [Atmospheric transmittance](/interpolants/modelling/BristowCampbell.html)
   * [Snowmelt](/interpolants/modelling/waterbudget/snowmeltCCF.html)
   * [Shallow groundwater](/interpolants/modelling/waterbudget/gw.html)
   * [Overland flow routing](/interpolants/modelling/waterbudget/overlandflow.html)
   * [References](/interpolants/modelling/waterbudgetmodel.html#references)
   