package gen

import (
	"github.com/google/go-cmp/cmp"
	"go/parser"
	"go/token"
	"testing"
)

func TestTraverse1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *File
	}{
		{
			name: "No Options",
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
				Declares: []*DeclareErr{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := token.NewFileSet()
			in, err := parser.ParseFile(fs, "example.go", tt.args, 0)
			if err != nil {
				t.Fatalf("invalid test input: %v", err)
			}

			got, err := Traverse(in)
			if err != nil {
				t.Errorf("Traverse() error = %v", err)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Traverse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
