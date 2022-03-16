---
title: "Tularam and Ilahee"
author: "M. Marchildon"
date: '2022-03-16'
output: html_document
---


## Original Form
Tularam, A.G., Ilahee, M., 2008. Exponential Smoothing Method of Base Flow Separation and its Impact on Continuous Loss Estimates. American Journal of Environmental Sciences 4(2):136-144.

$$b_t=ab_{t-1}+(1-a)q_t$$	

## General form

$$
  b_t = \alpha b_{t-1} + \beta q_t + \gamma q_{t-1}
$$
$$\alpha = a \qquad \beta = 1-a \qquad \gamma=0.0$$