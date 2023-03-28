package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MatheusAbdias/automata-maze/domain"
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

func ConverToTableIntegerMatrix(table [][]string) [][]int {
	var result [][]int

	for _, line := range table {
		convertedLine := ConvertStringToInt(strings.Split(line[0], " "))
		result = append(result, convertedLine)
	}

	return result

}

func ConvertStringToInt(values []string) []int {
	var result []int
	for _, value := range values {
		convert_value, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, convert_value)
	}

	return result
}

func ConvertIntMatrixToAutomataTable(matrix [][]int) [][]*domain.Cell {
	automata_table := make([][]*domain.Cell, 0)

	for lat, line := range matrix {
		for lgn, value := range line {
			cells := make([]*domain.Cell, 0)
			cells = append(cells, domain.InitCell(lat, lgn, value))

			if lgn == len(line)-1 {
				automata_table = append(automata_table, cells)

			}
		}
	}

	return automata_table
}
