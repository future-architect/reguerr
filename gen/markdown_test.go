package gen

import (
	"bytes"
	"gitlab.com/osaki-lab/reguerr"
	"testing"
)

func TestGenerateMarkdown(t *testing.T) {
	type args struct {
		decls []*Decl
		opts  []Option
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name:    "normal",
			args:    args{
				decls: []*Decl{
					{
						Name:             "InvalidInputErr",
						Code:             "100",
						Format:           "invalid input(%s)",
						LogLevelEnable:   false,
						LogLevel:         reguerr.Error,
						StatusCodeEnable: false,
						StatusCode:       500,
					},
				},
				opts:  nil,
			},
			wantW:   `# Error Code List

| CODE |      NAME       | LOGLEVEL | STATUSCODE |      FORMAT       |
|------|-----------------|----------|------------|-------------------|
|  100 | InvalidInputErr | Error    |        500 | invalid input(%s) |
`,
			wantErr: false,
		},
		{
			name:    "long_format",
			args:    args{
				decls: []*Decl{
					{
						Name:             "InvalidInputErr",
						Code:             "100",
						Format:           "invalid input(%s)",
						LogLevelEnable:   false,
						LogLevel:         reguerr.Error,
						StatusCodeEnable: false,
						StatusCode:       500,
					},
					{
						Name:             "InvalidInputErr2",
						Code:             "101",
						Format:           "reported time format is invalid",
						LogLevelEnable:   false,
						LogLevel:         reguerr.Error,
						StatusCodeEnable: false,
						StatusCode:       500,
					},
				},
				opts:  nil,
			},
			wantW:   `# Error Code List

| CODE |       NAME       | LOGLEVEL | STATUSCODE |             FORMAT              |
|------|------------------|----------|------------|---------------------------------|
|  100 | InvalidInputErr  | Error    |        500 | invalid input(%s)               |
|  101 | InvalidInputErr2 | Error    |        500 | reported time format is invalid |
`,
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := GenerateMarkdown(w, tt.args.decls, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("GenerateMarkdown() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
