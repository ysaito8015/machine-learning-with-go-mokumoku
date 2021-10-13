# Gathering and Organizing Data
- We need to make sure thatvant, clean data to poer out machine learning modles
- garbage in, garbage out
- data is often messy and hard to aggregate
- handling the missing values, mixedtypes, or corrupted values
- always consider data gathering, parsing, and organization as a key developed with a high level of integrity

## Handling data - Gopher syle
- Go is a static typing language


## Best practices for gathering and organizing data with Go
1. Check for and enforce expected types
2. Standardize and simplify your data ingress/egress
    - e.g., `database/sql`
3. Version your data
    - e.g., Pachydarm


## CSV files
- standard library
    - using `encoding/csv`
- third party
    - `go-gota/gota/dataframe`
        - https://github.com/go-gota/gota
    - `csvutil`
        - https://pkg.go.dev/go-hep.org/x/hep/csvutil

### Reading in CSV data from a file
- iris.csv
    - http://archive.ics.uci.edu/ml/datasets/iris


```
$ head iris.csv
5.1,3.5,1.4,0.2,Iris-setosa
4.9,3.0,1.4,0.2,Iris-setosa
4.7,3.2,1.3,0.2,Iris-setosa
4.6,3.1,1.5,0.2,Iris-setosa
5.0,3.6,1.4,0.2,Iris-setosa
5.4,3.9,1.7,0.4,Iris-setosa
4.6,3.4,1.4,0.3,Iris-setosa
5.0,3.4,1.5,0.2,Iris-setosa
4.4,2.9,1.4,0.2,Iris-setosa
4.9,3.1,1.5,0.1,Iris-setosa
```

### Handling unexpected fields
- `reader.FieldsPerRecord()` メソッドを使う
    - 期待するフィールド数を指定し, `reader.Read()` で一行ずつ読み込みながら, err でフィールド数にマッチしていない行を出力する

### Handling unexpected types
- 型が混ざった CSV ファイル
- 構造体を使ってフィールドと, 期待する型を指定する
    - 各行を, `CSVRecord` 構造体に格納する
        - 構造体のフィールドを, 各列に対応させ, `error` フィールドに, その行で起きたエラーを格納する


### Manipulating CSV data with data frames
- 手動での行ごとの CSV のパースは, ちょっと冗長
    - でも, それだけが理由で非標準ライブラリを使ってはいけない

#### go-gota/dataframe ライブラリ
- gonum と golang.org/x/net に依存
- https://pkg.go.dev/github.com/go-gota/gota/dataframe

```go
package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer irisFile.Close()

	irisDF := dataframe.ReadCSV(irisFile)

	fmt.Println(irisDF)

	// Create a filter
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// Filter the dataframe to see only the rows where
	// the iris species is "Iris-versicolor".
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	fmt.Println("==== Filter only versicolor ====")
	fmt.Println(versicolorDF)

	// Filter the dataframe.
	// Select out the sepal width and species columns.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})

	fmt.Println("==== Filter only sepal width of versicolor ====")
	fmt.Println(versicolorDF)

	// Filter and select the dataframe.
	// only display the first three results.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})

	fmt.Println("==== Filter only the first three results of sepal width of versicolor ====")
	fmt.Println(versicolorDF)
}
```

- 出力

```
[150x5] DataFrame

    sepal_length sepal_width petal_length petal_width species
 0: 5.100000     3.500000    1.400000     0.200000    Iris-setosa
 1: 4.900000     3.000000    1.400000     0.200000    Iris-setosa
 2: 4.700000     3.200000    1.300000     0.200000    Iris-setosa
 3: 4.600000     3.100000    1.500000     0.200000    Iris-setosa
 4: 5.000000     3.600000    1.400000     0.200000    Iris-setosa
 5: 5.400000     3.900000    1.700000     0.400000    Iris-setosa
 6: 4.600000     3.400000    1.400000     0.300000    Iris-setosa
 7: 5.000000     3.400000    1.500000     0.200000    Iris-setosa
 8: 4.400000     2.900000    1.400000     0.200000    Iris-setosa
 9: 4.900000     3.100000    1.500000     0.100000    Iris-setosa
    ...          ...         ...          ...         ...
    <float>      <float>     <float>      <float>     <string>

==== Filter only versicolor ====
[50x5] DataFrame

    sepal_length sepal_width petal_length petal_width species
 0: 7.000000     3.200000    4.700000     1.400000    Iris-versicolor
 1: 6.400000     3.200000    4.500000     1.500000    Iris-versicolor
 2: 6.900000     3.100000    4.900000     1.500000    Iris-versicolor
 3: 5.500000     2.300000    4.000000     1.300000    Iris-versicolor
 4: 6.500000     2.800000    4.600000     1.500000    Iris-versicolor
 5: 5.700000     2.800000    4.500000     1.300000    Iris-versicolor
 6: 6.300000     3.300000    4.700000     1.600000    Iris-versicolor
 7: 4.900000     2.400000    3.300000     1.000000    Iris-versicolor
 8: 6.600000     2.900000    4.600000     1.300000    Iris-versicolor
 9: 5.200000     2.700000    3.900000     1.400000    Iris-versicolor
    ...          ...         ...          ...         ...
    <float>      <float>     <float>      <float>     <string>

==== Filter only sepal width of versicolor ====
[50x2] DataFrame

    sepal_width species
 0: 3.200000    Iris-versicolor
 1: 3.200000    Iris-versicolor
 2: 3.100000    Iris-versicolor
 3: 2.300000    Iris-versicolor
 4: 2.800000    Iris-versicolor
 5: 2.800000    Iris-versicolor
 6: 3.300000    Iris-versicolor
 7: 2.400000    Iris-versicolor
 8: 2.900000    Iris-versicolor
 9: 2.700000    Iris-versicolor
    ...         ...
    <float>     <string>

==== Filter only the first three results of sepal width of versicolor ====
[3x2] DataFrame

    sepal_width species
 0: 3.200000    Iris-versicolor
 1: 3.200000    Iris-versicolor
 2: 3.100000    Iris-versicolor
    <float>     <string>

```

## JSON
- データの内部交換などで使う場合の出力
- `encoding/json` パッケージ


### Parsing JSON
- Citi Bike API からのデータをサンプルに
    - info: https://ride.citibikenyc.com/system-data
        - JSON: https://gbfs.citibikenyc.com/gbfs/en/station_status.json

- JSON と構造体を紐付けるために
    - ```` バッククォートで囲んだ構造体タグを使用する
    - 構造体タグの中身は, `JSON:"JSONのキー名をスネークケースで"`
    - 構造体のフィールド名を大文字で始める

- (%+v: 構造体のフィールド名を同時に出力する)


### JSON output
- JSON ファイル形式によるアウトプット (marshal)
- コード例

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData struct is used to unmarshal the JSON document
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station is used to unmarshal each of the station documents in
// stationData.
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_abailable_keys"`
}

func main() {
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sd stationData

	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n\n", sd.Data.Stations[0])
}
```


```go
const citiBikeURL = "https://gbfs.citibikenyc.com/gbfs/en/station_status.json"

// stationData struct is used to unmarshal the JSON document
type stationData struct {
	LastUpdated int `json:"last_updated"`
	TTL         int `json:"ttl"`
	Data        struct {
		Stations []station `json:"stations"`
	} `json:"data"`
}

// station is used to unmarshal each of the station documents in
// stationData.
type station struct {
	ID                string `json:"station_id"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumBikesDisabled  int    `json:"num_bike_disabled"`
	NumDocksAvailable int    `json:"num_docks_available"`
	NumDocksDisabled  int    `json:"num_docks_disabled"`
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	IsReturning       int    `json:"is_returning"`
	LastReported      int    `json:"last_reported"`
	HasAvailableKeys  bool   `json:"eightd_has_abailable_keys"`
}

func main() {
	response, err := http.Get(citiBikeURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var sd stationData

	if err := json.Unmarshal(body, &sd); err != nil {
		log.Fatal(err)
	}

	// Marshal the data.
	outputData, err := json.MarshalIndent(sd, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	// Save the marshalled data toa file.
	if err := ioutil.WriteFile("citibike.json", outputData, 0644); err != nil {
		log.Fatal(err)
	}
}
```


## SQL-like databases
- 標準パッケージの `database/sql` を利用する
    - https://pkg.go.dev/database/sql
- `db.Query()` の戻り値は, `&sql.Rows{}` 構造体のポインタ型
    - `row.Close()` を呼ぶ必要がある
- `Rows{}` 構造体をパースして, 期待する型に当てはめていく
    - `Scan()` メソッドを利用する

#### コードとアウトプット

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	//	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "goml"
	password = "go-ml"
	dbname   = "goml"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	// Query the database.
	rows, err := db.Query(`
	SELECT
		sepal_length as sLength,
		sepal_width as sWidth,
		petal_length as pLength,
		petal_width as pWidth
	FROM iris
	WHERE species = $1`, "Iris-setosa")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Printf("%#v\n", rows)

	// Iterate over the rows, sending the results to
	// standard out.
	for rows.Next() {
		var (
			sLength float64
			sWidth  float64
			pLength float64
			pWidth  float64
		)

		if err := rows.Scan(&sLength, &sWidth, &pLength, &pWidth); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)
	}

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
```


```
&{0 {host=localhost port=5432 user=goml password=go-ml dbname=goml sslmode=disable 0x81bdb0} 0 {0 0} [0xc00017e000] map[] 0 1 0xc000022420 false map[0xc00017e000:map[
0xc00017e000:true]] map[] 0 0 0 0 <nil> 0 0 0 0 0x47f320}
&sql.Rows{dc:(*sql.driverConn)(0xc00017e000), releaseConn:(func(error))(0x4ca360), rowsi:(*pq.rows)(0xc00011e780), cancel:(func())(nil), closeStmt:(*sql.driverStmt)(n
il), closemu:sync.RWMutex{w:sync.Mutex{state:0, sema:0x0}, writerSem:0x0, readerSem:0x0, readerCount:0, readerWait:0}, closed:false, lasterr:error(nil), lastcols:[]dr
iver.Value(nil)}
5.10, 3.50, 1.40, 0.20
4.90, 3.00, 1.40, 0.20
4.70, 3.20, 1.30, 0.20
4.60, 3.10, 1.50, 0.20
5.00, 3.60, 1.40, 0.20
5.40, 3.90, 1.70, 0.40
4.60, 3.40, 1.40, 0.30
5.00, 3.40, 1.50, 0.20
4.40, 2.90, 1.40, 0.20
4.90, 3.10, 1.50, 0.10
5.40, 3.70, 1.50, 0.20
4.80, 3.40, 1.60, 0.20
4.80, 3.00, 1.40, 0.10
4.30, 3.00, 1.10, 0.10
5.80, 4.00, 1.20, 0.20
5.70, 4.40, 1.50, 0.40
5.40, 3.90, 1.30, 0.40
5.10, 3.50, 1.40, 0.30
5.70, 3.80, 1.70, 0.30
5.10, 3.80, 1.50, 0.30
5.40, 3.40, 1.70, 0.20
5.10, 3.70, 1.50, 0.40
4.60, 3.60, 1.00, 0.20
5.10, 3.30, 1.70, 0.50
4.80, 3.40, 1.90, 0.20
5.00, 3.00, 1.60, 0.20
5.00, 3.40, 1.60, 0.40
5.20, 3.50, 1.50, 0.20
5.20, 3.40, 1.40, 0.20
4.70, 3.20, 1.60, 0.20
4.80, 3.10, 1.60, 0.20
5.40, 3.40, 1.50, 0.40
5.20, 4.10, 1.50, 0.10
5.50, 4.20, 1.40, 0.20
4.90, 3.10, 1.50, 0.20
5.00, 3.20, 1.20, 0.20
5.50, 3.50, 1.30, 0.20
4.90, 3.60, 1.40, 0.10
4.40, 3.00, 1.30, 0.20
5.10, 3.40, 1.50, 0.20
5.00, 3.50, 1.30, 0.30
4.50, 2.30, 1.30, 0.30
4.40, 3.20, 1.30, 0.20
5.00, 3.50, 1.60, 0.60
5.10, 3.80, 1.90, 0.40
4.80, 3.00, 1.40, 0.30
5.10, 3.80, 1.60, 0.20
4.60, 3.20, 1.40, 0.20
5.30, 3.70, 1.50, 0.20
5.00, 3.30, 1.40, 0.20
```

### PostgreSQL のメモ
- docker
    - `docker run --rm --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_INITDB_ARGS="--encoding=UTF8 --no-locale" -e TZ=UTC -v postgresdb:/var/lib/postgresql/data -p 5432:5432 -d postgres:14`
    - `docker exec -it postgres /bin/bash`

- psql
- ユーザを作る
    - `CREATE USER goml WITH PASSWORD 'goml';`
- DB を作る
    - `CREATE DATABASE goml WITH OWNER goml ENCODING utf8;`
- テーブルを作る
    - `CREATE TABLE iris (sepal_length double precision, sepal_width double precision, petal_length double precision, petal_width double precision, species varchar);`
- CSV を読み込む
    - `\copy iris FROM './iris.csv' DELIMITER ',' CSV HEADER;`
- ロールの確認
    - `\du`
    - `SELECT * FROM pg_roles;`
- ロールの変更
    - `ALTER ROLE goml WITH Superuser;`

- 参考サイト
    - Docker PostgreSQL を起動するコマンド https://qiita.com/kazokmr/items/48200cc2188d27d5d536
    - Docker の主なコマンド一覧 https://qiita.com/TaaaZyyy/items/4ecf21f23e6730faf696
    - docker run https://docs.docker.com/engine/reference/commandline/run/
    - PostgreSQL add or create a user account and grant permission for database
        - https://www.cyberciti.biz/faq/howto-add-postgresql-user-account/
    - データベースを作成する https://www.dbonline.jp/postgresql/database/index2.html
    - How to import CSV file data into a PostgreSQL table?
        - https://stackoverflow.com/questions/2987433/how-to-import-csv-file-data-into-a-postgresql-table
    - PostgreSQL データ型
        - https://www.postgresql.jp/document/9.3/html/datatype-numeric.html#DATATYPE-FLOAT
    - How to copy from CSV file to PostgreSQL table with headers in CSV file?
        - https://stackoverflow.com/questions/17662631/how-to-copy-from-csv-file-to-postgresql-table-with-headers-in-csv-file
    - PostgreSQL 権限操作まとめ
        - https://katsusand.dev/posts/postgresql-auth/
    - Connecting to a PostgreSQL database with Go's database/sql package
        - https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
    - GoLang PostgreSQL Example
        - https://golangdocs.com/golang-postgresql-example


### Modifying the database
- UPDATE ステートメントとともに, `Exec()` メソッドを呼ぶことでデータベースの内容を変更できる
    - `db.Query` の代わりに `db.Exec` を呼ぶ


```go
	// Update some values.
	res, err := db.Exec(
		"UPDATE iris SET species = 'setosa' WHERE species = 'Iris-setosa'")
	if err != nil {
		log.Fatal(err)
	}

	// See how many rows where updated.
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Output the number of rows to standard out.
	log.Printf("affected = %d\n", rowCount)
```

## Caching
- data from external sources (e.g., API)
- it might be accessed frequently
- to cache data in memory or embed the data locally


### Caching data in memory
- to cache a series of values in memory
    - `go-cache` パッケージ
        - https://github.com/patrickmn/go-cache
    - `rapidash` パッケージ
        - Qiita https://qiita.com/goccy/items/5f3a9c18a7ef25f09338
        - https://github.com/blastrain/rapidash

#### go-cache パッケージ
- キーバリュー型のインメモリキャッシュを作る
- キャッシュの保存期限も設定できる


```go
package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	// Key: foo, Value: bar
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Printf("Key: foo, Value: %#v\n", foo)
	}

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
	// values which will only be used once--e.g. for passing to another
	// function--is:
	baz, found := c.Get("baz")
	if found {
		fmt.Println(baz.(int) * 2.0)
	}
}
```


### Caching data locally on disk
- embedded cache
- `BoltDB` パッケージ archived
    - https://github.com/boltdb/bolt
- forked `bblot` パッケージ
    - https://github.com/etcd-io/bbolt
- local key-value store


```go
package main

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
	"log"
)

func main() {
	// Open an embedded.db data file in current directory
	db, err := bolt.Open("embedded.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a "bucket" in the boltdb file.
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// Put the map keys and values into the BoltDB file.
	if err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("mykey"), []byte("myvalue"))
		return err
	}); err != nil {
		log.Fatal(err)
	}

	// Output he keys and values.
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key: %s, value: %s\n", k, v)
			// Outoput: key: mykey, value: myvalue
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
```


## Data versioning
- ML models produce extremely different results depending on the training data you use, the choices of parameters, and the input data.
- **Collaboration**
    - need to have our colleagues' reviews and improve on our work.
    - for review, it is important that reviewers can reproduce our model results and analyses.
- **Creativity**
    - need to track exactly
        - what data we are using,
        - what results we created, and
        - how we created them.
    - This is only way we will be able to continually improve our models and techniques.
- **Compliance**
    - Laws give uses a right to an explanation for algorithmically made decisions.
    - data versioning and reproducibility


### Data versioning projects
- Pachyderm
    - https://pachyderm.io/
    - https://github.com/pachyderm/pachyderm
    - An open source framework for data versioning and data pipelining.
#### Article
- https://neptune.ai/blog/best-7-data-version-control-tools-that-improve-your-workflow-with-machine-learning-projects
 
| name  | URL | feature |
|------|------|---------|
| Neptune | https://neptune.ai/ | Data versioning, Experiment tracking, Model registry |
| Pachyderm | https://www.pachyderm.com/ | to control an end-to-end machine learning life cycle |
| Delta Lake | https://delta.io/ | an open-source storage layer, fully compatible with Apache Spark APIs |
| Git LFS | https://git-lfs.github.com/ | an open-source project, to version large files with Git |
| Dolt | https://github.com/dolthub/dolt | a SQL database that you can fork, clone, branch, merge, push, and pull just like a git |
| LakeFS | https://lakefs.io/ | an open-source platform that provides a Git-like branching and committing model, |
| DVC | https://dvc.org/ | an open-source version control system for machine learning projects, data pipeline|

- ACID: Atomicity, Consistency, Reliability, and Durability

### Pachyderm jargon
- Repositories
- Commits
- Branches
- Files


### Deploying/installing Pachyderm
- it's on top of Kubernetes
- local installation
    - https://docs.pachyderm.com/latest/getting_started/local_installation/

### Creating data repositories for data versioning
#### instration of minikube
- How to Install Minikube on Ubuntu 18.04 / 20.04
    - https://phoenixnap.com/kb/install-minikube-on-ubuntu

```shell
$ wget https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
--2021-10-13 20:27:18--  https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
Resolving storage.googleapis.com (storage.googleapis.com)... 172.217.23.208, 172.217.23.240, 216.58.201.80, ...
Connecting to storage.googleapis.com (storage.googleapis.com)|172.217.23.208|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 69041843 (66M) [application/octet-stream]
Saving to: ‘minikube-linux-amd64’

minikube-linux-amd64                      100%[===================================================================================>]  65.84M  2.41MB/s    in 27s

2021-10-13 20:27:46 (2.42 MB/s) - ‘minikube-linux-amd64’ saved [69041843/69041843]

$ sudo cp -a ./minikube-linux-amd64 /usr/local/bin/minikube
$ sudo chmod 755 ./minikube-linux-amd64 /usr/local/bin/minikube

$ minikube version
minikube version: v1.23.2
commit: 0a0ad764652082477c00d51d2475284b5d39ceed

$ curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 44.7M  100 44.7M    0     0  2420k      0  0:00:18  0:00:18 --:--:-- 2452k

$ chmod 755 ./kubectl

$ ls -al ./kubectl
-rwxr-xr-x 1 ysaito ysaito 46907392 Oct 13 20:29 ./kubectl

$ sudo cp -a ./kubectl /usr/local/bin/

$ kubectl --version
Error: unknown flag: --version
See 'kubectl --help' for usage.
/home/ysaito/src% kubectl version
Client Version: version.Info{Major:"1", Minor:"22", GitVersion:"v1.22.2", GitCommit:"8b5a19147530eaac9476b0ab82980b4088bbc1b2", GitTreeState:"clean", BuildDate:"2021-09-15T21:38:50Z", GoVersion:"go1.16.8", Compiler:"gc", Platform:"linux/amd64"}
The connection to the server localhost:8080 was refused - did you specify the right host or port?
/home/ysaito/src% kubectl version -o json
{
  "clientVersion": {
    "major": "1",
    "minor": "22",
    "gitVersion": "v1.22.2",
    "gitCommit": "8b5a19147530eaac9476b0ab82980b4088bbc1b2",
    "gitTreeState": "clean",
    "buildDate": "2021-09-15T21:38:50Z",
    "goVersion": "go1.16.8",
    "compiler": "gc",
    "platform": "linux/amd64"
  }
}
The connection to the server localhost:8080 was refused - did you specify the right host or port?

$ minikube start
😄  Ubuntu 20.04 上の minikube v1.23.2
✨  virtualboxドライバーが自動的に選択されました。他の選択肢:  none, ssh
💿  VM ブートイメージをダウンロードしています...
    > minikube-v1.23.1.iso.sha256: 65 B / 65 B [-------------] 100.00% ? p/s 0s
    > minikube-v1.23.1.iso: 225.22 MiB / 225.22 MiB  100.00% 2.42 MiB p/s 1m33s
👍  コントロールプレーンのノード minikube を minikube 上で起動しています
💾  Kubernetes v1.22.2 のダウンロードの準備をしています
    > preloaded-images-k8s-v13-v1...: 511.84 MiB / 511.84 MiB  100.00% 2.41 MiB
🔥  virtualbox VM (CPUs=2, Memory=6000MB, Disk=20000MB) を作成しています...
🐳  Docker 20.10.8 で Kubernetes v1.22.2 を準備しています...
    ▪ 証明書と鍵を作成しています...
    ▪ Control Plane を起動しています...
    ▪ RBAC のルールを設定中です...
    ▪ イメージ gcr.io/k8s-minikube/storage-provisioner:v5 を使用しています
🔎  Kubernetes コンポーネントを検証しています...
🌟  有効なアドオン: storage-provisioner, default-storageclass
🏄  完了しました！ kubectl が「"minikube"」クラスタと「"default"」ネームスペースを使用するよう構成されました

$ minikube status
minikube
type: Control Plane
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured

$ curl -o /tmp/pachctl.deb -L https://github.com/pachyderm/pachyderm/releases/download/v1.13.4/pachctl_1.13.4_amd64.deb && sudo dpkg -i /tmp/pachctl.deb
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   630  100   630    0     0   2693      0 --:--:-- --:--:-- --:--:--  2703
100 35.1M  100 35.1M    0     0  2411k      0  0:00:14  0:00:14 --:--:-- 2442k
dpkg: warning: parsing file '/var/lib/dpkg/tmp.ci/control' near line 10 package 'pachctl':
 missing 'Maintainer' field
Selecting previously unselected package pachctl.
(Reading database ... 274131 files and directories currently installed.)
Preparing to unpack /tmp/pachctl.deb ...
Unpacking pachctl (1.13.4) ...
Setting up pachctl (1.13.4) ...

$ pachctl version --client-only
1.13.4
/home/ysaito/src% pachctl deploy local
serviceaccount/pachyderm created
serviceaccount/pachyderm-worker created
clusterrole.rbac.authorization.k8s.io/pachyderm created
clusterrolebinding.rbac.authorization.k8s.io/pachyderm-default created
role.rbac.authorization.k8s.io/pachyderm-worker created
rolebinding.rbac.authorization.k8s.io/pachyderm-worker created
deployment.apps/etcd created
service/etcd created
service/pachd created
service/pachd-peer created
deployment.apps/pachd created
service/dash created
deployment.apps/dash created
secret/pachyderm-storage-secret created

Pachyderm is launching. Check its status with "kubectl get all"
Once launched, access the dashboard by running "pachctl port-forward"
```


#### Create repository

```shell
$ pachctl create-repo myrepo
WARNING: 'pachctl create-repo' is deprecated and will be removed in a future release, use 'pachctl create repo' instead.
$ pachctl create repo myrepo
repo myrepo already exists

$ pachctl list-repo
WARNING: 'pachctl list-repo' is deprecated and will be removed in a future release, use 'pachctl list repo' instead.
NAME   CREATED        SIZE (MASTER) DESCRIPTION
myrepo 25 seconds ago 0B

```


### Putting data into data repositories


```shell
$ cat ./blah.txt 
This is an example file.

$ pachctl put-file myrepo master -c -f blah.txt
WARNING: 'pachctl put-file' is deprecated and will be removed in a future release, use 'pachctl put file' instead.
flag --commit / -c is deprecated; as of 1.7.2, you will get the same behavior without it
$ pachctl put file myrepo@master -c -f blah.txt
flag --commit / -c is deprecated; as of 1.7.2, you will get the same behavior without it
blah.txt 25.00b / 25.00 b [==============================================================================================================================] 0s 0.00 b/s
$ pachctl put file myrepo@master -f blah.txt 
blah.txt 25.00b / 25.00 b [==============================================================================================================================] 0s 0.00 b/s
$ pachctl list-repo
WARNING: 'pachctl list-repo' is deprecated and will be removed in a future release, use 'pachctl list repo' instead.
NAME   CREATED       SIZE (MASTER) DESCRIPTION 
myrepo 8 minutes ago 75B                       
$ pachctl list repo
NAME   CREATED       SIZE (MASTER) DESCRIPTION 
myrepo 8 minutes ago 75B                       
$ pachctl list file myrepo@master
NAME      TYPE SIZE 
/blah.txt file 75B  
```


### Getting data out of versioned data repositories

```shell
$ pachctl get file myrepo@master:blah.txt 
This is an example file.
```
