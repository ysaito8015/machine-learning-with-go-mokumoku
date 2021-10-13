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
	file := openCsv("iris.csv")
	defer file.Close()
	records := readCsv(file)
	Max := getMaxFloat(records)
	fmt.Println(Max)
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
	// set FieldsPerRecord to negative
	r.FieldsPerRecord = -1
	// Read in all of the CSV records
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func readCsvByLine(file *os.File) [][]string {
	r := csv.NewReader(file)
	// set FieldsPerRecord to negative
	r.FieldsPerRecord = -1

	var rawCSVData [][]string

	for {
		// Read in a row of the CSV records
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		rawCSVData = append(rawCSVData, record)
	}
	return rawCSVData
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

func getMaxFloat(records [][]string) float64 {
	var fMax float64

	for _, record := range records {

		fVal, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		if fVal > fMax {
			fMax = fVal
		}
	}

	return fMax
}
