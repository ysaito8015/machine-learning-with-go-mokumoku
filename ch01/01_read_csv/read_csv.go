package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file := openCsv("myfile.csv")
	defer file.Close()
	records := readCsv(file)
	intMax := getMax(records)
	fmt.Println(intMax)
}

func openCsv(file string) *os.File {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func readCsv(file *os.File) [][]string {
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func getMax(records [][]string) int {
	var intMax int

	for _, record := range records {

		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		if intVal > intMax {
			intMax = intVal
		}
	}

	return intMax
}
