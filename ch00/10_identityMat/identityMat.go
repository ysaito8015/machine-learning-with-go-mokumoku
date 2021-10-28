package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	i := mat.NewDiagonalRect(2, 2, []float64{1, 1})

	fmt.Println(mat.Formatted(i))
	// Output:
	// ⎡1  0⎤
	// ⎣0  1⎦

	i = mat.NewDiagonalRect(3, 3, []float64{1, 1, 1})

	fmt.Println(mat.Formatted(i))
	// Output:
	// ⎡1  0  0⎤
	// ⎢0  1  0⎥
	// ⎣0  0  1⎦

	i = mat.NewDiagonalRect(4, 4, []float64{1, 1, 1, 1})

	fmt.Println(mat.Formatted(i))
	// Output:
	// ⎡1  0  0  0⎤
	// ⎢0  1  0  0⎥
	// ⎢0  0  1  0⎥
	// ⎣0  0  0  1⎦
}
