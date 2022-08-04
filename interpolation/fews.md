---
title: ORMGP-FEWS
author: M.Marchildon
output: html_document
---

# ORMGP-FEWS
[Delft-FEWS](https://www.deltares.nl/app/uploads/2015/01/Delft-FEWS_brochure-2017.pdf) (**F**lood **E**arly **W**arning **S**ystem) is an industry-standard "platform for real time forecasting and water resources management." Its strength is that it, as a stand-alone product, can web-scrape, interpolate, aggregate and re-generate data from any source; it is truly a generalized water data management tool.

FEWS can even integrate with hydrological models, by creating input data, running the model, and imports model output — all at the click of a button.

But most importantly, the FEWS platform offers a streamlined user interface that allows professionals the ability to inspect their data, prior to use.

At the ORMGP, the "ORMGP-FEWS" system is used to organize data collected from a variety of sources (international, federal, provincial, municipal, academia and private organizations) and consolidates them to a single spatially-distributed, temporally-aggregated data product we offer to our partners and utilized internally.

![ORMGP-FEWS sample](fig/ORMGP-FEWS-sample1.gif)

## Data types
The types of timeseries data hosted in ORMGP-FEWS mainly come in 2 flavours:

1. **Scalar data** — These are data most common to environmental databases. They are data collected at stations, and represent phenomena measured at a point.  These data require spatial interpolation to convert them to a distributed field.
1. **Vector data** — There data come in a distributed (i.e., raster) form, and thus may not require further interpolation.  This data format is relatively new given reduced technological constraints and have yet to make significant gains in practice. More and more open international sources offer such data, free of charge. Difficulties with the data are it's management, as they do not lend themselves well to standard "normalized" database schemas.  FEWS, on the other hand, is especially tailored to handle and manipulate vector data.