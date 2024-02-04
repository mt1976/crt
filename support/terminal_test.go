package support

import (
	"os"
	"reflect"
	"testing"
)

func TestNewCrt(t *testing.T) {
	tests := []struct {
		name string
		want Crt
	}{
		// TODO: Add test cases.
		// Create new Crt
		// Set width and height
		{"TestNewCrt", Crt{isTerminal: false, width: 0, height: 0, firstRow: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCrt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHostName(t *testing.T) {
	x, _ := os.Hostname()

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		// Get Hostname
		{"TestGetHostName", x, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHostName()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHostName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetHostName() = %v, want %v", got, tt.want)
			}
		})
	}
}
