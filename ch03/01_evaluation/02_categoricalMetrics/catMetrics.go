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

	// 真陽性と真陰性を格納する変数
	var truePosNeg int

	// 真陽性と真陰性の値を集める
	for idx, oVal := range observed {
		if oVal == predicted[idx] {
			truePosNeg++
		}
	}

	// Accuracy を計算する
	accuracy := float64(truePosNeg) / float64(len(observed))

	// 出力する
	fmt.Printf("Accuracy = %0.2f\n", accuracy)
	// Output: Accuracy = 0.97
}
