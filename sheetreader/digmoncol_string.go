// Code generated by "stringer -type=digmonCol -trimprefix=digmon"; DO NOT EDIT.

package sheetreader

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[digmonTag-0]
	_ = x[digmonAddress-1]
	_ = x[digmonDescription-2]
	_ = x[digmonInvert-3]
	_ = x[digmonAlarm-4]
	_ = x[digmonInvertAlarm-5]
}

const _digmonCol_name = "TagAddressDescriptionInvertAlarmInvertAlarm"

var _digmonCol_index = [...]uint8{0, 3, 10, 21, 27, 32, 43}

func (i digmonCol) String() string {
	if i < 0 || i >= digmonCol(len(_digmonCol_index)-1) {
		return "digmonCol(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _digmonCol_name[_digmonCol_index[i]:_digmonCol_index[i+1]]
}
