---
title: Water Budget Modelling
subtitle: Total Evaporation
author: M.Marchildon
output: html_document
---

* TOC
{:toc}

# Total Evaporation

Total evaporation, i.e., loss from the computational element to the atmosphere, including plant transpiration, evaporation from land surface, soil pores and interception stores, is dependent on the current soil moisture storage ($S$), soil moisture capacity ($S_\text{max}$) and potential evapotranspiration ($E_p$), which is interpreted here as the capacity for the atmosphere to remove moisture from a saturated and replenishable surface.

Four varieties of potential evapotranspiration is explored, all varying in simplicity that depends on data availability. The first is dependent only on the day of year and an estimate of long-term average annual $E_p$. Based on assumed southern Ontario conditions (i.e., evaporation generally limited by soil moisture), the simplest method that can be used is the sine curve function given by:

$$
	E_p = \left(\overline{E}_p-E_\text{base}\right)\left[1+\sin\left(\frac{2\pi}{365}\left(i-\varphi\right)-\frac{\pi}{2}\right)\right]+E_\text{base},
$$

where $\overline{E}_p$ is the average annual $E_p$, $E_\text{base}$ is the minimum annual $E_p$, $i$ is the day of year from January 1st and $\varphi$ is the day offset from $i$ when $E_\text{base}$ occurs.

The next means of computing potential evaporation is based on the empirical Makkink (1957) method, which is functionally identical to the more common Priestly-Taylor (1972) approach. The main difference is that the Makkink method utilizes total incoming short-wave radiation, or "global radiation" ($K^\downarrow$), which is more readily available than net radiation $(Q^\text{*})$ used in the Priestly-Taylor approach and other so-called combination approaches. In either case, both $K^\downarrow$ and $Q^\text{*}$ require observational measurements from sources that are rare and not adequately distributed at the regional scale the model is applied to. Instead, global radiation is approximated using a Prescott-type equation (Nov\'ak, 2012): <!-- pg.232 -->

$$
	K^\downarrow = \left(a+b\frac{n}{N}\right)K_e,
$$

where $K_e$ is the extra-terrestrial short-wave flux (MJ/m²) that can be computed on the basis of solar irradiation theory that takes account of latitude, time of year and surface slope and aspect, $a$ and $b$ are empirical coefficients, ranging from 0.18 to 0.4 (mean $\approx$ 0.27) and 0.42 to 0.56 (mean $\approx$ 0.52), respectively, $n$ is the number of sunshine hours and $N$ is the total possible number of sunshine hours on a clear/cloudless day. The Makkink (1957) equation for $E_p$ [m/d] is given by:

$$
	E_p = \alpha\frac{\Delta}{\Delta+\gamma}\frac{K^\downarrow}{\lambda}+\beta,
$$

where $\lambda$ is the latent heat of vapouration for water (MJ/m³), $\Delta$ is the slope of the saturation vapour pressure vs. temperature curve (kPa/K), $\gamma$ is the psychrometric constant (kPa/K), and $\alpha$ and $\beta$ are empirical coefficients determined equal to 0.61 and $-1.2\times 10^{-4}$ m/d, respectively (Makkink, 1957). $\lambda$ and $\Delta$ are dependent on air temperature alone and $\gamma$ on air temperature and atmospheric pressure.

Realistically, southern Ontario models at this scale can only rely on daily min/max temperatures, rainfall and snowfall accumulations. Hourly data do exists that include wind speeds, humidity, pressure and visibility, but unfortunately, these data are not as readily available. For this reason, the current iteration of the model (version 1.0) will solely utilize daily data, and assuming that atmospheric pressure remains constant at 101.3 kPa, and the number of sunshine hours ($n=0$) on days with recorded precipitation, $n=1$ on days without. Future versions will compile more regional data that will relax these assumptions.

While it is understood great simplification is made by choosing such an empirical model (i.e., neglecting the impacts from wind speed, surface roughness, atmospheric moisture capacity, etc.) it is deemed appropriate for the purpose of this regional recharge model on the basis that in southern Ontario climates, $E_p$ is negligible during the late fall, winter, and early spring seasons, whereas the remaining season, actual evaporation is limited by moisture availability and not by evaporative potential, meaning that evaporative flux has a greater dependence on soil moisture accounting rather than the method chosen to estimate $E_p$.

Future models may also incorporate more sophisticated methods such as the combination approaches. These methods require knowledge of net radiation ($Q^\text{*}=K^\text{*}+L^\text{*}$), where $K^\text{*}=(1-\alpha)K_\text{in}$ is the net shortwave radiation, $\alpha$ is the surface albedo, and $L^\text{*}$ is the net long-wave radiation that can be approximated by (Budyko, 1974; Nov\'ak, 2012):

$$
\begin{align*}
	L^* &= f(T)f(e)f(n) \\
	&= \underbrace{\left(\varepsilon\sigma T^4\right)}_{\substack{\text{thermal} \\\text{radiation}}}
	\cdot \underbrace{\left(0.254-.005e\right)}_{\substack{\text{humidity} \\\text{effect}}} % Brunt-type equation pg.233
	\cdot \underbrace{\left[1-c'\left(1-\frac{n}{N}\right)\right]}_{\substack{\text{cloudiness} \\\text{effect}}},
\end{align*}
$$

where $\varepsilon$ is the emissivity of the evaporating surface ($0.96\leq\varepsilon\leq 0.98$), $\sigma=5.67\times 10^{-8}$ W m $^{-2}$ K $^{-4}$ is the Stefan-Boltzmann constant, $T$ is the temperature of the evaporating surface, $e$ is the water vapour pressure [hPa], and $c^\prime$ is an empircal constant dependent on latitude; $c'\approx0.7$ at $45^\circ$ latitude (Budyko, 1974). <!-- pg.59 -->

With net radiation, $E_p$ can be approximated first by using the Priestly-Taylor equation:

$$
	\lambda E_p = \alpha_{_{PT}}\frac{\Delta}{\Delta+\gamma}\left(Q^*-Q_G\right),
$$

where $\alpha_{_{PT}}$ is the Priestly-Taylor coefficient ($\approx 1.26$, but can vary depending on atmospheric state and surface saturation), and $Q_G$ is the gound heat flux, which is typically deemed negligeable, but can be approximated by:

$$
	Q_G = -k_s\frac{\partial T}{\partial z},
$$

where $k_s$ is the thermal conductivity of the soil, and $\partial T/\partial z$ is the vertical temperature gradient near the surface.

Lastly, the most sophisticated method not only incorperates atmospheric conditions, but surface conditions aswell that can be parameterized with the model. The Penman-Monteith equation... 

_**TODO**_


## Land surface corrections

From the high resolution of the model domain, each of the above $E_p$ functions are adjusted further by considering the slope and aspect of the model cell.

_**TODO**_