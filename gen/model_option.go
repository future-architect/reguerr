package gen

import "github.com/future-architect/reguerr"

type Setting struct {
	Level      reguerr.Level
	StatusCode int
}

func NewSetting() *Setting {
	return &Setting{
		Level:      reguerr.DefaultErrorLevel,
		StatusCode: reguerr.DefaultStatusCode,
	}
}

type Option func(*Setting)

func DefaultErrorLevel(level reguerr.Level) Option {
	return func(s *Setting) {
		s.Level = level
	}
}

func DefaultStatusCode(code int) Option {
	return func(s *Setting) {
		s.StatusCode = code
	}
}

func (o Setting) EnableInit() bool {
	return !o.Level.Equals(reguerr.DefaultErrorLevel) || o.StatusCode != reguerr.DefaultStatusCode
}

func (o Setting) IsOverwriteErrorLevel() bool {
	return !o.Level.Equals(reguerr.DefaultErrorLevel)
}

func (o Setting) IsOverwriteStatusCode() bool {
	return o.StatusCode != reguerr.DefaultStatusCode
}
