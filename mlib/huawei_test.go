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
		{"", args{Duplicate(2, 100), 99}, 98},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBest2(tt.args.a, tt.args.pMax); got != tt.want {
				t.Errorf("findBest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findShortestPath(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{grid: [][]int{{1, 0, 1, 1, 1}, {1, 1, 1, 0, 1}, {1, 1, 0, 0, 1}, {1, 0, 1, 1, 1}, {1, 1, 1, 0, 1}}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestPath(tt.args.grid); got != tt.want {
				t.Errorf("findShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
