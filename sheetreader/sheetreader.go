package sheetreader

import (
	"github.com/bruyss/go-object-generator/logger"
	"github.com/bruyss/go-object-generator/plc"
	"github.com/xuri/excelize/v2"
)

// getTable reads in an excel file and returns a table starting on cell A1.
// Appends empty string in the case of empty last columns for certain rows to get around default behaviour.
func getTable(f *excelize.File, sheetName string) ([][]string, error) {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return nil, err
	}
	logger.Sugar.Debugw("Column names",
		"sheet name", sheetName,
		"columns", rows[0])
	table := make([][]string, len(rows))

	table[0] = rows[0]
	for i, row := range rows[1:] {
		diff := len(rows[0]) - len(row)
		for j := 0; j < diff; j++ {
			row = append(row, "")
		}
		table[i+1] = row
	}
	return table, nil
}

// getStandardData reads a table of PLC objects and returns the columns defined in the default per object.
// Returns the column names and values as seperate values.
func getStandardData(table [][]string, nrStandardColumns int) ([]string, [][]string) {
	columnNames := table[0][:nrStandardColumns]
	standardData := make([][]string, len(table)-1)
	for i := range standardData {
		standardData[i] = table[i+1][:nrStandardColumns]
	}
	return columnNames, standardData
}

// getCustomData reads a table of PLC objects and returns the columns not defined in the default per object.
// These columns should appear after the default columns.
// Returns the column names and values as seperate values.
func getCustomData(table [][]string, nrStandardColumns int) ([]string, [][]string) {
	columnNames := table[0][nrStandardColumns:]
	if len(table[0]) == nrStandardColumns {
		return columnNames, [][]string{}
	}
	customData := make([][]string, len(table)-1)
	for i := range customData {
		customData[i] = table[i+1][nrStandardColumns:]
	}
	return columnNames, customData
}

// makeCustomDataMap takes in a 2D list of strings with columns names
// and return a list of maps mapping the column name to the values per row of data
//
// | Col 1 | Col 2 | Col 3 |
//
// |   10  |   20  |   30  |
//
// |   11  |   21  |   31  |
//
// |   12  |   22  |   32  |
//
// transforms to =>
//
// [
//
//	{"Col 1": "10", "Col 2": 20, "Col 3": 30},
//	{"Col 1": "11", "Col 2": 21, "Col 3": 31},
//	{"Col 1": "12", "Col 2": 22, "Col 3": 32},
//
// ]
func makeCustomDataMap(columnNames []string, data [][]string, dataMap *[]map[string]string) {
	for i := range *dataMap {
		(*dataMap)[i] = make(map[string]string, len(columnNames))
		for j, column := range columnNames {
			// fmt.Printf("Row %d column %d-%s: %s", i, j, column, data[i][j])
			(*dataMap)[i][column] = data[i][j]
		}
	}
	// fmt.Println(*data_map)
}

// ReadMeasmons reads the "Measmon" sheet and returns a slice of PLC objects containing the created measmons
func ReadMeasmons(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetMeasmons)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standardData := getStandardData(table, len(measmonCols))
	if len(standardData) == 0 {
		return
	}
	customColumns, customData := getCustomData(table, len(measmonCols))
	customMaps := make([]map[string]string, len(standardData))
	makeCustomDataMap(customColumns, customData, &customMaps)
	for n, row := range standardData {
		m, err := plc.NewMeasmon(
			row[measmonTag],
			row[measmonDescription],
			row[measmonUnit],
			row[measmonAddress],
			row[measmonDirect],
			row[measmonMin],
			row[measmonMax],
			customMaps[n],
		)
		if err != nil {
			logger.Sugar.Errorw(err.Error(),
				"measmon", row[measmonTag],
			)
		} else {
			o = append(o, m)
			logger.Sugar.Infow("Object added to generator",
				"measmon", m.Tag)
		}
	}
	return
}

// ReadDigmons reads the "Measmon" sheet and returns a slice of PLC objects containing the created digmons
func ReadDigmons(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetDigmons)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standardData := getStandardData(table, len(digmonCols))
	if len(standardData) == 0 {
		return
	}
	customColumns, customData := getCustomData(table, len(digmonCols))
	customMaps := make([]map[string]string, len(standardData))
	makeCustomDataMap(customColumns, customData, &customMaps)
	for n, row := range standardData {
		d, err := plc.NewDigmon(
			row[digmonTag],
			row[digmonDescription],
			row[digmonAddress],
			row[digmonInvert],
			row[digmonAlarm],
			row[digmonInvertAlarm],
			customMaps[n],
		)
		if err != nil {
			logger.Sugar.Errorw(err.Error(),
				"digmon", row[digmonTag],
			)
		} else {
			o = append(o, d)
			logger.Sugar.Infow("Object added to generator",
				"digmon", d.Tag)
		}
	}
	return
}

// ReadValves reads the "Measmon" sheet and returns a slice of PLC objects containing the created valves
func ReadValves(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetValves)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standardData := getStandardData(table, len(valveCols))
	if len(standardData) == 0 {
		return
	}
	customColumns, customData := getCustomData(table, len(valveCols))
	customMaps := make([]map[string]string, len(standardData))
	makeCustomDataMap(customColumns, customData, &customMaps)
	for n, row := range standardData {
		v, err := plc.NewValve(
			row[valveTag],
			row[valveDescription],
			row[valveOutputAddress],
			row[valveFeedbackOpenTag],
			row[valveFeedbackClosedTag],
			row[valveFeedbackOpenAddress],
			row[valveFeedbackClosedAddress],
			row[valveMonitoringTimeOpen],
			row[valveMonitoringTimeClose],
			customMaps[n],
		)
		if err != nil {
			logger.Sugar.Errorw(err.Error(),
				"valve", row[valveTag],
			)
		} else {
			o = append(o, v)
			logger.Sugar.Infow("Object added to generator",
				"valve", v.Tag)
		}
	}
	return
}

// ReadControlValves reads the "Measmon" sheet and returns a slice of PLC objects containing the created control valves
func ReadControlValves(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetControlValves)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standardData := getStandardData(table, len(controlValveCols))
	if len(standardData) == 0 {
		return
	}
	customColumns, customData := getCustomData(table, len(controlValveCols))
	customMaps := make([]map[string]string, len(standardData))
	makeCustomDataMap(customColumns, customData, &customMaps)
	for n, row := range standardData {
		c, err := plc.NewControlValve(
			row[controlValveTag],
			row[controlValveDescription],
			row[controlValveOutput],
			row[controlValveOutputAddress],
			row[controlValveFeedbackTag],
			row[controlValveFeedbackAddress],
			row[controlValveMonitoringTime],
			customMaps[n],
		)
		if err != nil {
			logger.Sugar.Errorw(err.Error(),
				"control valve", row[controlValveTag],
			)
		} else {
			o = append(o, c)
			logger.Sugar.Infow("Object added to generator",
				"control valve", c.Tag)
		}
	}
	return
}

// ReadMotors reads the "Measmon" sheet and returns a slice of PLC objects containing the created motors
func ReadMotors(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetMotors)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standardData := getStandardData(table, len(motorCols))
	if len(standardData) == 0 {
		return
	}
	customColumns, customData := getCustomData(table, len(motorCols))
	customMaps := make([]map[string]string, len(standardData))
	makeCustomDataMap(customColumns, customData, &customMaps)
	for n, row := range standardData {
		m, err := plc.NewMotor(
			row[motorTag],
			row[motorDescription],
			row[motorOutputAddress],
			row[motorFeedbackTag],
			row[motorFeedbackAddress],
			row[motorBreakerTag],
			row[motorBreakerAddress],
			row[motorSwitchTag],
			row[motorSwitchAddress],
			customMaps[n],
		)
		if err != nil {
			logger.Sugar.Errorw(err.Error(),
				"motor", row[motorTag])
		} else {
			o = append(o, m)
			logger.Sugar.Infow("Object added to generator",
				"motor", m.Tag)
		}
	}
	return
}

// ReadFreqMotors reads the "Measmon" sheet and returns a slice of PLC objects containing the created frequency motors
func ReadFreqMotors(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetFreqMotors)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standardData := getStandardData(table, len(freqMotorCols))
	if len(standardData) == 0 {
		return
	}
	customColumns, customData := getCustomData(table, len(freqMotorCols))
	customMaps := make([]map[string]string, len(standardData))
	makeCustomDataMap(customColumns, customData, &customMaps)
	for n, row := range standardData {
		fm, err := plc.NewFreqMotor(
			row[freqMotorTag],
			row[freqMotorDescription],
			row[freqMotorOutputAddress],
			row[freqMotorPqwAddress],
			row[freqMotorFeedbackTag],
			row[freqMotorFeedbackAddress],
			row[freqMotorBreakerTag],
			row[freqMotorBreakerAddress],
			row[freqMotorSwitchTag],
			row[freqMotorSwitchAddress],
			row[freqMotorAlarmTag],
			row[freqMotorAlarmAddress],
			row[freqMotorDanfoss],
			customMaps[n],
		)
		if err != nil {
			logger.Sugar.Errorw(err.Error(),
				"freqMotor", row[freqMotorTag])
		} else {
			o = append(o, fm)
			logger.Sugar.Infow("Object added to generator",
				"freqency motor", fm.Tag)
		}
	}
	return
}
