package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewVecDense(2, []float64{2, 3})

	A := mat.NewDense(2, 2, []float64{
		2, -1,
		2, -2,
	})

	var b mat.Dense
	b.Mul(A, a)
	fmt.Println(mat.Formatted(&b))
	// Output:
	// ⎡ 1⎤
	// ⎣-2⎦
}
