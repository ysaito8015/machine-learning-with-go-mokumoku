package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	//Create a flat representation of our matrix.
	components := []float64{
		1.2, -5.7,
		-2.4, 7.3,
	}

	// Form our matrix.
	// mat.NewDense(Rows, Columns, slice)
	a := mat.NewDense(2, 2, components)
	fmt.Printf("a = %v\n\n", a)
	// Output: a = &{{2 2 [1.2 -5.7 -2.4 7.3] 2} 2 2}

	// As a sanity check, output the matrix to stdout.
	fa := mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("mat = \n%v\n\n", fa)
	// Output: mat =
	// ⎡ 1.2  -5.7⎤
	// ⎣-2.4   7.3⎦

	// Get a single value from the matrix.
	val := a.At(0, 1)
	fmt.Printf("The value of a at (0,1) is: %.2f\n\n", val)
	// Output: The value of a at (0,1) is: -5.70

	// Get the values in a specific column.
	col := mat.Col(nil, 0, a)
	fmt.Printf("The values in the 1st column are: %v\n\n", col)
	// Output: The values in the 1st column are: [1.2 -2.4]

	// Get the values in a specific row.
	row := mat.Row(nil, 1, a)
	fmt.Printf("The values in the 2nd row are: %v\n\n", row)
	// Output: The values in the 2nd row are: [-2.4 7.3]

	// Modify a single element.
	a.Set(0, 1, 11.2)
	fmt.Printf("The value of a at (0,1) is: %.2f\n\n", a.At(0, 1))
	// Output: The value of a at (0,1) is: 11.20

	// Modify an entire row.
	a.SetRow(0, []float64{14.3, -4.2})
	fmt.Printf("The values in the 1st row are: %v\n\n", mat.Row(nil, 0, a))
	// Output: The values in the 1st row are: [14.3 -4.2]

	// Modify an entire column.
	a.SetCol(0, []float64{1.7, -0.3})
	fmt.Printf("The values in the 1st column are: %v\n\n", mat.Col(nil, 0, a))
	// Output: The values in the 1st column are: [1.7 -0.3]

	fa = mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("mat = \n%v\n\n", fa)
	// Output: mat =
	// ⎡ 1.7  -4.2⎤
	// ⎣-0.3   7.3⎦
}
