package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bruyss/go-object-generator/plc"
	"github.com/bruyss/go-object-generator/sheetreader"
	"github.com/bruyss/go-object-generator/utils"
)

func init() {
	utils.InitializeCustomLogger()
}

func main() {
	// sheetreader.InitializeWorkbook("excelsource_go.xlsx")
	f, _ := excelize.OpenFile("excelsource_go.xlsx")
	o := make([]*plc.PlcObject, 30)
	sheetreader.ReadMeasmons(f, o)

}
