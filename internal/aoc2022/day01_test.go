package aoc2022

import (
	"testing"
)

type testData struct {
	input []int
}

var tests = testData{
	[]int{
		1000,
		2000,
		3000,
		0,
		4000,
		0,
		5000,
		6000,
		0,
		7000,
		8000,
		9000,
		0,
		10000,
	},
}

func TestDay01part01(t *testing.T) {
	tests := []struct {
		name string
		args testData
		want int
	}{
		{
			name: "Calories",
			args: tests,
			want: 24000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day01part01(tt.args.input); got != tt.want {
				t.Errorf("Day01part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay01part02(t *testing.T) {
	tests := []struct {
		name string
		args testData
		want int
	}{
		{
			name: "Top3Calories",
			args: tests,
			want: 45000,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day01part02(tt.args.input); got != tt.want {
				t.Errorf("Day01part02() = %v, want %v", got, tt.want)
			}
		})
	}
}
