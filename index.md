---
title:  ORMGP Interpolants
author: Oak Ridges Moraine Groundwater Program
output: html_document
---

* TOC
{:toc}

# Introduction
There's a need for an introduction here, highlighting the need to:
* coordinate disparate datasets so that correlation among environmental phenomena can be investigated
* data are sourced primarily from open datasets and are ingested into our database
* by virtue of being in our database, web applications allowing users to analyze the data become available
* this is a system of delivering information needed for water resources and planning.





# Data Sources
* **[Automated/near real-time](/interpolants/sources/sources.html)**


# Processes

## Tools & Technologies
* [Python](https://www.python.org/)
* [Go](https://go.dev/)
* FORTRAN/C++
* VB.NET/C#
* [R-Shiny](https://shiny.rstudio.com/)
   * [Leaflet for R](https://rstudio.github.io/leaflet/)
   * [ggplot2](https://ggplot2.tidyverse.org/)
   * [dygraphs for R](https://rstudio.github.io/dygraphs/)

### Data storage
* [MSSQL](https://www.microsoft.com/en-us/sql-server/sql-server-2019)
* [Delft-FEWS](https://www.deltares.nl/app/uploads/2015/01/Delft-FEWS_brochure-2017.pdf)

### Web scraping
* cron: a linux server with scheduled tasks written in Python
* ORMGP-FEWS: our operational Delft-FEWS system


## Interpolation
* **[Water Table](https://owrc.github.io/watertable/)**
* **[Barometry](/interpolants/interpolation/barometry.html)**
* **[Land use](/interpolants/interpolation/landuse.html)**
    * impervious cover
    * canopy cover
    * wetland cover
    * open water
    * relative permeability
* **Drainage and Topology**
    * **[Overland](/interpolants/interpolation/overland.html)**
    * **[Watercourses](/interpolants/interpolation/watercourses.html)**
    * **[Subwatersheds](/interpolants/interpolation/subwatershed.html)**



## Modelling
* Numerical Model Custodianship Program
* **[Hydrograph separation](/interpolants/modelling/hydrographseparation.html)**
   * **[Hydrograph disaggregation](/interpolants/modelling/hydroparse.html)**
* **[Water-budget model](/interpolants/modelling/waterbudgetmodel.html)**
   * **[Input data](/interpolants/modelling/waterbudget/data.html)**
   * **[Soil moisture accounting](/interpolants/modelling/waterbudget/sma.html)**
   * **[Total evaporation](/interpolants/modelling/waterbudget/pet.html)**
   * Potential Solar Irradiation
      * **[Atmospheric transmittance](/interpolants/modelling/BristowCampbell.html)**
   * **[Snowmelt](/interpolants/modelling/waterbudget/snowmeltCCF.html)**
   * **[Shallow groundwater](/interpolants/modelling/waterbudget/gw.html)**
   * **[Overland flow routing](/interpolants/modelling/waterbudget/overlandflow.html)**

   

# Products
* 6hr, 24hr
* mean min max temp
* atmospheric Pressure
* rainfall
* snowfall
* snowmelt

$$ e^{i\pi} = - 1 $$
