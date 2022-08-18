package sheetreader

const (
	sheetMeasmons      = "Measmons"
	sheetDigmons       = "Digmons"
	sheetValves        = "Valves"
	sheetControlValves = "ControlValves"
	sheetMotors        = "Motors"
	sheetFreqMotors    = "FreqMotors"
)

type measmonCol int

const (
	measmonTag measmonCol = iota
	measmonAddress
	measmonDescription
	measmonUnit
	measmonDirect
	measmonMin
	measmonMax
)

//go:generate stringer -type=measmonCol -trimprefix=measmon

var measmonCols = []string{
	measmonTag.String(),
	measmonAddress.String(),
	measmonDescription.String(),
	measmonUnit.String(),
	measmonDirect.String(),
	measmonMin.String(),
	measmonMax.String(),
}

type motorCol int

const (
	motorTag motorCol = iota
	motorAddress
	motorDescription
	motorFeedbackTag
	motorFeedbackAddress
	motorBreakerTag
	motorBreakerAddress
	motorSwitchTag
	motorSwitchAddress
)

//go:generate stringer -type=motorCol -trimprefix=motor

var motorCols = []string{
	motorAddress.String(),
	motorDescription.String(),
	motorFeedbackTag.String(),
	motorFeedbackAddress.String(),
	motorBreakerTag.String(),
	motorBreakerAddress.String(),
	motorSwitchTag.String(),
	motorSwitchAddress.String(),
}

type freqMotorCol int

const (
	freqMotorTag freqMotorCol = iota
	freqMotorAddress
	freqMotorDescription
	freqMotorDanfoss
	freqMotorFeedbackTag
	freqMotorFeedbackAddress
	freqMotorBreakerTag
	freqMotorBreakerAddress
	freqMotorSwitchTag
	freqMotorSwitchAddress
)

//go:generate stringer -type=freqMotorCol -trimprefix=freqMotor

var freqMotorCols = []string{
	freqMotorTag.String(),
	freqMotorAddress.String(),
	freqMotorDescription.String(),
	freqMotorDanfoss.String(),
	freqMotorFeedbackTag.String(),
	freqMotorFeedbackAddress.String(),
	freqMotorBreakerTag.String(),
	freqMotorBreakerAddress.String(),
	freqMotorSwitchTag.String(),
	freqMotorSwitchAddress.String(),
}

type digmonCol int

const (
	digmonTag digmonCol = iota
	digmonAddress
	digmonDescription
	digmonInvert
	digmonAlarm
	digmonInvertAlarm
)

//go:generate stringer -type=digmonCol -trimprefix=digmon

var digmonCols = []string{
	digmonTag.String(),
	digmonAddress.String(),
	digmonDescription.String(),
	digmonInvert.String(),
	digmonAlarm.String(),
	digmonInvertAlarm.String(),
}

type valveCol int

const (
	valveTag valveCol = iota
	valveDescription
	valveAddress
	valveFeedbackOpenTag
	valveFeedbackOpenAddress
	valveFeedbackClosedTag
	valveFeedbackClosedAddress
	valveMonitoringTimeOpen
	valveMonitoringTimeClose
)

//go:generate stringer -type=valveCol -trimprefix=valve

var valveCols = []string{
	valveTag.String(),
	valveDescription.String(),
	valveAddress.String(),
	valveFeedbackOpenTag.String(),
	valveFeedbackOpenAddress.String(),
	valveFeedbackClosedTag.String(),
	valveFeedbackClosedAddress.String(),
	valveMonitoringTimeOpen.String(),
	valveMonitoringTimeClose.String(),
}

type controlValveCol int

const (
	controlValveTag controlValveCol = iota
	controlValveDescription
	controlValveAddress
	controlValveFeedbackTag
	controlValveFeedbackAddress
	controlValveMonitoringTime
)

//go:generate stringer -type=controlValveCol -trimprefix=controlValve

var controlValveCols = []string{
	controlValveTag.String(),
	controlValveDescription.String(),
	controlValveAddress.String(),
	controlValveFeedbackTag.String(),
	controlValveFeedbackAddress.String(),
	controlValveMonitoringTime.String(),
}
