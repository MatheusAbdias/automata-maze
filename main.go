package main

import (
	"fmt"

	"github.com/MatheusAbdias/automata-maze/pkg/utils"
)

func main() {
	fmt.Println(utils.ReadCSVFile("automata.csv"))
}
