// Code generated by "stringer -type=Level"; DO NOT EDIT.

package errcdgen

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TraceLevel-1]
	_ = x[DebugLevel-2]
	_ = x[InfoLevel-3]
	_ = x[WarnLevel-4]
	_ = x[ErrorLevel-5]
	_ = x[FatalLevel-6]
}

const _Level_name = "TraceLevelDebugLevelInfoLevelWarnLevelErrorLevelFatalLevel"

var _Level_index = [...]uint8{0, 10, 20, 29, 38, 48, 58}

func (i Level) String() string {
	i -= 1
	if i < 0 || i >= Level(len(_Level_index)-1) {
		return "Level(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Level_name[_Level_index[i]:_Level_index[i+1]]
}
