package main

import (
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

	// 目的変数の列を抽出する
	yVals := advertDF.Col("Sales").Float()

	// データセットの各変数との散布図を生成する
	for _, colName := range advertDF.Names() {
		// pts 変数は, プロット用の値を保持する
		pts := make(plotter.XYs, advertDF.Nrow())

		// pts 変数にデータを埋める
		for i, floatVal := range advertDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		// plot を生成する
		p := plot.New()
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Radius = vg.Points(3)

		// plot を PNG ファイルへ保存する
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}

	}
}
