package main

import (
	"fmt"
	"time"

	trimmedMean "github.com/miguswong/trimmedmean"
)

func main() {
	//Create slice of ints 1:100
	sliceOInt := make([]int, 100)
	for i := 0; i < 100; i++ {
		sliceOInt[i] = i + 1
	}

	//Same for float slice
	sliceOFloat := make([]float64, 100)
	for i := 0; i < 100; i++ {
		sliceOFloat[i] = float64(i + 1)
	}

	//Start timer
	start := time.Now()

	//Print out int slice results
	fmt.Println("Int Trim Mean: ", trimmedMean.TrimMean(sliceOInt, 5))

	//Print out float slice results
	fmt.Println("Float Trim Mean: ", trimmedMean.TrimMean(sliceOFloat, 5))

	//End timer and print out how long execution took
	elapsed := time.Since(start)
	fmt.Printf("Process took %s", elapsed)
}
