

# Automatic hydrograph separation

> "Division of a hydrograph into direct and groundwater runoff as a basis for subsequent analysis is known as *hydrograph separation* or *hydrograph analysis*. Since there is no real basis for distinguishing between direct and groundwater flow in a stream at any instant, and since definitions of these two components are relatively arbitrary, the method of separation is usually equally arbitrary." - Linsley et.al., 1975

For hydrograph separation, it is generally assumed that total flow ($q$) at any particular time ($t$) of a streamflow hydrograph can be partitioned into two primary components:

1. The *__slow flow__* component $(b)$, which is itself composed of the gradual release of water from watershed stores in addition to groundwater discharging into streams (i.e., ``groundwater runoff'' in Linsley \etal, 1975). The slow flow component is commonly referred to as ``baseflow.'' and,
	\item The \emph{quick flow} component ($f$), which originates from rainfall and/or snow melt events (i.e., ``direct runoff'' in Linsley \etal, 1975).
\end{enumerate}

Together, the slow and quick flow components sum to total flow: $q=b+f$. Conceptually, after a period of time following a precipitation event, streamflow tends to decrease steadily as it is composed entirely of baseflow ($f=0$). Upon the onset of a heavy rain event, the hydrograph quickly rises, as quick flow is added to the baseflow signature. One could imagine that should this rain event never occur, the underlying baseflow would have continued uninterrupted (such as in Reed \etal, 1975). The difference between this ``underlying'' baseflow and actual flow is perceived as total runoff in the form of quick flow.

Hydrologists found the need to separate the hydrograph into its constitutive components as it was found that runoff created from a precipitation event (i.e., rainfall and/or snow melt) tended to correlate best with the quick flow component only, as opposed to the total flow hydrograph (Beven, 2012). Consequently, a number of automatic hydrograph separation routines were proposed, all (as in Linsley \etal, 1975) being ``equally arbitrary.''

In general, two metrics can be determined from hydrograph separation that need definition. The first is the baseflow index ($BFI$), which is the ratio of long term baseflow discharge to total discharge:

\begin{equation}
	\text{BFI}=\frac{\sum b}{\sum q}
\end{equation}

The second is the baseflow recession coefficient ($k$), which describes the withdrawal of water from storage within the watershed (Linsley \etal, 1975). The recession coefficient is a means of determining the amount baseflow recedes after a given period of time:

\begin{equation}
	\label{eq:recessioncoefficient}
	b_t=kb_{t-1},
\end{equation}

\noindent where $b_{t-1}$ represents the slow flow calculated at one timestep prior to $b_t$. (Note, this assumes that total flow measurements are reported at equal time intervals, when unequal intervals are used, $k^{\Delta t}$ must be used, where $\Delta t$ is the time interval between successive $b$ calculations---for now, constant intervals are assumed.)

Linsley \etal (1975) also offered an approximate means of determining the time (in days) after peak flow discharge to when quick flow ceases ($f=0$) and thus total flow is entirely composed of the slow flow component and thus can be predicted using Equation \ref{eq:recessioncoefficient}. As a ``rule of thumb'' (Linsley \etal, 1975) the number of days ($N$) when quick flow terminates is approximated by:

\begin{equation}
	\label{eq:quickflowterminate}
	N=0.827A^{0.2},
\end{equation}   

\noindent where $A$ is the watershed area [km\textsuperscript{2}]. The above empirical relation is included here as a number of automatic hydrograph separation algorithms discussed below utilizes this approximation. Hydrograph components and ``quick flow cessation'' ($N$) is implicitly conceptualized when performing automatic hydrograph separation routines (Figure \ref{fig:conceptual_hydrograph}).

\begin{figure}[h]
	\centering
	\includegraphics[width=0.85\textwidth]{images/pinder1.png}
	\caption{The conceptual hydrograph separated into its quick (direct) and slow (groundwater) components (after Linsley and Franzini, 1964).}
	\label{fig:conceptual_hydrograph}
\end{figure}


\subsection{Digital filters}

Digital filters represent a set of automatic hydrograph separation algorithms that require no input other than the measured stream flow signal ($q$). Considering the streamflow hydrograph as a signal is quite apt when dealing with digital filters, as they themselves were inspired from signal processing of Lyne and Hollick, 1979 (Nathan and McMahon, 1990). With respect to the quick and slow hydrograph components (Figure \ref{fig:conceptual_hydrograph}), hydrograph separation is nothing more than the application of a low-pass filter to the total streamflow signal.

Another point to note is that many authors have applied these digital filters in multiple passes, either in two-passes (forward--backward) or three-passes (forward--backward--forward) to increase the smoothing of the resulting slow flow signal (Chapman, 1991).

\bigskip

With digital filters, there is no physical interpretation to the algorithm, it only produces a baseflow signal that resembles what one would expect. The general form of all digital filters used for hydrograph separation follows:

\begin{equation} \label{eq:gdf}
	b_t = \alpha b_{t-1} + \beta\left(q_t + \gamma q_{t-1}\right),
\end{equation}

\noindent where $q_{t-1}$ represents the total flow measured at one timestep prior to $q_t$, and $\alpha$, $\beta$ and $\gamma$ are parameters. The above equation is a three-parameter equation, however most implementations do not require every parameter be specified or, in other cases, two or more parameters can be specified as a function of another.

For example, the Lyne and Hollick (1979) equation (the earliest of digital filters used for hydrograph separation), is a one-parameter equation found by a single smoothing parameter $a$ suggested to be set between the values of 0.9--0.95 (Nathan and McMahon, 1990), where:

\begin{equation}
	\alpha = a \qquad \beta = \frac{1-a}{2} \qquad \gamma=1.0
\end{equation}

Chapman (1991), after noting some conceptual discrepancies with the Lyne and Hollick (1979) equation, modified the equation into a parameter-less form as a function of the recession coefficient $k$, which can be determined from the $q$ time series (equation \ref{eq:recessioncoefficient}). The Chapman (1991) algorithm takes the form:

\begin{equation}
	\alpha = \frac{3k-1}{3-k} \qquad \beta = \frac{1-k}{3-k} \qquad \gamma=1.0
\end{equation}

Chapman and Maxwell (1996) later simplified the above equation by assuming that slow flow is the weighted average of quick flow and the slow flow from the previous timestep (Chapman, 1999), that is $b_t=kb_{t-1}+(1-k)f_t$, leading to:

\begin{equation}
	\alpha = \frac{k}{2-k} \qquad \beta = \frac{1-k}{2-k} \qquad \gamma=0.0
\end{equation}

Boughton (1993) used a similar approach to Chapman and Maxwell (1996), except added an adjustment parameter $C$, such that $b_t=kb_{t-1}+Cf_t$. The Boughton (1993) form of the digital filter thus requires:

\begin{equation}
	\alpha = \frac{k}{1+C} \qquad \beta = \frac{C}{1+C} \qquad \gamma=0.0
\end{equation}

While also investigating the generalized digital filter, Eckhardt (2005) discovered an interpretation of the Boughton (1993) algorithm that eliminated the $C$ parameter and introduced the concept of $\text{BFI}_\text{max}$: the maximum value of the baseflow index that can be achieved using the digital filter. The Eckhardt (2005) digital filter is found by:

\begin{equation}
	\alpha = \frac{(1-\text{BFI}_\text{max})k}{1-k\text{BFI}_\text{max}} \qquad \beta = \frac{(1-k)\text{BFI}_\text{max}}{1-k\text{BFI}_\text{max}} \qquad \gamma=0.0
\end{equation}	

\noindent or made equivalent to Boughton (1993) by setting:

\begin{equation}
	C = \frac{(1-k)\text{BFI}_\text{max}}{1-\text{BFI}_\text{max}}
\end{equation}	

The commonly cited Jakeman and Hornberger (1993) algorithm closely follows that of Boughton (1993) and Chapman and Maxwell (1996), except it was formulated from a component of the IHACRES data-based model rather than being intended strictly for hydrograph separation (Chapman, 1999). Nonetheless, the IHACRES model can be shown to fit the general digital filter of equation \ref{eq:gdf} above, using 3 parameters, where:

\begin{equation}
	\alpha = \frac{a}{1+C} \qquad \beta = \frac{C}{1+C} \qquad \gamma=\alpha_s
\end{equation}	

\noindent Note that setting $\alpha_s<0$ is conceptually correct, as it implies that the rate of change of slow flow is positively correlated the rate of change of total flow (Chapman, 1999).

Lastly, Tularam and Ilahee (2008) most recently presented an digital filter that also resembled that of Chapman and Maxwell (1996), with the slight difference of assuming that slow flow is the weighted average of the slow flow of the previous timestep and total flow, not quick flow (i.e., $b_t=ab_{t-1}+(1-a)q_t$). This formulation is essentially the same as Lyne and Hollick (1979) with the exception that Tularam and Ilahee (2008) does not average total flow of the current and previous timestep. The one-parameter Tularam and Ilahee (2008) form yields:

\begin{equation}
	\alpha = a \qquad \beta = 1-a \qquad \gamma=0.0
\end{equation}

\bigskip
\subsubsection{Digital filter equations in their published form:}

\noindent Lyne and Hollick (1979):

	\begin{equation}
		b_t = ab_{t-1} + \frac{1-a}{2}\left(q_t + q_{t-1}\right)
	\end{equation}

\noindent Chapman (1991):

	\begin{equation}
		b_t = \frac{3k-1}{3-k}b_{t-1} + \frac{1-k}{3-k}\left(q_t + q_{t-1}\right)
	\end{equation}

\noindent Chapman and Maxwell (1996):

	\begin{equation}
		b_t = \frac{k}{2-k}b_{t-1} + \frac{1-k}{2-k}q_t
	\end{equation}

\noindent Boughton (1993):

	\begin{equation}
		b_t = \frac{k}{1+C}b_{t-1} + \frac{C}{1+C}q_t
	\end{equation}
	
\noindent Eckhardt (2005):

	\begin{equation}
		b_t = \frac{(1-\text{BFI}_\text{max})kb_{t-1} + (1-k)\text{BFI}_\text{max}q_t}{1-k\text{BFI}_\text{max}}
	\end{equation}
	
\noindent Jakeman and Hornberger (1993):

\begin{equation}
	b_t = \frac{a}{1+C}b_{t-1} + \frac{C}{1+C}\left(q_t + \alpha_s q_{t-1}\right)
\end{equation}	

\noindent Tularam and Ilahee (2008):

\begin{equation}
	b_t=ab_{t-1}+(1-a)q_t
\end{equation}	


\subsection{Moving-window methods}

A second class of hydrograph separation schemes are here considered ``moving window methods'' also known as ``manual separation techniques'' in Arnold and Allen (1999). These methods do not follow an equation, per se, rather a methodology based on the explicit/manual selection of discharge values assumed representative of baseflow discharge within a window of a set number of days.

In total, 10 estimates of baseflow discharge are computed using variants of 4 methods. Many of these methods are included in stand-alone software packages and have been re-coded here. The methods include:

\begin{enumerate}
	\item The UKIH/Wallingford technique (Institute of Hydrology, 1980). This method operates by locating minimum discharges in a (user specified) $N$-day window. This set of minimum discharge is then further screened, automatically, for discharges that are considered representative of ``baseflow,'' which are deemed ``turning points.'' Linear interpolation is then conducted between subsequent turning points yielding the final (baseflow) discharge. In a similar fashion to the digital filters, this method extracts a filtered/smoothed hydrograph of total flow minima, and is therefore often also referred to as the ``smoothed minima technique.'' 
	
	Piggott \etal (2005) discussed how the UKIH technique can yield alternate baseflow estimates depending on the origin of the $N$-day window. They proposed staggering $N$-sets of UHIK baseflow estimates to create an overall aggregate baseflow hydrograph. Three versions of this modification are included here:
	
	\begin{enumerate}
		\item Sweeping minimum: returns the daily minimum of the staggered hydrographs;
		\item Sweeping maximum: returns the daily maximum of the staggered hydrographs; and,
		\item Sweeping median:  returns the median of the $N$-staggered hydrographs.
	\end{enumerate}
	
	
	\item The HYSEP technique (Sloto and Crouse, 1996). This method depends on the computed days of quick flow termination $N$ (equation \ref{eq:quickflowterminate}). Like the UHIK method, the HYSEP techniques then proceed to determine minimum discharges within the $N$-day window. Three methods of producing baseflow estimates are computed in HYSEP and are reproduced here, they include:
	
	\begin{enumerate}
		\item Fixed interval: where baseflow is assumed to be the minimum discharge reported within sequential, non-overlapping $N$-day windows. Like the UKIH method, results from the fixed interval method is dependent on the (``fixed'') window origin;
		\item Sliding interval: where baseflow is assumed to be the minimum discharge found within a moving $N$-day window. In contrast, this method tends to yield a higher BFI; and,
		\item Local minimum: linearly-interpolates total flow minima within a moving $N$-day window.
	\end{enumerate}
	
	
	\item The PART technique (Rutledge, 1998). This method aims to reproduce the conceptual hydrograph represented in Figure \ref{fig:conceptual_hydrograph}. Using a combination of quick flow termination estimates (equation \ref{eq:quickflowterminate}), recession coefficients (equation \ref{eq:recessioncoefficient}), and a parameter termed the ``antecedent requirement,'' a combination of forward and backward filtering techniques are used in producing the final hydrograph separation estimates. Three estimates using the PART method are produced here, based on the suggested antecedent requirement choices offered by Rutledge (1998):
	
	\begin{enumerate}
		\item $\text{antecedent requirement} = 1$;
		\item $\text{antecedent requirement} = 2$; and, 
		\item $\text{antecedent requirement} = 3$;
	\end{enumerate}
	
	
	\item The Clarifica Inc., (2002) technique. This method separates the total flow hydrograph by performing two sweeps on the hydrograph. The first is a 6-day moving minimum, followed by a 5-day moving average (3-days previous, 1-day ahead). This method was designed for use within southern Ontario watersheds and tends to produce higher estimates of baseflow during peak flow events.
	
\end{enumerate}



%\subsubsection{Physically-based digital filter}

%Furey and Gupta (2001) presented a digital filter that was formulated TO FINISH THIS PART





