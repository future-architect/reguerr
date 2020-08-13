package gen

import (
	"reflect"
	"testing"
)

func TestAnalyze(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want FmtVerbs
	}{
		{
			name: "zero_arg",
			arg:  "invalid request",
			want: FmtVerbs{
				Verb: nil,
			},
		},
		{
			name: "one_arg",
			arg:  "invalid request: %v",
			want: FmtVerbs{
				Verb: []string{"%v"},
			},
		},
		{
			name: "multiple_args",
			arg:  "invalid request(key=%s): %v",
			want: FmtVerbs{
				Verb: []string{"%s", "%v"},
			},
		},
		{
			name: "continuous",
			arg:  "%s%v%+v",
			want: FmtVerbs{
				Verb: []string{"%s", "%v", "%+v"},
			},
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
