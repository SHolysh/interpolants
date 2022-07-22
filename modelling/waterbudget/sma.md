---
title: Water Budget Modelling
subtitle: Soil moisture accounting
author: M.Marchildon
date: 2019
output: html_document
---


# Soil moisture accounting

The explicit soil moisture accounting (SMA—Todini and Wallis, 1977; O'Connell, 1991) algorithm used is intentionally simplified. Where this conceptualization differs from typical SMA schemes is that here there is no distinction made among a variety of storage reservoirs commonly conceptualized in established models. For example, all of interception storage, depression storage and soil storage (i.e., soil moisture below field capacity) is assumed representative by a single aggregate store, thus significantly reducing model parameterization. This conceptualization is allowed by the fact that this model is intended for regional-scale application ( $\gg$ 1000km²) with large computational elements ( $>$ 50 $\times$ 50m² cells). Users must consider the results of the model from a "bird's-eye view," where water retained at surface is indistinguishable from the many stores typically conceptualized in hydrological models. The SMA overall water balance is given as:

$$
	\Delta S = P-E-R-G,
$$

where $\Delta S$ is the change in (moisture) storage, $P$ is "active" precipitation (which includes direct inputs rainfall $+$ snowmelt, otherwise termed snowpack drainage ($P_\text{drn}$) and indirectect inputs such as runoff originating from upslope areas), which is the source of water to the computational element. $E$ is total evaporation (transpiration $+$ evaporation from soil and interception stores), $G$ is groundwater recharge, $R$ is excess water that is allowed to translate laterally. All of $E$, $R$, and $G$ are dependent on moisture storage ($S$) and $S_\text{max}$, the water storage capacity of the computational element. Both $E$ and $G$ have functional relationships to $S$ and $S_\text{max}$, and are dependent on other variables, whereas $R$ only occurs when there is excees water remaining (i.e., $S>S_\text{max}$), strictly speaking:

<!-- $$
    \Delta S = P-\left(E+G\right)\propto f\left(S,S_\text{max}\right)-R|_{S>S_\text{max}}
$$ -->

$$
    \Delta S = \underbrace{P}_{\substack{\text{sources}}}
        - \underbrace{f\left(E_p,K_\text{sat},S,S_\text{max}\right)}_{\text{sinks}}
        - \underbrace{R|_{S>S_\text{max}}}_{\text{excess}}
$$
