---
title: Water Budget Modelling
subtitle: Shallow groundwater model
author: M.Marchildon
date: 2019
output: html_document
---

* TOC
{:toc}

# Shallow groundwater model

Typically, the handling of groundwater in rainfall-runoff models involves the use of a groundwater storage reservoir with a linear/exponential decay function. When calibrating these models to long-term stream flow records, this model structure can be quite successful at matching baseflow recession. Given its success and it relative simplicity, the linear decay groundwater storage model is widely used and preferred for model calibration to stream flow gauging stations. However, implied in the linear decay groundwater storage model is the assumption that groundwater recharge is unidirectional and uniformly distributed basin-wide, or, at most, among common surficial material types. Consequently, catchment topography is neglected, which has a large role in distributing recharge, both in terms of focused recharge at overland flow convergence points and rejected recharge at low-lying areas where the groundwater table is close to, at, or above ground surface. For the purpose of this model, which is primarily intended for distributed recharge estimation, the linear decay groundwater storage model on its own will not suffice.

The groundwater model structure employed here follows the TOPMODEL structure of Beven and Kirkby (1979). TOPMODEL employs a distribution function that projects predicted groundwater discharge at surface from a simple groundwater storage reservoir. The distribution function is referred to as the "soil-topologic index" (Beven, 1986) which considers distributed surficial material properties, topography and upslope contributing areas. TOPMODEL does not explicitly model integrated groundwater/surface water processes, however it does provide the ability, albeit an indirect one, to include the influence of near-surface watertables. Most importantly, low-lying areas close to drainage features will be prevented from accepting recharge resulting in saturated overland flow conditions. 

Assumptions to TOPMODEL are that the groundwater reservoir is lumped meaning it is equivalently and instantaneously connected to all locations of the catchment constrained by topographic divides. This assumption also implies that recharge computed over a watershed is uniformly applied to the water table. The groundwater reservoir is thought to be a shallow, single-layered, unconfined aquifer whose hydraulic gradient can be approximated by local surface slope. 

Lateral (downslope) transmissivity ($T$) is then assumed to scale exponentially with soil moisture deficit ($D$) or depth to the watertable ($z_\text{wt}$) by:

$$
	T = T_o\exp{\left(\frac{-D}{m}\right)} = T_o\exp{\left(\frac{-\phi z_\text{wt}}{m}\right)},  
$$

where $T_o$ is the saturated lateral transmissivity per unit width [m $^2$/s] that occurs when no soil moisture deficit exists (i.e., $D=0$), $\phi$ is mean porosity, and $m$ is a scaling factor [m]. Consequently, recharge computed in the distributed sense is aggregated when added to the lumped reservoir, yet discharge to surface remains distributed according to the soil-topographic index. For additional discussion on theory and assumptions, please refer to Beven et.al., (1995), Ambroise et.al., (1996) and Beven (2012).

At the regional scale, multiple TOPMODEL reservoirs have been established to represent groundwater dynamics in sub-catchments (or groundwatersheds) of a pre-specified threshold area, here set at 10 km $^2$. Groundwater discharge estimates from each TOPMODEL reservoir instantaneously contributes lateral discharge to stream channel cells within every sub-catchment according to Darcy's law:

$$
	q = T\tan\beta,
$$

where $q$ is interpreted here as groundwater discharge per unit length of stream [m $^2$/s], and $\tan\beta$ is the surface slope angle in the downslope direction, assumed representatitve of the saturated zone's hydraulic gradient. The sub-surface storage deficit ($D_i$) at some location $i$ is determined using the distribution function:

$$
	D_i = \overline{D} + m \left[\gamma - \ln\left(\frac{a}{T_o \tan \beta}\right)_i\right],
$$

where $a$ is the unit contributing area to cell $i$ [m], defined here as the total contributing area to cell $i$ divided by the cell's width. The expression $\ln(a/T_o \tan \beta)$ is referred to as the soil-topologic index, as discussed above. $\gamma$ is the catchment average soil-topologic index and is determined by:

$$
	\gamma = \frac{1}{A}\sum_i A_i\ln\left(\frac{a}{T_o \tan \beta}\right)_i,
$$

where $A$ and $A_i$ is the sub-watershed and computational element (i.e., cell) areas, respectively. Before every time step, basin-wide average moisture deficit is updated by:

$$
	\overline{D}_t = \frac{1}{A}\sum_i A_i D_i - G_{t-1} + B_{t-1},
$$

where $G_{t-1}$ is the total sub-watershed-wide groundwater recharge computed during the previous time step [m], and $B_{t-1}$ is the normalized groundwater discharge to streams, also computed during the previous time step [m]. The volumetric rate of groundwater discharge to streams $Q_b$ [m$^3$/s] is given by:

$$
	Q_b = \frac{AB}{\Delta t} = \sum_{i=1}^Ml_iq_i,
$$

and $l_i$ is length of channel in stream cell $i$, here assumed constant and equivalent to the uniform cell width $w$, $\Delta t$ is the model time step and $M$ is the total number of cells containing mapped streams within the sub-watershed.

It should be noted that the above formulation is the very similar to the linear decay groundwater storage model, except here, TOPMODEL allows for an approximation of spatial soil moisture distribution, which will, in turn, determine spatial recharge patterns, as $D_i\leq 0$ will prevent recharge from occurring at cell $i$. Initial watershed average soil moisture deficit can also be determined from streamflow records by re-arranging equations in Beven (2012):

$$
	%\overline{D}_{t=0} = -m \ln\left(\frac{Q_{t=0}}{Q_o}\right),
	%\overline{D}_{t=0} = -m\left[\gamma +\ln\left(\frac{Q_{t=0}}{\sum_{i=1}^Ml_ia_i}\right)\right],
	\overline{D}_{t=0} = -m\left[\gamma +\ln\left(\frac{Q_{t=0}}{A}\right)\right],
$$

where $Q_{t=0}$ is the measured stream flow known at the beginning of the model run. The parameter $m$ can be pre-determined from baseflow recession analysis (Beven et.al.,, 1995; Beven, 2012).

Although plausible (and argued for), TOPMODEL can be allowed to have negative soil moisture deficits (i.e., $D_i<0$), meaning that water has ponded and overland flow is being considered. For the current model, ponded water and overland flow is being handled by the explicit soil moisture accounting (SMA) system and the lateral shallow water movement module discussed below. Conditions where $D_i<0$ for any computational element will have the excess water removed and added to the SMA system and $D_i$ will be set to zero. Under these conditions, no recharge will occur at this cell and the TOPMODEL portion of the model will only be used to represent groundwater discharge to streams.



<!-- Lateral flux is approximated with steady-state Darcian flows under a uniform recharge rate with an effictive hydraulic gradient approximated by local surface slope (Ambroise et.al.,, 1996).

$$
	q = aR = T\tan\beta,
$$

where $a$ is the contributing area [m $^2$], $R$ is the  -->

