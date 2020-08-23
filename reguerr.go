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
	code       string        // error code that you can define each error for your error handling.
	level      Level         // error Level. default:error
	statusCode int           // set http-status-code or exit-code. default:500
	format     string        // error message template. You can user fmt package placeholder style
	args       []interface{} // message argument
	err        error         // wrapped error that you hope
}

func (e *Error) WithArgs(args ...interface{}) *Error {
	return &Error{
		code:       e.code,
		level:      e.level,
		statusCode: e.statusCode,
		format:     e.format,
		args:       args,
		err:        e.err,
	}
}

func (e *Error) WithError(err error) *Error {
	return &Error{
		code:       e.code,
		level:      e.level,
		statusCode: e.statusCode,
		format:     e.format,
		args:       e.args,
		err:        err,
	}
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) StatusCode() int {
	return e.statusCode
}

func (e *Error) Message() string {
	return fmt.Sprintf(e.format, e.args)
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[%s]%s: %v", e.code, e.Message(), e.err)
	}
	return fmt.Sprintf("[%s]%s", e.code, e.Message())
}
func (e *Error) IsTraceLevel() bool {
	return e.level == TraceLevel
}

func (e *Error) IsDebugLevel() bool {
	return e.level == DebugLevel
}

func (e *Error) IsInfoLevel() bool {
	return e.level == InfoLevel
}

func (e *Error) IsWarnLevel() bool {
	return e.level == WarnLevel
}

func (e *Error) IsErrorLevel() bool {
	return e.level == ErrorLevel
}

func (e *Error) IsFatalLevel() bool {
	return e.level == FatalLevel
}

func (e *Error) withLevel(lvl Level) *Error {
	return &Error{
		code:       e.code,
		level:      lvl,
		statusCode: e.statusCode,
		format:     e.format,
		err:        e.err,
	}
}
