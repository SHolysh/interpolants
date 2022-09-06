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

Assumptions to TOPMODEL are that the groundwater reservoir is lumped meaning it is equivalently and instantaneously connected over its entire spatial extent. This assumption implies that recharge computed to a groundwater reservoir is uniformly applied to its water table. The groundwater reservoir is thought to be a shallow, single-layered, unconfined aquifer whose hydraulic gradient can be approximated by local surface slope. 

Furthermore, lateral (downslope) transmissivity ($T$) is then assumed to scale exponentially with soil moisture deficit ($D$) or depth to the watertable ($z_\text{wt}$) by:

$$
	T = T_o\exp{\left(\frac{-D}{m}\right)} = T_o\exp{\left(\frac{-\phi z_\text{wt}}{m}\right)}, 
$$

where $T_o=wK_\text{sat}$ is the saturated lateral transmissivity perpendicular to an isopotential contour of unit length ($w$) [m$^2$/s] that occurs when no soil moisture deficit exists (i.e., $D=0$), $\phi$ is effective porosity, and $m$ is a scaling factor [m]. Consequently, recharge computed in the distributed sense is aggregated when added to the lumped reservoir, yet discharge to surface remains distributed according to the soil-topographic index. For additional discussion on theory and assumptions, please refer to Beven etal (1995), Ambroise etal, (1996) and Beven (2012).

At the regional scale, multiple TOPMODEL reservoirs have been established to represent groundwater dynamics in greater physiographic regions, where it is assumed that material properties are functionally similar and groundwater dynamics are locally in sync. Groundwater discharge estimates from each TOPMODEL reservoir instantaneously contributes lateral discharge to stream channel cells according to Darcy's law:

$$
	q = T\tan\beta,
$$
<!-- $$
	q = T\tan\beta = T_o\tan\beta\cdot e^{\left(\frac{-D}{m}\right)},
$$ -->

where $q$ is interpreted here as groundwater discharge per unit length of stream [m²/s], and $\tan\beta$ is the surface slope angle in the downslope direction, assumed representative of the saturated zone's hydraulic gradient. The sub-surface storage deficit ($D_i$) at some location $i$ is determined using the distribution function:


$$
	D_i = \overline{D} + m \left(\gamma - \zeta_i\right),
$$
<!-- $$
	D_i = \overline{D} + m \left[\gamma - \ln\left(\frac{a}{T_o \tan \beta}\right)_i\right],
$$ -->

where $\zeta_i$ is the soil-topologic index at cell $i$, defined by $\zeta=\ln(a/T_o \tan \beta)$ and $a_i$ is the unit contributing area defined here as the total contributing area to cell $i$ divided by the cell's width.  $\gamma$ is the catchment average soil-topologic index:

$$
	\gamma = \frac{1}{A}\sum_i A_i\zeta_i,
$$
<!-- $$
	\gamma = \frac{1}{A}\sum_i A_i\ln\left(\frac{a}{T_o \tan \beta}\right)_i,
$$ -->

where $A$ and $A_i$ is the sub-watershed and computational element (i.e., cell) areas, respectively. 

<!-- Rearranging the above terms, keeping non-variable terms on the $RHS$:
$$
  \delta D_i=D_i-\overline{D}=m\left(\gamma - \zeta_i\right)
$$
where $\delta D_i$ is the groundwater deficit relative to the regional mean. From this, at any time $t$, the local deficit is thus:
$$
 D_{i,t} = \delta D_i+\overline{D}_t
$$ -->

Before every time step, average moisture deficit of every groundwater reservoir is updated at time $t$ by:

$$
	\overline{D}_t = \frac{1}{A}\sum_i A_i D_i - G_{t-1} + B_{t-1},
$$

where $G_{t-1}$ is the total groundwater recharge computed during the previous time step [m], and $B_{t-1}$ is the normalized groundwater discharge to streams, also computed during the previous time step [m]. The volumetric rate of groundwater discharge to streams $Q_b$ [m³/s] is given by:

$$
	Q_b = \frac{AB}{\Delta t} = \sum_{i=1}^Ml_iq_i,
$$

and $l_i$ is length of channel in stream cell $i$, here assumed constant and equivalent to the uniform cell width ($w$) times a sinuosity factor ($\Omega$), $\Delta t$ is the model time step and $M$ is the total number of model cells containing mapped streams within the sub-watershed. 

<!-- For every stream cell, groundwater flux to stream cells $h_b$ [m/s] at time $t$ is given by:

$$
  h_{b,i}=\frac{l_iq_i}{A_i}=Q_o\exp\left(\frac{D_\text{inc}-D}{m}\right) \qquad i \in \text{streams}
$$
where discharge to stream cell $i$ at saturated conditions $(D_\text{inc}-D=0)$:

$$
  Q_o=\Omega\cdot \frac{T_o\tan\beta}{w}
$$

and thus basin-wide groundwater discharge to streams:

$$
  B = \Delta t \sum_{i=1}^M h_{b,i}
$$

where $D_\text{inc}$ is an offset, meant to accommodate the degree of channel incision (i.e., the difference between channel elevation and cell elevation). -->

It should be noted that the above formulation is the very similar to the linear decay groundwater storage model, except here, TOPMODEL allows for an approximation of spatial soil moisture distribution, which will, in turn, determine spatial recharge patterns, as $D_i\leq 0$ will prevent recharge from occurring at cell $i$. 

Initial watershed average soil moisture deficit can also be determined from streamflow records by re-arranging equations in Beven (2012):

$$
	%\overline{D}_{t=0} = -m \ln\left(\frac{Q_{t=0}}{Q_o}\right),
	%\overline{D}_{t=0} = -m\left[\gamma +\ln\left(\frac{Q_{t=0}}{\sum_{i=1}^Ml_ia_i}\right)\right],
	\overline{D}_{t=0} = -m\left[\gamma +\ln\left(\frac{Q_{t=0}}{A}\right)\right],
$$

where $Q_{t=0}$ is the measured stream flow known at the beginning of the model run. The parameter $m$ can be pre-determined from baseflow recession analysis (Beven et.al., 1995; Beven, 2012).

Although plausible (and argued for), TOPMODEL can be allowed to have negative soil moisture deficits (i.e., $D_i<0$), meaning that water has ponded and overland flow is being considered. For the current model, ponded water and overland flow is being handled by the [explicit soil moisture accounting (SMA) system](/interpolants/modelling/waterbudget/sma.html) and the [lateral shallow water movement module](/interpolants/modelling/waterbudget/overlandflow.html). 

Conditions where $D_i<0$ for any computational element will have the excess water $(x)$ removed and added to the SMA system and $D_i$ will be set to zero. Under these conditions, no recharge will occur at this cell. The TOPMODEL portion of the model will only be used to represent groundwater discharge to streams.

This exchange, in addition to groundwater discharge to streams, represents the distributed interaction the groundwater system has on the surface. Given cell-based recharge ($g$) computed by the SMA, basin recharge $(G)$ is updated to determine net groundwater exchange:

$$
  G_{t+1} = G_i+\sum_i(g-x)_i
$$



<!-- Lateral flux is approximated with steady-state Darcian flows under a uniform recharge rate with an effective hydraulic gradient approximated by local surface slope (Ambroise et.al., 1996).
$$
	q = aR = T\tan\beta,
$$
where $a$ is the contributing area [m²], $R$ is the  -->


