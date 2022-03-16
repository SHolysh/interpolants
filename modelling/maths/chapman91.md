---
title: "Chapman 1991"
author: "M. Marchildon"
date: '2021-09-09'
output: html_document
---


## Original Form

Chapman, T.G., 1991. Comment on the evaluation of automated techniques for base flow and recession analyses, by R.J. Nathan and T.A. McMahon. Water Resource Research 27(7): 1783-1784

$$
  f_k = \frac{3\alpha-1}{3-\alpha}f_{k-1} + \frac{2}{3-\alpha}\left(y_k - \alpha y_{k-1}\right)
$$


$$
  y_k = f_k + b_k
$$
$$
  b_k = y_k - \frac{3\alpha-1}{3-\alpha}f_{k-1} - \frac{2}{3-\alpha}\left(y_k - \alpha y_{k-1}\right)
$$

$$
  b_k =  \frac{3\alpha-1}{3-\alpha} b_{k-1} + \left(1- \frac{2}{3-\alpha}\right) y_k + \left(\frac{2\alpha}{3-\alpha} - \frac{3\alpha-1}{3-\alpha}\right)y_{k-1} \qquad f_{k-1}=y_{k-1}-b_{k-1}
$$
$$
  b_k =  \frac{3\alpha-1}{3-\alpha} b_{k-1} + \left(\frac{1-\alpha}{3-\alpha}\right) y_k + \left(\frac{1-\alpha}{3-\alpha}\right)y_{k-1} \qquad f_{k-1}=y_{k-1}-b_{k-1}
$$
where $\alpha=k$, the recession coeficient.

## General form

$$
  b_t = \alpha b_{t-1} + \beta q_t + \gamma q_{t-1}
$$
$$
  \alpha = \frac{3k-1}{3-k} \qquad \beta = \frac{1-k}{3-k} \qquad \gamma=\frac{1-k}{3-k}
$$