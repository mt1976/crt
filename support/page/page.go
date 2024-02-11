package menu

import (
	"log"
	"strconv"
	"strings"

	"github.com/mt1976/admin_me/support"
)

const MaxPageRows = 15
const RowLength = 999
const TitleLength = 25

// The "Page" type represents a Page with a title, rows of data, a prompt, a list of actions, and
// information about the number of rows and pages.
// @property {string} title - The title property represents the title of the Page.
// @property {[]pageRow} pageRows - The `pageRows` property is a slice of `pageRow` objects. It
// represents the rows of content on the Page.
// @property {int} noRows - The `noRows` property represents the number of rows in the Page. It
// indicates how many rows of data can be displayed on a single Page.
// @property {string} prompt - The prompt is a string that represents the message or question displayed
// to the user to prompt them for input or action.
// @property {[]string} actions - The "actions" property is a slice of strings that represents the
// available actions for the Page. Each string in the slice represents an action that the user can take
// on the Page.
// @property {int} actionMaxLen - The property "actionMaxLen" represents the maximum length of the
// actions in the "actions" slice.
// @property {int} noPages - The "noPages" property represents the total number of pages in the Page
// structure.
type Page struct {
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

// The type "pageRow" represents a row of data for a page, with an ID and content.
// @property {int} ID - An integer representing the unique identifier of a page row.
// @property {string} Content - The "Content" property of the "pageRow" struct is a string that
// represents the content of a page row.
type pageRow struct {
	ID         int
	Content    string
	PageNumber int
}

// The New function creates a new page with a truncated title and initializes other properties.
func New(title string) *Page {
	// truncate title to 25 characters
	if len(title) > TitleLength {
		title = title[:TitleLength] + "..."
	}
	m := Page{title: title, pageRows: []pageRow{}, noRows: 0, prompt: promptString, actions: []string{}, actionMaxLen: 0, noPages: 0, CurrentPageNumber: 1, counter: 0}
	m.AddAction("Q") // Add Quit action
	m.AddAction("F") // Add Next action
	m.AddAction("B") // Add Previous action
	return &m
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (m *Page) Add(pageRowNumber int, rowContent string, altID string, dateTime string) {
	m.counter++
	if m.counter >= MaxPageRows {
		m.counter = 0
		m.noPages++
	}
	if len(rowContent) > RowLength {
		rowContent = rowContent[:RowLength] + "..."
	}
	// if rowContent != "" {
	// 	m.AddAction(fmt.Sprintf("%v", pageRowNumber))
	// }
	mi := pageRow{pageRowNumber, rowContent, m.noPages + 1}
	m.pageRows = append(m.pageRows, mi)
	m.noRows++
}

// The `AddAction` function is used to add a valid action to the page. It takes a `validAction` string
// as a parameter.
func (m *Page) AddAction(validAction string) {
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

// The `Display` function is responsible for displaying the page content to the user and handling user
// input.
func (m *Page) Display(crt *support.Crt, pageNumber int) (nextAction string, selected pageRow) {
	crt.Clear()
	m.AddAction("Q") // Add Quit action
	crt.Header(m.title)
	for i := range m.pageRows {
		if m.pageRows[i].Content == "" {
			crt.Println("")
			continue
		}
		crt.Println(format(crt, m.pageRows[i]))
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
			if support.Upcase(nextAction) == support.Upcase(m.actions[i]) {
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
	if support.IsInt(nextAction) {
		pos, _ := strconv.Atoi(nextAction)
		return support.Upcase(nextAction), m.pageRows[pos]
	}
	//spew.Dump(m)
	return support.Upcase(nextAction), pageRow{}
}

// The format function returns the first 50 characters of the content in a pageRow object.
func format(crt *support.Crt, m pageRow) string {
	return m.Content[:RowLength]
}
