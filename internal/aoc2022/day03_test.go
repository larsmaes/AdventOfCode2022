package aoc2022

import (
	"testing"
)

type testDataDay3 struct {
	input []string
}

var testsDay3 = testDataDay3{
	[]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	},
}

func TestDay03part01(t *testing.T) {
	tests := []struct {
		name string
		args testDataDay3
		want int
	}{
		{
			name: "Score",
			args: testsDay3,
			want: 157,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day03part01(tt.args.input); got != tt.want {
				t.Errorf("Day03part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay03part02(t *testing.T) {
	tests := []struct {
		name string
		args testDataDay3
		want int
	}{
		{
			name: "Score",
			args: testsDay3,
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day03part02(tt.args.input); got != tt.want {
				t.Errorf("Day03part02() = %v, want %v", got, tt.want)
			}
		})
	}
}
