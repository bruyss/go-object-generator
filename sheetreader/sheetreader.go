package sheetreader

import (
	"github.com/bruyss/go-object-generator/logger"
	"github.com/bruyss/go-object-generator/plc"
	"github.com/xuri/excelize/v2"
)

func getTable(f *excelize.File, sheet_name string) ([][]string, error) {
	rows, err := f.GetRows(sheet_name)
	if err != nil {
		return nil, err
	}
	logger.Sugar.Debugw("Column names",
		"sheet name", sheet_name,
		"columns", rows[0])
	table := make([][]string, 0)
	for i, row := range rows[1:] {
		diff := len(rows[0]) - len(row)
		for j := 0; j < diff; j++ {
			row = append(row, "")
		}
		rows[i] = row
		table = append(table, row)
	}
	return table, nil
}

func getStandardData(table [][]string, nr_columns int) ([]string, [][]string) {
	column_names := table[0][:nr_columns]
	standard_data := make([][]string, len(table)-1)
	for i := range standard_data {
		standard_data[i] = table[i+1][:nr_columns]
	}
	return column_names, standard_data
}

func getCustomData(table [][]string, nr_columns int) ([]string, [][]string) {
	column_names := table[0][nr_columns:]
	if len(table[0]) == nr_columns {
		return column_names, [][]string{}
	}
	custom_data := make([][]string, len(table)-1)
	for i := range custom_data {
		custom_data[i] = table[i+1][nr_columns:]
	}
	return column_names, custom_data
}

func makeCustomDataMap(column_names []string, data [][]string) []map[string]string {
	data_map := make([]map[string]string, len(data))
	for i := range data_map {
		data_map[i] = make(map[string]string)
		for j, column := range column_names {
			data_map[i][column] = data[i][j]
		}
	}
	return data_map
}

func ReadMeasmons(f *excelize.File) (o []plc.PlcObject) {
	table, err := getTable(f, sheetMeasmons)
	if err != nil {
		logger.Sugar.Fatalln(err)
	}
	_, standard_data := getStandardData(table, len(measmonCols))
	custom_columns, custom_data := getCustomData(table, len(measmonCols))
	custom_map := makeCustomDataMap(custom_columns, custom_data)
	for n, row := range standard_data {
		m, err := plc.NewMeasmon(
			row[measmonTag],
			row[measmonDescription],
			row[measmonUnit],
			row[measmonAddress],
			row[measmonDirect],
			row[measmonMin],
			row[measmonMax],
			custom_map[n],
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
	for _, row := range table {
		d, err := plc.NewDigmon(
			row[digmonTag],
			row[digmonDescription],
			row[digmonAddress],
			row[digmonInvert],
			row[digmonAlarm],
			row[digmonInvertAlarm],
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
	for _, row := range table {
		c, err := plc.NewControlValve(
			row[controlValveTag],
			row[controlValveDescription],
			row[controlValveAddress],
			row[controlValveFeedbackTag],
			row[controlValveFeedbackAddress],
			row[controlValveMonitoringTime],
		)
		if err != nil {
			logger.Sugar.Errorw(err.Error(),
				"valve", row[controlValveTag],
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
