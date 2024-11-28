package mlib

import "testing"

func Test_findBest(t *testing.T) {
	type args struct {
		a    []int
		pMax int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{50, 20, 20, 60}, 90}, 90},
		{"", args{[]int{50, 40}, 30}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBest(tt.args.a, tt.args.pMax); got != tt.want {
				t.Errorf("findBest() = %v, want %v", got, tt.want)
			}
		})
	}
}
