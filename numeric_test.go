package crt

import "testing"

func Test_isInt(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"All Numeric", args{"123"}, true},
		{"All Alpha", args{"abc"}, false},
		{"Numeric & Alpha", args{"1a2b3c"}, false},
		{"Empty", args{""}, false},
		{"Long Numeric", args{"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901  "}, false},
		{"Space", args{" "}, false},
		{"Newline", args{"\n"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInt(tt.args.in); got != tt.want {
				t.Errorf("isInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ToInt(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantOut int
	}{
		{
			name:    "Valid input",
			in:      "123",
			wantOut: 123,
		},
		{
			name:    "Invalid input",
			in:      "abc",
			wantOut: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut := toInt(tt.in)
			if gotOut != tt.wantOut {
				t.Errorf("toInt() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_ToString(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want string
	}{
		{
			name: "basic test",
			in:   123,
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.in); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
