// Code generated by "stringer -type=controlValveCol -trimprefix=controlValve"; DO NOT EDIT.

package sheetreader

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[controlValveTag-0]
	_ = x[controlValveDescription-1]
	_ = x[controlValveOutput-2]
	_ = x[controlValveOutputAddress-3]
	_ = x[controlValveFeedbackTag-4]
	_ = x[controlValveFeedbackAddress-5]
	_ = x[controlValveMonitoringTime-6]
}

const _controlValveCol_name = "TagDescriptionOutputOutputAddressFeedbackTagFeedbackAddressMonitoringTime"

var _controlValveCol_index = [...]uint8{0, 3, 14, 20, 33, 44, 59, 73}

func (i controlValveCol) String() string {
	if i < 0 || i >= controlValveCol(len(_controlValveCol_index)-1) {
		return "controlValveCol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _controlValveCol_name[_controlValveCol_index[i]:_controlValveCol_index[i+1]]
}
