package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// Define observed and expected values.
	observed := []float64{317, 168, 230, 85}
	expected := []float64{296, 176, 256, 72}

	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)

	// Calculate the degree of freedom
	df := len(observed) - 1

	// Chi-squared statistic
	totalA := observed[0] + expected[0]
	totalB := observed[1] + expected[1]
	totalO := observed[2] + expected[2]
	totalAB := observed[3] + expected[3]
	totalObs := sum(observed)
	totalExp := sum(expected)
	total := totalA + totalB + totalO + totalAB

	subAsq := (observed[0] - expected[0]) * (observed[0] - expected[0])
	subBsq := (observed[1] - expected[1]) * (observed[1] - expected[1])
	subOsq := (observed[2] - expected[2]) * (observed[2] - expected[2])
	subABsq := (observed[3] - expected[3]) * (observed[3] - expected[3])
	chiSq := subAsq/expected[0] + subBsq/expected[1] + subOsq/expected[2] + subABsq/expected[3]

	// Create a Chi-Squared distribution
	dist := distuv.ChiSquared{
		K: float64(df),
	}

	pvalue := 1 - dist.CDF(chiSquare)

	// Calculate Chi-Squate(df=3, 0.05)
	// Quantile returns the inverse of the cumulative distribution function.
	quant := dist.Quantile(0.95)

	fmt.Println(">>   Pearson's chi-squared test")
	fmt.Println("|       | Observed | Expected |  Total |")
	fmt.Println("|-------|----------|----------|--------|")
	fmt.Printf("|   A   | %8.2f | %8.2f | %6.1f |\n", observed[0], expected[0], totalA)
	fmt.Printf("|   B   | %8.2f | %8.2f | %6.1f |\n", observed[1], expected[1], totalB)
	fmt.Printf("|   O   | %8.2f | %8.2f | %6.1f |\n", observed[2], expected[2], totalO)
	fmt.Printf("|  AB   | %8.2f | %8.2f | %6.1f |\n", observed[3], expected[3], totalAB)
	fmt.Printf("| total | %8.2f | %8.2f | %6.1f |\n", totalObs, totalExp, total)
	fmt.Printf("Chi-Square statistic: %.4f\n", chiSquare)
	fmt.Printf("Chi-Square statistic: %.4f\n", chiSq)
	fmt.Printf("degree of freedom: %v\n", df)
	fmt.Printf("Chi-Squre(%v,0.05): %.4f\n", df, quant)
	fmt.Printf("p-value: %.4f\n", pvalue)

	/*
	  // output from R
	  > vx <- c(317,168,230,85)
	  > chisq.test(x=vx, p=c(0.37,0.22,0.32,0.09))

	          Chi-squared test for given probabilities

	  data:  vx
	  X-squared = 6.8413, df = 3, p-value = 0.07713

	*/
	/*
		>>   Pearson's chi-squared test
		|       | Observed | Expected |  Total |
		|-------|----------|----------|--------|
		|   A   |   317.00 |   296.00 |  613.0 |
		|   B   |   168.00 |   176.00 |  344.0 |
		|   O   |   230.00 |   256.00 |  486.0 |
		|  AB   |    85.00 |    72.00 |  157.0 |
		| total |   800.00 |   800.00 | 1600.0 |
		Chi-Square statistic: 6.8413
		Chi-Square statistic: 6.8413
		degree of freedom: 3
		Chi-Squre(3,0.05): 7.8147
		p-value: 0.0771
	*/
}

func sum(slice []float64) float64 {
	sum := 0.0
	for _, v := range slice {
		sum += v
	}
	return sum
}
