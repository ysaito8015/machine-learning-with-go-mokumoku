package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(3, 3, []float64{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	})

	b := mat.NewDense(3, 3, []float64{
		0, 1, 2,
		2, 0, 1,
		1, 2, 0,
	})

	var c mat.Dense
	c.MulElem(a, b)
	fmt.Println(mat.Formatted(&c))
	// Output:
	// ⎡ 0   1   4⎤
	// ⎢ 6   0   5⎥
	// ⎣ 6  14   0⎦

}
