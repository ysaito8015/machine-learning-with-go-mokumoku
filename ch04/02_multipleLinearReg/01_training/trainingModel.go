package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"log"
	"os"
	"strconv"
)

func main() {
	// 学習データセットを開く
	f, err := os.Open("./training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader を生成する
	reader := csv.NewReader(f)

	// CSV レコードをすべて読み込む
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 目的変数 Sales (y) を説明変数1 TV (x1) と
	// 説明変数2 Radio (x2), 切片でモデル化する
	// sajari/regression を使うモデルに必要な構造体を生成する
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")

	// 学習データを線形モデルの値に対応させるためにループさせる
	for i, record := range trainingData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// 目的変数を Sales にパースする
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 説明変数 x1を TV にパースする
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 説明変数 x2 を Radio にパースする
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 目的変数と説明変数を, 線形モデルに対応させる
		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))

	}

	// 学習を実行する
	r.Run()

	// 学習したモデルパラメタを出力する
	fmt.Printf("Regression Formula: %v\n", r.Formula)
	// Output: Regression Formula: Predicted = 2.9318 + TV*0.0473 + Radio*0.1794
}
