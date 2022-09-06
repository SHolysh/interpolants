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

where $b$ is the empirical response coefficient. Runoff leaving the cell at the end of the timestep is defined as:

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


# Lateral water movement
## ...and soil moisture accounting scheme

Each model cell consists of a retention reservoir $S_h$ (where water has the potential to drain) and a detention reservoir $S_k$ (where water is held locally); both with a prefined capacity and susceptible to evaporation loss. In addition, any number of cells can be grouped which share a semi-infinite conceptual groundwater store $(S_g)$.

Although conceptual, storage capacities may be related to common hydrologic concepts, for instance:

$$
  S_h=(\phi-\theta_\text{fc}) z_\text{ext}
$$
and

$$
  S_k=\theta_\text{fc} z_\text{ext}+f_\text{imp} h_\text{dep}+f_\text{can}\cdot h_\text{can}\cdot\text{LAI}
$$
where $\phi$ and $\theta_\text{fc}$ are the water contents at saturation (porosity) and field capacity, respectively; $f_\text{imp}$ and $f_\text{can}$ are the fractional cell coverage of impervious area and tree canopy, respectively; $h_\text{dep}$ and $h_\text{can}$ are the capacities of impervious depression and interception stores, respectively; $\text{LAI}$ is the leaf area index and $z_\text{ext}$ is the extinction depth (i.e., the depth of which evaporative loss becomes negligeable). Change in detention storage at any time is defined by:

$$
	\Delta S_k=k_\text{in}+f_h+h_b\Delta t-\left(a_k+f_k+k_\text{out}\right),
$$

where $k = q\frac{\Delta t}{w}$ is the volumetric discharge in ($_\text{in}$) and out ($_\text{out}$) of the computational element, $x$ is the excess runoff (i.e., saturated excess) casued by a high watertable, and $f_k$ is the volume of mobile storage infiltrating the soil zone; all units are [m]. (note that groundwater discharge to streams---$h_b\Delta t$---only occurs at stream cells an is only a gaining term.) Also note that $k_\text{out}$ only occurs when water stored in $S_k$ exceeds its capacity.

Change in retention storage [m]:

$$
	\Delta S_h = y+f_k+x-\left(a_h+f_h+g\right),
$$

where $y$ is the atmospheric yield (rainfall + snowmelt), $x$ is groundwater infiltration into the soil zone from a high watertable, $a$ is evapotranspiration, and $g$ is recharge.

and change in groundwater storage [m] (i.e., Net Groundwater Exchange---positive: recharge; negative: discharge):

$$
	\Delta S_g = g-\left(x+h_b\Delta t\right).
$$

The overall water balance, with $a=a_h+a_k$, over each CE is then given by:

$$
	\Delta S_k+\Delta S_h+\Delta S_g=y+k_\text{in}-\left(a+k_\text{out}\right).
$$


![](fig/sma1.svg)
Conceptual soil moisture accounting scheme.


## Overland Flow

Laterally moving water leaving the cell $k_\text{out}$ is corrected (primarily) based on landsurface slope and roughness. Water leaving a cell can only be directed to a single cell, but any cell can recieve from multiple cells. A "cascade factor", $f_\text{casc}=f(\text{slope},\text{roughness},\text{etc.})$, $0\leq f_\text{casc} \leq 1$, is applied to the current volume stored in the mobile store (in the form of a linear reservoir):

$$
  k_\text{out}=f_\text{casc}S_k^+,
$$
where $S_k^+>0$ is water in the detention store in excess of the store's capacity, and:

$$
  f_\text{casc}=1-\exp\left(-a\frac{\beta^2}{r^2}\right),
$$

where $\beta$ is land surface gradient, $r$ is called the "range" (note that the above equation is identical to the Gaussian variogram model), and $a$ is a scaling factor applied to the range such that it's value approaches unity at $r$; Below are examples with $a\approx 5$:

<!-- ```{r fcasc, echo=FALSE, fig.width=6,fig.height=4,fig.align='center'}
a = 5
fun.1 <- function(x) 1-exp(-a*x^2/0.25^2)
fun.2 <- function(x) 1-exp(-a*x^2/0.5^2)
fun.3 <- function(x) 1-exp(-a*x^2/0.75^2)
fun.4 <- function(x) 1-exp(-a*x^2)
fun.5 <- function(x) 1-exp(-a*x^2/1.5^2)

ggplot(data.frame(x = 0),aes(x=x)) + # dummy dataframe
  theme_bw() +
  stat_function(fun = fun.1) + 
  stat_function(fun = fun.2) + 
  stat_function(fun = fun.3) + 
  stat_function(fun = fun.4) + 
  stat_function(fun = fun.5) +
  annotate("text", label = "r=.25", x = .07, y = .75) +
  annotate("text", label = "r=.5", x = .2, y = .72) +
  annotate("text", label = "r=.75", x = .3, y = .69) +
  annotate("text", label = "r=1", x = .415, y = .66) +
  annotate("text", label = "r=1.5", x = .555, y = .57) +
  xlim(0,1) + labs(x="gradient",y=expression(f["casc"]))
``` -->

Special conditions are set for $f_\text{casc}$: All stream cells have $f_\text{casc}=1$, meaning that the the mobile water store remains 100% mobile. Here $b_\text{casc}=1$.

## Groundwater Recharge

Groundwater recharge at cell $i$ is computed as:

$$
  g=\min\left(S_h,K_\text{sat}\Delta t\right)+S_k^+\left[1-\exp\left(-\frac{K_\text{sat}}{L}\Delta t\right)\right].
$$

It is important that the second term of the groundwater recharge equation remain... **blah blah cascade towers**. Here the interface length $L$=10cm globally.