package sheetreader

const (
	sheetMeasmons      = "Measmons"
	sheetDigmons       = "Digmons"
	sheetValves        = "Valves"
	sheetControlValves = "ControlValves"
	sheetMotors        = "Motors"
	sheetFreqMotors    = "FreqMotors"
	sheetDigouts       = "DigitalOut"
)

type measmonCol int

const (
	measmonTag measmonCol = iota
	measmonDescription
	measmonAddress
	measmonUnit
	measmonDirect
	measmonMin
	measmonMax
)

//go:generate stringer -type=measmonCol -trimprefix=measmon

var measmonCols = []string{
	measmonTag.String(),
	measmonDescription.String(),
	measmonAddress.String(),
	measmonUnit.String(),
	measmonDirect.String(),
	measmonMin.String(),
	measmonMax.String(),
}

type motorCol int

const (
	motorTag motorCol = iota
	motorDescription
	motorOutput
	motorOutputAddress
	motorFeedbackTag
	motorFeedbackAddress
	motorBreakerTag
	motorBreakerAddress
	motorSwitchTag
	motorSwitchAddress
)

//go:generate stringer -type=motorCol -trimprefix=motor

var motorCols = []string{
	motorTag.String(),
	motorDescription.String(),
	motorOutput.String(),
	motorOutputAddress.String(),
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
	freqMotorDescription
	freqMotorOutput
	freqMotorOutputAddress
	freqMotorPqwAddress
	freqMotorDanfoss
	freqMotorFeedbackTag
	freqMotorFeedbackAddress
	freqMotorBreakerTag
	freqMotorBreakerAddress
	freqMotorSwitchTag
	freqMotorSwitchAddress
	freqMotorAlarmTag
	freqMotorAlarmAddress
)

//go:generate stringer -type=freqMotorCol -trimprefix=freqMotor

var freqMotorCols = []string{
	freqMotorTag.String(),
	freqMotorDescription.String(),
	freqMotorOutput.String(),
	freqMotorOutputAddress.String(),
	freqMotorPqwAddress.String(),
	freqMotorDanfoss.String(),
	freqMotorFeedbackTag.String(),
	freqMotorFeedbackAddress.String(),
	freqMotorBreakerTag.String(),
	freqMotorBreakerAddress.String(),
	freqMotorSwitchTag.String(),
	freqMotorSwitchAddress.String(),
	freqMotorAlarmTag.String(),
	freqMotorAlarmAddress.String(),
}

type digmonCol int

const (
	digmonTag digmonCol = iota
	digmonDescription
	digmonAddress
	digmonInvert
	digmonAlarm
	digmonInvertAlarm
)

//go:generate stringer -type=digmonCol -trimprefix=digmon

var digmonCols = []string{
	digmonTag.String(),
	digmonDescription.String(),
	digmonAddress.String(),
	digmonInvert.String(),
	digmonAlarm.String(),
	digmonInvertAlarm.String(),
}

type valveCol int

const (
	valveTag valveCol = iota
	valveDescription
	valveOutput
	valveOutputAddress
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
	valveOutput.String(),
	valveOutputAddress.String(),
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
	controlValveOutput
	controlValveOutputAddress
	controlValveFeedbackTag
	controlValveFeedbackAddress
	controlValveMonitoringTime
)

//go:generate stringer -type=controlValveCol -trimprefix=controlValve

var controlValveCols = []string{
	controlValveTag.String(),
	controlValveDescription.String(),
	controlValveOutput.String(),
	controlValveOutputAddress.String(),
	controlValveFeedbackTag.String(),
	controlValveFeedbackAddress.String(),
	controlValveMonitoringTime.String(),
}

type digoutCol int

const (
	digoutTag digoutCol = iota
	digoutDescription
	digoutOutput
	digoutOutputAddress
	digoutFeedbackTag
	digoutFeedbackAddress
	digoutBreakerTag
	digoutBreakerAddress
	digoutMonitoringTime
)

//go:generate stringer -type=digoutCol -trimprefix=digout

var digoutCols = []string{
	digoutTag.String(),
	digoutDescription.String(),
	digoutOutput.String(),
	digoutOutputAddress.String(),
	digoutFeedbackTag.String(),
	digoutFeedbackAddress.String(),
	digoutBreakerTag.String(),
	digoutBreakerAddress.String(),
	digoutMonitoringTime.String(),
}
