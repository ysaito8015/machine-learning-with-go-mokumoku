package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	// CSV ファイルを開く
	advertFile, err := os.Open("./Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer advertFile.Close()

	// CSV ファイルからデータフレームを作成する
	advertDF := dataframe.ReadCSV(advertFile)

	// Describe メソッドを使用して, サマリを作る
	advertSummary := advertDF.Describe()

	// 出力する
	fmt.Println(advertSummary)
	// Output:
	/*
		[8x5] DataFrame

		    column   TV         Radio     Newspaper  Sales
		 0: mean     147.042500 23.264000 30.554000  14.022500
		 1: median   149.750000 22.900000 25.750000  12.900000
		 2: stddev   85.854236  14.846809 21.778621  5.217457
		 3: min      0.700000   0.000000  0.300000   1.600000
		 4: 25%      73.400000  9.900000  12.600000  10.300000
		 5: 50%      149.700000 22.500000 25.600000  12.900000
		 6: 75%      218.500000 36.500000 45.100000  17.400000
		 7: max      296.400000 49.600000 114.000000 27.000000
		    <string> <float>    <float>   <float>    <float>
	*/

}
