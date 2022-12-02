package aoc2022

import (
	"testing"
)

type testDataDay2 struct {
	input []string
}

var testsDay2 = testDataDay2{
	[]string{
		"A Y",
		"B X",
		"C Z",
	},
}

func TestDay02part01(t *testing.T) {
	tests := []struct {
		name string
		args testDataDay2
		want int
	}{
		{
			name: "RockPaperScissors",
			args: testsDay2,
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day02part01(tt.args.input); got != tt.want {
				t.Errorf("Day02part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay02part02(t *testing.T) {
	tests := []struct {
		name string
		args testDataDay2
		want int
	}{
		{
			name: "RockPaperScissorsPart2",
			args: testsDay2,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day02part02(tt.args.input); got != tt.want {
				t.Errorf("Day02part02() = %v, want %v", got, tt.want)
			}
		})
	}
}
