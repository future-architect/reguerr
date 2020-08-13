package gen

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		args    File
		want    string
		wantErr bool
	}{
		{
			name:    "multiple declare",
			args:    File{
				PkgName:  "example",
				Decls: []*Decl{
					{
						Name:   "InvalidInputParameterErr",
						Code:   "1003",
						Format: "invalid input parameter: %v",
					},
					{
						Name:   "UpdateConflictErr",
						Code:   "1004",
						Format: "other user updated: key=%s",
					},
				},
			},
			want:    `// generated by errcdgen; DO NOT EDIT
package example

import (
	"gitlab.com/osaki-lab/errcdgen"
)

func NewInvalidInputParameterErr(err error) *errcdgen.CodeError {
	return InvalidInputParameterErr.WithError(err)
}

func NewUpdateConflictErr(err error) *errcdgen.CodeError {
	return UpdateConflictErr.WithError(err)
}
`,
			wantErr: false,
		},
		{
			name:    "DisableErr=true",
			args:    File{
				PkgName:  "example",
				Decls: []*Decl{
					{
						Name:             "InvalidInputParameterErr",
						Code:             "1003",
						Format:           "invalid input parameter: %v",
						DisableErr:       true,
					},
				},
			},
			want:    `// generated by errcdgen; DO NOT EDIT
package example

import (
	"gitlab.com/osaki-lab/errcdgen"
)

func NewInvalidInputParameterErr() *errcdgen.CodeError {
	return InvalidInputParameterErr
}
`,
			wantErr: false,
		},
		{
			name:    "Label",
			args:    File{
				PkgName:  "example",
				Decls: []*Decl{
					{
						Name:             "InvalidInputParameterErr",
						Code:             "1003",
						Format:           "invalid input parameter: %v",
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
			want:    `// generated by errcdgen; DO NOT EDIT
package example

import (
	"gitlab.com/osaki-lab/errcdgen"
)

func NewInvalidInputParameterErr(err error, payload []string) *errcdgen.CodeError {
	return InvalidInputParameterErr.WithError(err).Args(payload)
}
`,
			wantErr: false,
		},
		{
			name:    "Multiple_Label",
			args:    File{
				PkgName:  "example",
				Decls: []*Decl{
					{
						Name:             "InvalidInputParameterErr",
						Code:             "1003",
						Format:           "invalid input parameter: str:%v inv:%v",
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
			want:    `// generated by errcdgen; DO NOT EDIT
package example

import (
	"gitlab.com/osaki-lab/errcdgen"
)

func NewInvalidInputParameterErr(err error, strArg1 string, intArg1 int) *errcdgen.CodeError {
	return InvalidInputParameterErr.WithError(err).Args(strArg1, intArg1)
}
`,
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.PkgName, tt.args.Decls)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}


			if diff := cmp.Diff(tt.want, string(got)); diff != "" {
				t.Errorf("Traverse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}