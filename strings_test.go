package crt

import (
	"testing"

	styl "github.com/mt1976/crt/styles"
)

func Test_upcase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"All Lowercase", args{"abc"}, "ABC"},
		{"All Uppercase", args{"ABC"}, "ABC"},
		{"Mixed Case", args{"aBc"}, "ABC"},
		{"Mixed Case", args{"aBcD"}, "ABCD"},
		{"Empty", args{""}, ""},
		{"Space", args{" "}, " "},
		{"Newline", args{"\n"}, "\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := upcase(tt.args.s); got != tt.want {
				t.Errorf("upcase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downcase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"All Lowercase", args{"abc"}, "abc"},
		{"All Uppercase", args{"ABC"}, "abc"},
		{"Mixed Case", args{"aBc"}, "abc"},
		{"Mixed Case", args{"aBcD"}, "abcd"},
		{"Empty", args{""}, ""},
		{"Space", args{" "}, " "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := downcase(tt.args.s); got != tt.want {
				t.Errorf("downcase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimRepeatingCharacters(t *testing.T) {
	type args struct {
		s string
		c string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"All Lowercase", args{"abc", "a"}, "abc"},
		{"All Uppercase", args{"ABC", "A"}, "ABC"},
		{"Mixed Case", args{"aBc", "a"}, "aBc"},
		{"Mixed Case", args{"aBcD", "a"}, "aBcD"},
		{"Empty", args{"", ""}, ""},
		{"Space", args{"  ", " "}, " "},
		{"ReplaceX", args{"1X1XX1XXXX1", "X"}, "1X1X1X1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimRepeatingCharacters(tt.args.s, tt.args.c); got != tt.want {
				t.Errorf("trimRepeatingCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Bold(t *testing.T) {
	input := "hello world"
	expected := styl.Bold + input + styl.Reset
	actual := bold(input)
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
