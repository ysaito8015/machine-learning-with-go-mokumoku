package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// スライスでベクトルを表現
	a := []float64{1, 2, 3}
	b := []float64{-2.4, 0.25, -1.3, 1.8, 0.61}
	fmt.Println(a)
	// Output: [1 2 3]
	fmt.Println(b)
	// Output: [-2.4 0.25 -1.3 1.8 0.61]

	// gonum/mat パッケージの, NewVecDense コンストラクタ関数でベクトルを定義する
	c := mat.NewVecDense(3, []float64{1, 2, 3})
	d := mat.NewVecDense(5, []float64{11.0, 5.2, -1.3, -7.2, 4.2})
	fmt.Println(c)
	// Output: &{{3 [1 2 3] 1}}
	fmt.Println(d)
	// Output: &{{5 [11 5.2 -1.3 -7.2 4.2] 1}}
}
