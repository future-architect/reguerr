package errcdgen

import (
	"fmt"
)

type level int

// 書き換え用にExportしておく
var (
	DefaultErrorLevel = ErrorLevel
	DefaultExitCode   = 1
)

const (
	TraceLevel level = iota + 1
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

type CodeError struct {
	Code     string // error code
	Level    level  // default-level:error
	ExitCode int    // set http-status-code or exit-code. default:1
	format   string
	args     []interface{}
	err      error
}

type Arg struct {
	name   string
	goType string
}

func NewCodeError(code, format string) *CodeError {
	return &CodeError{
		Code:     code,
		Level:    DefaultErrorLevel,
		format:   format,
		ExitCode: DefaultExitCode,
	}
}

func (e *CodeError) Arg(name string, goType interface{}) *CodeError {
	// コード解析用の関数なので、内部的には何もしないしなくて良い
	return e
}

func (e *CodeError) ArgPath(name, goType string) *CodeError {
	// コード解析用の関数なので、内部的には何もしないしなくて良い
	return e
}

func (e *CodeError) DisableError() *CodeError {
	// 解析用途なのにで、何もしない
	return e
}

func (e *CodeError) Debug() *CodeError {
	// 解析用途なのにで、何もしない
	return e
}

func (e *CodeError) WithLevel(lvl level) *CodeError {
	return &CodeError{
		Code:     e.Code,
		Level:    lvl,
		ExitCode: e.ExitCode,
		format:   e.format,
		err:      e.err,
	}
}

func (e *CodeError) WithExitCode(exitCode int) *CodeError {
	return &CodeError{
		Code:     e.Code,
		Level:    e.Level,
		ExitCode: exitCode,
		format:   e.format,
		err:      e.err,
	}
}

func (e *CodeError) WithArgs(args ...interface{}) *CodeError {
	return &CodeError{
		Code:   e.Code,
		Level:  e.Level,
		format: e.format,
		args:   args,
	}
}

func (e *CodeError) WithError(err error) *CodeError {
	return &CodeError{
		Code:   e.Code,
		Level:  e.Level,
		format: e.format,
		args:   e.args,
		err:    err,
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
