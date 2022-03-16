---
title: "Chapman and Maxwell"
author: "M. Marchildon"
date: '2021-09-09'
output: html_document
---


## Original Form

Chapman, T.G. and A.I. Maxwell, 1996. Baseflow separation - comparison of numerical methods with tracer experiments. Institute Engineers Australia National Conference. Publ. 96/05, 539-545.

from: Chapman, T.G., 1999. A comparison of algorithms for stream flow recession and baseflow separation. Hydrological Processes 13: 701-714.

$$
  Q_b(i) = kQ_b(i-1) + (1-k)Q_d(i)
$$


$$
  Q(i) = Q_d(i) + Q_b(i)
$$
$$
  Q_b(i) + (1-k)Q_b(i) = kQ_b(i-1) +(1-k)Q(i)
$$

$$
  Q_b(i) = \frac{k}{2-k}Q_b(i-1) + \frac{1-k}{2-k}Q(i)
$$

## General form

$$
  b_t = \alpha b_{t-1} + \beta q_t + \gamma q_{t-1}
$$
$$
  \alpha = \frac{k}{2-k} \qquad \beta = \frac{1-k}{2-k} \qquad \gamma=0
$$