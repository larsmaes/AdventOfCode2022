package aoc2022

import "testing"

func TestDay06Part01(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getMarker1",
			args: args{
				input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 7,
		},
		{
			name: "getMarker2",
			args: args{
				input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 5,
		},
		{
			name: "getMarker3",
			args: args{
				input: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 6,
		},
		{
			name: "getMarker4",
			args: args{
				input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 10,
		},
		{
			name: "getMarker5",
			args: args{
				input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day06Part01(tt.args.input); got != tt.want {
				t.Errorf("Day06Part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay06Part02(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getMarker1",
			args: args{
				input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			},
			want: 19,
		},
		{
			name: "getMarker2",
			args: args{
				input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			},
			want: 23,
		},
		{
			name: "getMarker3",
			args: args{
				input: "nppdvjthqldpwncqszvftbrmjlhg",
			},
			want: 23,
		},
		{
			name: "getMarker4",
			args: args{
				input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			},
			want: 29,
		},
		{
			name: "getMarker5",
			args: args{
				input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			},
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day06Part02(tt.args.input); got != tt.want {
				t.Errorf("Day06Part01() = %v, want %v", got, tt.want)
			}
		})
	}
}
