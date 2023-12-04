package main

import (
	"testing"
)

/* note to self, setup unit tests:
* create file utils.go & add code (as its all within a single package, it need not be capitalised)
* right click on function and Generate=>test for function
* write some tests
* in the project window select utils_test.go rmb=>run file
* (it will fail, its ok)
* on top tight, select the dropdown and edit configurations
* change the config for utils_test.go to include both the test file and the utils.go file
* you can now run the utils_test.go utils.go test
*/

func TestLoadLines(t *testing.T) {

	lines,err := LoadLines("data/test01.txt")
	if err!= nil{
		t.Errorf("LoadLines() error = %v", err)
		return
	}

	if len(lines) != 4 {
		t.Errorf("LoadLines() wanted 4 files got #{len(lines)}")
		return
	}

	tests := []struct {
		expected string
	} {
		{"1abc2"},
		{"pqr3stu8vwx"},
		{"a1b2c3d4e5f"},
		{"treb7uchet"},
	}

	for idx, line := range tests {
		if lines[idx] != line.expected {
			t.Errorf("LoadLines() error line #{idx} got = #{lines[idx]} expected = #{line.expected}")
		}
	}
}

func TestLoadLinesOnEmpty(t *testing.T) {
	_, err := LoadLines("data/no_such_file")
	if err == nil {
		t.Errorf("LoadLines() was able to read file when not supposed to")
	}
}

func TestINear(t *testing.T) {
	type args struct {
		a    int
		b    int
		dist int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"simple", args{10,11, 1}, true},
		{"simple-false", args{10,12, 1}, false},
		{"reverse", args{11,10, 1}, true},
		{"reverse-false", args{12,10, 1}, false},
		{"zero", args{0,1, 2}, true},
		{"zero-false", args{0,3, 2}, false},
		{"negative", args{0,-1, 2}, true},
		{"negative-false", args{0,-3, 2}, false},
		{"across-zero", args{1,-1, 2}, true},
		{"across-zero-false", args{1,-3, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := INear(tt.args.a, tt.args.b, tt.args.dist); got != tt.want {
				t.Errorf("INear() = %v, want %v", got, tt.want)
			}
		})
	}
}