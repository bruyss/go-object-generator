// Code generated by "stringer -type=freqMotorCol -trimprefix=freqMotor"; DO NOT EDIT.

package sheetreader

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[freqMotorTag-0]
	_ = x[freqMotorAddress-1]
	_ = x[freqMotorPqwAddress-2]
	_ = x[freqMotorDescription-3]
	_ = x[freqMotorDanfoss-4]
	_ = x[freqMotorFeedbackTag-5]
	_ = x[freqMotorFeedbackAddress-6]
	_ = x[freqMotorBreakerTag-7]
	_ = x[freqMotorBreakerAddress-8]
	_ = x[freqMotorSwitchTag-9]
	_ = x[freqMotorSwitchAddress-10]
	_ = x[freqMotorAlarmTag-11]
	_ = x[freqMotorAlarmAddress-12]
}

const _freqMotorCol_name = "TagAddressPqwAddressDescriptionDanfossFeedbackTagFeedbackAddressBreakerTagBreakerAddressSwitchTagSwitchAddressAlarmTagAlarmAddress"

var _freqMotorCol_index = [...]uint8{0, 3, 10, 20, 31, 38, 49, 64, 74, 88, 97, 110, 118, 130}

func (i freqMotorCol) String() string {
	if i < 0 || i >= freqMotorCol(len(_freqMotorCol_index)-1) {
		return "freqMotorCol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _freqMotorCol_name[_freqMotorCol_index[i]:_freqMotorCol_index[i+1]]
}
