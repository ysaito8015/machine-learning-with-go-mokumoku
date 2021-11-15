package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	// CSV からデータフレームを作成する
	advertDF := dataframe.ReadCSV(f)

	// データセットの各列から, ヒストグラムを生成する
	for _, colName := range advertDF.Names() {

		// plotter.Values の値を生成し, データフレームの列から
		// 値を埋める
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// プロットを作成し, タイトルを設定する
		p := plot.New() // err 値は返さなくなった
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// ヒストグラムを生成する
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// ヒストグラムを標準化する
		h.Normalize(1)

		// プロットスペースに, ヒストグラムを追加する
		p.Add(h)

		// プロットを PNG ファイルに保存する
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
