package support

import (
	"testing"
)

func TestTrimRepeatingCharacters(t *testing.T) {
	type args struct {
		s string
		c string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test 1", args{"", ""}, ""},
		{"Test 2", args{"", "a"}, ""},
		{"Test 3", args{"a", ""}, "a"},
		{"Test 4", args{"a", "a"}, "a"},
		{"Test 5", args{"aa", "a"}, "a"},
		{"Test 6", args{"aaa", "a"}, "a"},
		{"Test 7", args{"aaaa", "a"}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimRepeatingCharacters(tt.args.s, tt.args.c); got != tt.want {
				t.Errorf("TrimRepeatingCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundFloatToTwo(t *testing.T) {
	type args struct {
		input float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"Test 1", args{0.0}, 0.0},
		{"Test 2", args{0.1}, 0.1},
		{"Test 3", args{0.01}, 0.01},
		{"Test 4", args{0.001}, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundFloatToTwo(tt.args.input); got != tt.want {
				t.Errorf("roundFloatToTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiskSizeHuman(t *testing.T) {
	type args struct {
		input uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test 1", args{0}, "0.00GB (0.00TB)"},
		{"Test 2", args{1}, "0.00GB (0.00TB)"},
		{"Test 3", args{1024}, "0.00GB (0.00TB)"},
		{"Test 4", args{1024 * 1024}, "0.00GB (0.00TB)"},
		{"Test 5", args{1024 * 1024 * 1024}, "1.00GB (0.00TB)"},
		{"Test 6", args{1024 * 1024 * 1024 * 1024}, "1024.00GB (1.00TB)"},
		{"Test 7", args{1024 * 1024 * 1024 * 1024 * 1024}, "1048576.00GB (1024.00TB)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiskSizeHuman(tt.args.input); got != tt.want {
				t.Errorf("DiskSizeHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}
