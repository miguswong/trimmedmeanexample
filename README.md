# Assignment 9: Create a Go Package
## Management Problem

A firm specializing in Go consulting wants to promote its services by developing open-source Go packages and offering these packages on GitHub. Its first offering will be a package for computing a trimmed mean.

### Assignment Requirements 

This assignment involves setting up two GitHub repositories.

Create a Go package for computing a trimmed mean and make it available as a GitHub repository: 

Follow these instructions for creating a Go package: https://www.tutorialspoint.com/how-to-create-your-own-package-in-golangLinks to an external site.
The trimmed mean package should provide a function to read a slice of integers and/or floats and compute a trimmed mean, where the degree-of-trimming arguments represent the proportions of slice elements to be taken from the lower and upper ends of the slice after ordering from the lowest value to the highest value. If only one degree-of-trimming argument is provided, assume that the trimming is symmetric, with equal proportions taken from both the lower and upper ends of the sorted observations. 
Set up a second GitHub repository for using the trimmed mean package:

Develop a main.go program that imports the trimmed mean package and uses the trimmed mean function to compute a trimmed mean. 
Using samples of at least 100 integers and 100 floats, test the Go trimmed mean function against symmetric trimming results from the base mean function in R with 0.05 observations taken from both ends of the distribution: mean(x, trim = 0.05) .
In the READMe.md file, provide complete instructions for getting and using the package within a Go program. Summarize results from your Go package, comparing those results to trimmed mean results from R.
(Optional) Run a bootstrap study showing the value of the trimmed mean as a robust estimator of central tendency. See the R programming example at the end of this assignment write-up.

## Using/Importing the Package
The [trimmedmean](https://github.com/miguswong/trimmedmean) package contains the function *TrimMean* to calculate the trimmed mean (float64) of a slice of floats or ints. 2 Variables are passed to the function:
* The slice of ints/floats you would like trimmed
* The amount (in terms of numbers) you would like trimmed from each side of the sorted slice.

To utilize the package in your go program, execute the follwing command in your program's directory.
```
go get -u github.com/miguswong/trimmedmean
```
This will pull the latest version of the package into your program. Ensure that you are using version 1.0.5 **at a minimimum** for the function to work properly.

This repository's [main.go]() file contains an example of utilizing the package to calculate the trimmed mean. There is only one function needed to handle int and float slices. This was accomplished via the use of generics. The slice is trimmed, summed, then averaged with the result being returned as a float. The program was written so that execution speed could be captured

## Comparison to R
The TrimMean function was built in reference to R's mean function which takes in a vector and a trim variable. *trim* represents the proportion to be trimmed from each side.

### Execution Speed
[trimmean.R]() contains the file with the replicated functionality as main.go. It was found that R took nearly twice as long. While Go took a time of 544.5 microseconds to execute. R, on average, took around 1037 microsecondss. These results show that R, while being simplistic and easy to learn in nature, also means that the program will significantly struggle with larger sets of data.