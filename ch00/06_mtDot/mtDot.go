package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 3, []float64{
		0, 1, 2,
		1, 2, 3,
	})

	b := mat.NewDense(3, 2, []float64{
		2, 1,
		2, 1,
		2, 1,
	})

	var c mat.Dense
	c.Mul(a, b)
	fmt.Println(mat.Formatted(&c))
	// Output:
	// ⎡ 6   3⎤
	// ⎣12   6⎦
}
