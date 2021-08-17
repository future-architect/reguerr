package gen

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/future-architect/reguerr"
)

func TestGenerate(t *testing.T) {
	type arg struct {
		file *File
		opts []Option
	}
	tests := []struct {
		name    string
		args    arg
		want    string
		wantErr bool
	}{
		{
			name: "multiple_declare",
			args: arg{
				file: &File{
					PkgName: "example",
					Decls: []*Decl{
						{
							Name:   "InvalidInputParameterErr",
							Code:   "1003",
							Format: "invalid input parameter",
						},
						{
							Name:   "UpdateConflictErr",
							Code:   "1004",
							Format: "other user updated",
						},
					},
				},
			},
			want: `// generated by reguerr; DO NOT EDIT
package example

import (
	"errors"
	"github.com/future-architect/reguerr"
)

// NewInvalidInputParameterErr is the error indicating [1003] invalid input parameter: $err.
func NewInvalidInputParameterErr(err error) *reguerr.ReguError {
	return InvalidInputParameterErr.WithError(err)
}

// IsInvalidInputParameterErr indicates if the passed in error is from the error with code [1003].
func IsInvalidInputParameterErr(err error) bool {
	var cerr *reguerr.ReguError
	if as := errors.As(err, &cerr); as {
		if cerr.Code() == InvalidInputParameterErr.Code() {
			return true
		}
	}
	return false
}

// NewUpdateConflictErr is the error indicating [1004] other user updated: $err.
func NewUpdateConflictErr(err error) *reguerr.ReguError {
	return UpdateConflictErr.WithError(err)
}

// IsUpdateConflictErr indicates if the passed in error is from the error with code [1004].
func IsUpdateConflictErr(err error) bool {
	var cerr *reguerr.ReguError
	if as := errors.As(err, &cerr); as {
		if cerr.Code() == UpdateConflictErr.Code() {
			return true
		}
	}
	return false
}
`,
			wantErr: false,
		},
		{
			name: "DisableErr=true",
			args: arg{
				file: &File{
					PkgName: "example",
					Decls: []*Decl{
						{
							Name:       "InvalidInputParameterErr",
							Code:       "1003",
							Format:     "invalid input parameter: %v",
							DisableErr: true,
						},
					},
				},
			},
			want: `// generated by reguerr; DO NOT EDIT
package example

import (
	"errors"
	"github.com/future-architect/reguerr"
)

// NewInvalidInputParameterErr is the error indicating [1003] invalid input parameter: %v.
func NewInvalidInputParameterErr(arg1 interface{}) *reguerr.ReguError {
	return InvalidInputParameterErr.WithArgs(arg1)
}

// IsInvalidInputParameterErr indicates if the passed in error is from the error with code [1003].
func IsInvalidInputParameterErr(err error) bool {
	var cerr *reguerr.ReguError
	if as := errors.As(err, &cerr); as {
		if cerr.Code() == InvalidInputParameterErr.Code() {
			return true
		}
	}
	return false
}
`,
			wantErr: false,
		},
		{
			name: "Label",
			args: arg{
				file: &File{
					PkgName: "example",
					Decls: []*Decl{
						{
							Name:   "InvalidInputParameterErr",
							Code:   "1003",
							Format: "invalid input parameter: %v",
							Labels: []Label{
								{
									Index:  0,
									Name:   "payload",
									GoType: "[]string",
								},
							},
						},
					},
				},
			},
			want: `// generated by reguerr; DO NOT EDIT
package example

import (
	"errors"
	"github.com/future-architect/reguerr"
)

// NewInvalidInputParameterErr is the error indicating [1003] invalid input parameter: %v: $err.
func NewInvalidInputParameterErr(err error, payload []string) *reguerr.ReguError {
	return InvalidInputParameterErr.WithError(err).WithArgs(payload)
}

// IsInvalidInputParameterErr indicates if the passed in error is from the error with code [1003].
func IsInvalidInputParameterErr(err error) bool {
	var cerr *reguerr.ReguError
	if as := errors.As(err, &cerr); as {
		if cerr.Code() == InvalidInputParameterErr.Code() {
			return true
		}
	}
	return false
}
`,
			wantErr: false,
		},
		{
			name: "Multiple_Label",
			args: arg{
				file: &File{
					PkgName: "example",
					Decls: []*Decl{
						{
							Name:   "InvalidInputParameterErr",
							Code:   "1003",
							Format: "invalid input parameter: str:%v inv:%v",
							Labels: []Label{
								{
									Index:  0,
									Name:   "strArg1",
									GoType: "string",
								},
								{
									Index:  1,
									Name:   "intArg1",
									GoType: "int",
								},
							},
						},
					},
				},
			},
			want: `// generated by reguerr; DO NOT EDIT
package example

import (
	"errors"
	"github.com/future-architect/reguerr"
)

// NewInvalidInputParameterErr is the error indicating [1003] invalid input parameter: str:%v inv:%v: $err.
func NewInvalidInputParameterErr(err error, strArg1 string, intArg1 int) *reguerr.ReguError {
	return InvalidInputParameterErr.WithError(err).WithArgs(strArg1, intArg1)
}

// IsInvalidInputParameterErr indicates if the passed in error is from the error with code [1003].
func IsInvalidInputParameterErr(err error) bool {
	var cerr *reguerr.ReguError
	if as := errors.As(err, &cerr); as {
		if cerr.Code() == InvalidInputParameterErr.Code() {
			return true
		}
	}
	return false
}
`,
			wantErr: false,
		},
		{
			name: "No_Label_But_Exists_Args",
			args: arg{
				file: &File{
					PkgName: "example",
					Decls: []*Decl{
						{
							Name:   "InvalidInputParameterErr",
							Code:   "1003",
							Format: "invalid input parameter: key:%v value:%v",
							Labels: []Label{},
						},
					},
				},
			},
			want: `// generated by reguerr; DO NOT EDIT
package example

import (
	"errors"
	"github.com/future-architect/reguerr"
)

// NewInvalidInputParameterErr is the error indicating [1003] invalid input parameter: key:%v value:%v: $err.
func NewInvalidInputParameterErr(err error, arg1 interface{}, arg2 interface{}) *reguerr.ReguError {
	return InvalidInputParameterErr.WithError(err).WithArgs(arg1, arg2)
}

// IsInvalidInputParameterErr indicates if the passed in error is from the error with code [1003].
func IsInvalidInputParameterErr(err error) bool {
	var cerr *reguerr.ReguError
	if as := errors.As(err, &cerr); as {
		if cerr.Code() == InvalidInputParameterErr.Code() {
			return true
		}
	}
	return false
}
`,
			wantErr: false,
		},
		{
			name: "overwrite_default_statusCode_and_errorLevel",
			args: arg{
				file: &File{
					PkgName: "example",
					Decls: []*Decl{
						{
							Name:   "InvalidInputParameterErr",
							Code:   "1003",
							Format: "invalid input parameter",
						},
					},
				},
				opts: []Option{
					DefaultErrorLevel(reguerr.Warn),
					DefaultStatusCode(501),
				},
			},
			want: `// generated by reguerr; DO NOT EDIT
package example

import (
	"errors"
	"github.com/future-architect/reguerr"
)

func init() {
	reguerr.DefaultErrorLevel = reguerr.Warn
	reguerr.DefaultStatusCode = 501
}

// NewInvalidInputParameterErr is the error indicating [1003] invalid input parameter: $err.
func NewInvalidInputParameterErr(err error) *reguerr.ReguError {
	return InvalidInputParameterErr.WithError(err)
}

// IsInvalidInputParameterErr indicates if the passed in error is from the error with code [1003].
func IsInvalidInputParameterErr(err error) bool {
	var cerr *reguerr.ReguError
	if as := errors.As(err, &cerr); as {
		if cerr.Code() == InvalidInputParameterErr.Code() {
			return true
		}
	}
	return false
}
`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateCode(tt.args.file, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, string(got)); diff != "" {
				t.Errorf("Traverse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
