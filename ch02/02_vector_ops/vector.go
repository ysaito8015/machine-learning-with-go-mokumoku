package main

import (
	"fmt"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	//Initialize a couple of "vectors" represented as slices.
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	// Compute the dot product of A and B.
	// https://pkg.go.dev/gonum.org/v1/gonum@v0.9.3/floats#Dot
	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)
	// Output: The dot product of A and B is: -63.99

	// Scale each element of A by 1.5.
	floats.Scale(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)
	// Output: Scaling A by 1.5 gives: [16.5 7.800000000000001 -1.9500000000000002]

	// Compute the norm/length of B.
	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
	// Output: The norm/length of B is: 9.77

	// Initialize a couple of "vectors" represented as slices.
	vectorC := mat.NewVecDense(3, []float64{11.0, 5.2, -1.3})
	vectorD := mat.NewVecDense(3, []float64{-7.2, 4.2, 5.1})

	// Compute the dot product of C and D.
	dotProduct = mat.Dot(vectorC, vectorD)
	fmt.Printf("The dot product of C and D is: %0.2f\n", dotProduct)
	// Output: The dot product of C and D is: -63.99

	// Scale each element of C by 1.5.
	vectorC.ScaleVec(1.5, vectorC)
	fmt.Printf("Scaling C by 1.5 fibes: %v\n", vectorC)
	// Output: Scaling C by 1.5 fibes: &{{3 [16.5 7.800000000000001 -1.9500000000000002] 1}}

	// Compute the norm/length of D.
	// blas64.Nrm2() function is changed.
	// https://pkg.go.dev/gonum.org/v1/gonum@v0.9.3/blas/blas64#Nrm2
	normD := blas64.Nrm2(vectorD.RawVector())
	fmt.Printf("The norm/length of D is: %0.2f\n", normD)
	// Output: The norm/length of D is: 9.77
}
