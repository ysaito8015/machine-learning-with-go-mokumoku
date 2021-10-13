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
