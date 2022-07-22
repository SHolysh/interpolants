---
title: Water Budget Modelling
subtitle: workflow
author: M.Marchildon
date: 2020
output: html_document
---



# Workflow (deprecated)


## ORMGP Database query

There are 5 queries that are first applied in the workflow, all found in `rdrr_meteorology.sln`:

## climate station query

The simplest of queries, simply creates station-based `.met` files for every location in our database from 1915:

    YCDB_T3_DailyClimateToMET("M:\ORMGP\met\ORMGP_YCDB.met", New Date(1915, 10, 1), New Date(2019, 9, 30), True)

this creates a directory with 600-odd `.met` files and are used (primarily) for station-based snow ablation model calibration.

## Long-term coarse interpolation field:

Using a 5km grid covering the ORMGP jurisdiction, a `.met` file is created interpolating precipitation data (via Theissen polygon) and temperature data (via inverse-distance squared) from 1915:

    YCDB_T3_DailyClimateInterp(New Grid.Definition("M:\ORMGP\met\ORMGP_5000.gdef"), New Date(1915, 10, 1), New Date(2019, 9, 30))

## Short-term fine interpolation field:

Using a 500m active grid covering the ORMGP jurisdiction, a `.met` file is created similar to the coarse version above, but from 1999:

    YCDB_T3_DailyClimateInterp(New Grid.Definition("M:\ORMGP\met\ORMGP_500a.gdef"), New Date(1999, 10, 1), New Date(2019, 9, 30))

## Pan evaporation climate interpolation:

For the number of pan evaporation stations available, location from where these data are collected, climate is interpolated using the same procedure as above. These station-based `.met` files are used for potential evaporation $(E_p)$ model calibration.

    YCDB_T3_DailyClimateInterp("E:\Sync\@dat\PanET\calibration\", OnatarioPanET, New Date(1915, 10, 1), New Date(2019, 9, 30))

## Hydrometric gauge catchment interpolation:

Using hill-climbing technique from a 50m hydro-corrected DEM of the ORMGP jurisdiction, catchments to gauge locations are assigned climate dataset, using an area-weighted coverage of the Long-term coarse interpolation field described above.

    YCBD_T3_DailyClimateInterp()

These catchment-based `.met` files are used for lumped hydrologic model calibration.

## Subcatchment interpolation

An additional dataset has been created for the 10km² subcatchment delineation of the ORMGP jurisdiction. Here, climate was interpolated (as above) to subcatchment centroids. The intention for this dataset was for Raven modelling.

    YCDB_T3_DailyClimateInterp(New Grid.Indx("M:\RDRR\ORMGP_50_hydrocorrect_SWS10.indx", New Grid.Definition("M:\RDRR\ORMGP_50_hydrocorrect.uhdem.gdef")), New Date(1915, 10, 1), New Date(2019, 9, 30))

# Pan evaporation calibration

Reading pan evaporation station-based `.met` files created above, $E_p$ model parameters are calibrated using auto-calibration routines. $E_p$ is simulated using the Makkink (1957) formulation:

$$E_p = \alpha\frac{K_g}{\lambda}\frac{\Delta}{\Delta+\gamma}+\beta$$

Global radiation $(K_g)$ is approximated using extraterrestrial solar radiation $K_e$ approximated using solar irradiation theory, adjusted for transmittance using the Bristow and Campbell (1984) relationship:

$$K_g = aK_e\left[(1 - e^{-b\Delta T^c}\right]$$

Since the Makkink multiplicative factor $\alpha$ is applied to $K_g$, the Bristow and Campbell coefficient is effectively absorbed and thus set to $a=1$ and thus not calibrated for. The daily range of air temperature $\Delta T = T_\text{max}-T_\text{min}$. This leaves 4 free parameters: $b$, $c$, $\alpha$ and $\beta$. (For more details and auto-calibration results, see `E:\Sync\@dat\PanET\calibration\epCalibration.Rmd`.)

Auto-calibration is performed using `github.com\maseology\hydrologicModelling\epCalibration\main.go`.

# Lumped (catchment) calibration

Next, using the optimized $E_p$ parameters, catchment models are built to i) calibrate snow-pack model, and ii) determine any signal for lateral groundwater exchange (i.e., `gwsink`). 

Lumped modelling is performed using the GR4J model (Perrin etal., 2003) with a cold-content factor snowmelt model.

