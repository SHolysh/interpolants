---
title: Clip n' Ship
author: Oak Ridges Moraine Groundwater Program
output: html_document
---

The ORMGP is previewing a service from which partners can easily gain access to the spatial layers we produce.  Imaging a *cookie-cutter* taking a vertical section within which the data are provided.


# Data formats
The data format chosen for the exports are provided in a format that is most compressible, maintaining efficient file transfer. The files formats are thus less less common yet are general enough that they'll open on many platforms

## File Formats

### Rasters
Rasters are provided in a [**band interleaved by line** (\*.bil)](https://desktop.arcgis.com/en/arcmap/10.5/manage-data/raster-and-images/bil-bip-and-bsq-raster-files.htm) format.  This file is in a raw binary format and is accompanied by header file (\*.hdr) and a projection file (\*.prj).  Files will open in [ArcGIS](https://www.arcgis.com/index.html), [QGIS](https://www.qgis.org/en/site/), and load as a [NumPy](https://numpy.org/) array.

### Vectors
Vector files (in particular polylines and polygons) are provided in [**GeoJson**](https://geojson.org/) (\*.geojson) format.  This is an all-encompassing ascii-format file which is flexible, but slow. It is suggested that files be converted to shapefiles (\*.shp) when performance is desired.

### Points
Point data, database queries, are provided in [**comma-separated value**]() (\*.csv) format.