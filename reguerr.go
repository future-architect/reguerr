package reguerr

import (
	"fmt"
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

type Error struct {
	Code       string        // error code that you can define each error for your error handling.
	Level      Level         // error Level. default:error
	StatusCode int           // set http-status-code or exit-code. default:500
	format     string        // error message template. You can user fmt package placeholder style
	args       []interface{} // message argument
	err        error         // wrapped error that you hope
}

func New(code, format string) *Error {
	return &Error{
		Code:       code,
		format:     format,
		Level:      DefaultErrorLevel,
		StatusCode: DefaultStatusCode,
	}
}

func (e *Error) Label(index int, name string, goType interface{}) *Error {
	// コード解析用の関数なので、内部的には何もしないしなくて良い
	return e
}

func (e *Error) DisableError() *Error {
	// 解析用途なのにで、何もしない
	return e
}

func (e *Error) TraceLevel() *Error {
	return e.withLevel(TraceLevel)
}

func (e *Error) DebugLevel() *Error {
	return e.withLevel(DebugLevel)
}

func (e *Error) InfoLevel() *Error {
	return e.withLevel(InfoLevel)
}

func (e *Error) WarnLevel() *Error {
	return e.withLevel(WarnLevel)
}

func (e *Error) ErrorLevel() *Error {
	return e.withLevel(ErrorLevel)
}

func (e *Error) FatalLevel() *Error {
	return e.withLevel(FatalLevel)
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

func (e *Error) WithStatusCode(statusCode int) *Error {
	return &Error{
		Code:       e.Code,
		Level:      e.Level,
		StatusCode: statusCode,
		format:     e.format,
		err:        e.err,
		args:       e.args,
	}
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
