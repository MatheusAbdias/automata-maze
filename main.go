package main

import (
	"github.com/MatheusAbdias/automata-maze/pkg/utils"
)

func main() {
	table := utils.ReadCSVFile("automata.csv")

	matrix := utils.ConverToTableIntegerMatrix(table)
	utils.ConvertIntMatrixToAutomataTable(matrix)

}
