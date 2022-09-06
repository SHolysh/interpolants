---
title: "RDRR-MC"
author: "M. Marchildon"
date: '2021-03-25'
output: html_document
---


## Sampling

blah blah n realizations

returned (for each realization) n hydrographs corresponding to active stream gauges

average monthly distribution of recharge (ro, et, ...)  12 values per cell

## Selection

realization accepted if at any gauge, both NSE and KGE > 0




the posterior likelihood $L\left(\psi|\theta\right)$:

$$
  L\left(\psi|\theta\right)=\frac{L\left(\theta|\psi\right)}{\sum L\left(\theta|\psi\right)}
$$

where $\theta$ is the set of model parameters, $\psi$ is the set of obervations, and $L\left(\theta|\psi\right)$ is the so-called likelihood measure (Beven and Binley, 1992). Note that the prior distibution of parameter values $L\left(\theta\right)$ has been ommitted as uniform distributions between 2 predefined limits are defined for each parameter; therefore $L\left(\theta\right)=\text{constant}$. The denominator is a nomalization constant such that $\sum L\left(\psi|\theta\right)=1$, that is the sum over the set of bahavioural realizations.

The bounds introduced at the grid scale are made independently in that a raster of minimums does not represent a minimal realization. 