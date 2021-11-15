package main

import (
	"bufio"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	// CSV ファイルを開く
	f, err := os.Open("./Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV からデータフレームを生成する
	advertDF := dataframe.ReadCSV(f)

	// 学習セット、テストセットそれぞれの要素数を計算する
	trainingNum := (4 * advertDF.Nrow()) / 5
	testNum := advertDF.Nrow() / 5
	if trainingNum+testNum < advertDF.Nrow() {
		trainingNum++
	}

	// 各セットのインデクスを生成する
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// 学習セットのインデクスをもとにデータを挿入する
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// テストセットのインデクスをもとにデータを挿入する
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// 学習セット, テストセットのデータフレームを生成する
	trainingDF := advertDF.Subset(trainingIdx)
	testDF := advertDF.Subset(testIdx)

	// データをファイルに書き出す際に使用する map を生成する
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// 学習セット, テストセットのファイルを生成する
	for idx, setName := range []string{"training.csv", "test.csv"} {

		// データセットを保存するファイルを生成する
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// buffered writer を生成
		w := bufio.NewWriter(f)

		// データフレームを CSV として書き込む
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}
	}
}
