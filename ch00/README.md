# Chapter 00 Linear Algebra with Gonum
- Gonum で線形代数の計算例を実行してみる


## スカラー, ベクトル, 行列, テンソル

### スカラー

### ベクトル


```python
import numpy as np

a = np.array([1, 2, 3])
b = np.array([-2.4, 0.25, -1.3, 1.8, 0.61])

```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// スライスでベクトルを表現
	a := []float64{1, 2, 3}
	b := []float64{-2.4, 0.25, -1.3, 1.8, 0.61}
	fmt.Println(a)
	// Output: [1 2 3]
	fmt.Println(b)
	// Output: [-2.4 0.25 -1.3 1.8 0.61]

	// gonum/mat パッケージの, NewVecDense コンストラクタ関数でベクトルを定義する
	c := mat.NewVecDense(3, []float64{1, 2, 3})
	d := mat.NewVecDense(5, []float64{11.0, 5.2, -1.3, -7.2, 4.2})
	fmt.Println(c)
	// Output: &{{3 [1 2 3] 1}}
	fmt.Println(d)
	// Output: &{{5 [11 5.2 -1.3 -7.2 4.2] 1}}
}
```


```julia
julia> a = [1, 2, 3]
3-element Vector{Int64}:
 1
 2
 3

julia> b = [-2.4, 0.25, -1.3, 1.8, 0.61]
5-element Vector{Float64}:
 -2.4
  0.25
 -1.3
  1.8
  0.61

```


### 行列

```python
import numpy as np

A = np.array([[1, 2, 3],
              [4, 5, 6]])

B = np.array([[0.21, 0.14],
              [-1.3, 0.81],
              [0.12, -2.1]])
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	// gonum/mat パッケージの NewDense コンストラクタ関数で行列を定義する
	A := mat.NewDense(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})

	B := mat.NewDense(3, 2, []float64{
		0.21, 0.14,
		-1.3, 0.81,
		0.12, -2.1,
	})

	fmt.Println(A)
	// Output: &{{2 3 [1 2 3 4 5 6] 3} 2 3}
	fmt.Println(B)
	// Output: &{{3 2 [0.21 0.14 -1.3 0.81 0.12 -2.1] 2} 3 2}
}
```


```julia
julia> A = [1 2 3; 4 5 6]
2×3 Matrix{Int64}:
 1  2  3
 4  5  6

julia> B = [0.21 0.14; -1.3 0.81; 0.12 -2.1]
3×2 Matrix{Float64}:
  0.21   0.14
 -1.3    0.81
  0.12  -2.1

```


### テンソル

```python
import numpy as np

A = np.array([[[0, 1, 2, 3],
               [2, 3, 4, 5],
               [4, 5, 6, 7]],

              [[1, 2, 3, 4],
               [3, 4, 5, 6],
               [5, 6, 7, 8]]])
```


```go
package main

import (
	"fmt"
	"gorgonia.org/tensor"
)

func main() {
	// (2,3,4) 3-Tensor を作成する
	A := tensor.New(tensor.WithBacking([]float64{
		//
		0, 1, 2, 3,
		2, 3, 4, 5,
		4, 5, 6, 7,
		//
		1, 2, 3, 4,
		3, 4, 5, 6,
		5, 6, 7, 8,
	}), tensor.WithShape(2, 3, 4))

	fmt.Println(A)
	// Output:
	// ⎡0  1  2  3⎤
	// ⎢2  3  4  5⎥
	// ⎣4  5  6  7⎦
	//
	// ⎡1  2  3  4⎤
	// ⎢3  4  5  6⎥
	// ⎣5  6  7  8⎦

}
```


```julia
julia> A = zeros(3,4,2)
3×4×2 Array{Float64, 3}:
[:, :, 1] =
 0.0  0.0  0.0  0.0
 0.0  0.0  0.0  0.0
 0.0  0.0  0.0  0.0

[:, :, 2] =
 0.0  0.0  0.0  0.0
 0.0  0.0  0.0  0.0
 0.0  0.0  0.0  0.0

julia> A[:,:,1] = [0 1 2 3; 2 3 4 5; 4 5 6 7]
3×4 Matrix{Int64}:
 0  1  2  3
 2  3  4  5
 4  5  6  7

julia> A[:,:,2] = [1 2 3 4; 3 4 5 6; 5 6 7 8]
3×4 Matrix{Int64}:
 1  2  3  4
 3  4  5  6
 5  6  7  8

julia> A = Float64.(A)
3×4×2 Array{Float64, 3}:
[:, :, 1] =
 0.0  1.0  2.0  3.0
 2.0  3.0  4.0  5.0
 4.0  5.0  6.0  7.0

[:, :, 2] =
 1.0  2.0  3.0  4.0
 3.0  4.0  5.0  6.0
 5.0  6.0  7.0  8.0

```

#### install TensorFlow C library

```shell
$ wget https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.6.0.tar.gz
--2021-10-27 19:14:56--  https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-2.6.0.tar.gz
Resolving storage.googleapis.com (storage.googleapis.com)... 172.217.23.208, 172.217.23.240, 216.58.201.80, ...
Connecting to storage.googleapis.com (storage.googleapis.com)|172.217.23.208|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 94992541 (91M) [application/x-tar]
Saving to: ‘libtensorflow-cpu-linux-x86_64-2.6.0.tar.gz’

libtensorflow-cpu-linux-x86_64-2.6.0.tar. 100%[===================================================================================>]  90.59M  2.41MB/s    in 37s

2021-10-27 19:15:34 (2.44 MB/s) - ‘libtensorflow-cpu-linux-x86_64-2.6.0.tar.gz’ saved [94992541/94992541]


$ sudo ldconfig
```

```C
#include <stdio.h>
#include <tensorflow/c/c_api.h>

int main() {
	printf("Hello from TensorFlow C library version %s\n", TF_Version());
	return 0;
}
```



## ベクトルの内積とノルム

### 内積

```python
import numpy as np

a = np.array([1, 2, 3])
b = np.array([3, 2, 1])

print(np.dot(a, b))
print(np.sum(a * b))
```

```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := []float64{1, 2, 3}
	b := []float64{3, 2, 1}

	fmt.Println(floats.Dot(a, b))

	sum := 0.0
	for i, v := range a {
		sum += v * b[i]
	}

	fmt.Println(sum)

	c := mat.NewVecDense(3, []float64{1, 2, 3})
	d := mat.NewVecDense(3, []float64{3, 2, 1})

	fmt.Println(mat.Dot(c, d))
}
```


```julia
julia> using LinearAlgebra

julia> dot([1, 2, 3], [4, 5, 6])
32

```


### ノルム
- ベクトの「大きさ」を抽象化した量
- 正則化に使う
		- 過学習をパラメータの調節で予防する

#### L² ノルム
- ベクトルの各要素を２乗和し, 平方根を取る
- ユークリッド距離

![equation](http://latex.codecogs.com/gif.latex?%5C%7C%5Coverrightarrow%7Bx%7D%5C%7C_2%20%3D%20%5Csqrt%7Bx%5E2_1%20&plus;%20x%5E2_2%20&plus;%20...%20&plus;%20x%5E2_n%7D%20%3D%20%5Csum%5E%7Bn%7D_%7Bk%3D1%7Dx%5E2_k)


#### Lⁱ ノルム
- マンハッタン距離
- ベクトルの各要素の絶対値を足し合わせる

![equation](https://latex.codecogs.com/gif.latex?%5C%7C%5Coverrightarrow%7Bx%7D%5C%7C_1%20%3D%20%7Cx_1%7C%20&plus;%20%7Cx_2%7C%20&plus;%20...%20&plus;%20%7Cx_n%7C%20%3D%20%5Csum%5E%7Bn%7D_%7Bk%20%3D%201%7D%20%7Cx_k%7C)


#### Lᵖ ノルム
- 一般系


![equation](https://latex.codecogs.com/gif.latex?%5C%7C%5Coverrightarrow%7Bx%7D%5C%7C_p%20%3D%20%28%7Cx_1%7C%5Ep%20&plus;%20%7Cx_2%7C%5Ep%20&plus;%20...%20&plus;%20%7Cx_n%7C%5Ep%29%5E%5Cfrac%7B1%7D%7Bp%7D%20%3D%20%5Cleft%28%20%5Csum%5E%7Bn%7D_%7Bk%20%3D%201%7D%20%7Cx_k%7C%5Ep%20%5Cright%20%29%5E%5Cfrac%7B1%7D%7Bp%7D)


### 表現

```python
imoprt numpy as np

a = np.array([1, 1, -1, -1])

print(np.linalg.norm(a))
print(np.linalg.norm(a, 1))
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := []float64{1, 1, -1, -1}

	// L1
	fmt.Println(floats.Norm(a, 1))
	// Output: 4

	// L2
	fmt.Println(floats.Norm(a, 2))
	// Output: 2

	b := mat.NewVecDense(4, []float64{1, 1, -1, -1})

	// L1
	fmt.Println(mat.Norm(b, 1))
	// Output: 4

	// L2
	fmt.Println(mat.Norm(b, 2))
	// Output: 2
}
```


```julia
julia> a = [1, 1, -1, -1]
4-element Vector{Int64}:
  1
  1
 -1
 -1

julia> using LinearAlgebra

julia> norm(a, 1)
4.0

julia> norm(a, 2)
2.0

```


## 行列の積
- 前の行列の列の数と, 後ろの行列の行の数が一致する必要がある


![equation](https://latex.codecogs.com/gif.latex?%5Cmathbf%7BA%7D%20%5Ccdot%20%5Cmathbf%7BB%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%20a_%7B11%7D%20%26%20a_%7B12%7D%20%26%20a_%7B13%7D%20%5C%5C%20a_%7B21%7D%20%26%20a_%7B22%7D%20%26%20a_%7B23%7D%20%5Cend%7Bpmatrix%7D%20%5Ccdot%20%5Cbegin%7Bpmatrix%7D%20b_%7B11%7D%20%26%20b_%7B12%7D%20%5C%5C%20b_%7B21%7D%20%26%20b_%7B22%7D%20%5C%5C%20b_%7B31%7D%20%26%20b_%7B32%7D%20%5Cend%7Bpmatrix%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%20%5Csum%5E3_%7Bk%3D1%7Da_%7B1k%7Db_%7Bk1%7D%20%26%20%5Csum%5E3_%7Bk%3D1%7Da_%7B1k%7Db_%7Bk2%7D%20%5C%5C%20%5Csum%5E3_%7Bk%3D1%7Da_%7B2k%7Db_%7Bk1%7D%20%26%20%5Csum%5E3_%7Bk%3D1%7Da_%7B2k%7Db_%7Bk2%7D%20%5Cend%7Bpmatrix%7D)


### 行列積の表現


```python
import numpy as np

A = np.array([[0, 1, 2],
              [1, 2, 3]])

B = np.array([[2, 1],
              [2, 1],
              [2, 1]])

print(np.dot(a, b))

```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 3, []float64{
		0, 1, 2,
		1, 2, 3,
	})

	b := mat.NewDense(3, 2, []float64{
		2, 1,
		2, 1,
		2, 1,
	})

	var c mat.Dense
	c.Mul(a, b)
	fmt.Println(mat.Formatted(&c))
	// Output:
	// ⎡ 6   3⎤
	// ⎣12   6⎦
}
```


```julia
julia> a = [0 1 2; 1 2 3]
2×3 Matrix{Int64}:
 0  1  2
 1  2  3

julia> b = [2 1; 2 1; 2 1]
3×2 Matrix{Int64}:
 2  1
 2  1
 2  1

julia> a * b
2×2 Matrix{Int64}:
  6  3
 12  6

```


### アダマール積

```python
import numpy as np

a = np.array([[0, 1, 2],
              [3, 4, 5],
              [6, 7, 8]])

b = np.array([[0, 1, 2],
              [2, 0, 1],
              [1, 2, 0]])

print(a*b)
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(3, 3, []float64{
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	})

	b := mat.NewDense(3, 3, []float64{
		0, 1, 2,
		2, 0, 1,
		1, 2, 0,
	})

	var c mat.Dense
	c.MulElem(a, b)
	fmt.Println(mat.Formatted(&c))
	// Output:
	// ⎡ 0   1   4⎤
	// ⎢ 6   0   5⎥
	// ⎣ 6  14   0⎦

}
```


```julia
julia> a = [0 1 2; 3 4 5; 6 7 8]
3×3 Matrix{Int64}:
 0  1  2
 3  4  5
 6  7  8

julia> b = [0 1 2; 2 0 1; 1 2 0]
3×3 Matrix{Int64}:
 0  1  2
 2  0  1
 1  2  0

julia> a .* b
3×3 Matrix{Int64}:
 0   1  4
 6   0  5
 6  14  0

```


## 転置

![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BA%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%201%20%26%202%20%26%203%20%5C%5C%204%20%26%205%20%26%206%20%5Cend%7Bpmatrix%7D)


![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BA%7D%5ET%20%3D%20%5Cbegin%7Bpmatrix%7D%201%20%26%204%20%5C%5C%202%20%26%205%20%5C%5C%203%20%26%206%20%5Cend%7Bpmatrix%7D)


### 転置の表現

```python
import numpy as np

a = np.array([[1, 2, 3],
              [4, 5, 6]])

print(a.T)
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})

	fmt.Printf("%#v\n", a.T())
	// Output:
	// mat.Transpose{Matrix:(*mat.Dense)(0xc0000b6040)}
	fmt.Println(mat.Formatted(a.T()))
	// Output:
	// ⎡1  4⎤
	// ⎢2  5⎥
	// ⎣3  6⎦
}
```


```julia
julia> a = [1 2 3; 4 5 6]
2×3 Matrix{Int64}:
 1  2  3
 4  5  6

julia> transpose(a)
3×2 transpose(::Matrix{Int64}) with eltype Int64:
 1  4
 2  5
 3  6

julia> a'
3×2 adjoint(::Matrix{Int64}) with eltype Int64:
 1  4
 2  5
 3  6

```


### 転置と行列積


```python
import numpy as np

a = np.array([[0, 1, 2],
              [1, 2, 3]])

b = np.array([[0, 1, 2],
              [1, 2, 3]])

print(np.dot(a, b.T))
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 3, []float64{
		0, 1, 2,
		1, 2, 3,
	})

	b := mat.NewDense(2, 3, []float64{
		0, 1, 2,
		1, 2, 3,
	})

	var c mat.Dense
	c.Mul(a, b.T())
	fmt.Println(mat.Formatted(&c))
	// Output:
	// ⎡ 5   8⎤
	// ⎣ 8  14⎦
}
```


```julia
julia> a = [0 1 2; 1 2 3]
2×3 Matrix{Int64}:
 0  1  2
 1  2  3

julia> b = [0 1 2; 1 2 3]
2×3 Matrix{Int64}:
 0  1  2
 1  2  3

julia> a * transpose(b)
2×2 Matrix{Int64}:
 5   8
 8  14

julia> a * b'
2×2 Matrix{Int64}:
 5   8
 8  14

```


## 行列式と逆行列

### 単位行列
- 行と列の数が等しく, 左上から右下に 1 が並び, その他の要素は 0 になる.
- 行列 A に, 単位行列をかけても, 行列 A は変化しない


![equation](https://latex.codecogs.com/gif.latex?%5Cbegin%7Bpmatrix%7D%201%20%26%200%20%26%20%5Cdots%20%26%200%20%5C%5C%200%20%26%201%20%26%20%5Cdots%20%26%200%20%5C%5C%20%5Cvdots%20%26%20%5Cvdots%20%26%20%5Cddots%20%26%20%5Cvdots%20%5C%5C%200%20%26%200%20%26%20%5Cdots%20%26%201%20%5Cend%7Bpmatrix%7D)


### 単位行列の表現

```python
import numpy as np

print(np.eye(2)
print(np.eye(3)
print(np.eye(4)

```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	i := mat.NewDiagonalRect(2, 2, []float64{1, 1})

	fmt.Println(mat.Formatted(i))
	// Output:
	// ⎡1  0⎤
	// ⎣0  1⎦

	i = mat.NewDiagonalRect(3, 3, []float64{1, 1, 1})

	fmt.Println(mat.Formatted(i))
	// Output:
	// ⎡1  0  0⎤
	// ⎢0  1  0⎥
	// ⎣0  0  1⎦

	i = mat.NewDiagonalRect(4, 4, []float64{1, 1, 1, 1})

	fmt.Println(mat.Formatted(i))
	// Output:
	// ⎡1  0  0  0⎤
	// ⎢0  1  0  0⎥
	// ⎢0  0  1  0⎥
	// ⎣0  0  0  1⎦
}
```


```julia
julia> using LinearAlgebra

julia> Matrix{Int8}(I, 2, 2)
2×2 Matrix{Int8}:
 1  0
 0  1

julia> Matrix{Int8}(I, 3, 3)
3×3 Matrix{Int8}:
 1  0  0
 0  1  0
 0  0  1

julia> Matrix{Int8}(I, 4, 4)
4×4 Matrix{Int8}:
 1  0  0  0
 0  1  0  0
 0  0  1  0
 0  0  0  1

```


### 逆行列
- ある行列 A にかけると, 単位行列になる行列


![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BA%7D%20%5Ccdot%20%5Cmathit%28A%29%5E%7B-1%7D%20%3D%20%5Cmathit%7BA%7D%5E%7B-1%7D%20%5Ccdot%20%5Cmathit%28A%29%20%3D%20%5Cmathit%7BE%7D)


### 行列式
- 正方行列に対して定義される量
    - 正方行列, 行と列が同数
- 線形変換に対して線形空間の拡大率
- 行列式が 0 である場合, 逆行列は存在しない

![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BA%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%20a%20%26%20b%20%5C%5C%20c%20%26%20d%20%5Cend%7Bpmatrix%7D%20%5C%5C%20%7C%20%5Cmathit%7BA%7D%20%7C%20%3D%20%5Cmathrm%7Bdet%7D%20%5Cmathit%7BA%7D%20%3D%20ad%20-%20bc%20%5C%5C%20%5Cmathit%7BA%7D%5E%7B-1%7D%20%3D%20%5Cfrac%7B1%7D%7Bad-bc%7D%20%5Cbegin%7Bpmatrix%7D%20d%20%26%20-b%20%5C%5C%20-c%20%26%20a%20%5Cend%7Bpmatrix%7D)


### 行列式の表現

```python
import numpy as np

a = np.array([[1, 2],
              [3, 4]])

print(np.linalg.det(a))

b = np.array([[1, 2],
              [0, 0]])

print(np.linalg.det(b))
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewDense(2, 2, []float64{
		1, 2,
		3, 4,
	})

	fmt.Println(mat.Det(a))
	// Output: -2

	b := mat.NewDense(2, 2, []float64{
		1, 2,
		0, 0,
	})

	fmt.Println(mat.Det(b))
	// Output: 0
}
```

```julia
julia> using LinearAlgebra

julia> A = [1 2; 3 4]
2×2 Matrix{Int64}:
 1  2
 3  4

julia> det(A)
-2.0

julia> B = [1 2; 0 0]
2×2 Matrix{Int64}:
 1  2
 0  0

julia> det(B)
0.0

```


### 逆行列の表現

```python
import numpy as np

a = np.array([[1, 2],
              [3, 4]])

print(np.linalg.inv(a))

b = np.array([[1, 2],
              [0, 0]])

print(np.linalg.inv(b))
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {
	a := mat.NewDense(2, 2, []float64{
		1, 2,
		3, 4,
	})

	var aInverse mat.Dense
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fmt.Println(mat.Formatted(&aInverse))
	// Output:
	// ⎡-1.9999999999999996   0.9999999999999998⎤
	// ⎣ 1.4999999999999998  -0.4999999999999999⎦

	b := mat.NewDense(2, 2, []float64{
		1, 2,
		0, 0,
	})

	var bInverse mat.Dense
	if err := bInverse.Inverse(b); err != nil {
		log.Fatal(err)
		// Output:
		// 2021/10/28 19:52:01 matrix singular or near-singular with condition number +Inf
		// exit status 1
	}
	fmt.Println(mat.Formatted(&bInverse))

}
```


```julia
julia> using LinearAlgebra

julia> A = [1 2; 3 4]
2×2 Matrix{Int64}:
 1  2
 3  4

julia> inv(A)
2×2 Matrix{Float64}:
 -2.0   1.0
  1.5  -0.5

julia> B = [1 2; 0 0]
2×2 Matrix{Int64}:
 1  2
 0  0

julia> inv(B)
ERROR: LAPACKException(2)
Stacktrace:
 [1] chklapackerror(ret::Int64)
   @ LinearAlgebra.LAPACK /buildworker/worker/package_linux64/build/usr/share/julia/stdlib/v1.6/LinearAlgebra/src/lapack.jl:38
 [2] trtrs!(uplo::Char, trans::Char, diag::Char, A::Matrix{Float64}, B::Matrix{Float64})
   @ LinearAlgebra.LAPACK /buildworker/worker/package_linux64/build/usr/share/julia/stdlib/v1.6/LinearAlgebra/src/lapack.jl:3426
 [3] ldiv!
   @ /buildworker/worker/package_linux64/build/usr/share/julia/stdlib/v1.6/LinearAlgebra/src/triangular.jl:739 [inlined]
 [4] inv(A::UpperTriangular{Float64, Matrix{Float64}})
   @ LinearAlgebra /buildworker/worker/package_linux64/build/usr/share/julia/stdlib/v1.6/LinearAlgebra/src/triangular.jl:821
 [5] inv(A::Matrix{Int64})
   @ LinearAlgebra /buildworker/worker/package_linux64/build/usr/share/julia/stdlib/v1.6/LinearAlgebra/src/dense.jl:811
 [6] top-level scope
   @ REPL[6]:1

```


## 線形変換
- 線形写像
- 行列によって空間（線形空間）を変形させること
    - https://www.headboost.jp/what-is-linear-transformation/


### ベクトルの描画


```python
import numpy as np
import matplotlib.pyplot as plt

# 矢印を描画する
def arrow(start, size, color):
    plt.quiver(start[0], start[1], size[0], size[1],
        angles="xy", scale_units="xy", scale=1, color=color)

# 矢印の始点
s = np.array([0, 0])

# ベクトル
a = np.array([2, 3])

arrow(s, a, color="black")

# グラフ表示
plt.xlim([-3, 3])
plt.ylim([-3, 3])
plt.xlabel("x", size=14)
plt.ylabel("y", size=14)
plt.grid()
plt.axes().set_aspect("equal")
plt.show()
```


### 線形変換
![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BA%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%202%20%26%20-1%20%5C%5C%202%20%26%20-2%20%5Cend%7Bpmatrix%7D%20%5C%5C%20%5Cvec%7Ba%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%202%20%5C%5C%203%20%5Cend%7Bpmatrix%7D%20%5C%5C%20%5Cvec%7Bb%7D%20%3D%20%5Cmathit%7BA%7D%20%5Cvec%7Ba%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%202%20%26%20-1%20%5C%5C%202%20%26%20-2%20%5Cend%7Bpmatrix%7D%20%5Ccdot%20%5Cbegin%7Bpmatrix%7D%202%20%5C%5C%203%20%5Cend%7Bpmatrix%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%201%20%5C%5C%20-2%20%5Cend%7Bpmatrix%7D)


### 線形変換の表現

```python
import numpy as np
import matplotlib.pyplot as plt

# 矢印を描画する
def arrow(start, size, color):
    plt.quiver(start[0], start[1], size[0], size[1],
        angles="xy", scale_units="xy", scale=1, color=color)

s = np.array([0, 0])

a = np.array([2, 3])

A = np.array([[2, -1],
              [2, -2]])

b = np.dot(A, a)

print("変換前のベクトル (a):", a)
print("変換後のベクトル (b):", b)

arrow(s, a, color="black")
arrow(s, b, color="blue")

# グラフ表示
plt.xlim([-3, 3])
plt.ylim([-3, 3])
plt.xlabel("x", size=14)
plt.ylabel("y", size=14)
plt.grid()
plt.axes().set_aspect("equal")
plt.show()
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func main() {
	a := mat.NewVecDense(2, []float64{2, 3})

	A := mat.NewDense(2, 2, []float64{
		2, -1,
		2, -2,
	})

	var b mat.Dense
	b.Mul(A, a)
	fmt.Println(mat.Formatted(&b))
	// Output:
	// ⎡ 1⎤
	// ⎣-2⎦
}
```


```julia
julia> a = [2; 3]
2-element Vector{Int64}:
 2
 3

julia> A = [2 -1; 2 -2]
2×2 Matrix{Int64}:
 2  -1
 2  -2

julia> A * a
2-element Vector{Int64}:
  1
 -2

```


### 標準基底
![equation](https://latex.codecogs.com/gif.latex?%5Cvec%7Be%7D_x%20%3D%20%5Cbegin%7Bpmatrix%7D%201%20%5C%5C%200%20%5Cend%7Bpmatrix%7D%2C%20%5Cvec%7Be%7D_y%20%3D%20%5Cbegin%7Bpmatrix%7D%200%20%5C%5C%201%20%5Cend%7Bpmatrix%7D%20%5C%5C%20%5Cvec%7Ba%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%202%20%5C%5C%203%20%5Cend%7Bpmatrix%7D%20%3D%202%20%5Ccdot%20%5Cbegin%7Bpmatrix%7D%201%20%5C%5C%200%20%5Cend%7Bpmatrix%7D%20&plus;%203%20%5Ccdot%20%5Cbegin%7Bpmatrix%7D%200%20%5C%5C%201%20%5Cend%7Bpmatrix%7D%20%5C%5C%20%3D%202%20%5Cvec%7Be%7D_x%20&plus;%203%20%5Cvec%7Be%7D_y)


### 標準基底の表現


```python
import numpy as np
import matplotlib.pyplot as plt

# 矢印を描画する
def arrow(start, size, color):
    plt.quiver(start[0], start[1], size[0], size[1],
        angles="xy", scale_units="xy", scale=1, color=color)

s = np.array([0, 0])

a = np.array([2, 3])

e_x = np.array([1, 0])
e_y = np.array([0, 1])


print("(a):", a)
print("(e_x):", e_x)
print("(e_y):", e_y)

arrow(s, a, color="blue")
arrow(s, e_x, color="black")
arrow(s, e_y, color="black")

# グラフ表示
plt.xlim([-3, 3])
plt.ylim([-3, 3])
plt.xlabel("x", size=14)
plt.ylabel("y", size=14)
plt.grid()
plt.axes().set_aspect("equal")
plt.show()
```

## 固有値と固有値ベクトル

### 固有値, 固有ベクトル
- λ: スカラー, 行列 A の固有値
- x : 固有ベクトル
    - 線形変換により各要素が固有値倍になるベクトル


![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BA%7D%20%5Cvec%7Bx%7D%20%3D%20%5Clambda%20%5Cvec%7Bx%7D)

- 単位行列を考える



![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BE%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%201%20%26%200%20%5C%5C%200%20%26%201%20%5Cend%7Bpmatrix%7D)


- 単位行列をかけてもベクトルは変化しない


![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BA%7D%20%5Cvec%7Bx%7D%20%3D%20%5Clambda%20%5Cmathit%7BE%7D%20%5Cvec%7Bx%7D%20%5C%5C%20%28%5Cmathit%7BA%7D%20-%20%5Clambda%20%5Cmathit%7BE%7D%20%29%20%5Cvec%7Bx%7D%20%3D%20%5Cvec%7B0%7D)


- （）カッコ内の行列が逆行列を持つ場合
    - 逆行列を辺々, 左からかける


![equation](https://latex.codecogs.com/gif.latex?%5Cvec%7Bx%7D%20%3D%20%28%5Cmathit%7BA%7D%20-%20%5Clambda%20%5Cmathit%7BE%7D%20%29%5E%7B-1%7D%20%5Cvec%7B0%7D%20%5C%5C%20%3D%20%5Cvec%7B0%7D)


- （）カッコ内の行列が逆行列を持たない場合
    - 以下の関係が満たされる
    - これを行列 A の固有方程式という


![equation](https://latex.codecogs.com/gif.latex?%5Cmathrm%7Bdet%7D%20%28%5Cmathit%7BA%7D%20-%20%5Clambda%20%5Cmathit%7BE%7D%20%29%20%3D%200)


### 固有値を求める


![equation](https://latex.codecogs.com/gif.latex?%5Cmathit%7BA%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%203%20%26%201%20%5C%5C%202%20%26%204%20%5Cend%7Bpmatrix%7D%20%5C%5C%20%5C%5C%20%5C%5C%20%5Cmathrm%7Bdet%7D%20%5Cmathit%7BA%20-%20%5Clambda%20%5Cmathit%7BE%7D%7D%20%3D%200%20%5C%5C%20%5C%5C%20%5Cmathrm%7Bdet%7D%20%5Cleft%28%20%5Cbegin%7Bpmatrix%7D%203%20%26%201%20%5C%5C%202%20%26%204%20%5Cend%7Bpmatrix%7D%20-%20%5Clambda%20%5Cbegin%7Bpmatrix%7D%201%20%26%200%5C%5C%200%20%26%201%20%5Cend%7Bpmatrix%7D%20%5Cright%29%20%3D%200%20%5C%5C%20%5C%5C%20%5C%5C%20%5Cmathrm%7Bdet%7D%20%5Cleft%20%28%20%5Cbegin%7Bpmatrix%7D%203%20-%20%5Clambda%20%26%201%20%5C%5C%202%20%26%204-%20%5Clambda%20%5Cend%7Bpmatrix%7D%20%5Cright%20%29%20%5C%5C%20%5C%5C%20%5C%5C%20%283%20-%20%5Clambda%29%284%20-%20%5Clambda%29%20-%201%20%5Ctimes%202%20%3D%200%20%5C%5C%20%5C%5C%20%5Clambda%20%5E2%20-%207%20%5Clambda%20&plus;%2010%20%3D%200%20%5C%5C%20%28%5Clambda%20-%202%29%28%5Clambda%20-%205%29%20%3D%200)


- このとき λ の固有値は 2 もしくは 5 となる


### 固有ベクトルを求める
- λ = 2 の場合と, λ = 5 の場合に分ける
- λ = 2 の場合

![equation](https://latex.codecogs.com/gif.latex?%5Cvec%7Bx%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%20p%20%5C%5C%20q%20%5Cend%7B%7D%20%5C%5C%20%5C%5C%20%5C%5C%20%28%5Cmathit%7BA%7D%20-%202%20%5Cmathit%7BE%7D%29%20%5Cbegin%7Bpmatrix%7D%20p%20%5C%5C%20q%20%5Cend%7B%7D%20%3D%20%5Cvec%7B0%7D%20%5C%5C%20%5C%5C%20%5C%5C%20%5Cbegin%7Bpmatrix%7D%203%20-2%20%26%201%20%5C%5C%202%204%20-2%20%5Cend%7B%7D%20%5Cbegin%7Bpmatrix%7D%20p%20%5C%5C%20q%20%5Cend%7B%7D%20%3D%20%5Cvec%7B0%7D%20%5C%5C%20%5C%5C%20%5C%5C%20%5Cbegin%7Bpmatrix%7D%201%20%26%201%20%5C%5C%202%20%26%202%20%5Cend%7B%7D%20%5Cbegin%7Bpmatrix%7D%20p%20%5C%5C%20q%20%5Cend%7B%7D%20%3D%20%5Cvec%7B0%7D%20%5C%5C%20%5C%5C%20%5C%5C%20%5Cbegin%7Bpmatrix%7D%20p%20&plus;%20q%20%5C%5C%202p%20&plus;%202q%20%5Cend%7B%7D%20%3D%20%5Cvec%7B0%7D)

- このとき p + q = 0 となる
- この条件を満たすベクトルが, λ =2 の場合の A の固有ベクトルになる
- t は任意の実数


![equation](https://latex.codecogs.com/gif.latex?%5Cvec%7Bx%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%20t%20%5C%5C%20-t%20%5Cend%7B%7D)


- λ = 5 の場合


![equation](https://latex.codecogs.com/gif.latex?%28%5Cmathit%7BA%7D%20-%205%20%5Cmathit%7BE%7D%29%20%5Cbegin%7Bpmatrix%7D%20p%20%5C%5C%20q%20%5Cend%7B%7D%20%3D%20%5Cvec%7B0%7D%20%5C%5C%20%5C%5C%20%5C%5C%20%5Cbegin%7Bpmatrix%7D%203%20-5%20%26%201%20%5C%5C%202%20%26%204%20-5%20%5Cend%7B%7D%20%5Cbegin%7Bpmatrix%7D%20p%20%5C%5C%20q%20%5Cend%7B%7D%20%3D%20%5Cvec%7B0%7D%20%5C%5C%20%5C%5C%20%5C%5C%20%5Cbegin%7Bpmatrix%7D%20-2%20%26%201%20%5C%5C%202%20%26%20-1%20%5Cend%7B%7D%20%5Cbegin%7Bpmatrix%7D%20p%20%5C%5C%20q%20%5Cend%7B%7D%20%3D%20%5Cvec%7B0%7D%20%5C%5C%20%5C%5C%20%5C%5C%20%5Cbegin%7Bpmatrix%7D%20-2p%20&plus;%20q%20%5C%5C%202p%20-%20q%20%5Cend%7B%7D%20%3D%20%5Cvec%7B0%7D)


![equation](https://latex.codecogs.com/gif.latex?%5Cvec%7Bx%7D%20%3D%20%5Cbegin%7Bpmatrix%7D%20t%20%5C%5C%202t%20%5Cend%7B%7D)


### 固有値と固有ベクトルの計算

```python
import numpy as np

a = np.array([[3, 1],
              [2, 4]])

# 固有値と, 固有ベクトルを求める
ev = np.linalg.eig(a)

print(ev[0])
print(ev[1])

```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {
	a := mat.NewDense(2, 2, []float64{
		3, 1,
		2, 4,
	})

	var eig mat.Eigen
	ok := eig.Factorize(a, mat.EigenRight)
	if !ok {
		log.Fatal("Eigendecomposition failed")
	}

	fmt.Printf("%#v \n", eig)
	// Output:
	// mat.Eigen{n:2, kind:4, values:[]complex128{(2+0i), (5+0i)}, rVectors:(*mat.CDense)(0xc0000b60c0), lVectors:(*mat.CDense)(nil)}
	fmt.Printf("%#v \n", eig.Values(nil))
	// Output:
	// []complex128{(2+0i), (5+0i)}
	E := mat.NewCDense(2, 2, []complex128{0 + 0i, 0 + 0i, 0 + 0i, 0 + 0i})
	eig.VectorsTo(E)
	fmt.Println(E)
	// Output:
	// &{{2 2 2 [(-0.7071067811865475+0i) (-0.4472135954999579+0i) (0.7071067811865475+0i) (-0.8944271909999159+0i)]} 2 2}
}
```


```julia
julia> using LinearAlgebra

julia> A = [3 1; 2 4]
2×2 Matrix{Int64}:
 3  1
 2  4

julia> eigen(A)
Eigen{Float64, Float64, Matrix{Float64}, Vector{Float64}}
values:
2-element Vector{Float64}:
 2.0
 5.0
vectors:
2×2 Matrix{Float64}:
 -0.707107  -0.447214
  0.707107  -0.894427

```


## コサイン類似度
- コサイン類似度は, ベクトル同士の向きの近さを表す


### ノルムと三角関数で内積を表す

![equation](https://latex.codecogs.com/gif.latex?%5Cvec%7Ba%7D%20%3D%20%28a_1%2C%20a_2%29%2C%20%5C%20%5C%20%5Cvec%7Bb%7D%20%3D%20%28b_1%2C%20b_2%29%20%5C%5C%20%5Cvec%7Ba%7D%20%5Ccdot%20%5Cvec%7Bb%7D%20%3D%20a_1b_1%20&plus;%20a_2b_2%20%5C%5C%20%5Cvec%7Ba%7D%20%5Ccdot%20%5Cvec%7Bb%7D%20%3D%20%5C%7C%20%5Cvec%7Ba%7D%20%5C%7C_2%20%5C%7C%20%5Cvec%7Bb%7D%20%5C%7C_2%20%5Ccos%20%5Ctheta%20%5C%5C%20%5C%5C%20%3D%20%5Csqrt%7Ba%5E2_1%20&plus;%20a%5E2_2%7D%20%5Csqrt%7Bb%5E2_1%20&plus;%20b%5E2_2%7D%20%5Ccos%20%5Ctheta%20%5C%5C%20%5C%5C%20%5Ccos%20%5Ctheta%20%3D%20%5Cfrac%7Ba_1b_1%20&plus;%20a_2b_2%7D%7B%5Csqrt%7Ba%5E2_1%20&plus;%20a%5E2_2%7D%20%5Csqrt%7Bb%5E2_1%20&plus;%20b%5E2_2%7D%7D)


- cos θ は, ベクトル間の角度 θが 0 のとき最大値を取り, この角度が大きくなると小さくなる
    - 2 つのベクトルの向きがどれだけ揃っているかの指標


- n 次元へ拡張


![equation](https://latex.codecogs.com/gif.latex?%5Ccos%20%5Ctheta%20%5C%5C%20%5C%5C%20%3D%20%5Cfrac%7B%5Csum%5E%7Bn%7D_%7Bk%3D1%7D%20a_kb_k%7D%7B%5Csqrt%7B%5Csum%5E%7Bn%7D_%7Bk%3D1%7D%20a%5E2_k%7D%20%5Csqrt%7B%5Csum%5E%7Bn%7D_%7Bk%3D1%7D%20b%5E2_k%7D%7D%20%5C%5C%20%5C%5C%20%3D%20%5Cfrac%7B%5Cvec%7Ba%7D%20%5Ccdot%20%5Cvec%7Bb%7D%7D%7B%5C%7Ca%5C%7C_2%20%5C%7C%5Cvec%7Bb%7D%5C%7C_2%7D)


### コサイン類似度を計算する


```python
import numpy as np

def cos_sim(vec_1, vec_2):
    return np.dot(vec_1, vec_2) / (np.linalg.norm(vec_1) * np.linalg.norm(vec_2))

a = np.array([2, 2, 2, 2])
b = np.array([1, 1, 1, 1])
c = np.array([-1, -1, -1, -1])

print(cos_sim(a, b))
print(cos_sim(a, c))
```


```go
package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func cosSim(v1, v2 mat.Vector) float64 {
	return mat.Dot(v1, v2) / (mat.Norm(v1, 2) * mat.Norm(v2, 2))
}

func main() {
	a := mat.NewVecDense(4, []float64{2, 2, 2, 2})
	b := mat.NewVecDense(4, []float64{1, 1, 1, 1})
	c := mat.NewVecDense(4, []float64{-1, -1, -1, -1})

	fmt.Println(cosSim(a, b))
	// Output: 1
	fmt.Println(cosSim(a, c))
	// Output: -1
}
```


```julia
julia> using LinearAlgebra

julia> a = [2, 2, 2, 2]
4-element Vector{Int64}:
 2
 2
 2
 2

julia> b = [1, 1, 1, 1]
4-element Vector{Int64}:
 1
 1
 1
 1

julia> c = [-1, -1, -1, -1]
4-element Vector{Int64}:
 -1
 -1
 -1
 -1

julia> dot(a, b) / (norm(a, 2) * norm(b, 2))
1.0

julia> dot(a, c) / (norm(a, 2) * norm(c, 2))
-1.0

```
