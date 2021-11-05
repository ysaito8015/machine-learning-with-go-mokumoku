# Chapter 03, Evaluation and Validation
- 評価とバリデーション
- 作ったモデルがどのように機能するのか計る必要がある
    - Evaluation
- プロダクション環境で得られるデータに対して, 作ったモデルがどれだけ汎化されているかを確認する必要がある
    - Validation


## Evaluation
- どのように評価するか, その手段
- "There is no one-size-fits-all metric"
    - なんにでも適合する評価項目はない
- 評価項目は, 文脈による
    - 評価の測定項目は, 作ったモデルで何をなそうとしているかによる
- しかし, 多くの機械学習の文脈において, 必要とされる測定は, どの程度, 予測, 推定や結果が, 理想的なそれらに適合しているかという測定
- 結果の種類
    - コンテニュアス
        - 連続的な数値として結果が与えられる
    - カテゴリカル
        - 定形のカテゴリーの中から結果が与えられる


### Continuous metrics
- 連続値の予測を行うモデルを考える
- 観測値と予測値, 誤差 (error) を考える
    - 誤差は, 観測値と予測値の差
        - observation - prediction = error
- 誤差をまとめる指標が必要
    - MSE, Mean Squared Error
        - 平均二乗誤差
        - 二乗をとるので誤差の値が大きくなる
        - 外れ値 (outliers) に敏感
    - MAE, Mean Absolute Error
        - 平均絶対誤差
        - 絶対値を取るので, 観測値と同じ次元の値


```go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	// CSV ファイルを開く
	f, err := os.Open("./continuous_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV ファイルを読むための CSV Reader を作る
	reader := csv.NewReader(f)

	// 観測値と推測値を格納するスライスを宣言する
	var observed []float64
	var predicted []float64

	// 行数のカウントを格納する変数
	line := 1

	// 予期しないデータタイプがないかカラムの内容をチェックする
	for {

		// ファイルの終わりに来たらループを抜ける
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// ヘッダ行を飛ばす
		if line == 1 {
			line++
			continue
		}

		// 観測値と予測値を読み出す
		observedVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		predictedVal, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Parsing line %d failed, unexpected type\n", line)
			continue
		}

		// それぞれのスライスへ追加する
		observed = append(observed, observedVal)
		predicted = append(predicted, predictedVal)
		line++
	} // データ読出しおわり

	// MSE, MAE を計算する
	var mAE float64
	var mSE float64
	for idx, oVal := range observed {
		mAE += math.Abs(oVal-predicted[idx]) / float64(len(observed))
		mSE += math.Pow(oVal-predicted[idx], 2) / float64(len(observed))
	}

	// 値を出力する
	fmt.Printf("MAE = %0.2f\n", mAE)
  // Output: MAE = 2.55 
	fmt.Printf("MSE = %0.2f\n", mSE)
  // Output: MSE = 10.51

}
```


#### R-squard value
- 決定係数, coefficient of determination
- R², R2
- 決定係数は, 観測値の分散に対する予測値の分散の比を計測している
    - (予測値の分散) / (観測値の分散)
    - 観測値のばらつきの中で, 予測値のばらつきがどのぐらいを占めているのか
        - 予測値の分散が観測値の分散に比べて多きければ, 精度が高い
    - 全変動平方和に対する残差平方和の割合
- y ハット = 予測値
- y = 実測値
- y バー = 予測値の平均

- ![equation](https://latex.codecogs.com/gif.latex?%5Cdpi%7B100%7D%20%5Cfn_cm%20R%5E2%20%3D%20%5C%5C%20%5Cleft%28%20%5Csum%5E%7Bn%7D_%7Bi%20%3D%201%7D%28%5Chat%7By_i%7D%20-%20%5Cbar%7By%7D%29%28y_i%20-%20%5Cbar%7By%7D%29%20%5Cright%20%29%5E2%20/%20%5Cleft%20%28%20%5Csum%5E%7Bn%7D_%7Bi%20%3D%201%7D%20%28%5Chat%7By%7D%20-%20%5Cbar%7By%7D%29%5E2%20%5Csum%5E%7Bn%7D_%7Bi%20%3D%201%7D%28y_i%20-%20%5Chat%7By%7D%29%5E2%20%5Cright%20%29)


```shell
$ go doc gonum/stat.RSquaredFrom
package stat // import "gonum.org/v1/gonum/stat"

func RSquaredFrom(estimates, values, weights []float64) float64
    RSquaredFrom returns the coefficient of determination defined as

        R^2 = 1 - \sum_i w[i]*(estimate[i] - value[i])^2 / \sum_i w[i]*(value[i] - mean(values))^2

    and the data in estimates and values with the given weights.

    The lengths of estimates and values must be equal. If weights is nil then
    all of the weights are 1. If weights is not nil, then len(values) must equal
    len(weights).
```


- `RSquaredFrom` 関数


```go
// go doc -src gonum/stat.RSquaredFrom
func RSquaredFrom(estimates, values, weights []float64) float64 {
        if len(estimates) != len(values) {
                panic("stat: slice length mismatch")
        }
        if weights != nil && len(weights) != len(values) {
                panic("stat: slice length mismatch")
        }

        w := 1.0
        mean := Mean(values, weights)
        var res, tot, d float64
        for i, val := range values {
                if weights != nil {
                        w = weights[i]
                }
                d = val - estimates[i]
                // 残差平方和の計算
                res += w * d * d
                d = val - mean
                // 全変動平方和の計算
                tot += w * d * d
        }
        // 1 - (残差平方和 / 全変動平方和)
        return 1 - res/tot
}
```


```go
	// 決定係数を計算する
	rSquared := stat.RSquaredFrom(observed, predicted, nil)

	// 値を出力する
	fmt.Printf("R^2 = %0.2f\n", rSquared)
	// Output: R^2 = 0.37
```


### Categorical metrics
- 離散値をとるモデルを考えた場合, 観測値と予測値は, 定形の値から一つをとる


#### Individual evaluation metrics for categorical variables
- 連続値を取るモデルの評価と同じように, 銀の弾丸はない
- 毎回必ず, 評価のためのメトリックが問題とゴールに適合しているか確認する


##### 予測とその結果に対するいくつかのシナリオ
- True Positive (TP), 真陽性
    - あるカテゴリであると予測して, 観測結果もそのカテゴリであった
    - 陽性と予測して (Predicted Positive), 観測結果も陽性
- Flase Positive (FP), 偽陽性
    - あるカテゴリであると予測して, 観測結果は, 別のカテゴリだった
    - 陽性と予測して (Predicted Positive), 観測結果は陰性
- True Negative (TN), 真陰性
    - あるカテゴリでないと予測して, 観測結果もそのカテゴリではなかった
    - 陰性と予測して (Predicted Negative), 観測結果も陰性
- False Negative (FN), 偽陰性
    - あるカテゴリでないと予測して, 観測結果はそのカテゴリだった
    - 陰性と予測して (Predicted Negative), 観測結果は陽性


##### シナリオに対応した合計手法, 評価方法
- Accuracy (ACC)
    - 全体に対する正しい予測の比率
    - (真陽性 + 真陰性) / (真陽性 + 真陰性 + 偽陽性 + 偽陰性)
    - (TP + TN) / (TP + TN + FP + FN)
- Precision
    - also positive predictive value (PPV)
    - 陽性的中率
    - 陽性予測判定全体に対する真陽性の比率
    - 真陽性 / (真陽性 + 偽陽性)
    - TP / (TP + FP)
- Recall
    - also sensitivity, hit rate, or true positive rate (TPR)
    - 感度
    - 本当に陽性だった群に対する真陽性の比率
    - 真陽性 / (真陽性 + 偽陰性)
    - TP / (TP + FN)
- Selectivity
    - also specificity, or true negative rate (TNR)
    - 特異度
    - 本当に陰性だった群に対する真陰性の比率
    - 真陰性 / (真陰性 + 偽陽性)
    - TN / (TN + FP)
- False Positive Rate (FPR)
    - 偽陽性率
    - 1 - 特異度, 1 - Selectivity
    - 本当に陰性だった軍に対する偽陽性の比率
    - FP / (TN + FP)

- 詳しい説明
    - https://en.wikipedia.org/wiki/Precision_and_recall


##### go code

```go
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
```



##### ２要素以上の水準に対応
- 陽性的中率, Precision
    - 陽性と予測された群に対する, 真陽性の比
- 感度, Recall
    - 陽性と観測された群に対する, 真陽性の比
- 水準ごとの Precision, Recall を計算する
    - 全体に対する値が必要であれば平均をとる
- もしある水準が他の水準に比べて重要であれば, weighted average を利用する


```go
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
    // Output:
    // Precision (class 0) = 1.00
    // Recall (class 0) = 1.00
    // Precision (class 1) = 0.96
    // Recall (class 1) = 0.94
    // Precision (class 2) = 0.94
    // Recall (class 2) = 0.96

	}
}
```


#### Confusion matrices, AUC, and ROC
##### Confusion matrices
- 混同行列, 誤差行列
- 行: 分類機が予測したカテゴリ
- 列: 観測された真のカテゴリ
- 理想的な状態では, 対角成分のみに値がある, TP, TN
- カテゴリが２以上の場合に特に有用
    - すべてのカテゴリに対するモデルのパフォーマンスを計測できる
    - どのカテゴリに対してモデルの適合度が悪いかを計測できる


|                   | Predicted Positive | Predicted Negative |
|-------------------|--------------------|--------------------|
| Observed Positive | TP, true positive  | FN, false negative |
| Observed Negative | FP, false positive | TN, true negative  |


##### Receiver Operating Characteristic curves, (ROC)
- ROC 曲線は, 二値分類の際によく用いられる
- y axis: Recall vs x axis: False Positive Rate
    - 感度, Recall, 本当に陽性だった群に対する真陽性の比率
    - 偽陽性率, FPR, 本当に陰性だった群に対する偽陽性の比率
- 二値分類のしきい値, (Thresholds) は２つに分類する境界
- ROC 曲線によって評価されるモデルは, 確率, ランク付けなどによって２つのクラスの予測をするモデル
- ROC 曲線の書き方
    - テスト例に対して, 各スコアの感度, (Recall), 偽陽性率, (FPR) を計算しプロットする
    - プロット点をつなげて曲線を描く
    - 原点からカーブの右上終端に向けて直線を描く
    - この直線が分類機のリファレンスラインになる
        - つまり分類予測がランダムによるものだった場合のライン

![ROC curve](https://i.gyazo.com/2cfc746d846206f292461585269a3376.png)


- 良い ROC 曲線は, リファレンスラインより左上側にプロットがある
    - つまりモデル, 分類機の予測が, ランダムによる分類より検出力が上回っている
- 良い ROC 曲線は, 大きな AUC を持つ


##### Area under the curve, (AUC)
- ROC 曲線と x 軸の間の面積

![AUC](https://i.gyazo.com/7f10b72c45ab215b63c711eea3f56d67.png)


##### Go code
- `gonum/stat.ROC` 関数
		- 書籍の時点と戻り値が変わっている
		- Thresholds を返すようになっている


```go
package stat // import "gonum.org/v1/gonum/stat"

// ROC returns paired false positive rate (FPR) and true positive rate
// (TPR) values corresponding to cutoff points on the receiver operator
// characteristic (ROC) curve obtained when y is treated as a binary
// classifier for classes with weights. The cutoff thresholds used to
// calculate the ROC are returned in thresh such that tpr[i] and fpr[i]
// are the true and false positive rates for y >= thresh[i].
//
// The input y and cutoffs must be sorted, and values in y must correspond
// to values in classes and weights. SortWeightedLabeled can be used to
// sort y together with classes and weights.
//
// For a given cutoff value, observations corresponding to entries in y
// greater than the cutoff value are classified as true, while those
// less than or equal to the cutoff value are classified as false. These
// assigned class labels are compared with the true values in the classes
// slice and used to calculate the FPR and TPR.
//
// If weights is nil, all weights are treated as 1. If weights is not nil
// it must have the same length as y and classes, otherwise ROC will panic.
//
// If cutoffs is nil or empty, all possible cutoffs are calculated,
// resulting in fpr and tpr having length one greater than the number of
// unique values in y. Otherwise fpr and tpr will be returned with the
// same length as cutoffs. floats.Span can be used to generate equally
// spaced cutoffs.
//
// More details about ROC curves are available at
// https://en.wikipedia.org/wiki/Receiver_operating_characteristic
func ROC(cutoffs, y []float64, classes []bool, weights []float64) (tpr, fpr, thresh []float64) {
        if len(y) != len(classes) {
                panic("stat: slice length mismatch")
        }
        if weights != nil && len(y) != len(weights) {
                panic("stat: slice length mismatch")
        }
        if !sort.Float64sAreSorted(y) {
                panic("stat: input must be sorted ascending")
        }
        if !sort.Float64sAreSorted(cutoffs) {
                panic("stat: cutoff values must be sorted ascending")
        }
        if len(y) == 0 {
                return nil, nil, nil
        }
        if len(cutoffs) == 0 {
                if cutoffs == nil || cap(cutoffs) < len(y)+1 {
                        cutoffs = make([]float64, len(y)+1)
                } else {
                        cutoffs = cutoffs[:len(y)+1]
                }
                // Choose all possible cutoffs for unique values in y.
                bin := 0
                cutoffs[bin] = y[0]
                for i, u := range y[1:] {
                        if u == y[i] {
                                continue
                        }
                        bin++
                        cutoffs[bin] = u
                }
                cutoffs[bin+1] = math.Inf(1)
                cutoffs = cutoffs[:bin+2]
        } else {
                // Don't mutate the provided cutoffs.
                tmp := cutoffs
                cutoffs = make([]float64, len(cutoffs))
                copy(cutoffs, tmp)
        }

        tpr = make([]float64, len(cutoffs))
        fpr = make([]float64, len(cutoffs))
        var bin int
        var nPos, nNeg float64
        for i, u := range classes {
                // Update the bin until it matches the next y value
                // skipping empty bins.
                for bin < len(cutoffs)-1 && y[i] >= cutoffs[bin] {
                        bin++
                        tpr[bin] = tpr[bin-1]
                        fpr[bin] = fpr[bin-1]
                }
                posWeight, negWeight := 1.0, 0.0
                if weights != nil {
                        posWeight = weights[i]
                }
                if !u {
                        posWeight, negWeight = negWeight, posWeight
                }
                nPos += posWeight
                nNeg += negWeight
                // Count false negatives (in tpr) and true negatives (in fpr).
                if y[i] < cutoffs[bin] {
                        tpr[bin] += posWeight
                        fpr[bin] += negWeight
                }
        }

        invNeg := 1 / nNeg
        invPos := 1 / nPos
        // Convert negative counts to TPR and FPR.
        // Bins beyond the maximum value in y are skipped
        // leaving these fpr and tpr elements as zero.
        for i := range tpr[:bin+1] {
                // Prevent fused float operations by
                // making explicit float64 conversions.
                tpr[i] = 1 - float64(tpr[i]*invPos)
                fpr[i] = 1 - float64(fpr[i]*invNeg)
        }
        for i, j := 0, len(tpr)-1; i < j; i, j = i+1, j-1 {
                tpr[i], tpr[j] = tpr[j], tpr[i]
                fpr[i], fpr[j] = fpr[j], fpr[i]
        }
        for i, j := 0, len(cutoffs)-1; i < j; i, j = i+1, j-1 {
                cutoffs[i], cutoffs[j] = cutoffs[j], cutoffs[i]
        }

        return tpr, fpr, cutoffs
}
```


- `gonum/integrate.Trapezoidal` 関数


```go
package integrate // import "gonum.org/v1/gonum/integrate"

// Trapezoidal returns an approximate value of the integral
//  \int_a^b f(x) dx
// computed using the trapezoidal rule. The function f is given as a slice of
// samples evaluated at locations in x, that is,
//  f[i] = f(x[i]), x[0] = a, x[len(x)-1] = b
// The slice x must be sorted in strictly increasing order. x and f must be of
// equal length and the length must be at least 2.
//
// The trapezoidal rule approximates f by a piecewise linear function and
// estimates
//  \int_x[i]^x[i+1] f(x) dx
// as
//  (x[i+1] - x[i]) * (f[i] + f[i+1])/2
// More details on the trapezoidal rule can be found at:
// https://en.wikipedia.org/wiki/Trapezoidal_rule
func Trapezoidal(x, f []float64) float64 {
        n := len(x)
        switch {
        case len(f) != n:
                panic("integrate: slice length mismatch")
        case n < 2:
                panic("integrate: input data too small")
        case !sort.Float64sAreSorted(x):
                panic("integrate: input must be sorted")
        }

        integral := 0.0
        for i := 0; i < n-1; i++ {
                integral += 0.5 * (x[i+1] - x[i]) * (f[i+1] + f[i])
        }

        return integral
}
```


- AUC 計算


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/integrate"
	"gonum.org/v1/gonum/stat"
)

func main() {
	// スコアとクラスを定義
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}

	// Recall, 感度と FPR, 偽陽性率を計算する
	tpr, fpr, thresh := stat.ROC(nil, scores, classes, nil)

	// AUC を計算する
	auc := integrate.Trapezoidal(fpr, tpr)

	// 出力する
	fmt.Printf("true positive rate: %v\n", tpr)
	// Output:
	// true positive rate: [0 0 0.5 0.5 1]
	fmt.Printf("false positive rate: %v\n", fpr)
	// Output:
	// false positive rate: [0 0.5 0.5 1 1]
	fmt.Printf("Thresholds: %v\n", thresh)
	// Output:
	// Thresholds: [+Inf 0.8 0.4 0.35 0.1]
	fmt.Printf("AUC: %v\n", auc)
	// Output:
	// AUC: 0.25

}
```


##### R code

![ROC](https://i.gyazo.com/d2a4dcb412fedae9c486a1a76444b935.png)


```R
> library("ROCR")
> rocdata <- read.table("roc_data.txt")
> rocdata
    V1    V2
1 0.10  true
2 0.35 false
3 0.40  true
4 0.80 false
> pred <- prediction(rocdata[,1], rocdata[,2])
> perf <- performance(pred, "tpr", "fpr")
> plot(perf)
> auc.tmp <- performance(pred, "auc")
> auc <- as.numeric(auc.tmp@y.values)
> auc
[1] 0.25
> table <- data.frame(Cutoff=unlist(pred@cutoffs),
+     TP=unlist(pred@tp), FP=unlist(pred@fp),
+     FN=unlist(pred@fn), TN=unlist(pred@tn),
+     Recall=unlist(pred@tp)/(unlist(pred@tp)+unlist(pred@fn)),
+     Selectivity=unlist(pred@tn)/(unlist(pred@fp)+unlist(pred@tn)),
+     Accuracy=(unlist(pred@tp)+unlist(pred@tn))/nrow(rocdata)
+ )
> table
  Cutoff TP FP FN TN Recall Selectivity Accuracy
1    Inf  0  0  2  2    0.0         1.0     0.50
2   0.80  0  1  2  1    0.0         0.5     0.25
3   0.40  1  1  1  1    0.5         0.5     0.50
4   0.35  1  2  1  0    0.5         0.0     0.25
5   0.10  2  2  0  0    1.0         0.0     0.50
```
