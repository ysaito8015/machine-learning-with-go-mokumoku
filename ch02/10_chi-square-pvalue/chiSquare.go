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
