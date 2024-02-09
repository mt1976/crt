package menu

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/mt1976/admin_me/support"
)

const MaxPageRows = 15

type page struct {
	title        string
	pageRows     []pageRow
	noRows       int
	prompt       string
	actions      []string
	actionMaxLen int
	noPages      int
}

type pageRow struct {
	ID      int
	Content string
}

func NewPage(title string) *page {
	// truncate title to 25 characters
	if len(title) > 25 {
		title = title[:25] + "..."
	}
	m := page{title: title, pageRows: []pageRow{}, noRows: 0, prompt: promptString}
	return &m
}

func (m *page) AddPageRow(pageRowNumber int, rowContent string, altID string, dateTime string) {
	if m.noRows >= MaxPageRows {
		log.Fatal(m.title + " " + maxPageRowsError)
		return
	}
	if len(rowContent) > 50 {
		rowContent = rowContent[:50] + "..."
	}
	if rowContent != "" {
		m.AddAction(fmt.Sprintf("%v", pageRowNumber))
	}
	mi := pageRow{pageRowNumber, rowContent}
	m.pageRows = append(m.pageRows, mi)
	m.noRows++
}

func (m *page) AddAction(validAction string) {
	if validAction == "" {
		log.Fatal(invalidActionError)
		return
	}
	validAction = strings.ReplaceAll(validAction, " ", "")
	m.actions = append(m.actions, validAction)
	if len(validAction) > m.actionMaxLen {
		m.actionMaxLen = len(validAction)
	}
}

func (m *page) DisplayPage(crt *support.Crt) (nextAction string, selected pageRow) {
	crt.Clear()
	m.AddAction("Q") // Add Quit action
	crt.Header(m.title)
	for i := range m.pageRows {
		if m.pageRows[i].Content == "" {
			crt.Println("")
			continue
		}
		crt.Println(formatRow(crt, m.pageRows[i]))
		//m.AddAction(m.pageRows[i].Number) // Add action for each menu item
	}
	extraRows := (MaxPageRows - m.noRows) + 1
	//log.Println("Extra Rows: ", extraRows)
	for i := 0; i <= extraRows; i++ {
		crt.Print("\n")
	}
	crt.Break()
	//crt.Print(m.prompt)
	ok := false
	for !ok {
		nextAction = crt.Input(m.prompt, "")
		if len(nextAction) > m.actionMaxLen {
			crt.InputError(invalidActionError + "'" + nextAction + "'")
			//crt.Shout("Invalid action '" + crt.Bold(nextAction) + "'")
			continue
		}

		for i := range m.actions {
			if upcase(nextAction) == upcase(m.actions[i]) {
				ok = true
				break
			}
		}
		if !ok {
			//crt.Shout("Invalid action '" + crt.Bold(nextAction) + "'")
			crt.InputError(invalidActionError + " '" + nextAction + "'")

		}
	}
	// if nextAction is a numnber, find the menu item
	if IsInt(nextAction) {
		pos, _ := strconv.Atoi(nextAction)
		return upcase(nextAction), m.pageRows[pos]
	}
	//spew.Dump(m)
	return upcase(nextAction), pageRow{}
}

func upcase(s string) string {
	return strings.ToUpper(s)
}

func formatRow(crt *support.Crt, m pageRow) string {
	return m.Content[:50]
}

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
