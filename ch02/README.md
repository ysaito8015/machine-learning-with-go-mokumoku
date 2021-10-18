# Matrices, Probability, and Statistics
- a fundamental understanding of probability and statistics will allow us to match certain algorithms with relevant problems, understand our data and result, and apply neccessary transformations to our data.
- Matrices and a linear algebra will then allow us to prperly represent out data and implement optimizations, minimizations, and matrix-based transformations.

## Matrices and vectors
- `gonum` パッケージ
    - `https://github.com/gonum`


### Vectors
- A vector is an orderd collection of numbers arranged in either a row (left to right) or column (up and down).
    - Each of the numbers in a vector is called a component.

#### Slices
- Go slices to represent these ordered collections of data.
- Slices are indeed ordered collections.
- gonum provides
    - `gonum.org/v1/gonum/floats`
        - to operate on slices of `float64` values
    - `gonum.org/v1/gonum/mat`
        - which along with matrices


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// Initialize a "vector" via a slice.
	var myvector []float64

	// Add a couple of components to the vector.
	myvector = append(myvector, 11.0)
	myvector = append(myvector, 5.2)

	// Output the results to stdout.
	fmt.Println(myvector)
	// Output: [11 5.2]

	// Function name was changed
	//myvector2 := mat.NewVector(2, []float64{11.0, 5.2})
	myvector2 := mat.NewVecDense(2, []float64{11.0, 5.2})
	fmt.Println(myvector2)
	// Output: &{{2 [11 5.2] 1}}
}
```


### Vector operations
- dot products, sorting, and distance.


```go
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
```


#### mat.NewVecDense

```shell
$ go doc mat.NewVecDense
package mat // import "gonum.org/v1/gonum/mat"

func NewVecDense(n int, data []float64) *VecDense
    NewVecDense creates a new VecDense of length n. If data == nil, a new slice
    is allocated for the backing slice. If len(data) == n, data is used as the
    backing slice, and changes to the elements of the returned VecDense will be
    reflected in data. If neither of these is true, NewVecDense will panic.
    NewVecDense will panic if n is zero.

```

### Matrices
- Matrices are just rectangular organizations of numbers, and
- linear algebra dictates the rules associated with their manipulation.


```go
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
	a := mat.NewDense(2, 2, components)
	fmt.Printf("a = %v\n\n", a)
	// Output: a = &{{2 2 [1.2 -5.7 -2.4 7.3] 2} 2 2}

	// As a sanity check, output the matrix to stdout.
	fa := mat.Formatted(a, mat.Prefix(" "))
	fmt.Printf("mat = %v\n\n", fa)
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
```


### Matrix operations
- take special care when doing things such as multiplying matrices together or taking an inverse.
- `mat.Apply()` 関数は, 行列の要素に関数を適用できる
- 大きさ 0 の行列は作れない `mat.NewDense(0, 0, nil)`
- `Dense` 構造体を, `mat.Formatted()` に渡すときは, `&` でポインタ型にする

```go
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
	d.Add(a, b)
	// NewDense will panic if either r or c is zero.
	//fd := mat.Formatted(d, mat.Prefix("          "))
	// Output: ./matrix.go:18:21: cannot use d (type mat.Dense) as type mat.Matrix in argument to mat.Formatted:
	//  mat.Dense does not implement mat.Matrix (At method has pointer receiver)
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
  // Matrices don't always have inverses.
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
```


#### mat.Dense / mat.Formatted
- `var foo mat.Dense` で定義した変数は, `mat.Dence` 構造体型
- `bar := mat.NewDense()` コンストラクタ関数で作られた変数は, `mat.Matrix` インタフェース型を満たす

```shell
$ go doc mat.Dense    
package mat // import "gonum.org/v1/gonum/mat"

type Dense struct {
        // Has unexported fields.
}
    Dense is a dense matrix representation.

$ go doc mat.NewDense
package mat // import "gonum.org/v1/gonum/mat"

func NewDense(r, c int, data []float64) *Dense
    NewDense creates a new Dense matrix with r rows and c columns. If data ==
    nil, a new slice is allocated for the backing slice. If len(data) == r*c,
    data is used as the backing slice, and changes to the elements of the
    returned Dense will be reflected in data. If neither of these is true,
    NewDense will panic. NewDense will panic if either r or c is zero.

    The data must be arranged in row-major order, i.e. the (i*c + j)-th element
    in the data slice is the {i, j}-th element in the matrix.

$ go doc mat.Matrix
package mat // import "gonum.org/v1/gonum/mat"

type Matrix interface {
        // Dims returns the dimensions of a Matrix.
        Dims() (r, c int)

        // At returns the value of a matrix element at row i, column j.
        // It will panic if i or j are out of bounds for the matrix.
        At(i, j int) float64

        // T returns the transpose of the Matrix. Whether T returns a copy of the
        // underlying data is implementation dependent.
        // This method may be implemented using the Transpose type, which
        // provides an implicit matrix transpose.
        T() Matrix
}
    Matrix is the basic matrix interface type.

$ go doc mat.Formatted
package mat // import "gonum.org/v1/gonum/mat"

func Formatted(m Matrix, options ...FormatOption) fmt.Formatter
    Formatted returns a fmt.Formatter for the matrix m using the given options.

```

## Statistics
- understanding of statistics
    1. check the quality of data
    2. understanding of the data
    3. evaluation/validation of the results


### Distributions
### Statistical measures
- `gonum/stat` パッケージ
    - https://pkg.go.dev/gonum.org/v1/gonum@v0.9.3/stat
- `stats` パッケージ
    - https://github.com/montanaflynn/stats

#### Measures of central tendency
- Mean, Median, Mode


```go
package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	irisFile, err := os.Open("./iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)
	//fmt.Println(irisDF)

	fmt.Println("Sepal Length Sumamry Statistics")

	// Get the float values from the "sepal_length" column.
	sepalLength := irisDF.Col("sepal_length").Float()
	fmt.Println(sepalLength)
	// Output: [5.1 4.9 4.7 4.6 5 5.4 4.6 5 4.4 4.9 5.4 4.8 4.8 4.3 5.8 5.7 5.4 5.1 5.7 5.1 5.4 5.1 4.6 5.1 4.8 5 5 5.2 5.2 4.7 4.8 5.4 5.2 5.5 4.9 5 5.5
	// 4.9 4.4 5.1 5 4.5 4.4 5 5.1 4.8 5.1 4.6 5.3 5 7 6.4 6.9 5.5 6.5 5.7 6.3 4.9 6.6 5.2 5 5.9 6 6.1 5.6 6.7 5.6 5.8 6.2 5.6 5.9 6.1 6.3 6.1 6.4 6.6
	// 6.8 6.7 6 5.7 5.5 5.5 5.8 6 5.4 6 6.7 6.3 5.6 5.5 5.5 6.1 5.8 5 5.6 5.7 5.7 6.2 5.1 5.7 6.3 5.8 7.1 6.3 6.5 7.6 4.9 7.3 6.7 7.2 6.5 6.4 6.8 5.7
	// 5.8 6.4 6.5 7.7 7.7 6 6.9 5.6 7.7 6.3 6.7 7.2 6.2 6.1 6.4 7.2 7.4 7.9 6 .4 6.3 6.1 7.7 6.3 6.4 6 6.9 6.7 6.9 5.8 6.8 6.7 6.7 6.3 6.5 6.2 5.9]

	// Calculate the Mean of the variable.
	meanVal := stat.Mean(sepalLength, nil)
	fmt.Printf("mean: %0.2f\n", meanVal)
	// Output: mean: 5.84

	// Calculate the Mode of the variable.
	modeVal, modeCount := stat.Mode(sepalLength, nil)
	fmt.Printf("mode: %0.2f\n", modeVal)
	// Output: mode: 5.00
	fmt.Printf("mode count: %d\n", modeCount)
	// Output: mode count: 10

	// Calculate the Median of the variable.
	// this function come from montanaflynn/stats package
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("median: %0.2f\n", medianVal)
	// Output: median: 5.80

	fmt.Println("Petal Length Sumamry Statistics")
	// Get the float values from the "petal_length" column.
	petalLength := irisDF.Col("petal_length").Float()
	fmt.Println(petalLength)
	// Output: [1.4 1.4 1.3 1.5 1.4 1.7 1.4 1.5 1.4 1.5 1.5 1.6 1.4 1.1 1.2 1.5 1.3 1.4 1.7 1.5 1.7 1.5 1 1.7 1.9 1.6 1.6 1.5 1.4 1.6 1.6 1.5 1.5 1.4 1.5 1.2
	// 1.3 1.4 1.3 1.5 1.3 1.3 1.3 1.6 1.9 1.4 1.6 1.4 1.5 1.4 4.7 4.5 4.9 4 4.6 4.5 4.7 3.3 4.6 3.9 3.5 4.2 4 4.7 3.6 4.4 4.5 4.1 4.5 3.9 4.8 4 4.9 4.7 4.3
	// 4.4 4.8 5 4.5 3.5 3.8 3.7 3.9 5.1 4.5 4.5 4.7 4.4 4.1 4 4.4 4.6 4 3.3 4.2 4.2 4.2 4.3 3 4.1 6 5.1 5.9 5.6 5.8 6.6 4.5 6.3 5.8 6.1 5.1 5.3 5.5 5 5.1
	// 5.3 5.5 6.7 6.9 5 5.7 4.9 6.7 4.9 5.7 6 4.8 4.9 5.6 5.8 6 .1 6.4 5.6 5.1 5.6 6.1 5.6 5.5 4.8 5.4 5.6 5.1 5.1 5.9 5.7 5.2 5 5.2 5.4 5.1]

	meanVal = stat.Mean(petalLength, nil)
	fmt.Printf("mean: %0.2f\n", meanVal)
	// Output: mean: 3.76
	modeVal, modeCount = stat.Mode(petalLength, nil)
	fmt.Printf("mode: %0.2f\n", modeVal)
	// Output: mode: 1.40
	fmt.Printf("mode count: %v\n", modeCount)
	// Output: mode count: 13
	medianVal, err = stats.Median(petalLength)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("median: %0.2f\n", medianVal)
	// Output: median: 4.35
}
```


#### Measures of spread or dispersion


```go
package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	//"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	irisFile, err := os.Open("./iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Get the float values from the "sepal_length" column.
	sepalLength := irisDF.Col("petal_length").Float()

	// Calculate the Min of the variable.
	minVal := floats.Min(sepalLength)
	fmt.Printf("Minimum: %0.2f\n", minVal)
	// Output:  Minimam: 1.00

	// Calculate the Max of the variable.
	maxVal := floats.Max(sepalLength)
	fmt.Printf("Maximum: %0.2fv\n", maxVal)
	// Output: Maximum: 6.90

	// Calculate the range of the variable.
	rangeVal := maxVal - minVal
	fmt.Printf("Range: %0.2fv\n", rangeVal)
	// Range: 5.90

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)
	fmt.Printf("Variance: %0.2f\n", varianceVal)
	// Output: Variance: 3.11

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(sepalLength, nil)
	fmt.Printf("Standard Deviation: %0.2f\n", stdDevVal)
	// Output: Standard Deviation: 1.77

	// Sort the values.
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Quantiles
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	// Output: 25 Quantile: 1.60
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	// Output: 50 Quantile: 4.30
	fmt.Printf("75 Quantile: %0.2f\n", quant75)
	// Output: 75 Quantile: 5.10
}
```


### Visualizing distributions
- plot examples
    - https://github.com/gonum/plot/wiki/Example-plots

#### Histograms

```go
package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	irisFile, err := os.Open("./iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	/// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Create a histogram for each of the feature columns in the dataset.
	for _, colName := range irisDF.Names() {
		// If the column is one of the feature columns, create a histogram
		// of the values.
		if colName != "species" {

			// Create a plotter.Values value and fill it with the
			// values from the respective column of the dataframe.
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// Make a plot and set its title.
			// https://pkg.go.dev/gonum.org/v1/plot#New
			// p, err := plot.New()
			p := plot.New()
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			// Create a histogram of the values drawn
			// from the stanndard normal.
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			// Normalize the histogram.
			h.Normalize(1)

			// Add the histogram to the plot.
			p.Add(h)

			// Save the plot to a PNG file.
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
```

- ![](https://i.gyazo.com/f78048a75352f5d28b749d36a0f84aa1.png)
- ![](https://i.gyazo.com/ee811dc861288236888792319dba70a4.png)
- ![](https://i.gyazo.com/0b6448f6c48c3043086cb6f44a6d9a0f.png)
- ![](https://i.gyazo.com/f551288158801a3bf7ed4068ec86e9e2.png)


```shell
$ go doc plotter.Normalize
package plotter // import "gonum.org/v1/plot/plotter"

func (h *Histogram) Normalize(sum float64)
    Normalize normalizes the histogram so that the total area beneath it sums to
    a given value.

$ go doc plot.New  
package plot // import "gonum.org/v1/plot"

func New() *Plot
    New returns a new plot with some reasonable default settings.

```


#### Box plots


```go
package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
)

func main() {
	// Open the CSV file.
	irisFile, err := os.Open("./iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	// Create the plot and set its title and axis label.
	//
	// p, err := plot.New()
	// ./boxPlot.go|24 col 12 error| compiler:Error:WrongAssignCount:cannot initialize 2 variables with 1 values
	// func New() *Plot {}
	// type Plot struct { Title struct { Text string; Padding vg.Length; TextStyle text.Style; } BackgroundColor color.Color
	// X, Y Axis; Legend Legend; TextHandler text.Handler }
	p := plot.New()

	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	// Create the box for our data.
	//
	// func Points(pt float64) Length
	// Points returns a length for the given number of points.
	w := vg.Points(50)
	// fmt.Printf("w: %#v\n", w)

	// fmt.Printf("irisDF.Names(): %#v\n", irisDF.Names())
	// Output: irisDF.Names(): []string{"sepal_length", "sepal_width", "petal_length", "petal_width", "species"}

	// Create a plotter.Values value and fill it with the
	// values from the respective column of the dataframe.
	for idx, colName := range irisDF.Names() {
		// fmt.Printf("irisDF.Names() -> colName: %#v\n", colName)
		// Output: irisDF.Names() -> colName: "species"

		// If the column is one of the feature column, create
		// a histogram of the values.
		if colName != "species" {

			v := make(plotter.Values, irisDF.Nrow())
			// fmt.Printf("irisDF: %#v\n", irisDF)
			// Output: irisDF:
			// dataframe.DataFrame{
			//   columns:[]series.Series{
			//     series.Series{
			//         Name:"sepal_length",
			//         elements:series.floatElements{
			//             series.floatElement{e:5.1, nan:false}
			//
			// fmt.Printf("irisDF.Col(colName): %#v\n", irisDF.Col(colName))
			// Output: irisDF.Col(colName):
			// series.Series{
			//    Name:"species",
			//    elements:series.stringElements{
			//        series.stringElement{e:"Iris-setosa", nan:false},
			//
			// fmt.Printf("irisDF.Col(colName).Float(): [%v], %#v\n", irisDF.Col(colName).Name, irisDF.Col(colName).Float())
			// Output:
			// irisDF.Col(colName).Float(): [sepal_length], []float64{5.1, 4.9, 4.7, 4.6, 5, ...
			// irisDF.Col(colName).Float(): [sepal_width], []float64{3.5, 3, 3.2, 3.1, 3.6, 3.9, ...
			// irisDF.Col(colName).Float(): [petal_length], []float64{1.4, 1.4, 1.3, 1.5, 1.4, 1.7, ...
			// irisDF.Col(colName).Float(): [petal_width], []float64{0.2, 0.2, 0.2, 0.2, 0.2, ...
			// irisDF.Col(colName).Float(): [species], []float64{NaN, NaN, NaN, NaN, NaN, ...
			//
			// func (df DataFrame) Col(colname string) series.Series
			// Col returns a copy of the Series with the given column name contained in the DataFrame.
			//
			// func (s Series) Float() []float64
			// Float returns the elements of a Series as a []float64. If the elements
			// can not be converted to float64 or contains a NaN returns
			// the float representation of NaN.
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			//fmt.Printf("v: %#v\n", v)
			// Output: v: plotter.Values{NaN, NaN, NaN, NaN, NaN, NaN,...
			//fmt.Printf("w: %#v\n", w)
			// Output: w: 50
			//
			// Add the data to the plot.
			// func NewBoxPlot(w vg.Length, loc float64, values Valuer) (*BoxPlot, error)
			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)
		}
	}

	// Set the X axis of the plot to nominal with
	// the given names for x=0, x=1, etc.
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")

	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}

}
```


## Probability
- Probability has to do with the likelihood of events or observations.
- `gonum/stat/distuv` パッケージ
    - https://github.com/gonum/gonum/tree/master/stat/distuv
- `prob` パッケージ, 習作
    - https://github.com/atgjack/prob

### Random variables
- Define a variable whose value could be one of the outcomes.
- This variable is referred to as a random variable.

### Probability measures

### Independent and conditional probability

### Hypothesis tesing

#### Test statistics
- `gonum.org/v1/gonum/stat` パッケージ
    - chi-square statistic
        - https://pkg.go.dev/gonum.org/v1/gonum/stat#ChiSquare


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// Define observed and expected values.
	observed := []float64{48, 52}
	expected := []float64{50, 50}

	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)

	// Calculate the degree of freedom
	df := len(observed) - 1
	totalA := observed[0] + expected[0]
	totalB := observed[1] + expected[1]
	totalObs := observed[0] + observed[1]
	totalExp := expected[0] + expected[1]
	total := totalObs + totalExp

	// Chi-squared statistic
	subAsq := (observed[0] - expected[0]) * (observed[0] - expected[0])
	subBsq := (observed[1] - expected[1]) * (observed[1] - expected[1])
	chiSq := subAsq/expected[0] + subBsq/expected[1]

	// Create a Chi-Squared distribution
	dist := distuv.ChiSquared{
		K: float64(df),
	}

	// Calculate Chi-Squate(df=1, 0.05)
	// Quantile returns the inverse of the cumulative distribution function.
	// https://pkg.go.dev/github.com/ArkaGPL/gonum@v0.8.5/stat/distuv#ChiSquared.Quantile
	quant := dist.Quantile(0.95)

	prob := 1 - dist.CDF(chiSquare)

	fmt.Println(">>   Pearson's chi-squared test")
	fmt.Println("|       | Observed | Expected |  Total |")
	fmt.Println("|-------|----------|----------|--------|")
	fmt.Printf("|   A   | %8.2f | %8.2f | %6.2f |\n", observed[0], expected[0], totalA)
	fmt.Printf("|   B   | %8.2f | %8.2f | %6.2f |\n", observed[1], expected[1], totalB)
	fmt.Printf("| total | %8.2f | %8.2f | %6.2f |\n", totalObs, totalExp, total)
	fmt.Printf("Chi-Square statistic: %v\n", chiSquare)
	fmt.Printf("Chi-Square statistic: %v\n", chiSq)
	fmt.Printf("degree of freedom: %v\n", df)

	fmt.Printf("Chi-Squre(%v,0.05): %.4f\n", df, quant)
	fmt.Printf("p-value: %.4f\n", prob)

	/*
	  // output from R
	  > vx <- c(48, 52)
	  > chisq.test(x=vx, p=c(0.5,0.5))

	          Chi-squared test for given probabilities

	  data:  vx
	  X-squared = 0.16, df = 1, p-value = 0.6892

	*/

	/*
		>>   Pearson's chi-squared test
		|       | Observed | Expected |  Total |
		|-------|----------|----------|--------|
		|   A   |    48.00 |    50.00 |  98.00 |
		|   B   |    52.00 |    50.00 | 102.00 |
		| total |   100.00 |   100.00 | 200.00 |
		Chi-Square statistic: 0.16
		Chi-Square statistic: 0.16
		degree of freedom: 1
		Chi-Squre(1,0.05): 3.8415
	*/
}
```


##### `gonum/stat.ChiSquare` 関数
- ピアソンのχ二乗検定が実装されている

```go
// $ go doc --src stat.ChiSquare
package stat // import "gonum.org/v1/gonum/stat"

// ChiSquare computes the chi-square distance between the observed frequencies 'obs' and
// expected frequencies 'exp' given by:
//  \sum_i (obs_i-exp_i)^2 / exp_i
//
// The lengths of obs and exp must be equal.
func ChiSquare(obs, exp []float64) float64 {
        if len(obs) != len(exp) {
                panic("stat: slice length mismatch")
        }
        var result float64
        for i, a := range obs {
                b := exp[i]
                if a == 0 && b == 0 {
                        continue
                }
                result += (a - b) * (a - b) / b
        }
        return result
}
```


##### gonum/stat/distub.ChiSquared 構造体

```go
//$ go doc -src distuv.ChiSquared
package distuv // import "gonum.org/v1/gonum/stat/distuv"

// ChiSquared implements the χ² distribution, a one parameter distribution
// with support on the positive numbers.
//
// The density function is given by
//  1/(2^{k/2} * Γ(k/2)) * x^{k/2 - 1} * e^{-x/2}
// It is a special case of the Gamma distribution, Γ(k/2, 1/2).
//
// For more information, see https://en.wikipedia.org/wiki/Chi-squared_distribution.
type ChiSquared struct {
        // K is the shape parameter, corresponding to the degrees of freedom. Must
        // be greater than 0.
        K float64

        Src rand.Source
}

func (c ChiSquared) CDF(x float64) float64
func (c ChiSquared) ExKurtosis() float64
func (c ChiSquared) LogProb(x float64) float64
func (c ChiSquared) Mean() float64
func (c ChiSquared) Mode() float64
func (c ChiSquared) NumParameters() int
func (c ChiSquared) Prob(x float64) float64
func (c ChiSquared) Quantile(p float64) float64
func (c ChiSquared) Rand() float64
func (c ChiSquared) StdDev() float64
func (c ChiSquared) Survival(x float64) float64
func (c ChiSquared) Variance() float64
```


##### (ChiSquared) Quantile() メソッド

```go
// Quantile returns the inverse of the cumulative distribution function.
func (c ChiSquared) Quantile(p float64) float64 {
        if p < 0 || p > 1 {
                panic(badPercentile)
        }
        return mathext.GammaIncRegInv(0.5*c.K, p) * 2
}
```


##### gonum/mathext.GammaIncRegInv() 関数


```go
//$ go doc -src mathext.GammaIncRegInv
package mathext // import "gonum.org/v1/gonum/mathext"

// GammaIncRegInv computes the inverse of the regularized incomplete Gamma integral. That is,
// it returns the x such that:
//  GammaIncReg(a, x) = y
// The input argument a must be positive and y must be between 0 and 1
// inclusive or GammaIncRegInv will panic. GammaIncRegInv should return a positive
// number, but can return NaN if there is a failure to converge.
func GammaIncRegInv(a, y float64) float64 {
        return gammaIncRegInv(a, y)
}
```


##### mathext.gammaIncRegInv 関数


```go
// gammaIncRegInv is the inverse of the regularized incomplete Gamma integral. That is, it
// returns x such that:
//  Igam(a, x) = y
// The input argument a must be positive and y must be between 0 and 1
// inclusive or gammaIncRegInv will panic. gammaIncRegInv should return a
// positive number, but can return NaN if there is a failure to converge.
func gammaIncRegInv(a, y float64) float64 {
	// For y not small, we just use
	//  IgamI(a, 1-y)
	// (inverse of the complemented incomplete Gamma integral). For y small,
	// however, 1-y is about 1, and we lose digits.
	if a <= 0 || y <= 0 || y >= 0.25 {
		return cephes.IgamI(a, 1-y)
	}

	lo := 0.0
	flo := -y
	hi := cephes.IgamI(a, 0.75)
	fhi := 0.25 - y

	params := []float64{a, y}

	// Also, after we generate a small interval by bisection above, false
	// position will do a large step from an interval of width ~1e-4 to ~1e-14
	// in one step (a=10, x=0.05, but similar for other values).
	result, bestX, _, errEst := falsePosition(lo, hi, flo, fhi, 2*machEp, 2*machEp, 1e-2*a, gammaIncReg, params)
	if result == fSolveMaxIterations && errEst > allowedATol+allowedRTol*math.Abs(bestX) {
		bestX = math.NaN()
	}

	return bestX
}
```


##### cephes.IgamI 関数


```go
// IgamI computes the inverse of the incomplete Gamma function. That is, it
// returns the x such that:
//  IgamC(a, x) = p
// The input argument a must be positive and p must be between 0 and 1
// inclusive or IgamI will panic. IgamI should return a positive number, but
// can return 0 even with non-zero y due to underflow.
func IgamI(a, p float64) float64 {
	// Bound the solution
	x0 := math.MaxFloat64
	yl := 0.0
	x1 := 0.0
	yh := 1.0
	dithresh := 5.0 * machEp

	if p < 0 || p > 1 || a <= 0 {
		panic(paramOutOfBounds)
	}

	if p == 0 {
		return math.Inf(1)
	}

	if p == 1 {
		return 0.0
	}

	// Starting with the approximate value
	//  x = a y^3
	// where
	//  y = 1 - d - ndtri(p) sqrt(d)
	// and
	//  d = 1/9a
	// the routine performs up to 10 Newton iterations to find the root of
	//  IgamC(a, x) - p = 0
	d := 1.0 / (9.0 * a)
	y := 1.0 - d - Ndtri(p)*math.Sqrt(d)
	x := a * y * y * y

	lgm := lgam(a)

	for i := 0; i < 10; i++ {
		if x > x0 || x < x1 {
			break
		}
```


##### cephes.IgamC 関数


```go
// IgamC computes the complemented incomplete Gamma integral.
//  IgamC(a,x) = 1 - Igam(a,x)
//             = (1/ Γ(a)) \int_0^\infty e^{-t} t^{a-1} dt
// The input argument a must be positive and x must be non-negative or
// IgamC will panic.
func IgamC(a, x float64) float64 {
	// The integral is evaluated by either a power series or continued fraction
	// expansion, depending on the relative values of a and x.
	// Sources:
	// [1] "The Digital Library of Mathematical Functions", dlmf.nist.gov
	// [2] Maddock et. al., "Incomplete Gamma Functions",
	// http://www.boost.org/doc/libs/1_61_0/libs/math/doc/html/math_toolkit/sf_gamma/igamma.html

	switch {
	case x < 0, a <= 0:
		panic(paramOutOfBounds)
	case x == 0:
		return 1
	case math.IsInf(x, 0):
		return 0
	}

	// Asymptotic regime where a ~ x; see [2].
	absxmaA := math.Abs(x-a) / a
	if (igamSmall < a && a < igamLarge && absxmaA < igamSmallRatio) ||
		(igamLarge < a && absxmaA < igamLargeRatio/math.Sqrt(a)) {
		return asymptoticSeries(a, x, igamC)
	}

	// Everywhere else; see [2].
	if x > 1.1 {
		if x < a {
			return 1 - igamSeries(a, x)
		}
		return igamCContinuedFraction(a, x)
	} else if x <= 0.5 {
		if -0.4/math.Log(x) < a {
			return 1 - igamSeries(a, x)
		}
		return igamCSeries(a, x)
	}

	if x*1.1 < a {
		return 1 - igamSeries(a, x)
	}
	return igamCSeries(a, x)
}
```


#### Calculating p-values
- senario
    - `vx <- c(260, 135, 105)`
    - `prob <- c(0.60, 0.25, 0.16)`
- H0: The deviations from the previously observed percnetages are due to pure chance
- Ha: THe deviations are due to some underlying effect outside of pure chance


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// Define observed and expected values.
	observed := []float64{
		260.0,
		135.0,
		105.0,
	}

	totalObserved := sum(observed)

	expected := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)

	// Calculate the degree of freedom
	df := len(observed) - 1
	totalA := observed[0] + expected[0]
	totalB := observed[1] + expected[1]
	totalC := observed[2] + expected[2]
	totalObs := observed[0] + observed[1] + observed[2]
	totalExp := expected[0] + expected[1] + expected[2]
	total := totalObs + totalExp

	// Create a Chi-Squared distribution
	dist := distuv.ChiSquared{
		K: float64(df),
	}

	// Calculate Chi-Squate(df=1, 0.05)
	// Quantile returns the inverse of the cumulative distribution function.
	// https://pkg.go.dev/github.com/ArkaGPL/gonum@v0.8.5/stat/distuv#ChiSquared.Quantile
	quant := dist.Quantile(0.95)

	pvalue := 1 - dist.CDF(chiSquare)
	prob := dist.Prob(chiSquare)

	fmt.Println(">>   Pearson's chi-squared test")
	fmt.Println("|       | Observed | Expected |  Total |")
	fmt.Println("|-------|----------|----------|--------|")
	fmt.Printf("|   A   | %8.2f | %8.2f | %6.1f |\n", observed[0], expected[0], totalA)
	fmt.Printf("|   B   | %8.2f | %8.2f | %6.1f |\n", observed[1], expected[1], totalB)
	fmt.Printf("|   C   | %8.2f | %8.2f | %6.1f |\n", observed[2], expected[2], totalC)
	fmt.Printf("| total | %8.2f | %8.2f | %6.1f |\n", totalObs, totalExp, total)
	fmt.Printf("Chi-Square statistic: %.4f\n", chiSquare)
	fmt.Printf("degree of freedom: %v\n", df)

	fmt.Printf("Chi-Squre(%v,0.05): %.4f\n", df, quant)
	fmt.Printf("p-value: %.4f\n", pvalue)
	fmt.Printf("p-value: %.4f\n", prob)

	/*
		// output from R
		> vx <- c(260, 135, 105)
		> chisq.test(x=vx, p=c(0.60,0.25,0.15))

		        Chi-squared test for given probabilities

		data:  vx
		X-squared = 18.133, df = 2, p-value = 0.0001155
	*/

	/*
		>>   Pearson's chi-squared test
		|       | Observed | Expected |  Total |
		|-------|----------|----------|--------|
		|   A   |   260.00 |   300.00 |  560.0 |
		|   B   |   135.00 |   125.00 |  260.0 |
		|   C   |   105.00 |    75.00 |  180.0 |
		| total |   500.00 |   500.00 | 1000.0 |
		Chi-Square statistic: 18.1333
		degree of freedom: 2
		Chi-Squre(2,0.05): 5.9915
		p-value: 0.0001
	*/

}

func sum(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum
}
```


