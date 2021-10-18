package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Initialize a "vector" via a slice.
	var myvector []float64

	// Add a couple of components to the vector.
	myvector = append(myvector, 11.0)
	myvector = append(myvector, 5.2)

	// Output the results to stdout.
	fmt.Println(myvector)
	// Output: [11 5.2]

	// Function name was changed
	// https://pkg.go.dev/gonum.org/v1/gonum/mat#VecDense
	//myvector2 := mat.NewVector(2, []float64{11.0, 5.2})
	myvector2 := mat.NewVecDense(2, []float64{11.0, 5.2})
	fmt.Println(myvector2)
	// Output: &{{2 [11 5.2] 1}}
}
