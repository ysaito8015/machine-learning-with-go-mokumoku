package main

import (
	"fmt"
	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/stat"
)

func main() {
	// スコアとクラスを定義
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}

	// Recall, 感度と FPR, 偽陽性率を計算する
	tpr, fpr, thresh := stat.ROC(nil, scores, classes, nil)

	// AUC を計算する
	auc := integrate.Trapezoidal(fpr, tpr)

	// 出力する
	fmt.Printf("true positive rate: %v\n", tpr)
	// Output:
	// true positive rate: [0 0 0.5 0.5 1]
	fmt.Printf("false positive rate: %v\n", fpr)
	// Output:
	// false positive rate: [0 0.5 0.5 1 1]
	fmt.Printf("Thresholds: %v\n", thresh)
	// Output:
	// Thresholds: [+Inf 0.8 0.4 0.35 0.1]
	fmt.Printf("AUC: %v\n", auc)
	// Output:
	// AUC: 0.25

}
