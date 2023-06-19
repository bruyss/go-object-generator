package sheetreader

import (
	"github.com/xuri/excelize/v2"

	"github.com/bruyss/go-object-generator/logger"
)

type colFormula struct {
	column  int
	formula string
}

func addSheet(f *excelize.File, sheetName string, columns []string, colFormulas ...colFormula) int {
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
			"sheet", sheetName,
		)
	}

	for _, cf := range colFormulas {
		cell, _ := excelize.CoordinatesToCellName(cf.column+1, 2)
		f.SetCellFormula(
			sheetName,
			cell,
			cf.formula,
		)
	}

	return index
}

// InitializeWorkbook creates a new workbook for generation
func InitializeWorkbook(name string) {
	f := excelize.NewFile()
	_ = addSheet(f, sheetMeasmons, measmonCols)
	_ = addSheet(f, sheetDigmons, digmonCols)
	_ = addSheet(f, sheetValves, valveCols,
		colFormula{int(valveOutput), `=$A2`},
		colFormula{int(valveFeedbackOpenTag), `=$A2 & "_FBO"`},
		colFormula{int(valveFeedbackClosedTag), `=$A2 & "_FBC"`},
	)
	_ = addSheet(f, sheetControlValves, controlValveCols,
		colFormula{int(controlValveOutput), `=$A2`},
		colFormula{int(controlValveFeedbackTag), `=$A2 & "_FB"`},
	)
	_ = addSheet(f, sheetMotors, motorCols,
		colFormula{int(motorOutput), `=$A2`},
		colFormula{int(motorFeedbackTag), `=$A2 & "_FB"`},
		colFormula{int(motorBreakerTag), `=$A2 & "_TH"`},
		colFormula{int(motorSwitchTag), `=$A2 & "_WS"`},
	)
	_ = addSheet(f, sheetFreqMotors, freqMotorCols,
		colFormula{int(freqMotorOutput), `=$A2`},
		colFormula{int(freqMotorFeedbackTag), `=$A2 & "_FB"`},
		colFormula{int(freqMotorBreakerTag), `=$A2 & "_TH"`},
		colFormula{int(freqMotorSwitchTag), `=$A2 & "_WS"`},
		colFormula{int(freqMotorAlarmTag), `=$A2 & "_AL"`},
	)
	_ = addSheet(f, sheetDigouts, digoutCols,
		colFormula{int(digoutOutput), `=$A2`},
		colFormula{int(digoutFeedbackTag), `=$A2 & "_FB"`},
		colFormula{int(digoutBreakerTag), `=$A2 & "_TH"`},
	)

	f.DeleteSheet(f.GetSheetName(0))

	if err := f.SaveAs(name); err != nil {
		logger.Sugar.Error("Initializing workbook failed",
			"filename", name,
			"error", err)
	}
}
