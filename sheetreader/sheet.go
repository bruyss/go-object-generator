package sheetreader

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.uber.org/zap"

	"github.com/bruyss/go-object-generator/utils"
)

func AddSheet(f *excelize.File, sheetName string, columns []string) int {
	index := f.NewSheet(sheetName)

	f.SetSheetRow(sheetName, "A1", &columns)
	bottomRight := excelize.ToAlphaString(len(columns)-1) + "2"
	formatString := fmt.Sprintf(`{
		"table_name": "%s",
		"table_style": "TableStyleMedium2",
    	"show_first_column": true,
    	"show_last_column": true,
    	"show_row_stripes": false,
    	"show_column_stripes": true
	}`, sheetName)
	err := f.AddTable(sheetName, "A1", bottomRight, formatString)

	if err != nil {
		utils.Logger.Error("Couldn't add table to worksheet", zap.String("sheet", sheetName), zap.Error(err))
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
		utils.Logger.Error("Initializing workbook failed", zap.String("filename", name), zap.Error(err))
	}

}
