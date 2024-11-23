vec <- 1:100

start <- Sys.time()
mean(vec, trim = 0.05)
mean(vec, trim = 0.05)
end <- Sys.time()
execution_time <- end-start
print(paste("Execution Time: ", execution_time, "seconds."))
