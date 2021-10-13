package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	fmt.Println(irisDF)

	// Create a filter
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// Filter the dataframe to see only the rows where
	// the iris species is "Iris-versicolor".
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println("==== Filter only versicolor ====")
	fmt.Println(versicolorDF)

	// Filter the dataframe.
	// Select out the sepal width and species columns.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})

	fmt.Println("==== Filter only sepal width of versicolor ====")
	fmt.Println(versicolorDF)

	// Filter and select the dataframe.
	// only display the first three results.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})

	fmt.Println("==== Filter only the first three results of sepal width of versicolor ====")
	fmt.Println(versicolorDF)
}
