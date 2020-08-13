package gen

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gitlab.com/osaki-lab/errcdgen"
	"go/parser"
	"go/token"
	"testing"
)

func TestTraverse(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *File
	}{
		{
			name: "No_Options",
			args: `package example
		
		import (
			"gitlab.com/osaki-lab/errcdgen"
		)
		
		var (
			InvalidInputParameterErr = errcdgen.NewCodeError("1003", "invalid input parameter: %v")
			UpdateConflictErr        = errcdgen.NewCodeError("1004", "other user updated: key=%s")
		)
		`,
			want: &File{
				PkgName: "example",
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
		},
		{
			name: "Method_chained_define",
			args: `package example

import (
	"gitlab.com/osaki-lab/errcdgen"
)

var InvalidInputParameterErr = errcdgen.NewCodeError("1003", "invalid input parameter: %v").
		DisableError().WarnLevel().WithStatusCode(404)
`,
			want: &File{
				PkgName: "example",
				Decls: []*Decl{
					{
						Name:             "InvalidInputParameterErr",
						Code:             "1003",
						Format:           "invalid input parameter: %v",
						LogLevel:         errcdgen.WarnLevel,
						LogLevelEnable:   true,
						StatusCode:       404,
						StatusCodeEnable: true,
						DisableErr:       true,
					},
				},
			},
		},
		{
			name: "Label_parse",
			args: `package example
		
		import (
			"gitlab.com/osaki-lab/errcdgen"
		)
		
		var InvalidInputParameterErr = errcdgen.NewCodeError("1003", "invalid input parameter: %v").
				Label(0, "payload", []string{})
		`,
			want: &File{
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
		{
			name: "Multiple_Label_parse",
			args: `package example
		
		import (
			"gitlab.com/osaki-lab/errcdgen"
		)
		
		var InvalidInputParameterErr = errcdgen.NewCodeError("1003", "strArg:%v intArg:%v mapArg:%v").
				Label(0, "strArg", "dummy").
				Label(1, "intArg", int(199)).
				Label(2, "mapArg", map[string]interface{}{})
		`,
			want: &File{
				PkgName: "example",
				Decls: []*Decl{
					{
						Name:   "InvalidInputParameterErr",
						Code:   "1003",
						Format: "strArg:%v intArg:%v mapArg:%v",
						Labels: []Label{
							{
								Index:  0,
								Name:   "strArg",
								GoType: "string",
							},
							{
								Index:  1,
								Name:   "intArg",
								GoType: "int",
							},
							{
								Index:  2,
								Name:   "mapArg",
								GoType: "map[string]interface{}",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := token.NewFileSet()
			f, err := parser.ParseFile(fs, "example.go", tt.args, 0)
			if err != nil {
				t.Fatalf("invalid test input: %v", err)
			}

			got, err := Traverse(f)
			if err != nil {
				t.Errorf("Traverse() error = %v", err)
				return
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreUnexported(Decl{})); diff != "" {
				t.Errorf("Traverse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
