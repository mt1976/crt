package menu

import (
	"testing"

	"github.com/mt1976/crt/support"
)

func TestPage_NextPage(t *testing.T) {
	type fields struct {
		title             string
		pageRows          []pageRow
		noRows            int
		prompt            string
		actions           []string
		actionMaxLen      int
		noPages           int
		CurrentPageNumber int
		counter           int
	}
	type args struct {
		crt *support.Crt
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"Test1", fields{title: "Test1", pageRows: []pageRow{}, noRows: 0, prompt: "Test1", actions: []string{}, actionMaxLen: 0, noPages: 0, CurrentPageNumber: 1, counter: 0}, args{crt: &support.Crt{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Page{
				title:           tt.fields.title,
				pageRows:        tt.fields.pageRows,
				noRows:          tt.fields.noRows,
				prompt:          tt.fields.prompt,
				actions:         tt.fields.actions,
				actionMaxLen:    tt.fields.actionMaxLen,
				noPages:         tt.fields.noPages,
				ActivePageIndex: tt.fields.CurrentPageNumber,
				counter:         tt.fields.counter,
			}
			m.NextPage(tt.args.crt)
		})
	}
}
