package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func cosSim(v1, v2 mat.Vector) float64 {
	return mat.Dot(v1, v2) / (mat.Norm(v1, 2) * mat.Norm(v2, 2))
}

func main() {
	a := mat.NewVecDense(4, []float64{2, 2, 2, 2})
	b := mat.NewVecDense(4, []float64{1, 1, 1, 1})
	c := mat.NewVecDense(4, []float64{-1, -1, -1, -1})

	fmt.Println(cosSim(a, b))
	// Output: 1
	fmt.Println(cosSim(a, c))
	// Output: -1
}
