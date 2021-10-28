package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 2, []float64{
		1, 2,
		3, 4,
	})

	fmt.Println(mat.Det(a))
	// Output: -2

	b := mat.NewDense(2, 2, []float64{
		1, 2,
		0, 0,
	})

	fmt.Println(mat.Det(b))
	// Output: 0
}
