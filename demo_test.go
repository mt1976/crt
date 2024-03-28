package crt

import "testing"

func TestSample(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"Sample"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sample()
		})
	}
}
