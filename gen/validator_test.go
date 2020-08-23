package gen

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		args    []*Decl
		wantErr bool
		err     error
	}{
		{
			name:    "zero length",
			args:    []*Decl{},
			wantErr: false,
		},
		{
			name: "one length",
			args: []*Decl{
				{
					Name:   "InvalidInputParameterErr",
					Code:   "1003",
					Format: "invalid input parameter",
				},
			},
			wantErr: false,
		},
		{
			name: "two length",
			args: []*Decl{
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
			wantErr: false,
		},
		{
			name: "duplicated message code",
			args: []*Decl{
				{
					Name:   "InvalidInputParameterErr",
					Code:   "1003",
					Format: "invalid input parameter",
				},
				{
					Name:   "UpdateConflictErr",
					Code:   "1003", // duplicated
					Format: "other user updated",
				},
			},
			wantErr: true,
			err:     errors.New("duplicated message code: 1003"),
		},
		{
			name: "there_is_duplicated_code",
			args: []*Decl{
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
				{
					Name:   "XxxErr",
					Code:   "1003",
					Format: "test",
				},
			},
			wantErr: true,
			err:     errors.New("duplicated message code: 1003"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.args); err != nil {
				if tt.wantErr {
					if err.Error() != tt.err.Error() {
						t.Errorf("Validate() error = %v, wantErr %v", err, tt.err)
					}
				} else {
					t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
