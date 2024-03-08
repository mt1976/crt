package page

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	e "github.com/mt1976/crt/errors"
	t "github.com/mt1976/crt/language"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
)

var C = config.Configuration

// Page represents a page in a document or a user interface.
type Page struct {
	title           string    // The title of the page.
	pageRows        []pageRow // The rows of content on the page.
	noRows          int       // The number of rows on the page.
	prompt          string    // The prompt displayed to the user.
	actions         []string  // The available actions on the page.
	actionMaxLen    int       // The maximum length of an action.
	noPages         int       // The total number of pages.
	ActivePageIndex int       // The index of the active page.
	counter         int       // A counter used for tracking.
	pageRowCounter  int       // A counter used for tracking the page rows.
}

// The type "pageRow" represents a row of data for a page, with an ID and content.
// @property {int} ID - An integer representing the unique identifier of a page row.
// @property {string} Content - The "Content" property of the "pageRow" struct is a string that
// represents the content of a page row.
type pageRow struct {
	ID          int
	Content     string
	PageIndex   int
	Title       string
	AlternateID string
	DateTime    string
}

// The New function creates a new page with a truncated title and initializes other properties.
func New(title string) *Page {
	// truncate title to 25 characters
	if len(title) > C.TitleLength {
		title = title[:C.TitleLength] + t.SymTruncate
	}
	m := Page{title: title, pageRows: []pageRow{}, noRows: 0, prompt: t.TxtPrompt, actions: []string{}, actionMaxLen: 0, noPages: 0, ActivePageIndex: 0, counter: 0}
	m.AddAction(t.SymActionQuit)    // Add Quit action
	m.AddAction(t.SymActionForward) // Add Next action
	m.AddAction(t.SymActionBack)    // Add Previous action
	m.pageRowCounter = 0
	return &m
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (p *Page) Add(rowContent string, altID string, dateTime string) {
	//lets clean the rowContent
	rowContent = cleanContent(rowContent)

	if rowContent == "" {
		return
	}

	if strings.Trim(rowContent, t.Space) == "" {
		return
	}

	if rowContent == t.SymBlank {
		rowContent = ""
	}

	p.counter++
	if p.counter >= C.MaxContentRows {
		p.counter = 0
		p.noPages++
	}

	remainder := ""
	if len(rowContent) > C.TerminalWidth {
		remainder = rowContent[C.TerminalWidth:]
		rowContent = rowContent[:C.TerminalWidth]
	}

	p.pageRowCounter++
	mi := pageRow{p.pageRowCounter, rowContent, p.noPages, "", "", ""}
	p.pageRows = append(p.pageRows, mi)
	p.noRows++
	if remainder != "" {
		p.Add(remainder, altID, dateTime)
	}
}

// cleanContent removes unwanted characters from the rowContent string
func cleanContent(rowContent string) string {
	// replace \n, \r, \t, and " with empty strings
	rowContent = strings.Replace(rowContent, t.SymNewline, "", -1)
	rowContent = strings.Replace(rowContent, "\r", "", -1)
	rowContent = strings.Replace(rowContent, "\t", "", -1)
	rowContent = strings.Replace(rowContent, "\"", t.Space, -1)
	return rowContent
}

// The `AddAction` function is used to add a valid action to the page. It takes a `validAction` string
// as a parameter.
func (p *Page) AddAction(validAction string) {
	if validAction == "" {
		log.Fatal(e.ErrInvalidAction)
		return
	}
	validAction = strings.ReplaceAll(validAction, t.Space, "")
	p.actions = append(p.actions, validAction)
	if len(validAction) > p.actionMaxLen {
		p.actionMaxLen = len(validAction)
	}
}

// The `Display` function is responsible for displaying the page content to the user and handling user
// input.
func (p *Page) Display(crt *support.Crt) (nextAction string, selected pageRow) {
	exit := false
	for !exit {
		nextAction, _ := p.displayIt(crt)
		switch {
		case nextAction == t.SymActionQuit:
			exit = true
			return t.SymActionQuit, pageRow{}
		case nextAction == t.SymActionForward:
			p.NextPage(crt)
		case nextAction == t.SymActionBack:
			p.PreviousPage(crt)
		case inActions(nextAction, p.actions):
			// upcase the action
			exit = true
			if support.IsInt(nextAction) {
				return nextAction, p.pageRows[support.ToInt(nextAction)-1]
			}
			return support.Upcase(nextAction), pageRow{}
		default:
			crt.InputError(e.ErrInvalidAction + "'" + nextAction + "'")
		}
	}
	return "", pageRow{}
}

// inActions determines if the given action is in the list of available actions
func inActions(action string, actions []string) bool {
	// loop through each action in the list
	for i := range actions {
		// if the given action matches an action in the list, return true
		if action == actions[i] {
			return true
		}
	}
	// if no match was found, return false
	return false
}

// Display displays the page content to the user and handles user input.
func (p *Page) displayIt(crt *support.Crt) (nextAction string, selected pageRow) {
	crt.Clear()
	rowsDisplayed := 0
	p.AddAction(t.SymActionQuit) // Add Quit action
	p.AddAction(t.SymActionExit)
	crt.Header(p.title)
	for i := range p.pageRows {
		if p.ActivePageIndex == p.pageRows[i].PageIndex {
			rowsDisplayed++
			if p.pageRows[i].Content == "" {
				crt.Println("")
				continue
			}
			crt.Println(format(crt, p.pageRows[i]))
		}
	}
	extraRows := (C.MaxContentRows - rowsDisplayed) + 1
	if extraRows > 0 {
		for i := 0; i <= extraRows; i++ {
			crt.Print(t.SymNewline)
		}
	}
	crt.Break()

	crt.InputPageInfo(p.ActivePageIndex+1, p.noPages+1)
	ok := false
	for !ok {
		nextAction = crt.Input(p.prompt, "")
		if len(nextAction) > p.actionMaxLen {
			crt.InputError(e.ErrInvalidAction + "'" + nextAction + "'")
			continue
		}

		for i := range p.actions {
			if support.Upcase(nextAction) == support.Upcase(p.actions[i]) {
				ok = true
				break
			}
		}
		if !ok {
			crt.InputError(e.ErrInvalidAction + " '" + nextAction + "'")

		}
	}
	// if nextAction is a numnber, find the menu item
	if support.IsInt(nextAction) {
		pos, _ := strconv.Atoi(nextAction)
		return support.Upcase(nextAction), p.pageRows[pos-1]
	}

	if support.Upcase(nextAction) == t.SymActionExit {
		os.Exit(0)
	}
	return support.Upcase(nextAction), pageRow{}
}

// The format function returns the first 50 characters of the content in a pageRow object.
// format returns the first 50 characters of the content in a pageRow object.
func format(crt *support.Crt, m pageRow) string {
	return m.Content
}

// NextPage moves to the next page.
// If the current page is the last page, it returns an error.
func (p *Page) NextPage(crt *support.Crt) {
	if p.ActivePageIndex == p.noPages {
		crt.InputError(e.ErrNoMorePages)
		return
	}
	p.ActivePageIndex++
}

// PreviousPage moves to the previous page.
// If the current page is the first page, it returns an error.
func (p *Page) PreviousPage(crt *support.Crt) {
	if p.ActivePageIndex == 0 {
		crt.InputError(e.ErrNoMorePages)
		return
	}
	p.ActivePageIndex--
}

// GetDebugRow returns the pageRow at the specified index.
//
// This function is used for debugging purposes.
func (p *Page) GetDebugRow(rowNo int) pageRow {
	return p.pageRows[rowNo]
}

// GetRows returns the number of rows in the page.
func (p *Page) GetRows() int {
	return p.noRows
}

// AddFieldValuePair adds a field value pair to the page
//
// AddFieldValuePair takes two strings as arguments, where the first string represents the field name and the second string represents the field value. The function adds a row to the page with the field name on the left and the field value on the right, separated by a colon.
//
// Example:
//
//	page.AddFieldValuePair("Field Name", "Field Value")
func (p *Page) AddFieldValuePair(crt *support.Crt, key string, value string) {
	// format the field value pair
	format := "%-20s : %s\n"
	p.Add(fmt.Sprintf(format, key, value), "", "")
}

// AddColumns adds columns of data to the page
//
// AddColumns takes a variadic number of strings as arguments, where each string represents a column of data.
// The function calculates the optimal column width based on the terminal width, and then adds each column
// to the page, right-aligned.
//
// If the number of columns specified is greater than 10, an error is returned.
//
// Example:
//
//	page.AddColumns("Column 1", "Column 2", "Column 3")
func (p *Page) AddColumns(crt *support.Crt, cols ...string) {
	// Check the number of columns
	if len(cols) > 10 {
		crt.Error(e.ErrAddColumns, nil)
		os.Exit(1)
	}

	// Get the terminal width
	screenWidth := crt.Width()

	// Calculate the column width
	colSize := screenWidth / len(cols)

	// Loop through each column
	var output []string
	for i := 0; i < len(cols); i++ {
		// Get the current column
		op := cols[i]

		// Check if the column is longer than the column width
		if len(op) > colSize {
			// Truncate the column to the column width
			op = op[0:colSize]
		} else {
			// Calculate the number of spaces to add
			noToAdd := colSize - (len(op) + 1)

			// Add the spaces to the column
			op = op + strings.Repeat(t.Space, noToAdd)
		}

		// Add the column to the output slice
		output = append(output, op)
	}

	// Join the output slice into a single string and add it to the page
	p.Add(strings.Join(output, t.Space), "", "")
}

// AddColumnsTitle adds a ruler to the page, separating the columns
func (p *Page) AddColumnsTitle(crt *support.Crt, cols ...string) {
	p.AddColumns(crt, cols...)
	var output []string
	screenWidth := crt.Width()
	colSize := screenWidth / len(cols)

	for i := 0; i < len(cols); i++ {

		op := cols[i]
		if len(op) > colSize {
			op = op[0:colSize]
		} else {
			noToAdd := colSize - (len(op) + 1)
			op = op + strings.Repeat(t.Space, noToAdd)
		}

		noChars := len(op)
		op = strings.Repeat(support.TableCharacterUnderline, noChars)

		output = append(output, op)
	}

	// turn string array into sigle string
	p.Add(strings.Join(output, t.Space), "", "")
}

// SetPrompt sets the prompt for the page
func (p *Page) SetPrompt(prompt string) {
	p.prompt = prompt
}

// ResetPrompt resets the prompt to the default value
func (p *Page) ResetPrompt() {
	p.prompt = t.TxtPrompt
}

// BlankRow adds a blank row to the page
func (p *Page) BlankRow() {
	p.Add(t.SymBlank, "", "")
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (m *Page) AddOption(id int, rowContent string, altID string, dateTime string) {
	// lets clean the rowContent
	rowContent = cleanContent(rowContent)

	if rowContent == "" {
		return
	}

	if strings.Trim(rowContent, t.Space) == "" {
		return
	}

	m.counter++

	if m.counter >= C.MaxContentRows {
		m.counter = 0
		m.noPages++
	}

	if len(rowContent) > C.TerminalWidth {
		rowContent = rowContent[:C.TerminalWidth]
	}

	m.pageRowCounter++
	mi := pageRow{}
	mi.ID = id
	mi.PageIndex = m.noPages
	mi.AlternateID = altID
	mi.Title = rowContent
	mi.DateTime = dateTime
	mi.Content = formatOption(mi)
	m.AddActionInt(id)
	m.pageRows = append(m.pageRows, mi)
	m.noRows++
}

// AddActionInt adds an action to the page with the given integer value
func (m *Page) AddActionInt(validAction int) {
	m.AddAction(fmt.Sprintf("%v", validAction))
}

// The function "format" takes a pointer to a support.Crt object and a menuItem object, and returns a
// formatted string containing the menu item's ID, title, and date.
func formatOption(m pageRow) string {
	miNumber := fmt.Sprintf(support.Bold("%3v"), m.ID)

	//add Date to end of row
	miTitle := m.Title
	//padd out to 70 characters
	width := C.TerminalWidth - 7
	pad := width - (len(miTitle) + len(m.DateTime))
	if pad > 0 {
		miTitle = miTitle + strings.Repeat(t.Space, pad)
	} else {
		miTitle = miTitle[:width-(len(m.DateTime)+1)] + " | " + m.DateTime
	}

	miString := fmt.Sprintf(miNumber + ") " + miTitle)
	return miString
}
