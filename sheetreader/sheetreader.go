package sheetreader

import (
	"github.com/bruyss/go-object-generator/logger"
	"github.com/bruyss/go-object-generator/plc"
	"github.com/xuri/excelize/v2"
)

// getTable reads in an excel file and returns a table starting on cell A1.
// Appends empty string in the case of empty last columns for certain rows to get around default behaviour.
func getTable(f *excelize.File, sheet_name string) ([][]string, error) {
	rows, err := f.GetRows(sheet_name)
	if err != nil {
		return nil, err
	}
	logger.Sugar.Debugw("Column names",
		"sheet name", sheet_name,
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
func getStandardData(table [][]string, nr_standard_columns int) ([]string, [][]string) {
	column_names := table[0][:nr_standard_columns]
	standard_data := make([][]string, len(table)-1)
	for i := range standard_data {
		standard_data[i] = table[i+1][:nr_standard_columns]
	}
	return column_names, standard_data
}

// getCustomData reads a table of PLC objects and returns the columns not defined in the default per object.
// These columns should appear after the default columns.
// Returns the column names and values as seperate values.
func getCustomData(table [][]string, nr_standard_columns int) ([]string, [][]string) {
	column_names := table[0][nr_standard_columns:]
	if len(table[0]) == nr_standard_columns {
		return column_names, [][]string{}
	}
	custom_data := make([][]string, len(table)-1)
	for i := range custom_data {
		custom_data[i] = table[i+1][nr_standard_columns:]
	}
	return column_names, custom_data
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
func makeCustomDataMap(column_names []string, data [][]string, data_map *[]map[string]string) {
	for i := range *data_map {
		(*data_map)[i] = make(map[string]string, len(column_names))
		for j, column := range column_names {
			// fmt.Printf("Row %d column %d-%s: %s", i, j, column, data[i][j])
			(*data_map)[i][column] = data[i][j]
		}
	}
	// fmt.Println(*data_map)
}

func ReadMeasmons(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetMeasmons)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standard_data := getStandardData(table, len(measmonCols))
	custom_columns, custom_data := getCustomData(table, len(measmonCols))
	custom_maps := make([]map[string]string, len(standard_data))
	makeCustomDataMap(custom_columns, custom_data, &custom_maps)
	for n, row := range standard_data {
		m, err := plc.NewMeasmon(
			row[measmonTag],
			row[measmonDescription],
			row[measmonUnit],
			row[measmonAddress],
			row[measmonDirect],
			row[measmonMin],
			row[measmonMax],
			custom_maps[n],
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

func ReadDigmons(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetDigmons)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standard_data := getStandardData(table, len(digmonCols))
	custom_columns, custom_data := getCustomData(table, len(digmonCols))
	custom_maps := make([]map[string]string, len(standard_data))
	makeCustomDataMap(custom_columns, custom_data, &custom_maps)
	for n, row := range standard_data {
		d, err := plc.NewDigmon(
			row[digmonTag],
			row[digmonDescription],
			row[digmonAddress],
			row[digmonInvert],
			row[digmonAlarm],
			row[digmonInvertAlarm],
			custom_maps[n],
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

func ReadValves(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetValves)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	for _, row := range table {
		v, err := plc.NewValve(
			row[valveTag],
			row[valveDescription],
			row[valveAddress],
			row[valveFeedbackOpenTag],
			row[valveFeedbackClosedTag],
			row[valveFeedbackOpenAddress],
			row[valveFeedbackClosedAddress],
			row[valveMonitoringTimeOpen],
			row[valveMonitoringTimeClose],
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

func ReadControlValves(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetControlValves)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standard_data := getStandardData(table, len(controlValveCols))
	custom_columns_names, custom_data := getCustomData(table, len(controlValveCols))
	custom_maps := make([]map[string]string, len(standard_data))
	makeCustomDataMap(custom_columns_names, custom_data, &custom_maps)
	for n, row := range standard_data {
		c, err := plc.NewControlValve(
			row[controlValveTag],
			row[controlValveDescription],
			row[controlValveAddress],
			row[controlValveFeedbackTag],
			row[controlValveFeedbackAddress],
			row[controlValveMonitoringTime],
			custom_maps[n],
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

func ReadMotors(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetMotors)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	for _, row := range table {
		m, err := plc.NewMotor(
			row[motorTag],
			row[motorDescription],
			row[motorAddress],
			row[motorFeedbackTag],
			row[motorFeedbackAddress],
			row[motorBreakerTag],
			row[motorBreakerAddress],
			row[motorSwitchTag],
			row[motorSwitchAddress],
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

func ReadFreqMotors(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetFreqMotors)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	for _, row := range table {
		fm, err := plc.NewFreqMotor(
			row[freqMotorTag],
			row[freqMotorDescription],
			row[freqMotorAddress],
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
