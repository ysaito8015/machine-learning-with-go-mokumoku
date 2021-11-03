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



