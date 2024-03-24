package crt

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	gtrm "github.com/buger/goterm"
	beep "github.com/gen2brain/beeep"
	boxr "github.com/mt1976/crt/box"
	conf "github.com/mt1976/crt/config"
	errs "github.com/mt1976/crt/errors"
	lang "github.com/mt1976/crt/language"
)

var config = conf.Configuration

const (
	first = iota
	middle
	last
	lineBreak
)

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
	viewPort        *ViewPort // The viewPort object used for displaying the page.
}

// pageRow represents a row of content on a page.
type pageRow struct {
	ID          int    // The unique identifier of the page row.
	RowContent  string // The content of the page row.
	PageIndex   int    // The index of the page row.
	Title       string // The title of the page row.
	AlternateID string // The alternate identifier of the page row.
	DateTime    string // The date and time of the page row.
}

func (p *Page) ViewPort() ViewPort {
	return *p.viewPort
}

func NewPage() *Page {
	return &Page{}
}

// The NewTitledPage function creates a new page with a truncated title and initializes other properties.
func (t *ViewPort) NewTitledPage(title string) *Page {
	// truncate title to 25 characters
	if len(title) > config.TitleLength {
		title = title[:config.TitleLength] + lang.SymTruncate
	}
	m := Page{title: title, pageRows: []pageRow{}, noRows: 0, prompt: lang.TxtPagingPrompt, actions: []string{}, actionMaxLen: 0, noPages: 0, ActivePageIndex: 0, counter: 0}
	m.AddAction(lang.SymActionQuit) // Add Quit action
	//m.AddAction(lang.SymActionForward) // Add Next action
	//m.AddAction(lang.SymActionBack)    // Add Previous action
	m.pageRowCounter = 0
	m.viewPort = t
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
	if p.noRows > config.MaxContentRows {
		p.AddAction(lang.SymActionForward) // Add Next action
		p.AddAction(lang.SymActionBack)    // Add Previous action
	}
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
	//If the validAction is already in the list of actions, return
	if slices.Contains(p.actions, validAction) {
		//do nothing
		return
	}
	validAction = strings.ReplaceAll(validAction, lang.Space, "")
	p.actions = append(p.actions, validAction)
	if len(validAction) > p.actionMaxLen {
		p.actionMaxLen = len(validAction)
	}
}

// AddAction_int adds an action to the page with the given integer value
func (p *Page) AddAction_int(validAction int) {
	p.AddAction(fmt.Sprintf("%v", validAction))
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (p *Page) AddMenuOption(id int, rowContent string, altID string, dateTime string) {
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
	mi.RowContent = formatOption(mi)
	p.AddAction_int(id)
	p.pageRows = append(p.pageRows, mi)
	p.noRows++
}

// AddFieldValuePair adds a field value pair to the page
//
// AddFieldValuePair takes two strings as arguments, where the first string represents the field name and the second string represents the field value. The function adds a row to the page with the field name on the left and the field value on the right, separated by a colon.
//
// Example:
//
//	page.AddFieldValuePair("Field Name", "Field Value")
func (p *Page) AddFieldValuePair(key string, value string) {
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
func (p *Page) AddColumns(columns ...string) {
	// Check the number of columns
	if len(columns) > 10 {
		p.viewPort.Error(errs.ErrAddColumns)
		os.Exit(1)
	}

	// Get the terminal width
	screenWidth := p.viewPort.Width()

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
func (p *Page) AddColumnsTitle(columns ...string) {
	p.AddColumns(columns...)
	var output []string
	screenWidth := p.viewPort.Width()
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

// AddBlankRow adds a blank row to the page
func (p *Page) AddBlankRow() {
	p.Add(lang.SymBlank, "", "")
}

func (p *Page) AddParagraph(msg []string) {
	// make sure the lines are no longer than the screen width and wrap them if they are.
	for _, s := range msg {
		s = trimRepeatingCharacters(s, lang.Space)
		p.Add(s, "", "")
	}
}

func (p *Page) DisplayWithActions() (nextAction string, selected pageRow) {

	exit := false
	for !exit {
		nextAction, _ := p.displayIt()
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
			if p.pageRows[i].RowContent == "" {
				p.viewPort.Println("")
				continue
			}
			p.viewPort.Println(p.pageRows[i].RowContent)
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
func (p *Page) displayIt() (nextAction string, selected pageRow) {
	gtrm.Clear()
	//gtrm.MoveCursor(col, 0)
	// header row
	//hr := gtrm.Color(strings.Repeat("*", config.TerminalWidth), gtrm.CYAN)
	//gtrm.Println(hr)
	p.Header(p.title)
	//gtrm.MoveCursor(col, 2)
	//gtrm.Println(p.title)
	//gtrm.MoveCursor(col, 3)
	//gtrm.Println(hr)
	//p.viewPort.Clear()
	rowsDisplayed := 0
	offset := 4
	//p.AddAction(lang.SymActionQuit) // Add Quit action
	//p.AddAction(lang.SymActionExit)
	//p.viewPort.Header(p.title)
	for i := range p.pageRows {
		if p.ActivePageIndex == p.pageRows[i].PageIndex {
			rowsDisplayed++
			if p.pageRows[i].RowContent == "" {
				gtrm.MoveCursor(col, offset+i)
				gtrm.Println(p.viewPort.Format("", ""))
				//	p.viewPort.Println("")
				continue
			}
			//p.viewPort.Println(p.pageRows[i].RowContent)
			gtrm.MoveCursor(col, offset+i)
			gtrm.Println(p.viewPort.Format(p.pageRows[i].RowContent, ""))
		}
	}
	extraRows := (config.MaxContentRows - rowsDisplayed)
	if extraRows > 0 {
		for i := 0; i <= extraRows; i++ {
			//p.viewPort.Print(lang.SymNewline)
			gtrm.MoveCursor(col, rowsDisplayed+i+offset)
			gtrm.Println(p.viewPort.Format("", ""))
		}
	}
	//p.viewPort.Break()
	gtrm.MoveCursor(col, 21)
	gtrm.Println(p.row(middle))
	gtrm.MoveCursor(col, 24)
	gtrm.Println(p.row(last))
	//	gtrm.MoveCursor(2, 4)
	p.PagingInfo(p.ActivePageIndex+1, p.noPages+1)
	p.Hint(lang.TxtValidActions, strings.Join(p.actions, ","))
	gtrm.Flush()
	ok := false
	for !ok {
		nextAction = p.viewPort.Input(p.prompt, "")
		if len(nextAction) > p.actionMaxLen {
			p.viewPort.InputError(errs.ErrInvalidAction, nextAction)
			continue
		}

		for i := range p.actions {
			if upcase(nextAction) == upcase(p.actions[i]) {
				ok = true
				break
			}
		}
		if !ok {
			p.viewPort.InputError(errs.ErrInvalidAction, nextAction)

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

// The `Header` function is a method of the `Crt` struct. It is responsible for printing a banner
// message to the console.
func (p *Page) Header(msg string) {
	// Print Header Line
	gtrm.MoveCursor(col, 1)
	gtrm.Println(p.row(first))
	gtrm.MoveCursor(col, 2)
	width := p.viewPort.Width()
	gtrm.Println(p.viewPort.Format(lang.TxtApplicationName, ""))
	midway := (width - len(msg)) / 2
	gtrm.MoveCursor(midway, 2)
	gtrm.Print(msg)
	gtrm.MoveCursor(width-len(dateTimeString()), 2)
	gtrm.Print(dateTimeString())
	gtrm.MoveCursor(width, 2)
	//gtrm.Println(lang.BoxCharacterUpright)
	gtrm.MoveCursor(col, 3)
	gtrm.Println(p.row(middle))
}

// The `Input` function is a method of the `Crt` struct. It is used to display a prompt for the user for input on the
// terminal.
func (p *Page) Input(msg string, options string) (output string) {
	// gtrm.MoveCursor(col, 21)
	// gtrm.Print(p.row(Middle))
	gtrm.MoveCursor(col, 22)
	mesg := msg
	//T.Format(msg, "")
	if options != "" {
		mesg = (p.viewPort.Format(msg, "") + pQuote(bold(options)))
	}
	mesg = mesg + lang.SymPromptSymbol
	//mesg = p.viewPort.Format(mesg, "")
	//T.Print(mesg)
	gtrm.Print(mesg)
	gtrm.Flush()
	var out string
	fmt.Scan(&out)
	output = out
	return output
}

func (p *Page) row(which int) string {
	bar := strings.Repeat(boxr.Horizontal, p.viewPort.width-2)
	switch which {
	case first:
		return boxr.StartLeft + bar + boxr.StartRight
	case last:
		return boxr.EndLeft + bar + boxr.EndRight
	case middle, lineBreak:
		return boxr.DividerLeft + bar + boxr.DividerRight
	default:
		return boxr.Upright + bar + boxr.Upright
	}
}

func (p *Page) Break(row int) {
	gtrm.MoveCursor(col, row)
	gtrm.Println(p.row(middle))
}

func (p *Page) PagingInfo(page, ofPages int) {
	msg := fmt.Sprintf(lang.TxtPaging, page, ofPages)
	lmsg := len(msg)
	gtrm.MoveCursor(p.viewPort.width-lmsg-1, 22)
	gtrm.Print(""+gtrm.Color(msg, gtrm.YELLOW), "")
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

// SetPrompt sets the prompt for the page
func (p *Page) SetPrompt(prompt string) {
	p.prompt = prompt
}

// ResetPrompt resets the prompt to the default value
func (p *Page) ResetPrompt() {
	p.prompt = lang.TxtPagingPrompt
}

func (p *Page) Error(err error, msg ...string) {
	gtrm.MoveCursor(col, 23)
	//pp := t.SError(err, msg...)
	pp := p.viewPort.SENotice(err.Error(), lang.TxtError, p.viewPort.Styles.Red, msg...)
	gtrm.Print(pp)
	gtrm.Flush()
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := p.viewPort.Delay()
	p.viewPort.SetDelayInSec(config.DefaultErrorDelay)
	p.viewPort.DelayIt()
	p.viewPort.SetDelayInMs(oldDelay)
	p.Clearline(22)
	gtrm.Print(p.prompt)
}

func (p *Page) Info(info string, msg ...string) {
	gtrm.MoveCursor(col, 23)
	gtrm.Print(p.viewPort.Styles.ClearLine)
	pp := p.viewPort.SENotice(info, lang.TxtInfo, p.viewPort.Styles.Cyan, msg...)
	gtrm.Print(pp)
	p.Clearline(22)
	gtrm.Print(p.prompt)
	gtrm.Flush()
}

func (p *Page) Hint(info string, msg ...string) {
	gtrm.MoveCursor(col, 23)
	gtrm.Print(p.viewPort.Styles.ClearLine)
	pp := p.viewPort.SENotice(info, lang.TxtHint, p.viewPort.Styles.Reset, msg...)
	gtrm.Print(pp)
	p.Clearline(22)
	gtrm.Print(p.prompt)
	gtrm.Flush()
}

func (p *Page) Warning(warning string, msg ...string) {
	gtrm.MoveCursor(col, 23)
	pp := p.viewPort.SENotice(warning, lang.TxtWarning, p.viewPort.Styles.Cyan, msg...)
	gtrm.Print(p.viewPort.Styles.ClearLine)
	gtrm.Print(pp)
	gtrm.Flush()
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := p.viewPort.Delay()
	p.viewPort.SetDelayInSec(config.DefaultErrorDelay)
	p.viewPort.DelayIt()
	p.viewPort.SetDelayInMs(oldDelay)
	p.Clearline(22)
	gtrm.Print(p.prompt)
}

func (p *Page) Clearline(row int) {
	gtrm.MoveCursor(col, row)
	gtrm.Print(strings.Repeat(lang.Space, config.TerminalWidth))
	gtrm.MoveCursor(col, row)
}
func (p *Page) Success(message string, msg ...string) {
	gtrm.MoveCursor(col, 23)
	gtrm.Print(p.viewPort.Styles.ClearLine)
	pp := p.viewPort.SENotice(message, lang.TxtSuccess, p.viewPort.Styles.Cyan, msg...)
	gtrm.Print(pp)
	gtrm.MoveCursor(col, 22)
	gtrm.Print(p.viewPort.Styles.ClearLine)
	gtrm.MoveCursor(col, 22)
	gtrm.Print(p.prompt)
	gtrm.Flush()
}
