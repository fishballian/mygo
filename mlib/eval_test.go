package mlib

import "testing"

func Test_eval(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"400+5"}, 405},
		{"", args{"[1+22*444]*4 + 5*((34-4)/2-33)"}, 38986},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eval(tt.args.s); got != tt.want {
				t.Errorf("eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
