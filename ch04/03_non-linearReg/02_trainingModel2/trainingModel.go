package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ysaito8015/ridge"
	"gonum.org/v1/gonum/mat"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 4

	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	featureData := make([]float64, 4*len(rawCSVData))
	yData := make([]float64, len(rawCSVData))

	var featureIndex int
	var yIndex int

	for idx, record := range rawCSVData {
		if idx == 0 {
			continue
		}

		for i, val := range record {
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			if i < 3 {
				if i == 0 {
					featureData[featureIndex] = 1
					featureIndex++
				}

				featureData[featureIndex] = valParsed
				featureIndex++
			}
			if i == 3 {
				yData[yIndex] = valParsed
				yIndex++
			}
		}
	}

	features := mat.NewDense(len(rawCSVData), 4, featureData)
	y := mat.NewVecDense(len(rawCSVData), yData)

	r := ridge.New(features, y, 1.0)

	r.Regress()

	c1 := r.Coefficients.At(0, 0)
	c2 := r.Coefficients.At(1, 0)
	c3 := r.Coefficients.At(2, 0)
	c4 := r.Coefficients.At(3, 0)
	fmt.Printf("Regression formula:\n")
	fmt.Printf("y = %0.3f + %0.3f TV + %0.3f Radio + %0.3f Newspaper\n", c1, c2, c3, c4)
	// Output:
	// y = 3.038 + 0.047 TV + 0.177 Radio + 0.001 Newspaper

}
