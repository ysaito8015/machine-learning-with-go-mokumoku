package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := []float64{1, 2, 3}
	b := []float64{3, 2, 1}

	fmt.Println(floats.Dot(a, b))

	sum := 0.0
	for i, v := range a {
		sum += v * b[i]
	}

	fmt.Println(sum)

	c := mat.NewVecDense(3, []float64{1, 2, 3})
	d := mat.NewVecDense(3, []float64{3, 2, 1})

	fmt.Println(mat.Dot(c, d))
}
