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
	file := openCsv("iris_messy_types.csv")
	defer file.Close()
	records := readCsvToStruct(file)
	fmt.Println(records)
}

type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
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
	// expected 5 fields per line
	r.FieldsPerRecord = 5

	// rawCSVData will hold successfully parsed rows.
	var rawCSVData [][]string

	for {
		// Read in a row of the CSV records
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		// if we had a parsing error, log the error
		if err != nil {
			log.Println(err)
			continue
		}
		rawCSVData = append(rawCSVData, record)
	}
	return rawCSVData
}

func readCsvToStruct(file *os.File) []CSVRecord {
	r := csv.NewReader(file)
	// expected 5 fields per line
	r.FieldsPerRecord = 5

	// Create a slice value that will hold all of the successfully parsed
	// records from the CSV.
	var csvData []CSVRecord

	for {
		// Read in a row of the CSV records
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		// Create a CSVRecord value for the row.
		var csvRecord CSVRecord
		rowIdx := 0

		// Parse each of the values in the record based on an expected type.
		for columnIdx, value := range record {
			// Parse the value is not an empty string.
			// If the value is an empty string, break the parsing loop.
			if columnIdx == 4 {
				if value == "" {
					log.Printf("Unexpected type in row %d, column %d\n", rowIdx, columnIdx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				// Add the string value to the CSVRecord instance
				csvRecord.Species = value
				continue
			}
			var floatValue float64

			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Unexpected type in row %d, column %d\n", rowIdx, columnIdx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			// Add the float value to the respective field in the CSVRecord.
			switch columnIdx {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}

			rowIdx++
		}
		// Append successfully parsed records to the slice defined above.
		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}
	}
	return csvData
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
