package aoc2022

import "testing"

type testDataDay04 struct {
	input []string
}

var testsDay04 = testDataDay04{
	[]string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	},
}

func TestDay04Part01(t *testing.T) {
	tests := []struct {
		name string
		args testDataDay04
		want int
	}{
		{
			name: "FullyOverlap",
			args: testsDay04,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day04Part01(tt.args.input); got != tt.want {
				t.Errorf("Day04Part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay04Part02(t *testing.T) {
	tests := []struct {
		name string
		args testDataDay04
		want int
	}{
		{
			name: "PartiallyOverlap",
			args: testsDay04,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day04Part02(tt.args.input); got != tt.want {
				t.Errorf("Day04Part01() = %v, want %v", got, tt.want)
			}
		})
	}
}
