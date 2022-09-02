

library(ggplot2)
library(gridExtra)
library(scales)
library(dplyr)


df.rdpa <- read.csv("interpolation/calc/rdpa-verification/precip-202010010100-RDPAdaily-msc-snodas-tcrit.csv")
df.hrdpa <- read.csv("interpolation/calc/rdpa-verification/precip-202208221100-HRDPAdaily-msc-snodas-tcrit.csv")



df.hrdpa %>% ggplot(aes(pMSC,pRDPA)) + 
  geom_point(aes(color=as.factor(wy)), alpha=.1) + 
  geom_abline(slope = 1,intercept = 0, color='red', size=3, alpha=.5)






p1 <- df.rdpa %>% ggplot(aes(tc)) + geom_density(aes(fill=factor(wy)), position = "stack")
p2 <- df.hrdpa %>% ggplot(aes(tc)) + geom_density(aes(fill=factor(wy)), position = "stack")

grid.arrange(p1, p2, ncol=2)




mn.rdpa <- mean(df.rdpa$resid)
mn.hrdpa <- mean(df.hrdpa$resid)

# residuals histogram
ggplot() + 
  geom_density(data = df.rdpa, aes(resid), size=1) +
  geom_density(data = df.hrdpa, aes(resid), size=1,color="blue") +
  # geom_vline(xintercept = mn.rdpa) +
  # geom_vline(xintercept = mn.hrdpa,color="blue") +
  # geom_text(aes(x=mn-1, label=paste0("avg. = ",round(mn,1)," mm/yr"), y=.4), angle=90, text=element_text(size=11)) +
  scale_x_log10(labels = label_number(), limits=c(0.001,NA)) +
  labs(x="Residual (mm/yr)", y=NULL)







df.rdpa %>% ggplot(aes(smSNODAS)) + geom_histogram(aes(fill=factor(wy)))
df.hrdpa %>% ggplot(aes(smSNODAS)) + geom_histogram(aes(fill=factor(wy)))


