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
		logger.Sugar.Errorf(err.Error(),
			"sheet", sheetName
		)
	}

	return index
}

func InitializeWorkbook(name string) {
	f := excelize.NewFile()
	_ = addSheet(f, sheetMeasmons, measmonCols)
	_ = addSheet(f, sheetDigmons, digmonCols)
	_ = addSheet(f, sheetValves, valveCols)
	_ = addSheet(f, sheetControlValves, controlValveCols)
	_ = addSheet(f, sheetMotors, motorCols)
	_ = addSheet(f, sheetFreqMotors, freqMotorCols)
	_ = addSheet(f, sheetDigouts, digoutCols)

	f.DeleteSheet(f.GetSheetName(0))

	if err := f.SaveAs(name); err != nil {
		logger.Sugar.Error("Initializing workbook failed",
			"filename", name,
			"error", err)
	}
}
