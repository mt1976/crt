package menu

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mt1976/crt/support"
)

const MaxPageRows int = 15
const RowLength int = 80
const TitleLength int = 25

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
	m := Page{title: title, pageRows: []pageRow{}, noRows: 0, prompt: promptString, actions: []string{}, actionMaxLen: 0, noPages: 0, ActivePageIndex: 0, counter: 0}
	m.AddAction(Quit)    // Add Quit action
	m.AddAction(Forward) // Add Next action
	m.AddAction(Back)    // Add Previous action
	m.pageRowCounter = 0
	return &m
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (m *Page) Add(rowContent string, altID string, dateTime string) {

	//lets clean the rowContent
	rowContent = strings.Replace(rowContent, "\n", "", -1)
	rowContent = strings.Replace(rowContent, "\r", "", -1)
	rowContent = strings.Replace(rowContent, "\t", "", -1)
	rowContent = strings.Replace(rowContent, "  ", " ", -1)
	rowContent = strings.Replace(rowContent, "  ", " ", -1)
	rowContent = strings.Replace(rowContent, "  ", " ", -1)
	rowContent = strings.Replace(rowContent, "\"", " ", -1)

	//fmt.Println(fmt.Sprintf("%v", len(rowContent)) + "[" + rowContent + "]")
	if rowContent == "" {
		return
	}
	if strings.Trim(rowContent, " ") == "" {
		return
	}
	m.counter++
	if m.counter >= MaxPageRows {
		m.counter = 0
		m.noPages++
	}
	remainder := ""
	if len(rowContent) > RowLength {
		remainder = rowContent[RowLength:]
		rowContent = rowContent[:RowLength]

		//m.pageRows = append(m.pageRows, pageRow{ID: m.counter, Content: rowContent, PageNumber: prn})
		//m.noRows++
		//m.Add(prn, "***"+remainder, altID, dateTime)
	}
	// if rowContent != "" {
	// 	m.AddAction(fmt.Sprintf("%v", pageRowNumber))
	// }
	m.pageRowCounter++
	mi := pageRow{m.pageRowCounter, rowContent, m.noPages}
	m.pageRows = append(m.pageRows, mi)
	m.noRows++
	if remainder != "" {
		//m.pageRowCounter++
		m.Add(remainder, altID, dateTime)
	}
}

// The `AddAction` function is used to add a valid action to the page. It takes a `validAction` string
// as a parameter.
func (m *Page) AddAction(validAction string) {
	if validAction == "" {
		log.Fatal(InvalidActionError)
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
func (m *Page) Display(crt *support.Crt) (nextAction string, selected pageRow) {
	//spew.Dump(m)
	crt.Clear()
	rowsDisplayed := 0
	m.AddAction(Quit) // Add Quit action
	m.AddAction(Exit)
	crt.Header(m.title)
	for i := range m.pageRows {
		//if i <= MaxPageRows {
		if m.ActivePageIndex == m.pageRows[i].PageNumber {
			rowsDisplayed++
			if m.pageRows[i].Content == "" {
				crt.Println("")
				continue
			}
			crt.Println(format(crt, m.pageRows[i]))
		}
		//}
		//m.AddAction(m.pageRows[i].Number) // Add action for each menu item
	}
	extraRows := (MaxPageRows - rowsDisplayed) + 1
	//log.Println("Extra Rows: ", extraRows)
	if extraRows > 0 {
		for i := 0; i <= extraRows; i++ {
			crt.Print(newline)
		}
	}
	crt.Break()
	//spew.Dump(m)
	//spew.Dump(crt)
	crt.InputPageInfo(m.ActivePageIndex+1, m.noPages+1)
	//crt.Print(m.prompt)
	ok := false
	for !ok {
		nextAction = crt.Input(m.prompt, "")
		if len(nextAction) > m.actionMaxLen {
			crt.InputError(InvalidActionError + "'" + nextAction + "'")
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
			crt.InputError(InvalidActionError + " '" + nextAction + "'")

		}
	}
	// if nextAction is a numnber, find the menu item
	if support.IsInt(nextAction) {
		pos, _ := strconv.Atoi(nextAction)
		return support.Upcase(nextAction), m.pageRows[pos-1]
	}

	if support.Upcase(nextAction) == Exit {
		os.Exit(0)
	}
	//spew.Dump(m)
	return support.Upcase(nextAction), pageRow{}
}

// The format function returns the first 50 characters of the content in a pageRow object.
// format returns the first 50 characters of the content in a pageRow object.
func format(crt *support.Crt, m pageRow) string {
	return m.Content
}

// NextPage moves to the next page.
// If the current page is the last page, it returns an error.
func (m *Page) NextPage(crt *support.Crt) {
	if m.ActivePageIndex == m.noPages {
		crt.InputError(noMorePagesError)
		return
	}
	m.ActivePageIndex++
}

// PreviousPage moves to the previous page.
// If the current page is the first page, it returns an error.
func (m *Page) PreviousPage(crt *support.Crt) {
	if m.ActivePageIndex == 0 {
		crt.InputError(noMorePagesError)
		return
	}
	m.ActivePageIndex--
}

// GetDebugRow returns the pageRow at the specified index.
//
// This function is used for debugging purposes.
func (m *Page) GetDebugRow(rowNo int) pageRow {
	return m.pageRows[rowNo]
}

// GetRows returns the number of rows in the page.
func (m *Page) GetRows() int {
	return m.noRows
}
