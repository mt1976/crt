package crt

import "testing"

func Test_timeString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{"Sample", timeString()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeString(); got != tt.want {
				t.Errorf("timeString() = %v, want %v", got, tt.want)
			}
			t.Log(tt.want)
			//t.Log(got)
		})
	}
}
