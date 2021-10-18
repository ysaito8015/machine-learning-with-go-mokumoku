package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	//"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	irisFile, err := os.Open("./iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Get the float values from the "sepal_length" column.
	sepalLength := irisDF.Col("petal_length").Float()

	// Calculate the Min of the variable.
	minVal := floats.Min(sepalLength)
	fmt.Printf("Minimum: %0.2f\n", minVal)
	// Output:  Minimam: 1.00

	// Calculate the Max of the variable.
	maxVal := floats.Max(sepalLength)
	fmt.Printf("Maximum: %0.2fv\n", maxVal)
	// Output: Maximum: 6.90

	// Calculate the range of the variable.
	rangeVal := maxVal - minVal
	fmt.Printf("Range: %0.2fv\n", rangeVal)
	// Range: 5.90

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)
	fmt.Printf("Variance: %0.2f\n", varianceVal)
	// Output: Variance: 3.11

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(sepalLength, nil)
	fmt.Printf("Standard Deviation: %0.2f\n", stdDevVal)
	// Output: Standard Deviation: 1.77

	// Sort the values.
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Quantiles
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	// Output: 25 Quantile: 1.60
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	// Output: 50 Quantile: 4.30
	fmt.Printf("75 Quantile: %0.2f\n", quant75)
	// Output: 75 Quantile: 5.10
}
