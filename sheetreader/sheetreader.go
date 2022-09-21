package sheetreader

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bruyss/go-object-generator/plc"
	"github.com/bruyss/go-object-generator/utils"
)

func ReadMeasmons(f *excelize.File) (o []plc.PlcObject) {
	for _, row := range f.GetRows(sheetMeasmons)[1:] {
		m, err := plc.NewMeasmon(
			row[measmonTag],
			row[measmonDescription],
			row[measmonUnit],
			row[measmonAddress],
			row[measmonDirect],
			row[measmonMin],
			row[measmonMax],
		)
		if err != nil {
			utils.Sugar.Errorw(err.Error(),
				"measmon", row[measmonTag],
			)
		} else {
			o = append(o, m)
		}
	}
	return
}

func ReadDigmons(f *excelize.File) (o []plc.PlcObject) {
	for _, row := range f.GetRows(sheetDigmons)[1:] {
		d, err := plc.NewDigmon(
			row[digmonTag],
			row[digmonDescription],
			row[digmonAddress],
			row[digmonInvert],
			row[digmonAlarm],
			row[digmonInvertAlarm],
		)
		if err != nil {
			utils.Sugar.Errorw(err.Error(),
				"digmon", row[digmonTag],
			)
		} else {
			o = append(o, d)
		}
	}
	return
}
func ReadValves(f *excelize.File) (o []plc.PlcObject) {
	for _, row := range f.GetRows(sheetValves)[1:] {
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
			utils.Sugar.Errorw(err.Error(),
				"valve", row[valveTag],
			)
		} else {
			o = append(o, v)
		}
	}
	return
}
func ReadControlValves(f *excelize.File) (o []plc.PlcObject) {
	for _, row := range f.GetRows(sheetControlValves)[1:] {
		c, err := plc.NewControlValve(
			row[controlValveTag],
			row[controlValveDescription],
			row[controlValveAddress],
			row[controlValveFeedbackTag],
			row[controlValveFeedbackAddress],
			row[controlValveMonitoringTime],
		)
		if err != nil {
			utils.Sugar.Errorw(err.Error(),
				"valve", row[controlValveTag],
			)
		} else {
			o = append(o, c)
		}
	}
	return
}
func ReadMotors(f *excelize.File) (o []plc.PlcObject) {
	for _, row := range f.GetRows(sheetMotors)[1:] {
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
			utils.Sugar.Errorw(err.Error(),
				"motor", row[motorTag])
		} else {
			o = append(o, m)
		}
	}
	return
}
func ReadFreqMotors(f *excelize.File) (o []plc.PlcObject) {
	for _, row := range f.GetRows(sheetFreqMotors)[1:] {
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
			utils.Sugar.Errorw(err.Error(),
				"freqMotor", row[freqMotorTag])
		} else {
			o = append(o, fm)
		}
	}
	return
}
