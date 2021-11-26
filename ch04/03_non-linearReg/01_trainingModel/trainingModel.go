package main

import (
	"encoding/csv"
	"fmt"
	//"github.com/adam-hanna/ridge"
	//"gonum.org/v1/gonum/mat"
	"github.com/berkmancenter/ridge"
	"github.com/gonum/matrix/mat64"
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
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// featureData スライスは, gonum での行列計算用に
	// すべての float value を保持する
	featureData := make([]float64, 4*len(rawCSVData))
	yData := make([]float64, len(rawCSVData))

	// featureIndex と yIndex は, 行列の現在示すインデクスを保持する
	var featureIndex int
	var yIndex int

	// 学習データをスライスに対応させるためにループさせる
	for i, record := range rawCSVData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// 列全体をループさせる
		for i, val := range record {

			// csv の値を float に変換
			valParsed, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}

			// Sales 以外の列の時
			if i < 3 {

				// 切片をモデルに加える
				if i == 0 {
					featureData[featureIndex] = 1
					featureIndex++
				}

				// 対象のスライスに値を入れる
				featureData[featureIndex] = valParsed
				featureIndex++
			}

			if i == 3 {

				// 目的変数のスライスに値を入れる
				yData[yIndex] = valParsed
				yIndex++
			}
		}
	}

	// gonum matrix を生成する
	features := mat64.NewDense(len(rawCSVData), 4, featureData)
	y := mat64.NewVector(len(rawCSVData), yData)

	// RidgeRegression 構造体を New 関数で生成する
	r := ridge.New(features, y, 1.0)

	// 学習する
	r.Regress()

	// 出力する
	c1 := r.Coefficients.At(0, 0)
	c2 := r.Coefficients.At(1, 0)
	c3 := r.Coefficients.At(2, 0)
	c4 := r.Coefficients.At(3, 0)
	fmt.Printf("Regression formula: \t y = %0.3f + %0.3f TV + %0.3f Radio + %0.3f Newspaper\n", c1, c2, c3, c4)
	// Output:
	// Regression formula:      y = 3.038 + 0.047 TV + 0.177 Radio + 0.001 Newspaper
}
