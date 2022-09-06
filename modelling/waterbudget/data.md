---
title: Water Budget Modelling
subtitle: Input Data
author: M.Marchildon
date: 2019
output: html_document
---

* TOC
{:toc}


# Model Input Data Processing

Meteorological data aquisiton, management, aggregation and interpolation was accomplished using [Delft-FEWS](https://www.deltares.nl/en/software/flood-forecasting-system-delft-fews-2/) (ver.2019.02 build.39845) a flood forecasting system offered (at no cost, only license agreement) by [Deltares](https://www.deltares.nl/en/). Configuration files for the Delft-FEWS system build can be found here: BLAH.

Climate forcing data required by the model are atmospheric yield and potential evaporation on a 6-hourly timestep. Atmospheric yield is her defined as water added to land surface in its mobile (read: liquid) form.  Additional processing performed on the data exported from Delft-FEWS is described below.



#### Data sets:

Forcings data to the model processed by FEWS include:

1. [CaPA-RDPA](https://weather.gc.ca/grib/grib2_RDPA_ps10km_e.html) 10/15km gridded precipitation fields, yielding 6-hourly precipiation totals, aquired from [CaSPAr](https://caspar-data.ca/);
1. [SNODAS](https://nsidc.org/data/g02158) (NOHRSC, 2004) ~1km gridded 24-hour (UTC 06-06) snowmelt totals; and
1. Meteorological Service of Canada (MSC) hourly mean temperature and pressure accessed [here](http://climate.weather.gc.ca/historical_data/search_historic_data_e.html).

#### Transformations

The time step of the model was set to the 6-hour time step offered with the CaPA-RDPA data. Hourly temperatures were 

#### Time step (temporal scale)

The time step of the model has been set to 6 hour steps. All other data sets have been either aggregated or disaggregated to match this temporal scale.

The aquired data come in a variety of time steps...




<!-- CHECK -->
![](fig/met_timeline.svg)






#### Interpolation (spatial scale)

Once transformed to the set time step, both scalar (i.e., point) data and gridded data are then interpolated to the 10 km² sub-watersheds. Given that each watershed contains 





### Atmospheric Yield
#### Precipitation and Snowmelt






CaSPAr

... 10 km resolution, compare that to the XX average spacing among operational meteorological stations.


snodas ... can avoid the need to model snowmelt explicitly, and leverage online resources.

<!-- The daily SNODAS time series was then disaggregated to the 6-hourly-UTC time stepping scheme described above using the Delft-FEWS transformation [Disaggregation: MeanToMean](https://publicwiki.deltares.nl/display/FEWSDOC/meanToMean). -->



The precipitation and snowmelt time-series was then projected onto the model sub-watersheds using the transformation: [InverseDistance](https://publicwiki.deltares.nl/display/FEWSDOC/InterpolationSpatialInverseDistance).





#### Atmospheric Yield


The data collected include total precipitation and snowmelt. Summing the two together would double count precipitation fallen as snow; the model, however does not account for snow, rather it relies on snowmelt as a forcing. Precipition is parsed into rainfall and snowfall on the basis of a critical temperature ($T_\text{crit}$):

$$
\text{Rainfall}=
\begin{cases}
\text{Precipitation}, & T>T_\text{crit}\\
0 & \text{otherwise},
\end{cases}
$$
$$
\text{Snowfall}=
\begin{cases}
\text{Precipitation}, & T\leq T_\text{crit}\\
0 & \text{otherwise}.
\end{cases}
$$

<!-- $$ -->
<!-- \begin{split} -->

<!-- \text{Rainfall} &= -->
<!-- \begin{cases} -->
<!-- \text{Precipitation}, & T>T_\text{crit}\\ -->
<!-- 0 & \text{otherwise}, -->
<!-- \end{cases} \\ -->

<!-- \text{Snowfall} &= -->
<!-- \begin{cases} -->
<!-- \text{Precipitation}, & T\leq T_\text{crit}\\ -->
<!-- 0 & \text{otherwise}. -->
<!-- \end{cases} -->

<!-- \end{split} -->
<!-- $$ -->

An optimization routine is employed to determine $T_\text{crit}$ such that annual average snowfall is equal to annual average snowmelt to ensure minimal deviation from total precipitation. Furthermore, snowmelt, which is aquired at a daily timestep that represents the *"total of 24 per hour melt rates, 06:00 UTC-06:00 UTC ... integrated for the previous 24 hours, giving daily totals"* (NOHRSC, 2004), is disaggregated to a 6-hourly time step based on the following rules:

1. If any timesteps has temperatures greater than 10°C, snowmelt is equally divided amongst them;
1. The first time step having 6-hour rainfall $\geq$ 5 mm, all melt is assumed to occur during this event;
1. If any timesteps within 06:00-06:00 UTC has rainfall greater than 1 mm, snowmelt is proportioned according to (and added with) rainfall;
1. If any timesteps has temperatures greater than 0°C, snowmelt is equally divided amongst them; otherwise
1. Snowmelt is equally divided among the 2 daytime time steps (12:00-00:00 UTC---07:00-19:00 EST).

The final product, a single forcing termed "Atmospheric Yield" is inputted in the model.

$$\text{Rainfall} + \text{Snowmelt} = \textit{Atmospheric Yield}$$

The aim of the model design is to simultaneously reduce the amount of computational processes and leverage near-realtime data assimilation products. It is recognized from a hydrological model design perspective, that the dynamic processes that dictate the spatial distribution of watershed moisture is only affected by atmospheric yield, that is water sourced from the atmosphere in liquid form.




<!-- ### Temperature, Pressure, Humidity and Wind Speed -->
### Atmospheric Demand $(E_a)$
> $T_a$, $P_a$, $\text{rh}$ and $u$

Historical hourly measurements of air temperature, pressure, relative humidity and wind speed were retrieved from the ECCC website: https://climate.weather.gc.ca/historical_data/search_historic_data_e.html. In total, 46 stations with varying periods of record length and quality were collected. 

The hourly time-series collected from each station was then aggregated to the 6-hourly-UTC time stepping scheme described above. Aggregation of all metrics was accomplished using the Delft-FEWS transformation ["Aggregation: MeanToMean"](https://publicwiki.deltares.nl/display/FEWSDOC/Aggregation+MeanToMean).

The 6-hourly-UTC time-series was then projected onto the model sub-watersheds using the transformation ["InterpolationSpatial: InverseDistance"](https://publicwiki.deltares.nl/display/FEWSDOC/InterpolationSpatialInverseDistance).

Model elevations range from 75-400 masl and orographic effects were deemed negligible beyond the spatial distribution meteorological stations.


#### Atmospheric Demand


The model considers the greater role the atmosphere has on its 30,000 km² extent. The atmosphere, taking a top-down perspective, requires consideration of PBL (Oke) as it reprensents the barrier from which mass must transfer when surface evaporation is captured by the atmosphere. This is particularily so when considering mass transfer over rough surfaces, where surface evaporation becomes coupled with advective (vapour deficit) flux through the PBL (surf clim can).

This is evident when relating pan evaporation to strictly aerodynamic variables temperature and humidity, integrated at the 6-hourly timestep. For instance, using the advective term [kg/m²/s] of Penman (1948) is:

$$
E_a=\rho_a \frac{\varepsilon}{p_a} d_a \cdot f(u)
$$

where $d_a=(1-rh)e_s$, $e_s \propto T_a$ and the wind-function $f(u)=au^b$ [m/s] can safely reduced to an empirical form:

$$
E_a=7.46\times 10^{-9} \cdot au^b d_a
$$
where $E_a$ is now given in [m/s]. This is the power form of open water evaporation $(E_o)$ used by Penman (1948). It is worth noting that this is modified from Penman (1948) in that it is assumed $T_s \propto T_a$, that is the relationship between surface temperature and air temperature is captured in this empirical form.

<!-- In addition to the above "Power" form, Penman (1948) also offers the the most common (linear) form: -->

<!-- $$ -->
<!-- E_a=a(1+bu) d_a -->
<!-- $$ -->
<!-- However, this did not perform as well as well as the power form. Either form is dependent on temperature and horizontal wind speed. -->

For not it's simplicity, the power law does perform well against observation. 24,641 data-days from 17 MSC daily pan evaporation stations were gathered for validatation. With $u$ [m/s] and $d_a$ [Pa], $a=0.009$ and $b=0.26$ resulted in a global Kling-Gupta (2008) efficiency of 0.66 and 0.86 for daily and monthly pan evaporation estimation, respectively.

![](fig/4937_daily.svg)
Simulated daily evaporation (using the above equation) against observed pan evaporation.

![](fig/4937_mscatter.svg)
Monthly totals, scatter plot. Red line is the 1:1 line

![](fig/4937_mts.svg)
Monthly totals, timeseries. (Note: shown here are consecutive months, only December-March are not included.)







## Rainfall, Snowfall and Snowmelt
### Data Sources
#### CaPA-RDPA

_**todo**_

#### SNODAS

_**todo**_

<!-- moved to  -->
<!-- ### Optimized Critical Temperature Approach
The precipitation fields are proportioned into rainfall and snowfall amounts based using the "critical temperature" $(T_\text{crit})$ approach:

$$
\text{Rainfall}=
\begin{cases}
\text{Precipitation}, & T>T_\text{crit}\\
0 & \text{otherwise},
\end{cases}
$$

$$
\text{Snowfall}=
\begin{cases}
\text{Precipitation}, & T\leq T_\text{crit}\\
0 & \text{otherwise}.
\end{cases}
$$

An optimization routine is employed to determine $T_\text{crit}$ such that annual average snowfall is equal to annual average snowmelt to ensure minimal deviation from total precipitation.  -->

### Sub-daily Snowmelt
Furthermore, snowmelt, which is acquired at a daily timestep that represents the *"total of 24 per hour melt rates, 06:00 UTC-06:00 UTC ... integrated for the previous 24 hours, giving daily totals"* (NOHRSC, 2004), is disaggregated to a 6-hourly time step based on the following rules:

1. If any timesteps has temperatures greater than 10°C, snowmelt is equally divided amongst them;
1. The first time step having 6-hour rainfall $\geq$ 5mm, all melt is assumed to occur during this event;
1. If any timesteps within 06:00-06:00 UTC has rainfall greater than 1mm, snowmelt is proportionned according to (and added with) rainfall;
1. If any timesteps has temperatures greater than 0°C, snowmelt is equally divided amongst them; otherwise
1. Snowmelt is equally divided among the 2 daytime time steps (12:00-00:00 UTC---07:00-19:00 EST).

The final product, a single forcing termed "Atmospheric Yield" is inputted in the model.

$$\text{Rainfall} + \text{Snowmelt} = \textit{Atmospheric Yield}$$

The aim of the model design is to simulatneously reduce the amount of computational processes and leverage near-realtime data assimilation products. It is recognized from a hydrological model design perspective, that the dynamic processes that dictate the spatial distribution of watershed moisture is only affected by atmopheric yeild, that is water sourced from the atmosphere in liquid form.


# References

National Operational Hydrologic Remote Sensing Center. 2004. Snow Data Assimilation System (SNODAS) Data Products at NSIDC, Version 1. [Indicate subset used]. Boulder, Colorado USA. NSIDC: National Snow and Ice Data Center. doi: https://doi.org/10.7265/N5TB14TC. [Date Accessed]
