package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})

	fmt.Printf("%#v\n", a.T())
	// Output:
	// mat.Transpose{Matrix:(*mat.Dense)(0xc0000b6040)}
	fmt.Println(mat.Formatted(a.T()))
	// Output:
	// ⎡1  4⎤
	// ⎢2  5⎥
	// ⎣3  6⎦
}
