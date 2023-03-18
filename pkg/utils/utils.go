package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSVFile(fileName string) [][]string {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	table, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	return table
}
