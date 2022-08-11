

library(ggplot2)
library(dplyr)


df <- read.csv('O:/baro.csv')

ggplot(df, aes(Pa,RD_VALUE)) +
  geom_point(aes(color=as.factor(INT_ID))) +
  geom_smooth(aes(color=factor(INT_ID)))


df[df$RD_VALUE<500000 & df$RD_VALUE>50000,] %>%
  ggplot(aes(Pa,RD_VALUE)) + 
  geom_point(aes(color=as.factor(INT_ID)), alpha=0.05) +
  geom_smooth(aes(color=factor(INT_ID)), method='lm') +
  geom_abline(linetype="dashed",size=2) +
  coord_fixed(ylim=c(NA,105000))
  
