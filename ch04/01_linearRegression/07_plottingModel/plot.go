package main

import (
	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
)

func predict(tv float64) float64 {
	return 7.0688 + tv*0.0489
}

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

	// pts は plot の値を保持する
	pts := make(plotter.XYs, advertDF.Nrow())

	// ptsPred は, plot 用の予測値を保持する
	ptsPred := make(plotter.XYs, advertDF.Nrow())

	// pts 変数に値を埋め込む
	for i, floatVal := range advertDF.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}

	// plot を生成する
	p := plot.New()
	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	// 観測値の散布図プロットを追加する
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Radius = vg.Points(3)

	// 予測値のためラインプロットを追加する
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	// plot を PNG ファイルに保存する
	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "regression_line.png"); err != nil {
		log.Fatal(err)
	}
}
