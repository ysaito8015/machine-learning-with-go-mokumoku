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

	pvalue := 1 - dist.CDF(chiSquare)
	prob := dist.Prob(chiSquare)

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
	fmt.Printf("p-value: %.4f\n", pvalue)
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
		p-value: 0.6892
	*/
}
