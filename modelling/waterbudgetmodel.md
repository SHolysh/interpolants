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

<iframe src="https://golang.oakridgeswater.ca/pages/subwatersheds.html"></iframe>


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


## Time-stepping

The model runs using a variable time stepping approach. The soil moisture accounting portion of the model runs on a 6-hourly basis in order to capture hydrological dependence on the diurnal fluctuation of energy flux at the ground surface, and the local inertial approximation SWE simulation time step varies to a maximum of 6-hours based on its computational requirements (see \sref{sec:LIA}).


## Model Breakdown

The "model" described herein is actually two separate hydrological analysis tools. The first is a long-term data assimilation system that utilizes state-of-the-art data products and leverages the ORMGP's database. This data assimilation system (DAS) differs from a model (in the traditional sense) in that it is used to \emph{interpolate} the hydrological function of the landscape given the data that are available. The computational structure of the DAS is specifically tailored to run as fast as possible, as there are thousands of projections continuously made on ORMGP servers. The results from this analysis will only be offered as likely ranges in long term seasonal water balance estimates.

The second tool is intended as a predictive model (this time, in the traditional sense), but is intended for short-term (i.e., less than a month) projections. The design of this model is more physically rigorous but will always be dependent on other models (or the DAS described here) to prescribe antecedent conditions. While computational time for this model is orders of magnitudes greater that the DAS, it continuous to share the philosophy of designed whereby computation efficiency is paramount.

Combined, these two tools will provide all water groundwater resources needs for the ORMGP partners in order to:

1. utilize readily available datasets provided daily by government agencies of Canada and the United States
1. interpolate long-term hydrological water budget components with an emphasis on its certainty
1. extrapolate near-term hydrological function from a given/known antecedent state




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