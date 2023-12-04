package main

import (
	"reflect"
	"testing"
)

func Test_find_numbers(t *testing.T) {
	tests := []struct {
		name string
		source string
		want []Number
	}{
		{"single", ".123.", []Number{{1,3,123}}},
		{"double", "467..114..", []Number{{0,2,467},{5,7,114}}},
		//{"symbol", "...*......", []Number{}}, // doesn't work cannot compare empty with empty
		{"start num", "617*......", []Number{{0,2,617}}},
		{"end num", "......*716", []Number{{7,9,716}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := find_numbers(tt.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("find_numbers() = %v, want %v", got, tt.want)
			}
		})
	}

	if got:= find_numbers("...*......"); len(got)>0{
		t.Errorf("find_numbers(empty) = %v, want {}", got)
	}
}

func Test_is_valid_number(t *testing.T) {

	lines, _ := LoadLines("data/test03.txt")

	type args struct {
		n    Number
		y    int
		grid []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"467",args{Number{0,2,467},0,lines}, true},
		{"114",args{Number{5,7,114},0,lines}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is_valid_number(tt.args.n, tt.args.y, tt.args.grid); got != tt.want {
				t.Errorf("is_valid_number() = %v, want %v", got, tt.want)
			}
		})
	}
}