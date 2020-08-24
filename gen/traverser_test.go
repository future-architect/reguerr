package gen

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gitlab.com/osaki-lab/reguerr"
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
			"gitlab.com/osaki-lab/reguerr"
		)
		
		var (
			InvalidInputParameterErr = reguerr.New("1003", "invalid input parameter: %v").Build()
			UpdateConflictErr        = reguerr.New("1004", "other user updated: key=%s").Build()
		)
		`,
			want: &File{
				PkgName: "example",
				Decls: []*Decl{
					{
						Name:      "InvalidInputParameterErr",
						Code:      "1003",
						Format:    "invalid input parameter: %v",
						CallBuild: true,
					},
					{
						Name:      "UpdateConflictErr",
						Code:      "1004",
						Format:    "other user updated: key=%s",
						CallBuild: true,
					},
				},
			},
		},
		{
			name: "Method_chained_define",
			args: `package example

import (
	"gitlab.com/osaki-lab/reguerr"
)

var InvalidInputParameterErr = reguerr.New("1003", "invalid input parameter: %v").
		DisableError().Warn().WithStatusCode(404).Build()
`,
			want: &File{
				PkgName: "example",
				Decls: []*Decl{
					{
						Name:             "InvalidInputParameterErr",
						Code:             "1003",
						Format:           "invalid input parameter: %v",
						LogLevel:         reguerr.Warn,
						LogLevelEnable:   true,
						StatusCode:       404,
						StatusCodeEnable: true,
						DisableErr:       true,
						CallBuild:        true,
					},
				},
			},
		},
		{
			name: "Label_parse",
			args: `package example
		
		import (
			"gitlab.com/osaki-lab/reguerr"
		)
		
		var InvalidInputParameterErr = reguerr.New("1003", "invalid input parameter: %v").
				Label(0, "payload", []string{}).Build()
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
						CallBuild: true,
					},
				},
			},
		},
		{
			name: "Multiple_Label_parse",
			args: `package example
		
		import (
			"gitlab.com/osaki-lab/reguerr"
		)
		
		var InvalidInputParameterErr = reguerr.New("1003", "strArg:%v intArg:%v mapArg:%v").
				Label(0, "strArg", "dummy").
				Label(1, "intArg", int(199)).
				Label(2, "mapArg", map[string]interface{}{}).
				Build()
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
						CallBuild: true,
					},
				},
			},
		},
		{
			name: "no_build_func_call",
			args: `package example
		
		import (
			"gitlab.com/osaki-lab/reguerr"
		)
		
		var InvalidInputParameterErr = reguerr.New("1003", "invalid input param")
		`,
			want: &File{
				PkgName: "example",
				Decls: []*Decl{
					{
						Name:      "InvalidInputParameterErr",
						Code:      "1003",
						Format:    "invalid input param",
						CallBuild: false,
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
