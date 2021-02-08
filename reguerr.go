package reguerr

import (
	"errors"
	"fmt"
	"strings"
)

var (
	DefaultErrorLevel = Error
	DefaultStatusCode = 500
)

// Level represents error level that developer can handle error depending on each variable
//go:generate stringer -type=Level
type Level int

const (
	Trace Level = iota + 1
	Debug
	Info
	Warn
	Error
	Fatal
	Unknown
)

func (i Level) Equals(arg Level) bool {
	return i.String() == arg.String()
}

func NewLevel(s string) (Level, error) {
	switch strings.ToLower(s) {
	case strings.ToLower(Trace.String()):
		return Trace, nil
	case strings.ToLower(Debug.String()):
		return Debug, nil
	case strings.ToLower(Info.String()):
		return Info, nil
	case strings.ToLower(Warn.String()):
		return Warn, nil
	case strings.ToLower(Error.String()):
		return Error, nil
	case strings.ToLower(Fatal.String()):
		return Fatal, nil
	default:
		return Trace, errors.New("unknown error level")
	}
}

func ErrorOf(err error) (*ReguError, bool) {
	var cerr *ReguError
	if as := errors.As(err, &cerr); as {
		return cerr, true
	}
	return nil, false
}

func CodeOf(err error) (string, bool) {
	var cerr *ReguError
	if as := errors.As(err, &cerr); as {
		return cerr.Code(), true
	}
	return "", false
}

func LevelOf(err error) (Level, bool) {
	var cerr *ReguError
	if as := errors.As(err, &cerr); as {
		return cerr.Level(), true
	}
	return Unknown, false
}

func StatusOf(err error) (int, bool) {
	var cerr *ReguError
	if as := errors.As(err, &cerr); as {
		return cerr.StatusCode(), true
	}
	return 0, false
}

type Code string

type ReguError struct {
	code       string        // error code that you can define each error for your error handling.
	level      Level         // error Level. default:error
	statusCode int           // set http-status-code or exit-code. default:500
	format     string        // error message template. You can user fmt package placeholder style
	args       []interface{} // message argument
	err        error         // wrapped error that you hope
}

// ReguError must meet error interface
var _ error = &ReguError{}

func (e *ReguError) Unwrap() error {
	return e.err
}

func (e *ReguError) WithArgs(args ...interface{}) *ReguError {
	return &ReguError{
		code:       e.code,
		level:      e.level,
		statusCode: e.statusCode,
		format:     e.format,
		args:       args,
		err:        e.err,
	}
}

func (e *ReguError) WithError(err error) *ReguError {
	return &ReguError{
		code:       e.code,
		level:      e.level,
		statusCode: e.statusCode,
		format:     e.format,
		args:       e.args,
		err:        err,
	}
}

func (e *ReguError) Code() string {
	return e.code
}

func (e *ReguError) StatusCode() int {
	return e.statusCode
}

func (e *ReguError) Level() Level {
	return e.level
}

func (e *ReguError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.code, e.message(), e.err)
	}
	return fmt.Sprintf("[%s] %s", e.code, e.message())
}

func (e *ReguError) message() string {
	if e.args != nil {
		return fmt.Sprintf(e.format, e.args)
	}
	return e.format
}

func (e *ReguError) IsTraceLevel() bool {
	return e.level == Trace
}

func (e *ReguError) IsDebugLevel() bool {
	return e.level == Debug
}

func (e *ReguError) IsInfoLevel() bool {
	return e.level == Info
}

func (e *ReguError) IsWarnLevel() bool {
	return e.level == Warn
}

func (e *ReguError) IsErrorLevel() bool {
	return e.level == Error
}

func (e *ReguError) IsFatalLevel() bool {
	return e.level == Fatal
}

func (e *ReguError) withLevel(lvl Level) *ReguError {
	return &ReguError{
		code:       e.code,
		level:      lvl,
		statusCode: e.statusCode,
		format:     e.format,
		err:        e.err,
	}
}
