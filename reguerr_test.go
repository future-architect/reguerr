package reguerr

import (
	"errors"
	"testing"
)

func TestCodeError_Error(t *testing.T) {
	tests := []struct {
		name string
		in   ReguError
		want string
	}{
		{
			name: "",
			in: ReguError{
				code:       "1003",
				level:      Error,
				statusCode: 500,
				format:     "invalid input parameter: %v",
				args:       []interface{}{`{"key":"hoge"}`},
				err:        errors.New("internal error"),
			},
			want: `[1003]invalid input parameter: [{"key":"hoge"}]: internal error`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReguError_Unwrap(t *testing.T) {

	var (
		inErr = errors.New("internal error")
	)

	tests := []struct {
		name    string
		in      *ReguError
		wantErr error
	}{
		{
			name: "exists_err",
			in: &ReguError{
				code:       "1003",
				level:      Error,
				statusCode: 500,
				format:     "invalid input parameter: %v",
				args:       []interface{}{`{"key":"hoge"}`},
				err:        inErr,
			},
			wantErr: inErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := errors.Unwrap(tt.in); err != tt.wantErr {
				t.Errorf("errors.Unwrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
