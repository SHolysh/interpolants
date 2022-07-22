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

### Model components:

1. use of regional climate fields (model inputs---precipitation, temperature, snowmelt, etc.)
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