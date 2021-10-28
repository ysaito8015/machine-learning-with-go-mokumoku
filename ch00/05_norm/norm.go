package main

import (
	"fmt"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := []float64{1, 1, -1, -1}

	// L1
	fmt.Println(floats.Norm(a, 1))
	// Output: 4

	// L2
	fmt.Println(floats.Norm(a, 2))
	// Output: 2

	b := mat.NewVecDense(4, []float64{1, 1, -1, -1})

	// L1
	fmt.Println(mat.Norm(b, 1))
	// Output: 4

	// L2
	fmt.Println(mat.Norm(b, 2))
	// Output: 2

	// using blas L2
	fmt.Println(blas64.Nrm2(b.RawVector()))
	// Output: 2
}
