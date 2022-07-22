---
title: Water Budget Modelling
subtitle: Model structure
author: M.Marchildon
date: 2019
output: html_document
---


# Model structure

The model (i.e., structural data, parameters, etc.) is structured as follows:

### Model domain
This is the overarching placeholder for all necessary data to run a model. The idea here is to collect data from a greater region that can be parsed to smaller areas where the numerical model is to be applied, say to some given basin outlet cell ID. This allows for a single greater set of input data to be consolidated into fewer computer files

### Subdomain
When the user wishes to run the model, there may be interest to specific areas and not the entire domain. For instance, user has only to specify the cell ID from which a catchment area drains to; or a grid definition with pre-defined active cells can be inputted.

### Sample
While the Domain and Subdomain carry mostly structural (i.e., unchanging) data, the sample is where a Subdomain is parameterized. Samples are handy when one attempts to optimize or perform a Monte Carlo analysis of parameter space using the same Subdomain.

### Subsample
For larger catchments, it may be advantageous to sub-divide the Subdomain for a particular parameter sample, say by subwatersheds of a predefined catchment area. This can be leveraged when scaling the model across multi-processor computer architectures. Subsample also holds sample results once the model has been evaluated.