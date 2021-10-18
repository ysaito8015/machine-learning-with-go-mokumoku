package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
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
	//fmt.Println(irisDF)

	fmt.Println("Sepal Length Sumamry Statistics")

	// Get the float values from the "sepal_length" column.
	sepalLength := irisDF.Col("sepal_length").Float()
	fmt.Println(sepalLength)
	// Output: [5.1 4.9 4.7 4.6 5 5.4 4.6 5 4.4 4.9 5.4 4.8 4.8 4.3 5.8 5.7 5.4 5.1 5.7 5.1 5.4 5.1 4.6 5.1 4.8 5 5 5.2 5.2 4.7 4.8 5.4 5.2 5.5 4.9 5 5.5
	// 4.9 4.4 5.1 5 4.5 4.4 5 5.1 4.8 5.1 4.6 5.3 5 7 6.4 6.9 5.5 6.5 5.7 6.3 4.9 6.6 5.2 5 5.9 6 6.1 5.6 6.7 5.6 5.8 6.2 5.6 5.9 6.1 6.3 6.1 6.4 6.6
	// 6.8 6.7 6 5.7 5.5 5.5 5.8 6 5.4 6 6.7 6.3 5.6 5.5 5.5 6.1 5.8 5 5.6 5.7 5.7 6.2 5.1 5.7 6.3 5.8 7.1 6.3 6.5 7.6 4.9 7.3 6.7 7.2 6.5 6.4 6.8 5.7
	// 5.8 6.4 6.5 7.7 7.7 6 6.9 5.6 7.7 6.3 6.7 7.2 6.2 6.1 6.4 7.2 7.4 7.9 6 .4 6.3 6.1 7.7 6.3 6.4 6 6.9 6.7 6.9 5.8 6.8 6.7 6.7 6.3 6.5 6.2 5.9]

	// Calculate the Mean of the variable.
	meanVal := stat.Mean(sepalLength, nil)
	fmt.Printf("mean: %0.2f\n", meanVal)
	// Output: mean: 5.84

	// Calculate the Mode of the variable.
	modeVal, modeCount := stat.Mode(sepalLength, nil)
	fmt.Printf("mode: %0.2f\n", modeVal)
	// Output: mode: 5.00
	fmt.Printf("mode count: %d\n", modeCount)
	// Output: mode count: 10

	// Calculate the Median of the variable.
	// this function come from montanaflynn/stats package
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("median: %0.2f\n", medianVal)
	// Output: median: 5.80

	fmt.Println("Petal Length Sumamry Statistics")
	// Get the float values from the "petal_length" column.
	petalLength := irisDF.Col("petal_length").Float()
	fmt.Println(petalLength)
	// Output: [1.4 1.4 1.3 1.5 1.4 1.7 1.4 1.5 1.4 1.5 1.5 1.6 1.4 1.1 1.2 1.5 1.3 1.4 1.7 1.5 1.7 1.5 1 1.7 1.9 1.6 1.6 1.5 1.4 1.6 1.6 1.5 1.5 1.4 1.5 1.2
	// 1.3 1.4 1.3 1.5 1.3 1.3 1.3 1.6 1.9 1.4 1.6 1.4 1.5 1.4 4.7 4.5 4.9 4 4.6 4.5 4.7 3.3 4.6 3.9 3.5 4.2 4 4.7 3.6 4.4 4.5 4.1 4.5 3.9 4.8 4 4.9 4.7 4.3
	// 4.4 4.8 5 4.5 3.5 3.8 3.7 3.9 5.1 4.5 4.5 4.7 4.4 4.1 4 4.4 4.6 4 3.3 4.2 4.2 4.2 4.3 3 4.1 6 5.1 5.9 5.6 5.8 6.6 4.5 6.3 5.8 6.1 5.1 5.3 5.5 5 5.1
	// 5.3 5.5 6.7 6.9 5 5.7 4.9 6.7 4.9 5.7 6 4.8 4.9 5.6 5.8 6 .1 6.4 5.6 5.1 5.6 6.1 5.6 5.5 4.8 5.4 5.6 5.1 5.1 5.9 5.7 5.2 5 5.2 5.4 5.1]

	meanVal = stat.Mean(petalLength, nil)
	fmt.Printf("mean: %0.2f\n", meanVal)
	// Output: mean: 3.76
	modeVal, modeCount = stat.Mode(petalLength, nil)
	fmt.Printf("mode: %0.2f\n", modeVal)
	// Output: mode: 1.40
	fmt.Printf("mode count: %v\n", modeCount)
	// Output: mode count: 13
	medianVal, err = stats.Median(petalLength)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("median: %0.2f\n", medianVal)
	// Output: median: 4.35
}
