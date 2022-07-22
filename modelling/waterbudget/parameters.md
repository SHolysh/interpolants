---
title: Water Budget Modelling
subtitle: Model parameters
author: M.Marchildon
date: 2019
output: html_document
---

* TOC
{:toc}

# Model parameters

In total, the model consists of six parameters per computational element (CE), and one basin-wide paremeter. Five of the six parameters are distributed based on land use characteristics, while the reminder (`ksat`) is determine from surficial geology maping.

* Percolation storage (`DrnSto`): Reservoirs that require filling prior to adding to runoff. Sinks of this reservoir include drainage from percolation and evaporation [m].
* Surface storage (`SrfSto`): Reservoirs that require filling prior to adding to runoff. Sinks include only evaporation [m].
* Fraction of impervious area (`fimp`): A fractional area of the CE that prevents infiltration to and evaporation from the percolation storage reservoir [--].
* Saturated hydraulic conductivity (`ksat`): The potential rate of percolation from the percolation storage reservoir [m/s].
* Variable source area connectivity threshold (`rill`): Watertable height (relative to surface) above which groundwater discharge contributes to surface runoff [m].
* The Manning-Strickler roughness coefficient (`n`}) [s/m $^{1/3}$].
* TOPMODEL recession parameter (`m`) [m].




# Monte Carlo sampling scheme

An unbiased Monte Carlo (MC) sampling strategy has been employed to test the model's sampling space against observation. The strategy is termed "unbiased" in the sense that each sample is made independent of the results from previous samples (i.e., no parameter optimization/auto-calibration is being attempted). 

Sample space is derived using a set of predefined sampling distributions following either a uniform ($\text{U}[a,b]$), triangular ($\text{T}[a,m,b]$), or trapezoidal ($\text{Tr}[a,m,n,b]$) form, which can optionally be sampled in $\log_{10}$ space; in such cases the following notation is employed: $\text{U}_{10}[a,b]=\text{U}[\log_{10}(a),\log_{10}(b)]$.



## Sampling distributions

### Uniform

This distribution uniformly picks random samples between the upper ($b$) and lower ($a$) limits, such that $a<b$:

\todo{ SHOW FIGURE}

### Triangular

This distribution picks random samples between the upper ($b$) and lower ($a$) limits, but with a higher probability around the mode ($m$), where $a<m<b$:

\todo{ SHOW FIGURE}

### Trapezoidal

This distribution is a hybrid of the uniform and triangular where random picks between the upper ($b$) and lower ($a$) limits that are linearly tailed to the most-probable range of $m$ to $n$; here $a<m<n<b$:

\todo{ SHOW FIGURE}


## Applied sample space

The following is a description of the MC sampling applied to the DAS. There's no set number of samples to be applied, rather the DAS will perpetually re-sample this space thus ever-refining the model's certainty bounds. In sections to follow, a discussion on how the fitness of the DAS will be determined by observation constraints. Model parameters outlined above may be sampled directly, or derived from a set of samples. Also, samples may be taken as a basin-wide parameter, or will be distributed according to land use and surficial geology mapping.

The sample name is followed by the distribution used. A Latin Hypercube (LHC—Lemieux, 2009) sampling scheme is used to generate *i.i.d* samples on $\text{U}[0,1]$ which is then transformed according to the pre-define parameter distibution listed below:

<!-- \bigskip
\begin{minipage}{\textwidth}
\begin{tabular}{p{3.5cm}p{3.5cm}p{4cm}c}
	\multicolumn{4}{l}{\textbf{Basin-wide parameters}} \\ [0.5ex]
	\texttt{rill} & & $\text{U}_{10}[0.01, 1.0]$ & [m] \\ 
	\texttt{TOPMODEL $m$} & & $\text{U}_{10}[0.001,10.0]$ & [m] \\
	\texttt{soil depth\footnote{not to be taken literally, mostly representative of the evaporative depth} ($d_\text{soil}$)} & & $\text{U}[0.1,1.0]$ & [m] \\ 
	\multicolumn{2}{l}{\texttt{impervious depression storage ($I_i$)}} & $\text{U}_{10}[0.0001,0.001]$ & [m] \\
	\multicolumn{2}{l}{\texttt{vegetation interception capacity\footnote{weighting ($f_I$) is applied according to land use type, \% bare cover, presence of short/tall vegetaion, etc.} ($I_v$)}} & $\text{U}[0.001,0.004]$ & [m] \\ [1ex]
	\multicolumn{4}{l}{\textbf{Cell-based parameters}\footnote{according to land use and surficial geology mapping}} \\ [0.5ex] 
	\texttt{Manning's $n$} & & $\text{U}_{10}[0.0001,100.0]$ & [s/m\textsuperscript{1/3}] \\	
	\texttt{soil porosity ($\phi$)} & 
			Clay: & $\text{U}[0.4,0.7]$ & [--] \\
			& Loam/Silt: & $\text{U}[0.35,0.5]$ \\
			& Sand: & $\text{U}[0.25,0.5]$  \\
	\multicolumn{2}{l}{\texttt{soil field capacity\footnote{as a fraction of total porosity} ($f_c$)}} & $\text{U}[0.05,0.4]$ & [--] \\	
	\texttt{ksat} & 
			Low: & $\text{Tr}_{10}[10^{-11},10^{-9},10^{-7},10^{-6}]$ & [m/s] \\
			& Low-Med: & $\text{Tr}_{10}[10^{-9},10^{-7},10^{-6},10^{-5}]$ \\
			& Medium: & $\text{Tr}_{10}[10^{-8},10^{-6},10^{-5},10^{-4}]$ \\
			& Med-High: & $\text{Tr}_{10}[10^{-6},10^{-5},10^{-4},10^{-3}]$ \\
			& High: & $\text{Tr}_{10}[10^{-5},10^{-4},10^{-3},10^{-2}]$ \\
			& Streambed:\footnote{e.g., alluvium/unconsolidated/fluvial/floodplain material} & $\text{Tr}_{10}[10^{-8},10^{-7},10^{-5},10^{-4}]$ \\
			& Organics:\footnote{e.g., wetland sediments} & $\text{Tr}_{10}[10^{-8},10^{-7},10^{-5},10^{-4}]$ \\
			& Unknown/variable: & $\text{Tr}_{10}[10^{-9},10^{-7},10^{-5},10^{-3}]$ \\	
\end{tabular}
\end{minipage} -->


In relation to the model parameters the above samples are used to derive `DrnSto` and `SrfSto` by:

$$
\begin{align*}
	\texttt{DrnSto} &=\phi (1-f_c) d_\text{soil} \\
	\texttt{SrfSto} &=\phi f_c d_\text{soil} + I_v f_I (1-\texttt{fimp}) + I_i\cdot\texttt{fimp}
\end{align*}
$$




## Sample fitness

TODO

GLUE
constrained against daily streamflow measurements
top $\approx$500: 250: overall best perfoming (weighted by contributing area); 250/n catchments locally best performing