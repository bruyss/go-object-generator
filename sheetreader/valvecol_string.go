// Code generated by "stringer -type=valveCol -trimprefix=valve"; DO NOT EDIT.

package sheetreader

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[valveTag-0]
	_ = x[valveDescription-1]
	_ = x[valveOutput-2]
	_ = x[valveOutputAddress-3]
	_ = x[valveFeedbackOpenTag-4]
	_ = x[valveFeedbackOpenAddress-5]
	_ = x[valveFeedbackClosedTag-6]
	_ = x[valveFeedbackClosedAddress-7]
	_ = x[valveMonitoringTimeOpen-8]
	_ = x[valveMonitoringTimeClose-9]
}

const _valveCol_name = "TagDescriptionOutputOutputAddressFeedbackOpenTagFeedbackOpenAddressFeedbackClosedTagFeedbackClosedAddressMonitoringTimeOpenMonitoringTimeClose"

var _valveCol_index = [...]uint8{0, 3, 14, 20, 33, 48, 67, 84, 105, 123, 142}

func (i valveCol) String() string {
	if i < 0 || i >= valveCol(len(_valveCol_index)-1) {
		return "valveCol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _valveCol_name[_valveCol_index[i]:_valveCol_index[i+1]]
}
