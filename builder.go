package reguerr

type Builder struct {
	err *Error
}

func New(code, format string) *Builder {
	return &Builder{
		err: &Error{
			code:       code,
			format:     format,
			level:      DefaultErrorLevel,
			statusCode: DefaultStatusCode,
		},
	}
}

func (b *Builder) Label(index int, name string, goType interface{}) *Builder {
	// コード解析用の関数なので、内部的には何もしないしなくて良い
	return b
}

func (b *Builder) DisableError() *Builder {
	// 解析用途なのにで、何もしない
	return b
}

func (b *Builder) TraceLevel() *Builder {
	return &Builder{
		err: b.err.withLevel(TraceLevel),
	}
}

func (b *Builder) DebugLevel() *Builder {
	return &Builder{
		err: b.err.withLevel(DebugLevel),
	}
}

func (b *Builder) InfoLevel() *Builder {
	return &Builder{
		b.err.withLevel(InfoLevel),
	}
}

func (b *Builder) WarnLevel() *Builder {
	return &Builder{
		b.err.withLevel(WarnLevel),
	}
}

func (b *Builder) ErrorLevel() *Builder {
	return &Builder{
		b.err.withLevel(ErrorLevel),
	}
}

func (b *Builder) FatalLevel() *Builder {
	return &Builder{
		b.err.withLevel(FatalLevel),
	}
}

func (b *Builder) WithStatusCode(statusCode int) *Builder {
	return &Builder{
		err: &Error{
			code:       b.err.code,
			level:      b.err.level,
			statusCode: statusCode,
			format:     b.err.format,
			err:        b.err.err,
			args:       b.err.args,
		},
	}
}

func (b *Builder) Build() *Error {
	return b.err
}
