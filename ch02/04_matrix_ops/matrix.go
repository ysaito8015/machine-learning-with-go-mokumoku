package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
)

func main() {
	// Create matrices.
	// mat.NewDense(Rows, Columns, slice of elements)
	a := mat.NewDense(3, 3, []float64{
		1, 2, 3,
		0, 4, 5,
		0, 0, 6,
	})
	b := mat.NewDense(3, 3, []float64{
		8, 9, 10,
		1, 4, 2,
		9, 0, 2,
	})
	c := mat.NewDense(3, 2, []float64{
		3, 2,
		1, 4,
		0, 8,
	})

	// Add a and b.
	//d := mat.NewDense(0, 0, nil)
	// Output: panic: mat: zero length in matrix dimension
	var d mat.Dense
	fmt.Printf("%#v\n%#v\n", a, d)
	// Output: &mat.Dense{mat:blas64.General{Rows:3, Cols:3, Data:[]float64{1, 2, 3, 0, 4, 5, 0, 0, 6}, Stride:3}, capRows:3, capCols:3}
	// Output: mat.Dense{mat:blas64.General{Rows:0, Cols:0, Data:[]float64(nil), Stride:0}, capRows:0, capCols:0}
	d.Add(a, b)

	//fd := mat.Formatted(d, mat.Prefix("          "))
	// Output: ./matrix.go:18:21: cannot use d (type mat.Dense) as type mat.Matrix in argument to mat.Formatted:
	//  mat.Dense does not implement mat.Matrix (At method has pointer receiver)
	//
	// * mat.Formatted() uses with *mat.Dense
	// var d *mat.Dense
	// fd := mat.Formatted(d, mat.Prefix(""))
	// Output: (*mat.Dense)(nil)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x49fa58]
	fd := mat.Formatted(&d, mat.Prefix(""))
	fmt.Printf("d = a + b = \n%0.4v\n\n", fd)
	// Output:d = a + b =
	// ⎡ 9  11  13⎤
	// ⎢ 1   8   7⎥
	// ⎣ 9   0   8⎦

	// Multiply a and c.
	var f mat.Dense
	f.Mul(a, c)
	ff := mat.Formatted(&f, mat.Prefix(""))
	fmt.Printf("f = a・b = \n%0.4v\n\n", ff)
	// Output: f = a・b =
	// ⎡ 5  34⎤
	// ⎢ 4  56⎥
	// ⎣ 0  48⎦

	// Raising a matrix to a power.
	var g mat.Dense
	g.Pow(a, 5)
	fg := mat.Formatted(&g, mat.Prefix(""))
	fmt.Printf("g = a^5 = \n%0.4v\n\n", fg)
	// Output: g = a^5 =
	// ⎡        1        682  1.074e+04⎤
	// ⎢        0       1024  1.688e+04⎥
	// ⎣        0          0       7776⎦

	// Apply a function to each of the elements of a.
	var h mat.Dense
	sqrt := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
	h.Apply(sqrt, a)
	fh := mat.Formatted(&h, mat.Prefix(""))
	fmt.Printf("h = sqrt(a) = \n%0.4v\n\n", fh)
	// Output: h = sqrt(a) =
	// ⎡    1  1.414  1.732⎤
	// ⎢    0      2  2.236⎥
	// ⎣    0      0  2.449⎦

	// Compute and output the transpose of the matrix.
	ft := mat.Formatted(a.T(), mat.Prefix(""))
	fmt.Printf("a^T = \n%v\n\n", ft)
	// Ouput: a^T =
	// ⎡1  0  0⎤
	// ⎢2  4  0⎥
	// ⎣3  5  6⎦

	// Compute and output the determinant of a.
	deta := mat.Det(a)
	fmt.Printf("det(a) = %v\n\n", deta)
	// Output: det(a) = 23.999999999999993

	// Compute and output the inverse of a.
	var aInverse mat.Dense
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(&aInverse, mat.Prefix(""))
	fmt.Printf("a^-1 =\n%v\n\n", fi)
	// Output: a^-1 =
	// ⎡                   1                  -0.5  -0.08333333333333333⎤
	// ⎢                   0                  0.25  -0.20833333333333331⎥
	// ⎣                   0                     0   0.16666666666666666⎦
}
