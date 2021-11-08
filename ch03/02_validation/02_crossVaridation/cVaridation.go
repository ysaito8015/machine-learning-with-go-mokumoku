package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/filters"
	"github.com/sjwhitworth/golearn/trees"
	"log"
	"math"
)

func main() {

	iris, err := base.ParseCSVToInstances("./iris_headers.csv", true)
	if err != nil {
		log.Fatal(err)
	}

	// iris データセットを Chi-Merge 法で離散化する
	filt := filters.NewChiMergeFilter(iris, 0.999)
	for _, a := range base.NonClassFloatAttributes(iris) {
		filt.AddAttribute(a)
	}
	filt.Train()
	irisf := base.NewLazilyFilteredInstances(iris, filt)

	//
	// ID3 アルゴリズムを使用する
	//
	// デシジョンツリーモデルを定義する
	// 剪定 pruning パラメタは 0.6
	var tree base.Classifier
	param := 0.6
	tree = trees.NewID3DecisionTree(param)
	// (Parameter controls train-prune split.)

	// クロスバリデーションを k = 5 で実行する
	cfs, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(irisf, tree, 5)
	if err != nil {
		log.Fatal(err)
	}

	// メトリックスを計算する
	mean, variance := evaluation.GetCrossValidatedMetric(cfs, evaluation.GetAccuracy)
	stdev := math.Sqrt(variance)

	// 出力する
	fmt.Println("ID3 Performance (information gain)")
	fmt.Printf("Pruning: %0.2f\t\tAvg. accuracy: %.2f (+/-2SD: %.2f)\n", param, mean, stdev*2)
	// Output:
	// Pruning: 0.60           Avg. accuracy: 0.72 (+/-2SD: 0.19)
	for i, _ := range cfs {
		fmt.Println(evaluation.GetSummary(cfs[i]))
	}
	// Output:
	/*
	   ID3 Performance (information gain)
	   Reference Class True Positives  False Positives True Negatives  Precision       Recall  F1 Score
	   --------------- --------------  --------------- --------------  ---------       ------  --------
	   Iris-setosa     9               6               13              0.6000          1.0000  0.7500
	   Iris-versicolor 2               4               16              0.3333          0.2500  0.2857
	   Iris-virginica  7               0               17              1.0000          0.6364  0.7778
	   Overall accuracy: 0.6429

	   Reference Class True Positives  False Positives True Negatives  Precision       Recall  F1 Score
	   --------------- --------------  --------------- --------------  ---------       ------  --------
	   Iris-setosa     12              8               16              0.6000          1.0000  0.7500
	   Iris-versicolor 0               0               24              NaN             0.0000  NaN
	   Iris-virginica  12              4               20              0.7500          1.0000  0.8571
	   Overall accuracy: 0.6667

	   Reference Class True Positives  False Positives True Negatives  Precision       Recall  F1 Score
	   --------------- --------------  --------------- --------------  ---------       ------  --------
	   Iris-setosa     10              7               15              0.5882          1.0000  0.7407
	   Iris-versicolor 0               0               23              NaN             0.0000  NaN
	   Iris-virginica  13              2               17              0.8667          1.0000  0.9286
	   Overall accuracy: 0.7188

	   Reference Class True Positives  False Positives True Negatives  Precision       Recall  F1 Score
	   --------------- --------------  --------------- --------------  ---------       ------  --------
	   Iris-setosa     11              9               14              0.5500          1.0000  0.7097
	   Iris-versicolor 5               0               17              1.0000          0.2941  0.4545
	   Iris-virginica  6               3               25              0.6667          1.0000  0.8000
	   Overall accuracy: 0.6471

	   Reference Class True Positives  False Positives True Negatives  Precision       Recall  F1 Score
	   --------------- --------------  --------------- --------------  ---------       ------  --------
	   Iris-setosa     8               1               11              0.8889          1.0000  0.9412
	   Iris-versicolor 2               0               16              1.0000          0.5000  0.6667
	   Iris-virginica  8               1               11              0.8889          1.0000  0.9412
	   Overall accuracy: 0.9000
	*/
}
