# Chapter 04 Regression
- Regression is a process through which we can understand how one variable changes with respect to another variable.
    - regression is a process to analyze a relationship between one variable and snother variable
- 連続値に用いられることが多い
- 離散値に用いると Classification となる


## Understanding regression model jargon
- Response or dependent variable
    - 目的変数
    - 予想しようとしている変数
    - y とラベル付けられることが多い
- Explanatory variables, independent variables, features, attributes, or regressors
    - 説明変数
    - Response を予測するために用いる変数
    - x や x_1, x_2 などとラベル付けられることが多い
- Linear regression
    - this means that the dependent variable depends on the independent variable linearly.
    - 目的変数が, 説明変数に線形に依存している
- Nonlinear regression
    - the dependent variable depends on the independent variable in a relationship non-linearly.
    - 目的変数が, 説明変数に非線形に依存している
- Multiple regression
    - a regression with more than one independent variable
    - 一つ以上の変数に依存している線形モデル
- Fitting or training
    - 学習過程
    - 目的変数を予測するためにモデルをパラメタライジングする過程
    - the process of parameterizing a model so that it can predict a certain dependent variable.
- Prediction
    - 予測
    - 目的変数を予測するためにパラメタライズドモデルを使う過程
    - the process of using a parameterized model to predict a certain dependent variable.


## Linear regression
- 避けては通れない最も重要かつ基本的なモデル
- integrity in machine learning applcations is crucial, and the simpler and more interpretable a model is, the easier it is to maintain integrity.
- 機械学習アプリケーションにおいて整合性は非常に重要であり, モデルが単純で解釈しやすいほど、整合性を維持しやすくなる
- Mike Lee Williams from Fast Forward Labs
    - "The future is algorithmic. Interpretable models offer a safer, more productive, and ultimately more collaborative relationship between humans and intelligent machines."
- 線形モデルは、解釈可能であり, それ故安全で, 生産的な選択肢を提供する
- 連続値の予測を行いたい場合は, 検討し一度は試してみるべき


### Overview of linear regression
- attempt to model our dependent variable, y, by an independent variable, x
- m : slope of the line
- b : the intercept
- ![equation](https://latex.codecogs.com/gif.latex?y%20%3D%20mx%20&plus;%20b)
- e.g., sales and number of users
    - `sales = m * (number of users) + b`
    - ![](https://i.gyazo.com/cefe5fa0d90e4b4b20764b7351266dc7.png)


#### determining the values of m and b
- Ordinary Least Squares (OLS)
    - 最小二乗法
    - 各実測点から, 予測点を引いた値を残差 (residuals), エラー (errors) という
- ![](https://i.gyazo.com/936f1e2a21a54a299b0f20c57a965bf5.png)
    - calculate the sum of the squares of these residuals
        - Residual sum of squares
        - 残差平方和
    - ![equation](https://latex.codecogs.com/gif.latex?RSS%20%3D%20%5Csum%5En_%7Bi%3D1%7D%20%28y%20-%20%5Chat%7By%7D%29)
        - y hat : y の予測点
        - この残差平方和 RSS を最小化するように予測の線を引く
- RSS を最小化する一般的な最適化手法は, gradient descent
    - 再急降下法
    - 付録で説明される
        - Appendex, Algorithms/Techniques Related to Machine Learning


### Linear regression assumptions and pitfalls
- linear regression の得意な領域, 不得意な領域

#### Linear regression の想定する過程 
- Linear relationship
    - 線形関係
        - 目的変数が, 説明変数に線形に依存している
- Normality
    - this assumption means that your variables should be distributed according to a normal distribution
    - 扱う変数が, 正規分布に従っている
    - Wikipedia の記述では, 必須のではなくオプション扱いの仮定
- No multicollinearity
    - Multicollinearity, 多重共撰性
    - 説明変数が実際には説明変数足りえず, 目的変数とともに第三の変数に依存している
- No auto-correlation
    - Auto-correlation, 自己相関
    - 変数が, 時間的にシフトした変数それ自身と一致している
    - 変数のそれ自身との相互相関
- Homoscedasticity
    - 等分散性
    - residual の分散が, 一定であり, independent variable の値に依存しない
    - the variance of your data is about the same around the regression line for all values of your independent variable


#### Linear regression の落とし穴
- 学習データ範囲以外のデータの予測には注意を払うこと
    - なぜなら, regression line は適用できないことが多い
- 実際には関連のない 2 つの変数感の疑似関係を見つけることで, linear regression モデルを誤って適用する
    - 変数が関連している可能性について, 論理的な理由を検証する
- データの外れ値, または, 極端な値は, 学習に影響を与える
    - 外れ値に強いアルゴリズムなどを検討する


### Linear regression example
- メディアへの広告費用と, 売上の CSV
- モデルのゴールは, 広告費から売上を予測する

```shell
$ head ./Advertising.csv
TV,Radio,Newspaper,Sales
230.1,37.8,69.2,22.1
44.5,39.3,45.1,10.4
17.2,45.9,69.3,9.3
151.5,41.3,58.5,18.5
180.8,10.8,58.4,12.9
8.7,48.9,75,7.2
57.5,32.8,23.5,11.8
120.2,19.6,11.6,13.2
8.6,2.1,1,4.8

```

#### Profiling the data
- いずれの機械学習モデル作成も, data profiling から始まる
    - 各データ行の, 分散, Range, variability (変化) に対する理解を得る
- Ch.02 の記述統計を適用する


```go
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
```


##### Plotting data

```go
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
```


- ![Newspaper](https://i.gyazo.com/50df2536d7f3be7d77f9d4a891f0bc26.png)
- ![TV](https://i.gyazo.com/e506f28dd73729227c7b960ff4fdd1e3.png)
- ![Radio](https://i.gyazo.com/427cea45aee58e66901ea8bb1019da34.png)
- ![Sales](https://i.gyazo.com/5e41dc988bd3fb366e571bdceeca8670.png)
- 4 つの変数すべてが正規分布ではなく, Sales がややベル・カーブのように見える


##### Normal Distribution
- Quantile-quantile plot
    - 得られた分布がどの程度正規分布に近いかを決定する
- Statistical test
    - 変数が正規分布に近いかの確率を決定する


##### when not fit within the assumptions of linear regression model
1. 変数変換する
    - 線形モデルを当てはめられる
    - 解釈性が落ちてしまう
2. 別のデータを取得する
3. 線形モデルの仮定を無視する
    - 著者はこの選択肢をまず試してみたら、と


#### Choosing our independent variable
- 説明変数を選ぶ
- 簡単な方法
    - 視覚的に目的変数との相関があるかで選ぶ
    - scatter plot を各変数と, 目的変数の間で作成する


##### Plotting

```go
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
```


- ![Sales](https://i.gyazo.com/cca66e6aebd2a2321b5cdcdcadf9bc50.png)
- ![TV](https://i.gyazo.com/127a6e966d105dd9224cc32fc63ef1de.png)
- ![Radio](https://i.gyazo.com/2c312af7ea1320d3d95202aec7d1a57b.png)
- ![Newspaper](https://i.gyazo.com/225a97209601b492faaa9f126b508d27.png)


##### Deduce relationships
- Radio と TV が Sales と線形関係が見られそう. TV を説明変数に採用してみる
- ![equation](https://latex.codecogs.com/gif.latex?Sales%20%3D%20m%20%5Ccdot%20TV%20&plus;%20b)
- TV は, おそらく等分散性ではない
    - モデルの予測性能が低い場合は, ここに立ち戻って可能な説明を考える


#### Creating our training and test sets
- 学習用とテスト用のデータセットを作成する
- holdout set は作成しない
    - 学習とテストを繰り返しせずに, モデルの学習を一度通してやって見るだけだから
- 80/20 split を使用する

##### Go code

```go
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
```

```shell
$ wc -l ./*.csv
  201 ./Advertising.csv
   41 ./test.csv
  161 ./training.csv
  403 total
```


- 上記で作成したデータは, データセットごとに並べ替えていないが, ランダム割付が本来は必要


#### Training our model
- regression package
	- https://github.com/sajari/regression


```go
package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"log"
	"os"
	"strconv"
)

func main() {
	// 学習データセットを開く
	f, err := os.Open("./training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader を生成する
	reader := csv.NewReader(f)

	// CSV レコードをすべて読み込む
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 目的変数 Sales (y) を説明変数 TV (x) と切片でモデル化する
	// sajari/regression を使うモデルに必要な構造体を生成する
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")

	// 学習データを線形モデルの値に対応させるためにループさせる
	for i, record := range trainingData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// 目的変数を Sales にパースする
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 説明変数を TV にパースする
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 目的変数と説明変数を, 線形モデルに対応させる
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))

	}

	// 学習を実行する
	r.Run()

	// 学習したモデルパラメタを出力する
	fmt.Printf("Regression Formula: %v\n", r.Formula)
	// Output: Regression Formula: Predicted = 7.0688 + TV*0.0489
}
```


- 傾き, slope は, 0.0489
- 切片, intercept は, 7.0688


#### Evaluating the trained model
- measure the performance of the model to see if we really have any power to predict Sales using TV as in independent variable
- 今回は, MAE, Mean Absolute Error を指標にする
    - 外れ値や, 極端な値を気にすることなく, モデルの予測値と, 観測値の Sales の値を比較できる


##### Go code

```go
package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// 学習データセットを開く
	f, err := os.Open("./training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader を生成する
	reader := csv.NewReader(f)

	// CSV レコードをすべて読み込む
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 目的変数 Sales (y) を説明変数 TV (x) と切片でモデル化する
	// sajari/regression を使うモデルに必要な構造体を生成する
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")

	// 学習データを線形モデルの値に対応させるためにループさせる
	for i, record := range trainingData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// 目的変数を Sales にパースする
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 説明変数を TV にパースする
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 目的変数と説明変数を, 線形モデルに対応させる
		r.Train(regression.DataPoint(yVal, []float64{tvVal}))

	}

	// 学習を実行する
	r.Run()

	// 学習したモデルパラメタを出力する
	fmt.Printf("Regression Formula: %v\n", r.Formula)
	// Output: Regression Formula: Predicted = 7.0688 + TV*0.0489

	// テストデータセットを開く
	f, err = os.Open("./test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader を生成する
	reader = csv.NewReader(f)

	// CSV レコードをすべて読み込む
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// テストデータから, 目的変数を予測し, 予測結果を, MAE で評価する
	var mAE float64
	for i, record := range testData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// テストデータの観測値 Sales をパースする
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// テストデータの説明変数 TV をパースする
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 学習済みモデルから, 目的変数を予測する
		yPredicted, err := r.Predict([]float64{tvVal})

		// MAE に加算する
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// MAE を出力する
	fmt.Printf("MAE = %0.2f\n", mAE)
	// Output: MAE = 3.01

}
```


##### MAE = 3.01 について
- Sales の Summary を思い出すと...
    - Mean = 14.02
    - s.d. = 5.21
- 今回の MAE は, s.d. より小さく, Mean の 21% 程度である
    - ある程度の予測精度がある


##### Plot the linear regression line
- 作成した線形モデルを関数にしておく


```go
func predict(tv float64) float64 {
    return 7.0688 + tv * 0.0489
}
```

- plotting code


```go
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
```


- ![regression_line](https://i.gyazo.com/81637728393fdc7ea188cabf1e1a4686.png)



## Multiple linear regression
- 複数の説明変数を持つ
- ![equation](https://latex.codecogs.com/gif.latex?y%20%3D%20m_1x_1%20&plus;%20m_2x_2%20&plus;%20...%20&plus;m_%7BN%7Dx_%7BN%7D%20&plus;%20b)
    - x : 様々な独立変数
    - m : 様々な傾き
    - b : 切片


### Assumptions of multiple linear regression
- Overfitting
    - 説明変数を追加すればするほど, モデルは複雑になる
    - モデルが複雑に慣れば, 過学習のリスクが上がる
    - 過学習を予防するテクニックは, regulatization
        - 複雑なモデルに対してペナルティ値を生成する
- Relative Scale
    - ある説明変数のスケールが他の説明変数のスケールより大きい
    - 変数を標準化することを検討


### refine Sales model
- ![equation](https://latex.codecogs.com/gif.latex?Sales%20%3D%20m_1%20%5Ccdot%20TV%20&plus;%20m_2%20%5Ccdot%20Radio%20&plus;%20b)


#### Go code
- training Model


```go
package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"log"
	"os"
	"strconv"
)

func main() {
	// 学習データセットを開く
	f, err := os.Open("./training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader を生成する
	reader := csv.NewReader(f)

	// CSV レコードをすべて読み込む
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 目的変数 Sales (y) を説明変数1 TV (x1) と
	// 説明変数2 Radio (x2), 切片でモデル化する
	// sajari/regression を使うモデルに必要な構造体を生成する
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")

	// 学習データを線形モデルの値に対応させるためにループさせる
	for i, record := range trainingData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// 目的変数を Sales にパースする
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 説明変数 x1を TV にパースする
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 説明変数 x2 を Radio にパースする
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 目的変数と説明変数を, 線形モデルに対応させる
		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))

	}

	// 学習を実行する
	r.Run()

	// 学習したモデルパラメタを出力する
	fmt.Printf("Regression Formula: %v\n", r.Formula)
	// Output: Regression Formula: Predicted = 2.9318 + TV*0.0473 + Radio*0.1794
}
```


- testing Model


```go
package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sajari/regression"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// 学習データセットを開く
	f, err := os.Open("./training.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader を生成する
	reader := csv.NewReader(f)

	// CSV レコードをすべて読み込む
	reader.FieldsPerRecord = 4
	trainingData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 目的変数 Sales (y) を説明変数1 TV (x1) と
	// 説明変数2 Radio (x2), 切片でモデル化する
	// sajari/regression を使うモデルに必要な構造体を生成する
	var r regression.Regression
	r.SetObserved("Sales")
	r.SetVar(0, "TV")
	r.SetVar(1, "Radio")

	// 学習データを線形モデルの値に対応させるためにループさせる
	for i, record := range trainingData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// 目的変数を Sales にパースする
		yVal, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 説明変数 x1を TV にパースする
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 説明変数 x2 を Radio にパースする
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 目的変数と説明変数を, 線形モデルに対応させる
		r.Train(regression.DataPoint(yVal, []float64{tvVal, radioVal}))

	}

	// 学習を実行する
	r.Run()

	// 学習したモデルパラメタを出力する
	fmt.Printf("Regression Formula: %v\n", r.Formula)
	// Output: Regression Formula: Predicted = 2.9318 + TV*0.0473 + Radio*0.1794

	// 学習データセットを開く
	f, err = os.Open("./test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV reader を生成する
	reader = csv.NewReader(f)

	// CSV レコードをすべて読み込む
	reader.FieldsPerRecord = 4
	testData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// テストデータから, 目的変数を予測し, 予測結果を, MAE で評価する
	var mAE float64
	for i, record := range testData {

		// ヘッダ行を飛ばす
		if i == 0 {
			continue
		}

		// テストデータの観測値 Sales をパースする
		yObserved, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatal(err)
		}

		// テストデータの説明変数 TV をパースする
		tvVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// テストデータの説明変数 TV をパースする
		radioVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 学習済みモデルから, 目的変数を予測する
		yPredicted, err := r.Predict([]float64{tvVal, radioVal})

		// MAE に加算する
		mAE += math.Abs(yObserved-yPredicted) / float64(len(testData))
	}

	// MAE を出力する
	fmt.Printf("MAE = %0.2f\n", mAE)
	// Output: MAE = 1.26
}
```

- MAE = 1.26, improved

## Nonlinear and other types of regression
