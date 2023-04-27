package sheetreader

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"

	"github.com/bruyss/go-object-generator/logger"
)

func AddSheet(f *excelize.File, sheetName string, columns []string) int {
	index := f.NewSheet(sheetName)

	f.SetSheetRow(sheetName, "A1", &columns)
	bottomRight := excelize.ToAlphaString(len(columns)-1) + "2"
	formatString := fmt.Sprintf(`{
		"table_name": "%s",
		"table_style": "TableStyleMedium2",
    	"show_first_column": true,
    	"show_last_column": false,
    	"show_row_stripes": true,
    	"show_column_stripes": false
	}`, sheetName)
	err := f.AddTable(sheetName, "A1", bottomRight, formatString)

	if err != nil {
		logger.Sugar.Error(err.Error(),
			"sheet", sheetName)
	}

	return index
}

func InitializeWorkbook(name string) {

	f := excelize.NewFile()
	_ = AddSheet(f, sheetMeasmons, measmonCols)
	_ = AddSheet(f, sheetDigmons, digmonCols)
	_ = AddSheet(f, sheetValves, valveCols)
	_ = AddSheet(f, sheetControlValves, controlValveCols)
	_ = AddSheet(f, sheetMotors, motorCols)
	_ = AddSheet(f, sheetFreqMotors, freqMotorCols)

	f.DeleteSheet(f.GetSheetName(1))

	if err := f.SaveAs(name); err != nil {
		logger.Sugar.Error("Initializing workbook failed",
			"filename", name,
			"error", err)
	}

}
