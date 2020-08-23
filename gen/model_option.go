package gen

import "gitlab.com/osaki-lab/reguerr"

type Setting struct {
	ErrorLevel reguerr.Level
	StatusCode int
}

func NewSetting() *Setting {
	return &Setting{
		ErrorLevel: reguerr.DefaultErrorLevel,
		StatusCode: reguerr.DefaultStatusCode,
	}
}

type Option func(*Setting)

func DefaultErrorLevel(level reguerr.Level) Option {
	return func(s *Setting) {
		s.ErrorLevel = level
	}
}

func DefaultStatusCode(code int) Option {
	return func(s *Setting) {
		s.StatusCode = code
	}
}

func (o Setting) EnableInit() bool {
	return !o.ErrorLevel.Equals(reguerr.DefaultErrorLevel) || o.StatusCode != reguerr.DefaultStatusCode
}

func (o Setting) IsOverwriteErrorLevel() bool {
	return !o.ErrorLevel.Equals(reguerr.DefaultErrorLevel)
}

func (o Setting) IsOverwriteStatusCode() bool {
	return o.StatusCode != reguerr.DefaultStatusCode
}

