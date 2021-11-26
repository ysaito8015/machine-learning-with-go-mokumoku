package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// ridge regression model の学習結果を利用する
func predict(tv, radio, newspaper float64) float64 {
	return 3.038 + 0.047*tv + 0.177*radio + 0.001*newspaper
}

func main() {
	// テストデータセットを開く
	f, err := os.Open("./test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader を生成する
	reader := csv.NewReader(f)

	// CSV レコードをすべて読み込む
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// holdout data をループさせる. y の値を予測させて,
	// MAE, Mean Absolute Error の値を計算する
	var mAE float64
	for i, record := range testData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// Sales をパースする
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// TV をパースする
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Radio をパースする
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Newspaper をパースする
		newspaperVal, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 学習済みモデルから y を予測する
		yPredicted := predict(tvVal, radioVal, newspaperVal)

		// 残差を MAE に加算する
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// MAE を出力する
	fmt.Printf("MAE = %0.2f\n", mAE)
	// Output:
	// MAE = 1.26
}
