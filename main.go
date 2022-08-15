package main

import (
	"github.com/bruyss/go-object-generator/sheetreader"
	"github.com/bruyss/go-object-generator/utils"
)

func init() {
	utils.InitializeCustomLogger()
}

func main() {
	sheetreader.InitializeWorkbook("excelsource_go.xlsx")
}
