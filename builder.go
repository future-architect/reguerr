package reguerr

type Builder struct {
	err *ReguError
}

func New(code, format string) *Builder {
	return &Builder{
		err: &ReguError{
			code:       code,
			format:     format,
			level:      DefaultErrorLevel,
			statusCode: DefaultStatusCode,
		},
	}
}

func (b *Builder) Label(index int, name string, goType interface{}) *Builder {
	// ignore because this func is for analysis
	return b
}

func (b *Builder) DisableError() *Builder {
	// ignore because this func is for analysis
	return b
}

func (b *Builder) Trace() *Builder {
	return &Builder{
		err: b.err.withLevel(Trace),
	}
}

func (b *Builder) Debug() *Builder {
	return &Builder{
		err: b.err.withLevel(Debug),
	}
}

func (b *Builder) Info() *Builder {
	return &Builder{
		b.err.withLevel(Info),
	}
}

func (b *Builder) Warn() *Builder {
	return &Builder{
		b.err.withLevel(Warn),
	}
}

func (b *Builder) Error() *Builder {
	return &Builder{
		b.err.withLevel(Error),
	}
}

func (b *Builder) Fatal() *Builder {
	return &Builder{
		b.err.withLevel(Fatal),
	}
}

func (b *Builder) WithStatusCode(statusCode int) *Builder {
	return &Builder{
		err: &ReguError{
			code:       b.err.code,
			level:      b.err.level,
			statusCode: statusCode,
			format:     b.err.format,
			err:        b.err.err,
			args:       b.err.args,
		},
	}
}

func (b *Builder) Build() *ReguError {
	return b.err
}
