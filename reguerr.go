package reguerr

import (
	"errors"
	"fmt"
	"strings"
)

var (
	DefaultErrorLevel = ErrorLevel
	DefaultStatusCode = 500
)

// Level represents error level that developer can handle error depending on each variable
//go:generate stringer -type=Level
type Level int

const (
	TraceLevel Level = iota + 1
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

func (i Level) Equals(arg Level) bool {
	return i.String() == arg.String()
}

func NewLevel(s string) (Level, error) {
	switch strings.ToLower(s) {
	case strings.ToLower(TraceLevel.String()):
		return TraceLevel, nil
	case strings.ToLower(DebugLevel.String()):
		return DebugLevel, nil
	case strings.ToLower(InfoLevel.String()):
		return InfoLevel, nil
	case strings.ToLower(WarnLevel.String()):
		return WarnLevel, nil
	case strings.ToLower(ErrorLevel.String()):
		return ErrorLevel, nil
	case strings.ToLower(FatalLevel.String()):
		return FatalLevel, nil
	default:
		return TraceLevel, errors.New("unknown error level")
	}
}

type Error struct {
	Code       string        // error code that you can define each error for your error handling.
	Level      Level         // error Level. default:error
	StatusCode int           // set http-status-code or exit-code. default:500
	format     string        // error message template. You can user fmt package placeholder style
	args       []interface{} // message argument
	err        error         // wrapped error that you hope
}

func (e *Error) WithArgs(args ...interface{}) *Error {
	return &Error{
		Code:       e.Code,
		Level:      e.Level,
		StatusCode: e.StatusCode,
		format:     e.format,
		args:       args,
		err:        e.err,
	}
}

func (e *Error) WithError(err error) *Error {
	return &Error{
		Code:       e.Code,
		Level:      e.Level,
		StatusCode: e.StatusCode,
		format:     e.format,
		args:       e.args,
		err:        err,
	}
}

func (e *Error) Message() string {
	return fmt.Sprintf(e.format, e.args)
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[%s]%s: %v", e.Code, e.Message(), e.err)
	}
	return fmt.Sprintf("[%s]%s", e.Code, e.Message())
}
func (e *Error) IsTraceLevel() bool {
	return e.Level == TraceLevel
}

func (e *Error) IsDebugLevel() bool {
	return e.Level == DebugLevel
}

func (e *Error) IsInfoLevel() bool {
	return e.Level == InfoLevel
}

func (e *Error) IsWarnLevel() bool {
	return e.Level == WarnLevel
}

func (e *Error) IsErrorLevel() bool {
	return e.Level == ErrorLevel
}

func (e *Error) IsFatalLevel() bool {
	return e.Level == FatalLevel
}

func (e *Error) withLevel(lvl Level) *Error {
	return &Error{
		Code:       e.Code,
		Level:      lvl,
		StatusCode: e.StatusCode,
		format:     e.format,
		err:        e.err,
	}
}
