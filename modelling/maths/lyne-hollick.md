---
title: "Lyne-Hollick"
author: "M. Marchildon"
date: '2021-09-09'
output: html_document
---


## Original Form

Lyne, V. and M. Hollick, 1979. Stochastic time-variable rainfall-runoff modelling. Hydrology and Water Resources Symposium, Institution of Engineers Australia, Perth: 89-92.

$$
  f_k = af_{k-1} + \frac{(1+a)}{2}\left(y_k-y_{k-1}\right)
$$


$$
  y_k = f_k + b_k
$$
$$
  b_k = y_k-af_{k-1} - b\left(y_k-y_{k-1}\right) \qquad b=\frac{(1+a)}{2}
$$

$$
  b_k = ab_{k-1} + (1 - b)y_k + (b-a) y_{k-1} \qquad f_{k-1}=y_{k-1}-b_{k-1}
$$

$$
  1-b = b-a = \frac{(1-a)}{2}
$$

## General form

$$
  b_t = \alpha b_{t-1} + \beta q_t + \gamma q_{t-1}
$$
$$
  \alpha = a \qquad \beta = \frac{1-a}{2} \qquad \gamma=\frac{1-a}{2}
$$