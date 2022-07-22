---
title: Water Budget Modelling
subtitle: Overland flow routing
author: M.Marchildon
date: 2019
output: html_document
---




# Overland flow routing

The continuity-based overland flow routing assumes a maximum overland flow capacity ($h_\text{max}$) on every cell. Overland flow is computed as the sum of upslope runon, groundwater discharge, saturation excess and impervious runoff and is here deemed "potential runoff" ($R_o$). Depending on cell slope and the roughness characteristics of the cell, the depth of "mobile" water that remains on the cell at the end of the timestep is computed by:

$$
	h = h_\text{max}\left(1-\exp\frac{-R_o}{b}\right),
$$

where $b$ is the empircal response coeficient. Runoff leaving the cell at the end of the timestep is defined as:

$$
	R=
	\begin{cases}
		R_o-h  \qquad &\text{if $R_o>h$} \\
		0 &\text{otherwise}
	\end{cases}
$$

TODO

$$
	C = \frac{\log\frac{S_\text{min}}{S}}{\log\frac{S_\text{min}}{S_\text{max}}}\left(C_\text{max}-C_\text{min}\right) + C_\text{min}
$$




# Mobile water storage [m]:
TODO

$$
	\Delta S_k=k_i+r+x-f_k-k_o,
$$

where $k = q\frac{\Delta t}{w}$ is the volumetric discharge in ($i$) and out ($o$) of the computational element, $r$ is the runoff (i.e., infiltration excess) generated on the land surface, $x$ is the excess runoff (i.e., saturated excess) casued by a high watertable, and $f_k$ is the volume of mobile storage infiltrating the soil zone; all units are [m].


Land surface and soil zone storage [m]:

$$
	\Delta S_h = y+f_k+f_g-r-a-g,
$$

where $y$ is the atmospheric yield (rainfall $+$ snowmelt), $f_g$ is groundwater infiltration into the soil zone from a high watertable, $a$ is evapotranspiration, and $g$ is recharge.


Groundwater storage [m]:

$$
	\Delta S_g = g-f_g-x,
$$


The overall water balance over each CE is then given by:

$$
	\Delta S_k+\Delta S_h+\Delta S_g=y+k_i-\left(k_o+a\right).
$$


![Schematic diagram of a computational element.](fig/CE-WB_sketch.png)
Schematic diagram of a computational element.


Other useful metrics include:

Net Groundwater Exchange (positive: recharge; negative: discharge):

$$
	g_\text{net}=g-f_g-x
$$
