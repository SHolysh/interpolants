---
title: Water Budget Modelling
subtitle: Multiphase flow through porous media
author: M.Marchildon
date: 2019
output: html_document
---

* TOC
{:toc}

# Multiphase flow through porous media

## Liquid phase

Flux through a porous medium is defined as:

$$
	f_l=-K\frac{d\psi}{dz}
$$

where $f_l$ is liquid flux density [kg m $^{-2}$ s $^{-1}$], $\psi$ is the matric potential [J kg $^{-1}$] $\left(\text{note: }\psi=gh\right)$, and $K$ is the hydraulic conductivity [kg s m $^{-3}$].\footnote{note: $K_{kg\cdot s\cdot m^{-3}}=\frac{\rho_l}{g}K_{m\cdot s^{-1}}$}

When combined with the mass conservation equation, the result yields the transient vertical (1D) Richards (1931) equation: 

<!-- pg.166 Bittelli etal (2015); potential is being represented in the energy per unit mass form (J/kg) rather than the per unit volume form (J/m3=Pa) to avoid temperature dependence on liquid volume (see pg.98). -->

$$
	\rho_l\frac{\partial\theta}{\partial t}=\frac{\partial}{\partial z}\left[K(\psi)\left(\frac{\partial\psi}{\partial z}+g\right)\right],
$$

which is often expressed in its so-called $\psi$--based form:

$$
	\rho_lC(\psi)\frac{\partial\psi}{\partial t}=\frac{\partial}{\partial z}\left[K(\psi)\left(\frac{\partial\psi}{\partial z}+g\right)\right],
$$

where $\rho_l$ is the density of the liquid [kg m $^{-3}$], $g$ is the acceleration to to gravity ($\approx9.80665$ m s $^{-2}$ or J kg $^{-1}$ m $^{-1}$). The capacity [kg J] term is solved using the Campbell (1974) model: <!-- pg121 -->

$$
	C(\psi)=\frac{d\theta}{d\psi}=\frac{-\theta}{b\psi}
$$

and conductivity:

$$
	K(\psi)=
	\begin{cases}
		K_s\left(\frac{\psi_e}{\psi}\right)^{2+3/b}  \qquad &\text{if $\psi_m\leq\psi_e$} \\
		K_s  &\text{otherwise}
	\end{cases}
$$

where $K_s$ is the saturated conductivity [kg s m $^{-3}$], $b$ is a shape parameter and $\psi_e$ is the air-entry potential [J kg $^{-1}$]. Furthermore, water content ($\theta$) [m $^3$ m $^{-3}$] and $\psi$ are assumed related by the power-law relationship:

$$
	\theta(\psi)=
	\begin{cases}
		\theta_s\left(\frac{\psi}{\psi_e}\right)^{-1/b}  \qquad &\text{if $\psi_m\leq\psi_e$} \\
		\theta_s  &\text{otherwise}
	\end{cases}
$$

or conversely,

$$
	\psi(\theta)=
	\begin{cases}
		\psi_e\left(\frac{\theta}{\theta_s}\right)^{-b}  \qquad &\text{if $\theta\leq\theta_s$} \\
		\psi_e  &\text{otherwise}
	\end{cases}
$$

where $\theta_s$ is the water content at saturation. The water balance for cell $i$ can be solved using the its finite-difference form: <!-- linear method -->

$$
	\frac{\rho_l\left(\theta_i^{j+1}-\theta_i^j\right)\left(z_{i+1}-z_{i-1}\right)}{2\Delta t}=\frac{\overline{K}_i\left(\psi_{i+1}-\psi_i\right)}{z_{i+1}-z_i}-\frac{\overline{K}_{i-1}\left(\psi_i-\psi_{i-1}\right)}{z_i-z_{i-1}}+u_i
$$

or in its $\psi$-based form:
$$
	\frac{\rho_lC_i\left(\psi_i^{j+1}-\psi_i^j\right)\left(z_{i+1}-z_{i-1}\right)}{2\Delta t}=\frac{\overline{K}_i\left(\psi_{i+1}-\psi_i\right)}{z_{i+1}-z_i}-\frac{\overline{K}_{i-1}\left(\psi_i-\psi_{i-1}\right)}{z_i-z_{i-1}}+u_i
$$

The source term $u_i$ could include extraction by roots, evaporation, etc. In most cases, for vertical simulations, the source term represents gravitational flux, where:
$$
	u_i=g\left(\overline{K}_{i-1}-\overline{K}_i\right).
$$

<!-- The solution to the above equation is done implicitly, by solving a system equations. Transforming the above to:

$$
\begin{align*}
	\varrho(i) &=\rho_lC_i\frac{\left(z_{i+1}-z_{i-1}\right)}{2\Delta t} \\
	k(i) &=\frac{K_i}{z_{i+1}-z_i} \\
	c(i) &= a(i+1) = k(i) \\
	b(i) &= k(i-1)+k(i)+\varrho(i) \\
	d(i) &= \varrho(i)\psi_i^j+u(i)
\end{align*}
$$

the finite-difference scheme for $4$ layers is put in matrix form:
$$
	\begin{bmatrix}
		b(1) & c(1) & 0 & 0 \\
		a(2) & b(2) & c(2) & 0 \\
		0 & a(3) & b(3) & c(3) \\
		0 & 0 & a(4) & b(4)
	\end{bmatrix}
	\begin{bmatrix}
		\psi^{j+1}(1) \\
		\psi^{j+1}(2) \\
		\psi^{j+1}(3) \\
		\psi^{j+1}(4)
	\end{bmatrix}
	=
	\begin{bmatrix}
		d(1) \\
		d(2) \\
		d(3) \\
		d(4)
	\end{bmatrix}
$$ -->



# Vapour phase within the soil matrix

Vapour flux density [kg m $^{-2}$ s $^{-1}$] (i.e., molecular diffusion flux of water vapour in the air) is defined using Fick's Law: %pg.70 Novak
$$
	f_v=-\rho_a D(\theta) \frac{dq}{dz},
$$

where $\rho_a$ is the density of air, [kg m $^{-3}$], $D(\theta)$ is the effective coefficient of turbulent diffusion of water vapour in pore-space air [m $^2$ s $^{-1}$], often given the form attributed to Penman (1940):
$$
	D(\theta)=D_a\eta\left(\theta_s-\theta\right),
$$

where $D_a$ is coefficient of molecular diffusion of water vapour in air $\approx2.12\times10^{-5}$ m $^2$ s $^{-1}$, and $\eta$ is a dimensionless coefficient characterizing the turbulent diffusion in porous media, $\eta\approx0.66$ (Penman, 1940). Note that the term $\left(\theta_s-\theta\right)$ is included correct for the volume of gas-filled porosity, where vapour flux is occurring. $q$ is the specific humidity (moisture content) of air [kg kg $^{-1}$]:
$$
	q=\frac{\varepsilon e}{P}=wq^*(T,P),
$$

where $e$ and $P$ are the vapour and total pressures, respectively [kg m $^{-1}$ s $^{-2}$], $\varepsilon$ is the ratio of the molecular weight of water to the molecular weight of air ($\approx 0.622$), $w$ is the relative humidity [--], and $q^*$ is the saturated specific humidity determined using the August-Roche-Magnus-Tetens relationship:
$$
	q^*(T,P)=\frac{0.38}{P}\exp\left(\frac{17.625T}{T+243.04}\right),
$$

where $P$ is in [kPa] and $T$ is in [$^\circ$ C].

Within a porous medium, assuming isobaric (i.e., $P_\text{pores}\approx P_a$) and isothermal (constant temperature) conditions, the specific humidity gradient can be related to pore-space relative humidity ($w_p$) and soil temperature ($T_s$) by:
$$
	\frac{dq}{dz}=q^*(T_s,P)\frac{dw_p}{dz}.
$$

By further assuming that the liquid and vapour phases are in equilibrium (i.e., have equal potential and the partial pressure of water has reached its vapour pressure, $e$), according to the ideal gas law, pore-space relative humidity can be determined from liquid potential by (Philip, 1957):
$$ %eq:wp
	w_p=\frac{e}{e^*(T_s)}=\frac{q}{q^*(T_s,P)}=\exp{\frac{\omega_w\psi}{RT_s}},
$$

where the gas constant $R=8.3143$ J mol $^{-1}$ K $^{-1}$, the molecular mass of water $\omega_w=0.01802$ kg mol $^{-1}$, and $e^*$ is the saturated vapour pressure at the temperature of the soil surface $T_s$ [K]. Next, using the chain rule,
$$
	\frac{dw_p}{dz}=\frac{dw_p}{d\psi}\frac{d\psi}{dz}=w_p\frac{\omega_w}{RT_s}\frac{d\psi}{dz},
$$

Combining the above equations, results in:
$$
	f_v=-K_v\frac{d\psi}{dz},
$$

where
$$
	K_v=\frac{\omega_w\eta\rho_aD_aq}{RT_s}\left(\theta_s-\theta\right),
$$

By combining the flux of water vapour, the Richards equation can be re-stated as (Nov\'ak, 2012): % pg.71 
$$
	\rho_l\frac{\partial\theta}{\partial t}=\frac{\partial}{\partial z}\left[K(\psi)\left(\frac{\partial\psi}{\partial z}+g\right)+\rho_a D(\theta)\frac{\partial q}{\partial z}\right],
$$

and further:

$$
	\rho_l\frac{\partial\theta}{\partial t}=\frac{\partial}{\partial z}\left[\left[K(\psi)+K_v(\theta,q)\right]\frac{\partial\psi}{\partial z}+gK(\psi)\right].
$$




# Atmospheric exchange

At the soil surface--atmosphere interface, the flux density to the atmosphere remains (Nov\'ak, 2012):
$$
	f_v=\rho_a D \frac{dq}{dz}
$$

Assuming that vertical fluxes do not change in the vicinity of the evaporating soil surface, integrating the above equation from the soil surface ($z_s$) and some effective height above the surface ($z_a$) yields: %pg.45 in Novak

$$
	-\int\limits_{q_s}^{q_a}dq = \frac{f_v}{\rho_a}\int\limits_{z_s}^{z_a}\frac{dz}{D(z)},
$$

re-arranging:

$$
	f_v=\frac{1}{\int_{z_s}^{z_a}\frac{dz}{D(z)}}\rho_a(q_s-q_a).
$$

Next, by defining a water vapour turbulent transport coefficient ($k_v$) [m s $^{-1}$]:

$$
	k_v=\frac{1}{\int_{z_s}^{z_a}\frac{dz}{D(z)}},
$$

vapour flux density to the atmosphere from the gas-filled pores can be written as:
$$
\begin{align*}
	f_{v,g} &=\rho_ak_v(q_s-q_a) \\
	&=\rho_ak_v\left[q_s-w_aq^*(T_a,P_a)\right],
\end{align*}
$$

and
$$
	f_{v,l}=\rho_ak_v(q^*(T_s,P_a)-q_a),
$$

where $w_a$ is the atmospheric relative humidity, $T_a$ is air temperature [K], and $P_a$ is air pressure [Pa = kg m $^{-1}$ s $^{-2}$]. (Note that for water, the units for flux density [kg m $^{-2}$ s $^{-1}$] is equivalent to [mm s $^{-1}$].) The water vapour turbulent transport coefficient can be related to wind speed ($u(z)$) [m s $^{-1}$] using boundary-layer theory: % simplified from pg.49 Novak
$$
	k_v\approx\frac{\kappa^2u(z)}{\left[\ln\left(\frac{z}{z_0}\right)+\frac{\beta}{L_*}z\right]^2} \approx\frac{\kappa^2u(z)}{\ln\left(\frac{z-d_e}{z_0}\right)^2},
$$

where $\kappa$ is the von-K\'arm\'an constant ($\approx0.4$), $z_0$ is the roughness length [m], $\beta$ and $L_*$ are the Monin-Obukhov (1954) coefficient of atmospheric stability and characteristic length, respectively, and $d_e$ is the so-called zero-plane displacement height [m]. Alternatively, the Penman (1948)/Penman-Monteith (1965) relationship to aerodynamic resistance ($r_a$), soil surface resistance ($r_s$) and canopy resistance ($r_c$) [s m $^{-1}$] can be used, where:
$$
	k_v=\frac{1}{r_a+r_s+r_c}.
$$

Lastly, the effective mass flux evaporating from the soil surface must account for the portion of the soil surface area exposed to gas-filled pores and the area of direct liquid exposure:
$$
\begin{align*}
	f_e &=\left(\theta_s-\theta\right)f_{v,g}+\theta f_{v,l}. \\
		&=\rho_ak_v\left[\theta_s(q_s-q_a)+\theta(q^*_s-q_s)\right]
\end{align*}
$$





# Numerical solution to the Richards equation

One solution to the 1D Richards equation is the cell-centered finite volume solution scheme following Bittelli et.al. (2015). The cell-centered finite-volume form to the above equations gives the mass-balance at node $i$ as:
$$
	\rho_lV_i\frac{\partial\theta_i}{\partial t}=\sum_{j=1}^{n}F_{ij}+u_i \qquad\forall i\neq j,
$$

where $V_i$ is the finite volume of cell/node $i$, and the inter-nodal mass flux [kg s $^{-1}$]:
$$
	F_{ij}=-A_{ij}\overline{K}_i\frac{h_i-h_j}{L_{ij}},
$$

and $h_i=\psi_i+gz$ is the total hydrostatic potential [J kg $^{-1}$]. In the 1D vertical, The water balance in cell $i$ (i.e., implicit Euler method---evaluating at the end of the time step) is: %pg.63
$$
	\frac{\rho_lV_i\overline{C}_i}{\Delta t}\left(h_i^{k+1}-h_i^k\right)=F_i\left(h_i^{k+1}-h_{i+i}^{k+1}\right)-F_{i-1}\left(h_{i-1}^{k+1}-h_i^{k+1}\right)+u_i,
$$

where
$$
	F_i=-A_i\frac{\overline{K}_i}{z_{i+1}-z_i}.
$$
and
$$
	\overline{C}_i=\frac{d\theta}{dh}\approx\frac{\theta_i^{k+1}-\theta_i^k}{h_i^{k+1}-h_i^k}.
$$

The solution to the above equation is done implicitly, by solving a system equations, expanding the above and letting:
$$
\begin{align*}
	\varrho(i) &=\frac{\rho_l V_i\overline{C}_i}{\Delta t} \\
	f(i) &=\frac{A_i\overline{K}_i}{z_{i+1}-z_i} \\
	c(i) &= a(i+1) = -f(i) \\
	b(i) &= f(i-1)+f(i)+\varrho(i) \\
	d(i) &= \varrho(i)h_i^k+u_i
\end{align*}
$$

the finite-volume scheme for $4$ layers is put in matrix form:
$$
	\begin{bmatrix}
		b(1) & c(1) & 0 & 0 \\
		a(2) & b(2) & c(2) & 0 \\
		0 & a(3) & b(3) & c(3) \\
		0 & 0 & a(4) & b(4)
	\end{bmatrix}
	\begin{bmatrix}
		h^{k+1}(1) \\
		h^{k+1}(2) \\
		h^{k+1}(3) \\
		h^{k+1}(4)
	\end{bmatrix}
	=
	\begin{bmatrix}
		d(1) \\
		d(2) \\
		d(3) \\
		d(4)
	\end{bmatrix}
$$

which is a tri-diagonal matrix that is solved using the Thomas algorithm.


# Newton-Raphson transformation

A finite-difference solution to the 1D Richards equation is solved using the Newton-Raphson solution scheme following Bittelli et.al. (2015). The water balance in cell $i$ is:

$$
	F_i=\frac{\overline{K}_i\left(\psi_{i+1}-\psi_i\right)}{z_{i+1}-z_i}-\frac{\overline{K}_{i-1}\left(\psi_i-\psi_{i-1}\right)}{z_i-z_{i-1}}
		-g\left(\overline{K}_i-\overline{K}_{i-1}\right)-\frac{\rho_l}{2\Delta t}\left(\theta_i^{k+1}-\theta_i^k\right)\left(z_{i+1}-z_{i-1}\right),
$$

where $F$ is the mass balance residual and $\overline{K}$ is the mean elemental hydraulic conductivity. Applying integral transform methods (Bittelli et.al., 2015), $\overline{K}$ is defined by: <!-- eq.8.53 pg.178 -->
$$
	\overline{K}_i=\frac{K_{i+1}\psi_{i+1}-K_i\psi_i}{\left(1+3/b\right)\left(\psi_{i+1}-\psi_i\right)}.
$$

The solution to the above equation is then considered a minimization problem, whereby Newton-Raphson:

$$
	\frac{\partial F}{\partial\psi}\cdot \left(\psi^{k+1}-\psi^k\right)=-F\rightarrow 0 \\
$$

or in a matrix-form for a $3$—layer profile is:

$$
	\begin{bmatrix}
		\frac{\partial F_1}{\partial\psi_1} & \frac{\partial F_1}{\partial\psi_2} & \frac{\partial F_1}{\partial\psi_3} \\
		\frac{\partial F_2}{\partial\psi_1} & \frac{\partial F_2}{\partial\psi_2} & \frac{\partial F_2}{\partial\psi_3} \\
		\frac{\partial F_3}{\partial\psi_1} & \frac{\partial F_3}{\partial\psi_2} & \frac{\partial F_3}{\partial\psi_3}
	\end{bmatrix}
	\begin{bmatrix}
		\psi_1^{k+1}-\psi_1^k \\
		\psi_2^{k+1}-\psi_2^k \\
		\psi_3^{k+1}-\psi_3^k
	\end{bmatrix}
	=
	\begin{bmatrix}
		-F_1 \\
		-F_2 \\
		-F_3 \\
	\end{bmatrix}	
$$

where $k$ indicates the $k^{th}$ iteration of the minimization attempt. The mass-balance function $F$ is solved by:

$$
	F_i=f_i-f_{i-1}+u_i-u_{i-1}-\frac{\rho_l}{2\Delta t}\left(\theta_i^{k+1}-\theta_i^k\right)\left(z_{i+1}-z_{i-1}\right),
$$

where
$$
	f_i=-\frac{\overline{K}_i\left(\psi_{i+1}-\psi_i\right)}{z_{i+1}-z_i} = -\frac{K_{i+1}\psi_{i+1}-K_i\psi_i}{\left(1+3/b\right)\left(z_{i+1}-z_i\right)}	
$$

and letting
$$
\begin{align*}
	u_i &=-gK_i \\
	\varrho_i &=-\rho_l\theta_i\frac{z_{i+1}-z_{i-1}}{2\Delta t},
\end{align*}
$$

then substituting yields:
$$
	F_i=\frac{K_i\psi_i-K_{i-1}\psi_{i-1}}{\left(1+3/b\right)\left(z_i-z_{i-1}\right)}
	-\frac{K_{i+1}\psi_{i+1}-K_i\psi_i}{\left(1+3/b\right)\left(z_{i+1}-z_i\right)}
	-g\left(K_i-K_{i-1}\right)-
	\frac{\rho_l}{2\Delta t}\left(\theta_i^{k+1}-\theta_i^k\right)\left(z_{i+1}-z_{i-1}\right).
$$

Differentiating $F_i$ with respect to $\psi_i$, $\psi_{i-1}$ and $\psi_{i+1}$ yields:

$$
\begin{align*}
	f'_i &=\frac{\partial f_i}{\partial\psi_i}=\frac{K_i}{z_{i+1}-z_i} \\
	u'_i &=\frac{\partial u_i}{\partial\psi_i}=\frac{gK_i}{\psi_i}(2+3/b) \\
	\varrho'_i &=\frac{\partial \varrho_i}{\partial\psi_i}=\rho_l\frac{\theta_i}{b\psi_i}\frac{z_{i+1}-z_{i-1}}{2\Delta t}
\end{align*}
$$

and thus <!-- pg.180-181 -->
$$
\begin{align*}
	\frac{\partial F_i}{\partial\psi_i} &=\frac{K_i}{z_{i+1}-z_i}+\frac{K_{i-1}}{z_i-z_{i-1}}+u'_i+\varrho'_i \\
	\frac{\partial F_i}{\partial\psi_{i-1}} &=\frac{-K_{i-1}}{z_i-z_{i-1}}-u'_{i-1} \\
	\frac{\partial F_i}{\partial\psi_{i+1}} &=\frac{-K_i}{z_{i+1}-z_i}
\end{align*}
$$

Solving the Newton-Raphson system of equations is then accomplished like the linear methods described above.



# Boundary conditions

Boundary conditions for the top of the vertical profile can either be a constant flux (e.g., evaporation) or constant potential:

- For constant flux, the source is added to $F_1$; 
- For a constant potential, $\psi_1^{j+1}$ is known and is set at the start of the time step (since the Newton-Raphson method computes changes in $\psi$ to bring $F$ to zero, $F_1$ and $\frac{\partial F_1}{\partial\psi_i}$ are set to zero before solving; therefore the value of $\psi_1$ will remain constant throughout the time step. For infiltration, set $\psi_1=\psi_e$.

At the bottom of the profile, either a constant potential (i.e., water table) or a free drainage condition can be specified using a ghost below the profile that never becomes part of the solution. Free drainage is set by:
$$
\begin{align*}
	\psi_{n+1} &=\psi_n \\
	\theta_{n+1} &=\theta_n \\
	K_{n+1} &=K_n
\end{align*}
$$

after each outer iteration, where $n$ is the number of finite difference cells being solved.