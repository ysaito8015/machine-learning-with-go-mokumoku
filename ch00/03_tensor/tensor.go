package main

import (
	"fmt"
	//tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"gorgonia.org/tensor"
)

func main() {
	// (2,3,4) 3-Tensor を作成する
	A := tensor.New(tensor.WithBacking([]float64{
		//
		0, 1, 2, 3,
		2, 3, 4, 5,
		4, 5, 6, 7,
		//
		1, 2, 3, 4,
		3, 4, 5, 6,
		5, 6, 7, 8,
	}), tensor.WithShape(2, 3, 4))

	fmt.Println(A)
	// Output:
	// ⎡0  1  2  3⎤
	// ⎢2  3  4  5⎥
	// ⎣4  5  6  7⎦
	//
	// ⎡1  2  3  4⎤
	// ⎢3  4  5  6⎥
	// ⎣5  6  7  8⎦

	/*
		B := tf.NewTensor([][][]float64{
			{
				{
					{0.0, 1.0, 2.0, 3.0},
					{2.0, 3.0, 4.0, 5.0},
					{4.0, 5.0, 6.0, 7.0},
				},
			},
			{
				{
					{1.0, 2.0, 3.0, 4.0},
					{3.0, 4.0, 5.0, 6.0},
					{5.0, 6.0, 7.0, 8.0},
				},
			},
		})

		fmt.Println(B)
	*/
}
