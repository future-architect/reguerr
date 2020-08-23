package reguerr

import (
	"errors"
	"testing"
)

func TestCodeError_Error(t *testing.T) {
	tests := []struct {
		name string
		in   Error
		want string
	}{
		{
			name: "",
			in: Error{
				code:       "1003",
				level:      ErrorLevel,
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
