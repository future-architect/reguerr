package errcdgen

import (
	"fmt"
)

var (
	DefaultErrorLevel = ErrorLevel
	DefaultStatusCode = 500
)

// Level represents error level that developer can handle error depending on each variable
type Level int

const (
	TraceLevel Level = iota + 1
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type CodeError struct {
	Code       string        // error code that you can define each error for your error handling.
	Level      Level         // error Level. default:error
	StatusCode int           // set http-status-code or exit-code. default:500
	format     string        // error message template. You can user fmt package placeholder style
	args       []interface{} // message argument
	err        error         // wrapped error that you hope
}

func NewCodeError(code, format string) *CodeError {
	return &CodeError{
		Code:       code,
		format:     format,
		Level:      DefaultErrorLevel,
		StatusCode: DefaultStatusCode,
	}
}

func (e *CodeError) Label(index int, name string, goType interface{}) *CodeError {
	// コード解析用の関数なので、内部的には何もしないしなくて良い
	return e
}

func (e *CodeError) DisableError() *CodeError {
	// 解析用途なのにで、何もしない
	return e
}

func (e *CodeError) TraceLevel() *CodeError {
	return e.withLevel(TraceLevel)
}

func (e *CodeError) DebugLevel() *CodeError {
	return e.withLevel(DebugLevel)
}

func (e *CodeError) InfoLevel() *CodeError {
	return e.withLevel(InfoLevel)
}

func (e *CodeError) WarnLevel() *CodeError {
	return e.withLevel(WarnLevel)
}

func (e *CodeError) ErrorLevel() *CodeError {
	return e.withLevel(ErrorLevel)
}

func (e *CodeError) FatalLevel() *CodeError {
	return e.withLevel(FatalLevel)
}

func (e *CodeError) withLevel(lvl Level) *CodeError {
	return &CodeError{
		Code:       e.Code,
		Level:      lvl,
		StatusCode: e.StatusCode,
		format:     e.format,
		err:        e.err,
	}
}

func (e *CodeError) WithStatusCode(statusCode int) *CodeError {
	return &CodeError{
		Code:       e.Code,
		Level:      e.Level,
		StatusCode: statusCode,
		format:     e.format,
		err:        e.err,
		args:       e.args,
	}
}

func (e *CodeError) WithArgs(args ...interface{}) *CodeError {
	return &CodeError{
		Code:       e.Code,
		Level:      e.Level,
		StatusCode: e.StatusCode,
		format:     e.format,
		args:       args,
		err:        e.err,
	}
}

func (e *CodeError) WithError(err error) *CodeError {
	return &CodeError{
		Code:       e.Code,
		Level:      e.Level,
		StatusCode: e.StatusCode,
		format:     e.format,
		args:       e.args,
		err:        err,
	}
}

func (e *CodeError) Message() string {
	return fmt.Sprintf(e.format, e.args)
}

func (e *CodeError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[%s]%s: %v", e.Code, e.Message(), e.err)
	}
	return fmt.Sprintf("[%s]%s", e.Code, e.Message())
}
