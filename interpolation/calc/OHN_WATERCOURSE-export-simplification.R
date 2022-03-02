
library(ggplot2)
library(dplyr)


# copy table in excel, then run below
dfseg <- read.table(file = "clipboard", sep = "\t", header=TRUE)
View(dfseg)

dfvert <- read.table(file = "clipboard", sep = "\t", header=TRUE)
View(dfvert)



dfseg %>% ggplot() + geom_histogram(aes(x=length)) + xlim(c(NA,5000))

dfvert %>% ggplot() + geom_histogram(aes(x=length)) + xlim(c(NA,500))
