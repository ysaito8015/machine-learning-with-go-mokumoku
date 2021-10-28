package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {
	a := mat.NewDense(2, 2, []float64{
		1, 2,
		3, 4,
	})

	var aInverse mat.Dense
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fmt.Println(mat.Formatted(&aInverse))
	// Output:
	// ⎡-1.9999999999999996   0.9999999999999998⎤
	// ⎣ 1.4999999999999998  -0.4999999999999999⎦

	b := mat.NewDense(2, 2, []float64{
		1, 2,
		0, 0,
	})

	var bInverse mat.Dense
	if err := bInverse.Inverse(b); err != nil {
		log.Fatal(err)
		// Output:
		// 2021/10/28 19:52:01 matrix singular or near-singular with condition number +Inf
		// exit status 1
	}
	fmt.Println(mat.Formatted(&bInverse))

}
