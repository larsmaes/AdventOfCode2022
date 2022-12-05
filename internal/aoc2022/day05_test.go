package aoc2022

import "testing"

func TestDay05Part01(t *testing.T) {
	type args struct {
		state      map[int][]string
		directives []Directive
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Rearrange",
			args: args{
				state: map[int][]string{
					1: {
						"Z",
						"N",
					},
					2: {
						"M",
						"C",
						"D",
					},
					3: {
						"P",
					},
				},
				directives: []Directive{
					{
						Amount: 1,
						From:   2,
						To:     1,
					},
					{
						Amount: 3,
						From:   1,
						To:     3,
					},
					{
						Amount: 2,
						From:   2,
						To:     1,
					},
					{
						Amount: 1,
						From:   1,
						To:     2,
					},
				},
			},
			want: "CMZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day05Part01(tt.args.state, tt.args.directives); got != tt.want {
				t.Errorf("Day05Part01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay05Part02(t *testing.T) {
	type args struct {
		state      map[int][]string
		directives []Directive
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "RearrangePlus",
			args: args{
				state: map[int][]string{
					1: {
						"Z",
						"N",
					},
					2: {
						"M",
						"C",
						"D",
					},
					3: {
						"P",
					},
				},
				directives: []Directive{
					{
						Amount: 1,
						From:   2,
						To:     1,
					},
					{
						Amount: 3,
						From:   1,
						To:     3,
					},
					{
						Amount: 2,
						From:   2,
						To:     1,
					},
					{
						Amount: 1,
						From:   1,
						To:     2,
					},
				},
			},
			want: "MCD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Day05Part02(tt.args.state, tt.args.directives); got != tt.want {
				t.Errorf("Day05Part02() = %v, want %v", got, tt.want)
			}
		})
	}
}
