package crt

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	gtrm "github.com/buger/goterm"
	beep "github.com/gen2brain/beeep"
	conf "github.com/mt1976/crt/config"
	errs "github.com/mt1976/crt/errors"
	lang "github.com/mt1976/crt/language"
)

var config = conf.Configuration

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
	viewPort        *Crt      // The viewPort object used for displaying the page.
}

// pageRow represents a row of content on a page.
type pageRow struct {
	ID          int    // The unique identifier of the page row.
	Content     string // The content of the page row.
	PageIndex   int    // The index of the page row.
	Title       string // The title of the page row.
	AlternateID string // The alternate identifier of the page row.
	DateTime    string // The date and time of the page row.
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (p *Page) Add(rowContent string, altID string, dateTime string) {
	//lets clean the rowContent
	rowContent = cleanContent(rowContent)

	if rowContent == "" {
		return
	}

	if strings.Trim(rowContent, lang.Space) == "" {
		return
	}

	if rowContent == lang.SymBlank {
		rowContent = ""
	}

	p.counter++
	if p.counter >= config.MaxContentRows {
		p.counter = 0
		p.noPages++
	}

	remainder := ""
	if len(rowContent) > config.TerminalWidth {
		remainder = rowContent[config.TerminalWidth:]
		rowContent = rowContent[:config.TerminalWidth]
	}

	p.pageRowCounter++
	mi := pageRow{p.pageRowCounter, rowContent, p.noPages, "", "", ""}
	p.pageRows = append(p.pageRows, mi)
	p.noRows++
	if remainder != "" {
		p.Add(remainder, altID, dateTime)
	}
}

// AddAction takes a validAction string as a parameter. The function adds the validAction to the list of available actions on the page.
func (p *Page) AddAction(validAction string) {
	if validAction == "" {
		log.Fatal(errs.ErrInvalidAction)
		return
	}
	validAction = strings.ReplaceAll(validAction, lang.Space, "")
	p.actions = append(p.actions, validAction)
	if len(validAction) > p.actionMaxLen {
		p.actionMaxLen = len(validAction)
	}
}

func (p *Page) DisplayWithActions(t *Crt) (nextAction string, selected pageRow) {

	exit := false
	for !exit {
		nextAction, _ := p.displayIt(t)
		switch {
		case nextAction == lang.SymActionQuit:
			exit = true
			return lang.SymActionQuit, pageRow{}
		case nextAction == lang.SymActionForward:
			p.NextPage()
		case nextAction == lang.SymActionBack:
			p.PreviousPage()
		case isInList(nextAction, p.actions):
			// upcase the action
			exit = true
			if isInt(nextAction) {
				return nextAction, p.pageRows[toInt(nextAction)-1]
			}
			return upcase(nextAction), pageRow{}
		default:
			p.viewPort.InputError(errs.ErrInvalidAction, nextAction)
		}
	}
	return "", pageRow{}
}

func (p *Page) DisplayAndInput(minLen, maxLen int) (nextAction string, selected pageRow) {
	rowsDisplayed := 0
	p.viewPort.Clear()
	p.viewPort.Header(p.title)
	for i := range p.pageRows {
		if p.ActivePageIndex == p.pageRows[i].PageIndex {
			rowsDisplayed++
			if p.pageRows[i].Content == "" {
				p.viewPort.Println("")
				continue
			}
			p.viewPort.Println(format(p.viewPort, p.pageRows[i]))
		}
	}
	extraRows := (config.MaxContentRows - rowsDisplayed) + 1
	if extraRows > 0 {
		for i := 0; i <= extraRows; i++ {
			p.viewPort.Print(lang.SymNewline)
		}
	}
	p.viewPort.Break()
	for {

		if minLen > 0 || maxLen > 0 {
			p.Hint(lang.TxtMinMaxLength, strconv.Itoa(minLen), strconv.Itoa(maxLen))
		}

		p.viewPort.InputPagingInfo(p.ActivePageIndex+1, p.noPages+1)

		out := p.viewPort.Input("", "")
		if isActionIn(out, lang.SymActionQuit) {
			return lang.SymActionQuit, pageRow{}
		}

		if isActionIn(out, lang.SymActionExit) {
			os.Exit(0)
		}

		if minLen > 0 && len(out) < minLen {
			p.viewPort.InputError(errs.ErrInputLengthMinimum, out, strconv.Itoa(minLen))
		}

		if maxLen > 0 && len(out) > maxLen {
			p.viewPort.InputError(errs.ErrInputLengthMaximum, out, strconv.Itoa(maxLen), strconv.Itoa(len(out)))
		}

		if len(out) >= minLen && len(out) <= maxLen {
			return out, pageRow{}
		}
	}
}

// Display displays the page content to the user and handles user input.
func (p *Page) displayIt(t *Crt) (nextAction string, selected pageRow) {
	t.Clear()
	rowsDisplayed := 0
	p.AddAction(lang.SymActionQuit) // Add Quit action
	p.AddAction(lang.SymActionExit)
	t.Header(p.title)
	for i := range p.pageRows {
		if p.ActivePageIndex == p.pageRows[i].PageIndex {
			rowsDisplayed++
			if p.pageRows[i].Content == "" {
				t.Println("")
				continue
			}
			t.Println(format(t, p.pageRows[i]))
		}
	}
	extraRows := (config.MaxContentRows - rowsDisplayed) + 1
	if extraRows > 0 {
		for i := 0; i <= extraRows; i++ {
			t.Print(lang.SymNewline)
		}
	}
	t.Break()

	t.InputPagingInfo(p.ActivePageIndex+1, p.noPages+1)
	ok := false
	for !ok {
		nextAction = t.Input(p.prompt, "")
		if len(nextAction) > p.actionMaxLen {
			t.InputError(errs.ErrInvalidAction, nextAction)
			continue
		}

		for i := range p.actions {
			if upcase(nextAction) == upcase(p.actions[i]) {
				ok = true
				break
			}
		}
		if !ok {
			t.InputError(errs.ErrInvalidAction, nextAction)

		}
	}
	// if nextAction is a numnber, find the menu item
	if isInt(nextAction) {
		pos, _ := strconv.Atoi(nextAction)
		return upcase(nextAction), p.pageRows[pos-1]
	}

	if upcase(nextAction) == lang.SymActionExit {
		os.Exit(0)
	}
	return upcase(nextAction), pageRow{}
}

// NextPage moves to the next page.
// If the current page is the last page, it returns an error.
func (p *Page) NextPage() {
	if p.ActivePageIndex == p.noPages {
		p.viewPort.InputError(errs.ErrNoMorePages)
		return
	}
	p.ActivePageIndex++
}

// PreviousPage moves to the previous page.
// If the current page is the first page, it returns an error.
func (p *Page) PreviousPage() {
	if p.ActivePageIndex == 0 {
		p.viewPort.InputError(errs.ErrNoMorePages)
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
func (p *Page) AddFieldValuePair(t *Crt, key string, value string) {
	// format the field value pair
	format := "%-20s : %s" + lang.SymNewline
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
func (p *Page) AddColumns(t *Crt, columns ...string) {
	// Check the number of columns
	if len(columns) > 10 {
		t.Error(errs.ErrAddColumns)
		os.Exit(1)
	}

	// Get the terminal width
	screenWidth := t.Width()

	// Calculate the column width
	colSize := screenWidth / len(columns)

	// Loop through each column
	var output []string
	for i := 0; i < len(columns); i++ {
		// Get the current column
		op := columns[i]

		// Check if the column is longer than the column width
		if len(op) > colSize {
			// Truncate the column to the column width
			op = op[0:colSize]
		} else {
			// Calculate the number of spaces to add
			noToAdd := colSize - (len(op) + 1)

			// Add the spaces to the column
			op = op + strings.Repeat(lang.Space, noToAdd)
		}

		// Add the column to the output slice
		output = append(output, op)
	}

	// Join the output slice into a single string and add it to the page
	p.Add(strings.Join(output, lang.Space), "", "")
}

// AddColumnsTitle adds a ruler to the page, separating the columns
func (p *Page) AddColumnsTitle(t *Crt, columns ...string) {
	p.AddColumns(t, columns...)
	var output []string
	screenWidth := t.Width()
	colSize := screenWidth / len(columns)

	for i := 0; i < len(columns); i++ {

		op := columns[i]
		if len(op) > colSize {
			op = op[0:colSize]
		} else {
			noToAdd := colSize - (len(op) + 1)
			op = op + strings.Repeat(lang.Space, noToAdd)
		}

		noChars := len(op)
		op = strings.Repeat(lang.TableCharacterUnderline, noChars)

		output = append(output, op)
	}

	// turn string array into sigle string
	p.Add(strings.Join(output, lang.Space), "", "")
}

// SetPrompt sets the prompt for the page
func (p *Page) SetPrompt(prompt string) {
	p.prompt = prompt
}

// ResetPrompt resets the prompt to the default value
func (p *Page) ResetPrompt() {
	p.prompt = lang.TxtPagingPrompt
}

// BlankRow adds a blank row to the page
func (p *Page) BlankRow() {
	p.Add(lang.SymBlank, "", "")
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (p *Page) AddOption(id int, rowContent string, altID string, dateTime string) {
	// lets clean the rowContent
	rowContent = cleanContent(rowContent)

	if rowContent == "" {
		return
	}

	if strings.Trim(rowContent, lang.Space) == "" {
		return
	}

	p.counter++

	if p.counter >= conf.Configuration.MaxContentRows {
		p.counter = 0
		p.noPages++
	}

	if len(rowContent) > config.TerminalWidth {
		rowContent = rowContent[:config.TerminalWidth]
	}

	p.pageRowCounter++
	mi := pageRow{}
	mi.ID = id
	mi.PageIndex = p.noPages
	mi.AlternateID = altID
	mi.Title = rowContent
	mi.DateTime = dateTime
	mi.Content = formatOption(mi)
	p.AddActionInt(id)
	p.pageRows = append(p.pageRows, mi)
	p.noRows++
}

// AddActionInt adds an action to the page with the given integer value
func (p *Page) AddActionInt(validAction int) {
	p.AddAction(fmt.Sprintf("%v", validAction))
}

func (p *Page) Paragraph(msg []string) {
	// make sure the lines are no longer than the screen width and wrap them if they are.
	for _, s := range msg {
		s = trimRepeatingCharacters(s, lang.Space)
		p.Add(s, "", "")
	}
}

func (p *Page) Error(t *Crt, err error, msg ...string) {
	gtrm.MoveCursor(2, 23)
	//pp := t.SError(err, msg...)
	pp := t.SENotice(err.Error(), lang.TxtError, lang.TextColorRed, msg...)
	gtrm.Print(pp)
	gtrm.Flush()
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := t.Delay()
	t.SetDelayInSec(config.DefaultErrorDelay)
	t.DelayIt()
	t.SetDelayInMs(oldDelay)
	gtrm.MoveCursor(2, 22)
	gtrm.Print(lang.ConsoleClearLine)
	gtrm.MoveCursor(2, 22)
	gtrm.Print(p.prompt)
}

func (p *Page) Info(t *Crt, info string, msg ...string) {
	gtrm.MoveCursor(2, 23)
	gtrm.Print(lang.ConsoleClearLine)
	pp := t.SENotice(info, lang.TxtInfo, lang.TextColorCyan, msg...)
	gtrm.Print(pp)
	gtrm.MoveCursor(2, 22)
	gtrm.Print(p.prompt)
	gtrm.Flush()
}

func (p *Page) Hint(info string, msg ...string) {
	gtrm.MoveCursor(2, 23)
	gtrm.Print(lang.ConsoleClearLine)
	pp := p.viewPort.SENotice(info, lang.TxtHint, lang.TextStyleReset, msg...)
	gtrm.Print(pp)
	gtrm.MoveCursor(2, 22)
	gtrm.Print(p.prompt)
	gtrm.Flush()
}

func (p *Page) Warning(t *Crt, warning string, msg ...string) {
	gtrm.MoveCursor(2, 23)
	pp := t.SENotice(warning, lang.TxtWarning, lang.TextColorCyan, msg...)
	gtrm.Print(lang.ConsoleClearLine)
	gtrm.Print(pp)
	gtrm.Flush()
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := t.Delay()
	t.SetDelayInSec(config.DefaultErrorDelay)
	t.DelayIt()
	t.SetDelayInMs(oldDelay)
	gtrm.MoveCursor(2, 22)
	gtrm.Print(lang.ConsoleClearLine)
	gtrm.MoveCursor(2, 22)
	gtrm.Print(p.prompt)
}

func (p *Page) PageSuccess(t *Crt, message string, msg ...string) {
	gtrm.MoveCursor(2, 23)
	gtrm.Print(lang.ConsoleClearLine)
	pp := t.SENotice(message, lang.TxtSuccess, lang.TextColorCyan, msg...)
	gtrm.Print(pp)
	gtrm.MoveCursor(2, 22)
	gtrm.Print(lang.ConsoleClearLine)
	gtrm.MoveCursor(2, 22)
	gtrm.Print(p.prompt)
	gtrm.Flush()
}
