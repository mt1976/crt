package crt

import (
	"fmt"
	"testing"

	lang "github.com/mt1976/crt/language"

	goterm "github.com/buger/goterm"
)

func TestPage_Add(t *testing.T) {
	type fields struct {
		title           string
		pageRows        []pageRow
		noRows          int
		prompt          string
		actions         []string
		actionMaxLen    int
		noPages         int
		ActivePageIndex int
		counter         int
		pageRowCounter  int
		viewPort        *ViewPort
	}
	type args struct {
		rowContent string
		altID      string
		dateTime   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"Test 1", fields{}, args{"", "", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				title:           tt.fields.title,
				pageRows:        tt.fields.pageRows,
				noRows:          tt.fields.noRows,
				prompt:          tt.fields.prompt,
				actions:         tt.fields.actions,
				actionMaxLen:    tt.fields.actionMaxLen,
				noPages:         tt.fields.noPages,
				ActivePageIndex: tt.fields.ActivePageIndex,
				counter:         tt.fields.counter,
				pageRowCounter:  tt.fields.pageRowCounter,
				viewPort:        tt.fields.viewPort,
			}
			p.Add(tt.args.rowContent, tt.args.altID, tt.args.dateTime)
		})
	}
}

func Test_PrintIt(t *testing.T) {
	t.Run("No Baud Rate", func(t *testing.T) {
		vp := New()
		msg := "Hello, World!"
		vp.PrintIt(msg)
		if got, want := fmt.Sprintf("%s", goterm.Screen.String()), fmt.Sprintf("%s%s", msg, "\n"); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("With Baud Rate", func(t *testing.T) {
		vp := New()
		vp.SetBaud(9600)
		msg := "Hello, World!"
		vp.PrintIt(msg)
		if got, want := fmt.Sprintf("%s", goterm.Screen.String()), fmt.Sprintf("%c%s%c", lang.BoxCharacterNormal, msg, lang.BoxCharacterNormal); got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
