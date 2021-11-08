package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// CSV ファイルを開く
	f, err := os.Open("./labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV リーダーを宣言する
	reader := csv.NewReader(f)

	// 観測値と予測値を格納するスライスを宣言する
	var observed []int
	var predicted []int

	// 行数を格納する変数
	line := 1

	// 期待している以外の型が行内にあるかをチェックする
	for {

		// 行を読む. ファイルの終わりかどうかをチェックする.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// ヘッダ行を飛ばす
		if line == 1 {
			line++
			continue
		}

		// 観測値を予測値を読み込む
		observedVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// 格納するスライスへ追加する
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++

	}

	// classes 変数は, ラベルデータ内にありえる３つの選択肢をスライス要素として格納する
	classes := []int{0, 1, 2}

	// 各クラスごとにループさせる
	for _, class := range classes {

		// 真陽性, 偽陽性, 偽陰性を格納する変数
		var truePos int
		var falsePos int
		var falseNeg int

		for idx, oVal := range observed {

			// switch 文で, クラス値をチェックする
			switch oVal {

			// 観測値が, クラス値と等しい場合, 真陽性のカウントをアップする
			case class:
				if predicted[idx] == class {
					truePos++
					continue
				}
				falseNeg++

			// 観測値が, クラス値と違う場合, 偽陽性化チェックする
			default:
				if predicted[idx] == class {
					falsePos++
				}
			}
		}

		// precision を計算する
		precision := float64(truePos) / float64(truePos+falsePos)

		// recall を計算する
		recall := float64(truePos) / float64(truePos+falseNeg)

		// 出力する
		fmt.Printf("Precision (class %d) = %0.2f\n", class, precision)
		fmt.Printf("Recall (class %d) = %0.2f\n", class, recall)
	}
}
