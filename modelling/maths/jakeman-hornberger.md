---
title: "Jakeman and Hornberger"
author: "M. Marchildon"
date: '2022-03-16'
output: html_document
---


## Original Form
Jakeman, A.J. and Hornberger G.M., 1993. How much complexity is warranted in a rainfall-runoff model? Water Resources Research 29: 2637-2649.

$$b_t = \frac{a}{1+C}b_{t-1} + \frac{C}{1+C}\left(q_t + \alpha_s q_{t-1}\right)$$	

## General form

$$
  b_t = \alpha b_{t-1} + \beta q_t + \gamma q_{t-1}
$$
$$
\alpha = \frac{a}{1+C} \qquad \beta = \frac{C}{1+C} \qquad \gamma=\beta\alpha_s
$$