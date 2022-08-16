---
title: Water Budget Modelling
subtitle: Input Data
author: M.Marchildon
date: 2019
output: html_document
---

* TOC
{:toc}


# Input Data
6-hourly









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
