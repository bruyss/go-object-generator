// Code generated by "stringer -type=freqMotorCol -trimprefix=freqMotor"; DO NOT EDIT.

package sheetreader

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[freqMotorTag-0]
	_ = x[freqMotorAddress-1]
	_ = x[freqMotorDescription-2]
	_ = x[freqMotorDanfoss-3]
	_ = x[freqMotorFeedbackTag-4]
	_ = x[freqMotorFeedbackAddress-5]
	_ = x[freqMotorBreakerTag-6]
	_ = x[freqMotorBreakerAddress-7]
	_ = x[freqMotorSwitchTag-8]
	_ = x[freqMotorSwitchAddress-9]
}

const _freqMotorCol_name = "TagAddressDescriptionDanfossFeedbackTagFeedbackAddressBreakerTagBreakerAddressSwitchTagSwitchAddress"

var _freqMotorCol_index = [...]uint8{0, 3, 10, 21, 28, 39, 54, 64, 78, 87, 100}

func (i freqMotorCol) String() string {
	if i < 0 || i >= freqMotorCol(len(_freqMotorCol_index)-1) {
		return "freqMotorCol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _freqMotorCol_name[_freqMotorCol_index[i]:_freqMotorCol_index[i+1]]
}