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

	b := mat.NewDense(2, 3, []float64{
		0, 1, 2,
		1, 2, 3,
	})

	var c mat.Dense
	c.Mul(a, b.T())
	fmt.Println(mat.Formatted(&c))
	// Output:
	// ⎡ 5   8⎤
	// ⎣ 8  14⎦
}
