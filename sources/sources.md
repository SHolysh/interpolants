---
title: Data Sources
author: M.Marchildon
date: 2022-02-22
output: html_document
---

# Web Data Scrapers

On a nightly basis, scheduled tasks are used to automatically download and insert new data into the [ORMGP database](https://owrc.github.io/database-manual/Contents/TOC.html), a process called "[web scraping](https://en.wikipedia.org/wiki/Web_scraping)." In addition to hydrogeologic information, the database maintains hydrometeorological data at the **daily** time scale. This means that, for the most part, information scraped below do not enter the database if they are sub-daily resolution. Rather, the higher-resolution data are maintained using [Delft-FEWS](https://oss.deltares.nl/web/delft-fews/about-delft-fews) on a dedicated server.

## Current data sources

Below is a list data sources used in deriving many of the interpolation products. These sources are all free and open, but come in a variety of formats.


* **WSC** - Water Survey of Canada 
  - [MSC Datamart](https://eccc-msc.github.io/open-data/msc-datamart/readme_en/) *.csv file server for near real-time data
    - daily streamflow stage and discharge
  - [National Water Data Archive: HYDAT](https://www.canada.ca/en/environment-climate-change/services/water-overview/quantity/monitoring/survey/data-products-services/national-archive-hydat.html) are updated quarterly and overwrites any data recorded through Datamart.
    - daily streamflow stage and discharge

* **ECCC-MSC** - Environment and Climate Change Canada-Meteorological Service of Canada
    - [Historical Data](https://climate.weather.gc.ca/historical_data/search_historic_data_e.html) (html portal to meteorologic data)
        - hourly: air temperature, air pressure, relative humidity, wind speed, wind direction
        - daily: min/max air temperature, rainfall, snowfall, snow depth
    - Canadian Precipitation Analysis (CaPA)
        - [Regional Deterministic Precipitation Analysis (CaPA-RDPA)](https://weather.gc.ca/grib/grib2_RDPA_ps10km_e.html) is a gridded (~10km) 6-hourly near realtime continuous precipitation field for the past 20 years, albeit in many overlapping versions.
        - [High Resolution Deterministic Precipitation Analysis (CaPA-HRDPA)](https://eccc-msc.github.io/open-data/msc-data/nwp_hrdpa/readme_hrdpa_en/) is a refined version (~2.5km) of the product above, going back to 2019.

* **NOAA-NSIDC** - National Oceanic and Atmospheric Administration-National Snow & Ice Data Center
    * [Snow Data Assimilation System (SNODAS)](https://nsidc.org/data/g02158)
        - continuous ~30m resolution grid 24-hour accumulations of snowmelt, realtime from 2009 (when coverage was opened to Canada)

## Partners' data

A number of our partners maintain internal databases. Attempts are being made to integrate these sources to our workflow without the need for data duplication.

* TRCA
* CVC
* CLOCA
* York Region