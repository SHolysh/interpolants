---
title: "Boughton and Eckhardt"
author: "M. Marchildon"
date: '2022-03-16'
output: html_document
---


## Original Form

Boughton, W.C., 1993. A hydrograph-based model for estimating the water yield of ungauged catchments. Hydrology and Water Resources Symposium, Institution of Engineers Australia, Newcastle: 317-324.

$$b_t = \frac{k}{1+C}b_{t-1} + \frac{C}{1+C}q_t$$

Eckhardt, K., 2005. How to construct recursive digital filters for baseflow separation. Hydrological Processes 19, 507-515.

$$
  b_t = \frac{(1-\text{BFI}_\text{max})kb_{t-1} + (1-k)\text{BFI}_\text{max}q_t}{1-k\text{BFI}_\text{max}}
$$

where

$$
  C = \frac{(1-k)\text{BFI}_\text{max}}{1-\text{BFI}_\text{max}}
$$	



## General form

$$
  b_t = \alpha b_{t-1} + \beta q_t + \gamma q_{t-1}
$$
$$
  \alpha = \frac{k}{1+C} \qquad \beta = \frac{C}{1+C} \qquad \gamma=0
$$