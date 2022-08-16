package sheetreader

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bruyss/go-object-generator/plc"
)

func ReadMeasmons(f *excelize.File, o []*plc.PlcObject) {
	var m plc.PlcObject
	for _, row := range f.GetRows(sheetMeasmons)[1:] {
		m = plc.NewMeasmon(
			row[measmonTag],
			row[measmonDescription],
			row[measmonUnit],
			row[measmonAddress],
			row[measmonDirect],
			row[measmonMin],
			row[measmonMax],
		)
		o = append(o, &m)
	}
}
