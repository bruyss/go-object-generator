package sheetreader

import (
	"github.com/xuri/excelize/v2"

	"github.com/bruyss/go-object-generator/logger"
)

func AddSheet(f *excelize.File, sheetName string, columns []string) int {
	index, err := f.NewSheet(sheetName)
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	f.SetSheetRow(sheetName, "A1", &columns)
	bottomRight, err := excelize.CoordinatesToCellName(len(columns), 2)
	if err != nil {
		logger.Sugar.Fatal(err)
	}

	enable := true
	err = f.AddTable(sheetName, &excelize.Table{
		Range:             "A1:" + bottomRight,
		Name:              sheetName,
		StyleName:         "TableStyleMedium2",
		ShowColumnStripes: false,
		ShowFirstColumn:   true,
		ShowHeaderRow:     &enable,
		ShowLastColumn:    false,
		ShowRowStripes:    &enable,
	})
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

	f.DeleteSheet(f.GetSheetName(0))

	if err := f.SaveAs(name); err != nil {
		logger.Sugar.Error("Initializing workbook failed",
			"filename", name,
			"error", err)
	}
}
