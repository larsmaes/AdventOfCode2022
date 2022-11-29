package aoc2022

import "testing"

func TestDay1part1(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{[]int{1, 2, 3, 4}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day1part1(tt.args.input); got != tt.want {
				t.Errorf("Day1part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
