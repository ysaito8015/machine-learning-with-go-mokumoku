package main

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/gonum/stat"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// CSV ファイルを開く
	f, err := os.Open("./continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV ファイルを読むための CSV Reader を作る
	reader := csv.NewReader(f)

	// 観測値と推測値を格納するスライスを宣言する
	var observed []float64
	var predicted []float64

	// 行数のカウントを格納する変数
	line := 1

	// 予期しないデータタイプがないかカラムの内容をチェックする
	for {

		// ファイルの終わりに来たらループを抜ける
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// ヘッダ行を飛ばす
		if line == 1 {
			line++
			continue
		}

		// 観測値と予測値を読み出す
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// それぞれのスライスへ追加する
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	} // データ読出しおわり

	// MSE, MAE を計算する
	var mAE float64
	var mSE float64
	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal-predicted[idx], 2) / float64(len(observed))
	}

	// 値を出力する
	fmt.Printf("MAE = %0.2f\n", mAE)
	// Output: MAE = 2.55
	fmt.Printf("MSE = %0.2f\n", mSE)
	// Output: MSE = 10.51

	// 決定係数を計算する
	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	// 値を出力する
	fmt.Printf("R^2 = %0.2f\n", rSquared)
	// Output: R^2 = 0.37
}
