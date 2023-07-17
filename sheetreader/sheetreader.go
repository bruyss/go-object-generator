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

// readObjects generates a function to read an object type from a specified worksheet
func readObjects(objectName, sheetName string, columns []string, makeFunc func([]string, map[string]string) (plc.PlcObject, error)) func(*excelize.File) []plc.PlcObject {
	return func(f *excelize.File) []plc.PlcObject {
		table, err := getTable(f, sheetName)
		if err != nil {
			logger.Sugar.Fatalln(err)
		}
		_, standardData := getStandardData(table, len(columns))
		logger.Sugar.Debugf("Standard data length %d", len(standardData))

		if len(standardData) == 0 {
			return []plc.PlcObject{}
		}
		if standardData[0][0] == "" { // Predefined formulas can cause reader to think there is an object in the first row without a tag name
			return []plc.PlcObject{}
		}

		objects := make([]plc.PlcObject, len(standardData))

		customColumns, customData := getCustomData(table, len(columns))
		customMaps := make([]map[string]string, len(standardData))
		makeCustomDataMap(customColumns, customData, &customMaps)
		for n, row := range standardData {
			object, err := makeFunc(row, customMaps[n])
			if err != nil {
				logger.Sugar.Errorw(err.Error(),
					objectName, row[0]) // row 0 should contain the tag name
			} else {
				// objects = append(objects, object)
				objects[n] = object
				logger.Sugar.Infow("Object added to generator",
					objectName, object)
			}
		}
		return objects
	}
}

// ReadMeasmons reads the "Measmon" sheet and returns a slice of PLC objects containing the created measmons
var ReadMeasmons = readObjects(
	"measmon",
	sheetMeasmons,
	measmonCols,
	func(standard []string, custom map[string]string) (plc.PlcObject, error) {
		return plc.NewMeasmon(
			standard[measmonTag],
			standard[measmonDescription],
			standard[measmonUnit],
			standard[measmonAddress],
			standard[measmonDirect],
			standard[measmonMin],
			standard[measmonMax],
			custom,
		)
	},
)

// ReadDigmons reads the "Measmon" sheet and returns a slice of PLC objects containing the created digmons
var ReadDigmons = readObjects(
	"digmon",
	sheetDigmons,
	digmonCols,
	func(standard []string, custom map[string]string) (plc.PlcObject, error) {
		return plc.NewDigmon(
			standard[digmonTag],
			standard[digmonDescription],
			standard[digmonAddress],
			standard[digmonInvert],
			standard[digmonAlarm],
			standard[digmonInvertAlarm],
			custom,
		)
	},
)

// ReadValves reads the "Measmon" sheet and returns a slice of PLC objects containing the created valves
var ReadValves = readObjects(
	"valve",
	sheetValves,
	valveCols,
	func(standard []string, custom map[string]string) (plc.PlcObject, error) {
		return plc.NewValve(
			standard[valveTag],
			standard[valveDescription],
			standard[valveOutputAddress],
			standard[valveFeedbackOpenTag],
			standard[valveFeedbackClosedTag],
			standard[valveFeedbackOpenAddress],
			standard[valveFeedbackClosedAddress],
			standard[valveMonitoringTimeOpen],
			standard[valveMonitoringTimeClose],
			custom,
		)
	},
)

// ReadControlValves reads the "Measmon" sheet and returns a slice of PLC objects containing the created control valves
var ReadControlValves = readObjects(
	"control valve",
	sheetControlValves,
	controlValveCols,
	func(standard []string, custom map[string]string) (plc.PlcObject, error) {
		return plc.NewControlValve(
			standard[controlValveTag],
			standard[controlValveDescription],
			standard[controlValveOutput],
			standard[controlValveOutputAddress],
			standard[controlValveFeedbackTag],
			standard[controlValveFeedbackAddress],
			standard[controlValveMonitoringTime],
			custom,
		)
	},
)

// ReadMotors reads the "Measmon" sheet and returns a slice of PLC objects containing the created motors
var ReadMotors = readObjects(
	"motor",
	sheetMotors,
	motorCols,
	func(standard []string, custom map[string]string) (plc.PlcObject, error) {
		return plc.NewMotor(
			standard[motorTag],
			standard[motorDescription],
			standard[motorOutputAddress],
			standard[motorFeedbackTag],
			standard[motorFeedbackAddress],
			standard[motorBreakerTag],
			standard[motorBreakerAddress],
			standard[motorSwitchTag],
			standard[motorSwitchAddress],
			custom,
		)
	},
)

// ReadFreqMotors reads the "DigitalOut" sheet and returns a slice of PLC objects containing the created frequency motors
var ReadFreqMotors = readObjects(
	"frequency motor",
	sheetFreqMotors,
	freqMotorCols,
	func(standard []string, custom map[string]string) (plc.PlcObject, error) {
		return plc.NewFreqMotor(
			standard[freqMotorTag],
			standard[freqMotorDescription],
			standard[freqMotorOutputAddress],
			standard[freqMotorPqwAddress],
			standard[freqMotorFeedbackTag],
			standard[freqMotorFeedbackAddress],
			standard[freqMotorBreakerTag],
			standard[freqMotorBreakerAddress],
			standard[freqMotorSwitchTag],
			standard[freqMotorSwitchAddress],
			standard[freqMotorAlarmTag],
			standard[freqMotorAlarmAddress],
			standard[freqMotorDanfoss],
			custom,
		)
	},
)

// ReadDigouts reads the "DigitalOuts" sheet and returns a slice of PLC objects containing the generated digital outs
var ReadDigouts = readObjects(
	"digital out",
	sheetDigouts,
	digoutCols,
	func(standard []string, custom map[string]string) (plc.PlcObject, error) {
		return plc.NewDigout(
			standard[digoutTag],
			standard[digoutDescription],
			standard[digoutOutputAddress],
			standard[digoutFeedbackTag],
			standard[digoutFeedbackAddress],
			standard[digoutBreakerTag],
			standard[digoutBreakerAddress],
			standard[digoutMonitoringTime],
			custom,
		)
	},
)
