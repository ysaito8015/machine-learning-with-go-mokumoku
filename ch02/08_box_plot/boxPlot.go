package main

import (
	//"fmt"
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
