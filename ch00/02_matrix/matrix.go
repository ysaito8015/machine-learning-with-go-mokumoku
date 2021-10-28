package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// gonum/mat パッケージの NewDense コンストラクタ関数で行列を定義する
	A := mat.NewDense(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})

	B := mat.NewDense(3, 2, []float64{
		0.21, 0.14,
		-1.3, 0.81,
		0.12, -2.1,
	})

	fmt.Println(A)
	// Output: &{{2 3 [1 2 3 4 5 6] 3} 2 3}
	fmt.Println(B)
	// Output: &{{3 2 [0.21 0.14 -1.3 0.81 0.12 -2.1] 2} 3 2}
}
