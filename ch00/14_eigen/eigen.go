package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {
	a := mat.NewDense(2, 2, []float64{
		3, 1,
		2, 4,
	})

	var eig mat.Eigen
	ok := eig.Factorize(a, mat.EigenRight)
	if !ok {
		log.Fatal("Eigendecomposition failed")
	}

	fmt.Printf("%#v \n", eig)
	// Output:
	// mat.Eigen{n:2, kind:4, values:[]complex128{(2+0i), (5+0i)}, rVectors:(*mat.CDense)(0xc0000b60c0), lVectors:(*mat.CDense)(nil)}
	fmt.Printf("%#v \n", eig.Values(nil))
	// Output:
	// []complex128{(2+0i), (5+0i)}
	E := mat.NewCDense(2, 2, []complex128{0 + 0i, 0 + 0i, 0 + 0i, 0 + 0i})
	eig.VectorsTo(E)
	fmt.Println(E)
	// Output:
	// &{{2 2 2 [(-0.7071067811865475+0i) (-0.4472135954999579+0i) (0.7071067811865475+0i) (-0.8944271909999159+0i)]} 2 2}
}
