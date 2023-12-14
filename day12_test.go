package main

import "testing"

func Test_nono_arrangements(t *testing.T) {
	type args struct {
		line string
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"??", []int{1}}, 2},
		{"1-fixed", args{"#.", []int{1}}, 1},
		{"1-imposs", args{"#.", []int{2}}, 0},
		{"patchy", args{"???...??..", []int{2}}, 3},
		{"patchy-2", args{"?#?...??..", []int{2}}, 2},
		{"patchy-3", args{"??#...??..", []int{2}}, 1},
		{"line1", args{"#.#.###", []int{1,1,3}}, 1},
		{"line1a", args{"???.###", []int{1,1,3}}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nono_arrangements(tt.args.line, tt.args.vals); got != tt.want {
				t.Errorf("nono_arrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
