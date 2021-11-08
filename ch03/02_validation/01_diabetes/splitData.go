package main

import (
	"bufio"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	//糖尿病のデータセットファイルを開く
	f, err := os.Open("./diabetes.csv")
	if err != nil {
		log.Fatal(err)
	}

	// CSV ファイルからデータフレームを作成する
	// 列の型は推論される
	diabetesDF := dataframe.ReadCSV(f)

	// 学習セット, テストセットの要素の数を数える
	trainingNum := (4 * diabetesDF.Nrow()) / 5
	testNum := diabetesDF.Nrow() / 5
	if trainingNum+testNum < diabetesDF.Nrow() {
		trainingNum++
	}

	// サブセットのインデックスを生成する
	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	// 学習データのインデックスを列挙する
	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// テストデータのインデックスを列挙する
	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// 学習データとテストデータのサブセットを生成する
	trainingDF := diabetesDF.Subset(trainingIdx)
	testDF := diabetesDF.Subset(testIdx)

	// それぞれのデータセットをファイルに書き込むためのマップを宣言する
	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	// それぞれのファイルを生成する
	for idx, setName := range []string{"training.csv", "test.csv"} {
		// それぞれのデータセットファイルを生成
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// バッファ付きライタを生成
		w := bufio.NewWriter(f)
		// CSV としてデータフレームの値を書き出す
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}

		// log
		log.Printf("dump %v", setName)
	}
}
