---
title: Water Budget Modelling
author: M.Marchildon
date: 2022-07-22
output: html_document
---

* TOC
{:toc}


# Introduction

A regionally-distributed runoff/recharge model has been developed to simulate hydrologic processes at a fine (grid-based) scale. The model conceptualization has been written to a large-scale yet high-resolution distributed numerical model code optimized for implementation on parallel computer architectures. No process of the model is in any way novel, rather a suite of existing model structures have been combined into one platform chosen specifically for their ease of implementation, practical applicability and computational scalability.

The model's primary intention is to interpolate watershed moisture conditions for the purposes of estimating regional groundwater recharge, given all available data. While the model can project stream flow discharge and overland runoff at any point in space, the model is not optimized, nor intended to serve as your typical rainfall runoff model. Outputs from this model will most often be used to constrain regional groundwater models within the 30,000km² Oak Ridges Moraine Groundwater Program jurisdiction of southern Ontario.

This model is currently in beta mode, and its use for practical application should proceed with caution. Users should be aware that model results posted online will always be subject to change. Ultimately, the intent of this model is to produced ranges of long-term (monthly average) water budgeting as a hydrological reference for the partners of the ORMGP. 



The model is physically based in that mass is conserved and it is not constrained to any particular timestep. Most parameters are common, percent impervious, conductivity of surficial soils, etc.

## Summary
- model run from the hydrological water year 2010 (2009-10-01) through water year 2020. 
- Potential evaporation determined using the empirical wind functions of Penman (1948)
- Rainfall was collected from capa
- Snowmelt was collected from SNODAS
- T, rh from MSC
- 12.1M x 50x50 m cells x 6-hourly ts x 10 years

The following is a description of the water budget tool located on our website, hereinafter referred to as the "*model*".

### Model components:

1. use of regional climate fields (model inputs—precipitation, temperature, snowmelt, etc.)
1. explicit soil moisture accounting scheme
1. cold content energy balance snowpack model
1. distribution function-based shallow groundwater model with groundwater feedback mechanisms
1. a one-dimensional fully-implicit multi-layer solution to multi-phase flow through porous media
1. a one-dimensional first-order kinematic overland flow module	
1. local inertial approximation of the shallow water equation for lateral movement of water


## Computational Elements

Model grids in the model represent a homogenized area on the land surface, similar to the concept of a Hydrologic Response Unit, or HRU. The term HRU will be avoided here as the concept itself is not well defined and the term "computational element" (CE) will be used herein to avoid ambiguity. All processes within the computational element are assumed to be zero-dimensional, that is processes are modelled at the point-scale, which effectively represents the "average" condition within the grid cell. The spatial proximity of each computational element is maintained, meaning that runoff occurring from an upgradient CE will runon to an adjacent downgradient CE, thus preserving the spatial distribution of land surface moisture.

## Model Structure

Runoff is conceptualized as being generated through the saturation excess (Dunne, 1975 CHECK) mechanism. Land area that has the capacity to retain water (through interception, soil retention, depression/rill storage, etc.) must be satisfied. The saturation excess mechanism is dependent on topography and it's interaction with the groundwater system; thus the model is distributed (cell-based) and has an integrated (albeit conceptual) groundwater system.

Surface water-groundwater integration is viewed from the hydrologists' perspective: areas where is the groundwater system limiting infiltration (shallow groundwater table) and even contributing to overland runoff (springs/seepage areas). As a model output, this can be quantified as a net groundwater exchange field (groundwater recharge and discharge)



dunnian processes.



The purpose of the model is to have a 



The purpose of the model is to account for the water balance  estimate the spatial and temporal distribution of water on the landscape.



The basis of the model is that topography is paramount to the lateral movement of water yielding runoff. The model is deemed regional, in that it covers a large areal extent, yet is kept to a fine resolution to ensure that observed geomorphic flow patterns are represented in the model.


## Time-stepping

The model runs using a variable time stepping approach. The soil moisture accounting portion of the model runs on a 6-hourly basis in order to capture hydrological dependence on the diurnal fluctuation of energy flux at the ground surface, and the local inertial approximation SWE simulation time step varies to a maximum of 6-hours based on its computational requirements.

The time step of the model has been set to 6 hour steps on 00:00 UTC, yielding the local time steps of 01:00, 07:00, 13:00, and 19:00 EST. This step is chosen as it matches the received precipitation dataset described below.


## Model Breakdown

The "model" described herein is actually two separate hydrological analysis tools. The first is a long-term data assimilation system that utilizes state-of-the-art data products and leverages the ORMGP's database. This data assimilation system (DAS) differs from a model (in the traditional sense) in that it is used to \emph{interpolate} the hydrological function of the landscape given the data that are available. The computational structure of the DAS is specifically tailored to run as fast as possible, as there are thousands of projections continuously made on ORMGP servers. The results from this analysis will only be offered as likely ranges in long term seasonal water balance estimates.

The second tool is intended as a predictive model (this time, in the traditional sense), but is intended for short-term (i.e., less than a month) projections. The design of this model is more physically rigorous but will always be dependent on other models (or the DAS described here) to prescribe antecedent conditions. While computational time for this model is orders of magnitudes greater that the DAS, it continuous to share the philosophy of designed whereby computation efficiency is paramount.

Combined, these two tools will provide all water groundwater resources needs for the ORMGP partners in order to:

1. utilize readily available datasets provided daily by government agencies of Canada and the United States
1. interpolate long-term hydrological water budget components with an emphasis on its certainty
1. extrapolate near-term hydrological function from a given/known antecedent state



# Model structure and physical constraints

### Digital Elevation Model

The Greater Toronto Area 2002 DEM (OMNRF, 2015) was re-sampled to the model's 50x50m grid cell resolution. Surface depressions were removed using Wang and Liu (2006) and flat regions were corrected using Martz (1997).

Drainage directions and flowpaths of the now "hydrologically correct" DEM were were assigned based on the direction of steepest decent (D8).Cell gradients ($b$) and slope aspects were calculated based on a 9-cell planar interpolation routine. The unit contributing area $a=A/w$ topographic wetness index $ln\frac{a}{\tan b}$ (Beven and Kirkby, 1979--CHECK) were computed for every cell.

### Sub-basins

The 30,000 km² model area has been sub divided into 2,813 approximately 10 km² sub-basins. Within these sub-basins:
1. Meteorological forcings from external sources are aggregated applied uniformly within the sub-basin (via a pre-processing routine); and
1. Local shallow water response is assumed to act uniformly (the shallow subsurface catchment area).

<!-- ![](fig/fig-sws1.png) -->
<iframe src="https://golang.oakridgeswater.ca/pages/subwatersheds.html" width="100%" height="400" scrolling="no" allowfullscreen></iframe>





# Model Parameterization

Although a distributed model, the procedures applied at the cell scale are quite parsimonious. There is no separate treatment of interception, depression storage, nor soil water retention, rather it is assumed that these processes respond to environmental factors (e.g., evaporation) in parallel and thus can be treated in bulk.

From the top down perspective, viewing some 12.1 million 50x50m cells covering 30,000 km², it seems rather overcomplicated (possibly frivolous) to account water any more than to total mass present at any particular (lateral) location.

#### Global parameters

Cells with a contributing area ($A$) greater than 1 km² are deemed "stream cells" in which additional sources include groundwater discharge to streams.

#### Sub-basin parameters

Groundwater processes in the model are conceptualized at the sub-basin scale and so much of the groundwater parameterization is implimented here.

#### Cell-based parameters

Each cell is classified according to (i) surficial geology and (ii) land use mapping where each class is parameterized independently. "Look-up tables" are used to distribute model parameters accordin to their classification.

##### Parameters:
- impervious fraction
- retention/storage capacity
- depression storage
- percolation rates

## Parameters
> (merge with above)

#### globally applied, cell-base parameters

- $D_\text{inc}$: depth of incised channel relative to cell elevation [m] (note, while it is possibile to assing this parameter on a cell basis, it was treated here as a global "tuning" parameter.)
- $f_\text{casc}$: cascade fractions are based on a pre-determined relationship with local gradients.

#### regional (sub-watershed based)

- $m$: TOPMODEL groundwater scaling factor [m]

#### land use based

- $F_\text{imp}$: fraction of impervious cover
- $F_\text{can}$: fraction of canopy cover
- $\phi$: porosity
- $F_c$: field capacity
- $d_\text{ext}$: extinction depth of soil, where evapotranspiration cease to occur.

#### surficial geology based

- $K_\text{sat}$: hydraulic conductivity as saturation

## Glossary

- masl - metres above sea level

- ECCC - Environment Canada and Climate Change

- atmospheric yeild: term used to describe water (in liquid form) that is actively altering the hydrologic state at a particular location.




# Inputs

The model's structure is defined rather simply by at least 5 raster data sets. Given these data, the model's pre-processor will generate additional information based on these data:

1. digital elevation model
1. land use index (with parameter lookup table)
1. surficial geology index (with parameter lookup table)
1. groundwater system index


# Model Variables

One goal set for the model design was to leverage contemporary gridded data sets available from a variety of open and public sources. Products known as "*data re-analysis products*" or "*data-assimilation products*" attempt to merge meteorological information from a variety of sources, whether they be ground (station) measurements, remote sensing observations (e.g., radar, satellite, etc.), and archived near-cast weather model outputs.  When combined, the gridded data resemble the spatial distribution of meteorological forcings unlike what can be accomplished through standard interpolation practices using point measurements (e.g., thiessen polygons, inverse distance weighting, etc.).

An advantage to the data-assimilation products is that it removes the modeller from needing to model certain processes explicitly. Here, for example, the model does not account for a snowpack, rather inputs to the model include snowmelt derived from SNODAS.

The extent of the model combined with the resolution of the processes simulated lends itself best viewed from a top-down perspective (REF). This allows for model simplification by which many of the layered water stores (i.e., interception, depression, soil, etc.) may be handled procedurally as one unit. Viewing the model domain in it's 30,000 km² extents,one can imagine how difficult it would be to discern any vertical detail.




# Theory

## Long-term data assimilation system

* **[Input data](/interpolants/modelling/waterbudget/data.html)**
* **[Soil moisture accounting](/interpolants/modelling/waterbudget/sma.html)**
* **[Total evaporation](/interpolants/modelling/waterbudget/pet.html)**
* **[Snowmelt](/interpolants/modelling/waterbudget/snowmeltCCF.html)**
* **[Shallow groundwater](/interpolants/modelling/waterbudget/gw.html)**
* **[Overland flow routing](/interpolants/modelling/waterbudget/overlandflow.html)**
* **[Model Parameters and Sampling](/interpolants/modelling/waterbudget/parameters.html)**
* **[Model Structure](/interpolants/modelling/waterbudget/structure.html)**
* **[Future plans](/interpolants/modelling/waterbudget/future.html)**


## Short-term infiltration model

* **[Lateral water movement](/interpolants/modelling/infiltration/lateral.html)**
* **[Multiphase flow through porous media](/interpolants/modelling/infiltration/pmflow.html)**
* test


# References



<!-- from intro -->

- Todini, E., Wallis, J.R., 1977. Using CLS for daily or longer period rainfall-runoff modelling. In Mathematical Models for Surface Water Hydrology (edited by T.A. Ciriani, U. Maione and J.R. Wallis): John Wiley, London, UK. — (from Tucci, C.E.M., R.T. Clarke, 1980. Adaptive forecasting with a conceptual rainfall—runoff model.)

- O'Connell, P.E., 1991. A historical perspective. In Recent advances in the modeling of hydrologic systems ed. D.S. Bowles and P.E. O'Connell: 3—30.


<!-- from pet -->

- Priestley, C.H.B. and R.J. Taylor, 1972. On the assessment of surface heat flux and evaporation using large scale parameters. Monthly Weather Review 100:81--92.

<!-- - Makkink, G.F., 1957. Testing the Penman formula by means of lysimeters. International Journal of Water Engineering 11:277--288. -->
- Makkink, G.F., 1957. Ekzameno de la Formulo de Penman. Netherlands Journal of Agricultural Science 5:290—305.

- Nov\'ak, V., 2012. Evapotranspiration in the Soil-Plant-Atmosphere System. Springer Science and Business Media, Dordrecht. 253pp.

- Budyko, M.I., 1974. Climate and Life. Academic Press, Noew York and London. 508pp.


<!-- from snow -->

- DeWalle, D.R. and A. Rango, 2008. Principles of Snow Hydrology. Cambridge University Press, Cambridge. 410pp.

- Judson, A. and N. Doesken, 2000. Density of Freshly Fallen Snow in the Central Rocky Mountains. Bulletin of the American Meteorological Society 81(7):1577—1587.

- Martinec, J., 1960. The degree-day factor for snowmelt-runoff forecasting. in Proceesings General Assembly of Helsinki, Commission on Surface Waters, IASH Publication 51:468—477.


<!-- from GW -->
- Ambroise, B., K. Beven, J. Freer, 1996. Toward a generalization of the TOPMODEL concepts: Topographic indices of hydrological similarity, Water Resources Research 32(7): pp.2135—2145.
- Beven, K.J., 1986. Hillslope runoff processes and flood frequency characteristics. Hillslope Processes, edited by A.D. Abrahams, Allen and Unwin, Winchester, Mass. pp.187—202.
- Beven, K.J., R. Lamb, P.F. Quinn, R. Romanowicz, and J. Freer, 1995. TOPMODEL. In Singh V.P. editor, Computer Models of Watershed Hydrology. Water Resources Publications, Highland Ranch, CO: pp. 627—668.
- Beven, K.J. and M.J. Kirkby, 1979. A physically based variable contributing area model of basin hydrology. Hydrological Science Bulleton 24(1):43—69.
- Beven, K.J., 2012. Rainfall-Runoff modelling: the primer, 2nd ed. John Wiley & Sons, Ltd. 457pp.



<!-- from LIA -->

- Almeda et.al., (2012)


<!-- from overland routing -->

- Neitsch, S.L., J.G. Arnold, J.R., Kiniry, J.R. Williams, 2011. Soil and Water Assessment Tool: Theoretical Documentation Version 2009 (September 2011). 647pp.

- Ponce, V.M., 1989. Engineering Hydrology: Principles and Practices. Prentice Hall, NJ. 640pp.

- Williams J.R., 1969. Flood routing with variable travel time or variable storage coefficients. Transactions of the ASAE 12(1): 100—103.


<!-- porous media -->

- Bittelli, M., Campbell, G.S., and Tomei, F., 2015. Soil Physics with Python. Oxford University Press.

- Campbell, G.S., 1974. A simple method for determining unsaturated conductivity from moisture retention data. Soil Science, 117: 311—387.

- Monteith, J.L.,1965. Evaporation and environment. Symposia of the Society for Experimental Biology 19: 205—224.

- Nov\'ak, V., 2012. Evapotranspiration in the Soil-Plant-Atmosphere System. Springer. 253pp.

- Penman, H.L., 1940. Gas and vapour movements in the soil: I. The diffusion of vapours through porous solids. Journal of Agricultural Science 30(3): 437—462.

- Penman, H.L., 1948. Natural evaporation from open water, bare soil and grass. Proceedings of the Royal Society of London. Series A, Mathematical and Physical Sciences 193(1032): 120—145.

- Philip, J.R., 1957. Evaporation, and moisture and heat fields in the soil. Journal of Meteorology 14: 354—366.

- Richards, L.A., 1931. Capillary conduction of liquids through porous media. Physics 1: 318—333.



<!-- MC -->

- Lemieux, C., 2009. Monte Carlo and Quasi-Monte Carlo Sampling. Springer Science. 373pp.



<!-- EXTRA?? (need to check) -->
Beven, K.J.; Kirkby, M. J. (1979). "A physically based, variable contributing area model of basin hydrology". Hydrological Science Bulletin. 24 (1): 43–69. doi:10.1080/02626667909491834

Garbrecht Martz 1997 The assignment of drainage direction over flat surfaces in raster digital elevation models

Gupta etal 2009 Decomposition of the mean squared error and NSE performance criteria- Implications for improving hydrological modelling

Monteith, J.L., 1965. Evaporation and environment. Symposia of the Society for Experimental Biology 19: 205–224.

National Operational Hydrologic Remote Sensing Center. 2004. Snow Data Assimilation System (SNODAS) Data Products at NSIDC, Version 1. [Indicate subset used]. Boulder, Colorado USA. NSIDC: National Snow and Ice Data Center. doi: https://doi.org/10.7265/N5TB14TC. [Date Accessed]

Ontario Ministry of Natural Resources and Forestry, 2015. GTA DEM 2002 User Guide. Queen’s Printer for Ontario. 14pp.

Penman, H.L., 1948. Natural evaporation from open water, bare soil and grass. Proceedings of the Royal Society of London. Series A, Mathematical and Physical Sciences 193(1032): 120-145.

Wang, L., H. Liu, 2006. An efficient method for identifying and filling surface depressions in digital elevation models for hydrologic analysis and modelling. International Journal of Geographical Information Science 20(2): 193-213.
