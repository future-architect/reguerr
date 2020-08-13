package gen

import (
	"reflect"
	"testing"
)

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want []Verb
	}{
		{
			name: "zero_arg",
			arg:  "invalid request",
			want: nil,
		},
		{
			name: "one_arg",
			arg:  "invalid request: %v",
			want: []Verb{"%v"},
		},
		{
			name: "multiple_args",
			arg:  "invalid request(key=%s): %v",
			want: []Verb{"%s", "%v"},
		},
		{
			name: "continuous",
			arg:  "%s%v%+v",
			want: []Verb{"%s", "%v", "%+v"},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Analyze(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
