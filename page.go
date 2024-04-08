package crt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	//	disp "github.com/buger/goterm"
	spew "github.com/davecgh/go-spew/spew"
	beep "github.com/gen2brain/beeep"
	boxr "github.com/mt1976/crt/box"
	conf "github.com/mt1976/crt/config"
	disp "github.com/mt1976/crt/display"
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
	title            string    // The title of the page.
	pageRows         []pageRow // The rows of content on the page.
	noRows           int       // The number of rows on the page.
	prompt           string    // The prompt displayed to the user.
	showOptions      bool      // The text to be displayed to the user in the case options are possible
	actions          []string  // The available actions on the page.
	actionLen        int       // The maximum length of an action.
	noPages          int       // The total number of pages.
	ActivePageIndex  int       // The index of the active page.
	counter          int       // A counter used for tracking.
	pageRowCounter   int       // A counter used for tracking the page rows.
	viewPort         *ViewPort // The viewPort object used for displaying the page.
	headerBarTop     int       // The header row top row
	headerBarContent int       // The header row content row
	headerBarBotton  int       // The header row bottom row
	footerBarTop     int       // The row where the input box starts
	footerBarInput   int       // The row where the input box is
	footerBarMessage int       // The row where the info box is
	footerBarBottom  int       // The last row of the page
	textAreaStart    int       // The row where the text area starts
	textAreaEnd      int       // The row where the text area ends
	height           int       // The height of the page
	width            int       // The width of the page
	maxContentRows   int       // The maximum number of rows available for content on the page.
	helpText         []string  // The help text to be displayed to the user
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

//func NewPage() *Page {
//	return &Page{}
//}

// The NewPage function creates a new page with a truncated title and initializes other properties.
func (t *ViewPort) NewPage(title string) *Page {
	// truncate title to 25 characters
	if len(title) > config.TitleLength {
		title = title[:config.TitleLength] + lang.SymTruncate
	}
	p := Page{title: title, pageRows: []pageRow{}, noRows: 0, prompt: lang.TxtPagingPrompt, actions: []string{}, actionLen: 0, noPages: 0, ActivePageIndex: 0, counter: 0}
	p.AddAction(lang.SymActionQuit)    // Add Quit action
	p.AddAction(lang.SymActionForward) // Add Next action
	p.AddAction(lang.SymActionBack)    // Add Previous action
	p.showOptions = false
	p.pageRowCounter = 0
	p.viewPort = t
	// Setup viewport page info
	p.height = t.height
	p.width = t.width
	p.headerBarTop = 1
	p.headerBarContent = 2
	p.headerBarBotton = 3
	p.textAreaStart = 4
	p.textAreaEnd = t.height - 4
	p.footerBarTop = t.height - 3
	p.footerBarInput = t.height - 2
	p.footerBarMessage = t.height - 1
	p.footerBarBottom = t.height
	p.maxContentRows = (t.height - 4)       // Remove the number of rows used for the footer
	p.maxContentRows = p.maxContentRows - 3 // Remove the number of rows used for the header
	p.ResetSetHelp()
	p.Clear()

	return &p
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
	if p.counter >= p.maxContentRows {
		p.counter = 0
		p.noPages++
	}

	remainder := ""
	width := p.width - 5
	if len(rowContent) > width {
		remainder = rowContent[width:]
		rowContent = rowContent[:width]
	}

	p.pageRowCounter++
	mi := pageRow{p.pageRowCounter, rowContent, p.noPages, "", "", ""}
	p.pageRows = append(p.pageRows, mi)
	p.noRows++
	if p.noRows > p.maxContentRows {
		p.AddAction(lang.SymActionForward) // Add Next action
		p.AddAction(lang.SymActionBack)    // Add Previous action
	}
	if remainder != "" {
		p.Add(remainder, altID, dateTime)
	}
}

// AddAction takes a validAction string as a parameter. The function adds the validAction to the list of available actions on the page.
func (p *Page) AddAction(validAction string) {

	if validAction == "?" {
		p.Error(errs.ErrInvalidAction, validAction)
		return
	}

	validAction = strings.ReplaceAll(validAction, lang.Space, "")

	if validAction == "" {
		log.Fatal(errs.ErrNoActionSpecified)
		return
	}
	//If the validAction is already in the list of actions, return
	if slices.Contains(p.actions, validAction) {
		//do nothing
		return
	}
	p.actions = append(p.actions, validAction)
	if len(validAction) > p.actionLen {
		p.actionLen = len(validAction)
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

	if p.counter >= p.maxContentRows {
		p.counter = 0
		p.noPages++
	}

	visible := p.width - 10
	if len(rowContent) > visible {
		rowContent = rowContent[:visible]
	}

	p.pageRowCounter++
	mi := pageRow{}
	mi.ID = id
	mi.PageIndex = p.noPages
	mi.AlternateID = altID
	mi.Title = rowContent
	mi.DateTime = dateTime
	mi.RowContent = p.formatNumberedOptionText(mi)
	p.AddAction_int(id)
	p.pageRows = append(p.pageRows, mi)
	p.noRows++
}

func (p *Page) formatNumberedOptionText(row pageRow) string {
	si := strconv.Itoa(row.ID)
	if len(si) < 4 {
		si = si + strings.Repeat(lang.Space, 4-len(si))
	}
	seq := bold(si)

	miString := fmt.Sprintf("%v) %v", seq, row.Title)
	return miString
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
	format := "%-20s : %s"
	//+ disp.Printewline
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
		p.Error(errs.ErrAddColumns, strconv.Itoa(len(columns)), "10")
		os.Exit(1)
	}

	// Calculate the column width
	colSize := p.calcColSize(columns)
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
			if noToAdd > 0 {
				op = op + strings.Repeat(lang.Space, noToAdd)
			}
		}

		// Add the column to the output slice
		output = append(output, op)
	}

	// Join the output slice into a single string and add it to the page
	p.Add(strings.Join(output, lang.Space), "", "")
}

func (p *Page) calcColSize(cols []string) int {
	// Calculate the column width
	colSize := ((p.width - 2) / len(cols))
	return colSize
}

// AddColumnsTitle adds a ruler to the page, separating the columns
func (p *Page) AddColumnsTitle(columns ...string) {
	p.AddColumns(columns...)
	var output []string

	colSize := p.calcColSize(columns)

	for i := 0; i < len(columns); i++ {

		op := columns[i]
		if len(op) > colSize {
			op = op[0:colSize]
		} else {
			noToAdd := colSize - (len(op) + 1)
			if noToAdd > 0 {
				op = op + strings.Repeat(lang.Space, noToAdd)
			}
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
	disp.Clear()
	exit := false
	for !exit {
		nextAction, _ := p.displayIt()
		switch {
		case nextAction == lang.SymActionHelp:
			p.Help()
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
			p.Error(errs.ErrInvalidAction, nextAction)
		}
	}
	return "", pageRow{}
}

func (p *Page) Clear() {
	disp.Clear()
	p.Header(p.title)
	p.Body()
	p.Footer()
}
func (p *Page) DisplayAndInput(minLen, maxLen int) (nextAction string, selected pageRow) {
	if p.prompt == "" {
		p.Error(errs.ErrNoPromptSpecified, "Please set a prompt for the page")
		os.Exit(1)
	}
	if minLen > 0 || maxLen > 0 {
		//	p.Hint(lang.TxtMinMaxLength, strconv.Itoa(minLen), strconv.Itoa(maxLen))
		p.Add(lang.SymBlank, "", "")
		p.Add("Help:", "", "")
		p.Add(p.MinMaxHint(minLen, maxLen), "", "")
	}
	drawScreen(p)

	for {

		p.PagingInfo(p.ActivePageIndex+1, p.noPages+1)

		out := p.Input(p.prompt, "")
		if isActionIn(out, lang.SymActionQuit) {
			return lang.SymActionQuit, pageRow{}
		}

		if isActionIn(out, lang.SymActionExit) {
			os.Exit(0)
		}

		if isActionIn(out, lang.SymActionHelp) {
			p.Help()
		}

		if minLen > 0 && len(out) < minLen {
			p.Error(errs.ErrInputLengthMinimum, out, strconv.Itoa(minLen))
		}

		if maxLen > 0 && len(out) > maxLen {
			p.Error(errs.ErrInputLengthMaximum, out, strconv.Itoa(maxLen), strconv.Itoa(len(out)))
		}

		if len(out) >= minLen && len(out) <= maxLen {
			return out, pageRow{}
		}
	}
}

func drawScreen(p *Page) {

	rowsDisplayed := 0

	disp.Clear()
	p.Header(p.title)
	p.Body()

	for i := range p.pageRows {
		if p.ActivePageIndex == p.pageRows[i].PageIndex {
			rowsDisplayed++
			lineNumber := (p.textAreaStart + rowsDisplayed) - 1
			if p.pageRows[i].RowContent == "" || p.pageRows[i].RowContent == lang.SymBlank {
				continue
			}
			disp.PrintAt(p.pageRows[i].RowContent, inputColumn, lineNumber)
		}
	}

	p.Footer()
	p.PagingInfo(p.ActivePageIndex+1, p.noPages+1)

}

// The `Header` function is a method of the `Crt` struct. It is responsible for printing a banner
// message to the console.
func (p *Page) Header(msg string) {
	// Print Header Line
	disp.PrintAt(p.boxPartDraw(first), startColumn, p.headerBarTop)
	width := p.width
	disp.PrintAt(p.boxPartDraw(99), startColumn, p.headerBarContent)
	disp.PrintAt(lang.TxtApplicationName, inputColumn, p.headerBarContent)
	midway := (width - len(msg)) / 2
	disp.PrintAt(msg, midway, p.headerBarContent)
	disp.PrintAt(dateTimeString(), width-(len(dateTimeString())+1), p.headerBarContent)
	disp.PrintAt(p.boxPartDraw(middle), startColumn, p.headerBarBotton)
}
func (p *Page) Body() {
	for x := 4; x < p.footerBarMessage; x++ {
		disp.PrintAt(p.FormatRowOutput(""), 0, x)
	}
	return
}

func (p *Page) Footer() {
	disp.PrintAt(p.boxPartDraw(middle), startColumn, p.footerBarTop)
	disp.PrintAt(p.boxPartDraw(99), startColumn, p.footerBarInput)
	disp.PrintAt(p.FormatRowOutput(p.prompt), startColumn, p.footerBarMessage)
	disp.PrintAt(p.boxPartDraw(last), startColumn, p.footerBarBottom)
}

// Display displays the page content to the user and handles user input.
func (p *Page) displayIt() (string, pageRow) {

	drawScreen(p)

	inputAction := ""
	ok := false
	for !ok {
		inputAction = p.Input(p.prompt, "")

		if len(inputAction) > p.actionLen {
			p.Error(errs.ErrInvalidActionLen, inputAction, strconv.Itoa(len(inputAction)), strconv.Itoa(p.actionLen))
			continue
		}

		if inputAction == lang.SymActionHelp {
			p.Help()
			continue
		}

		ok = p.viewPort.Helpers.IsActionIn(upcase(inputAction), p.actions...)
		if !ok {
			p.Error(errs.ErrInvalidAction, inputAction)
		}
	}
	// if nextAction is a numnber, find the menu item
	if isInt(inputAction) {
		pos, _ := strconv.Atoi(inputAction)
		return upcase(inputAction), p.pageRows[pos-1]
	}

	if upcase(inputAction) == lang.SymActionExit {
		os.Exit(0)
	}
	return upcase(inputAction), pageRow{}
}

// The `Input` function is a method of the `Crt` struct. It is used to display a prompt for the user for input on the
// terminal.
func (p *Page) Input(msg string, options string) string {
	mesg := msg + lang.SymPromptSymbol + lang.Space
	if p.showOptions {
		mesg = msg + lang.Space + italic(p.GetOptions(true))
		p.showOptions = false
	}

	disp.PrintAt(mesg, inputColumn, p.footerBarMessage)
	p.PagingInfo(p.ActivePageIndex, p.noPages)

	input, err := p.getUserInput()
	if err != nil {
		p.Error(errs.ErrInputFailure, err.Error())
	}

	return input
}

func (p *Page) ShowOptions() {
	p.showOptions = true
}

func (p *Page) getUserInput() (string, error) {
	disp.MoveCursor(inputColumn, p.footerBarInput)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", errs.ErrInputScannerFailure
	}
	var input string
	fmt.Sscanf(scanner.Text(), "%s", &input)
	return input, nil
}

func (p *Page) Dump(in ...string) {

	time.Sleep(1 * time.Second)

	seconds := strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "")
	filename := fmt.Sprintf("dump_%v.txt", seconds)
	thisPath, _ := os.Getwd()
	currentpath := filepath.Join(thisPath, "dumps", filename)
	f, err := os.Create(currentpath)
	if err != nil {
		panic(err)
		//p.Error(err, "Unable to create file")
	}
	defer f.Close()
	for i := range in {
		f.WriteString(in[i] + "\n")
	}
	f.WriteString("\n")
	f.WriteString(spew.Sdump(p))
	f.WriteString("\n")
	f.WriteString(fmt.Sprintf("P=%+v\n", p))
	f.WriteString(fmt.Sprintf("T=%+v\n", p.ViewPort()))
	f.WriteString("END")
	//p.Info(fmt.Sprintf("Dumped to %v", filename))
	f.Close()
}

func (p *Page) FormatRowOutput(msg string) string {
	p.viewPort.DelayIt()
	xx := fmt.Sprintf("%s %s", boxr.Upright, msg)
	// place a upright at the end of the string at the last position based on screen width
	if len(xx) < p.width {
		addChars := (p.width - len(xx)) + 1
		xx = xx + strings.Repeat(" ", addChars) + boxr.Upright
	}
	return xx
}

func (p *Page) boxPartDraw(which int) string {
	bar := strings.Repeat(boxr.Horizontal, p.width-2)
	space := strings.Repeat(lang.Space, p.width-2)
	switch which {
	case first:
		return boxr.StartLeft + bar + boxr.StartRight
	case last:
		return boxr.EndLeft + bar + boxr.EndRight
	case middle, lineBreak:
		return boxr.DividerLeft + bar + boxr.DividerRight
	default:
		return boxr.Upright + space + boxr.Upright
	}
}

func (p *Page) Break(row int) {
	disp.MoveCursor(startColumn, row)
	disp.Print(p.boxPartDraw(middle))
}

func (p *Page) PagingInfo(page, ofPages int) {
	msg := fmt.Sprintf(lang.TxtPaging, page+1, ofPages+1)
	lmsg := len(msg)
	if ofPages == 0 {
		msg = strings.Repeat(" ", lmsg)
	}
	msg = yellow(msg)
	disp.PrintAt(msg, p.width-lmsg-1, p.footerBarMessage)
}

func (p *Page) InputHintInfo(msg string) {
	lmsg := len(msg)
	disp.PrintAt(msg, p.width-lmsg-1, p.footerBarMessage)
}

func (p *Page) MinMaxHint(min, max int) string {
	if min <= 0 && max <= 0 {
		return ""
	}
	msg := fmt.Sprintf(lang.TxtMinMax, min, max)
	return msg
}

// NextPage moves to the next page.
// If the current page is the last page, it returns an error.
func (p *Page) NextPage() {
	if p.ActivePageIndex == p.noPages {
		p.Error(errs.ErrNoMorePages)
		return
	}
	p.ActivePageIndex++
}

// PreviousPage moves to the previous page.
// If the current page is the first page, it returns an error.
func (p *Page) PreviousPage() {
	if p.ActivePageIndex == 0 {
		p.Error(errs.ErrNoMorePages)
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
	p.ClearContent(p.footerBarMessage)
	pp := p.SENotice(err.Error(), red(lang.TxtWarning), msg...)
	disp.PrintAt(pp, inputColumn, p.footerBarMessage)
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := p.viewPort.Delay()
	p.viewPort.SetDelayInSec(config.DefaultErrorDelay)
	p.viewPort.DelayIt()
	p.viewPort.SetDelayInMs(oldDelay)
	p.ClearContent(p.footerBarInput)
	p.ClearContent(p.footerBarMessage)
}

func (p *Page) Info(info string, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	p.PagingInfo(p.ActivePageIndex, p.noPages)
	pp := p.SENotice(info, white(lang.TxtInfo), msg...)
	disp.PrintAt(pp, inputColumn, p.footerBarMessage)
}

func (p *Page) Hint(info string, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	p.PagingInfo(p.ActivePageIndex, p.noPages)
	pp := p.SENotice(info, cyan(lang.TxtHint), msg...)
	disp.PrintAt(pp, inputColumn, p.footerBarMessage)
}

func (p *Page) Warning(warning string, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	pp := p.SENotice(warning, yellow(lang.TxtWarning), msg...)
	disp.PrintAt(pp, inputColumn, p.footerBarMessage)
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := p.viewPort.Delay()
	p.viewPort.SetDelayInSec(config.DefaultErrorDelay)
	p.viewPort.DelayIt()
	p.viewPort.SetDelayInMs(oldDelay)
	p.ClearContent(p.footerBarInput)
	p.ClearContent(p.footerBarMessage)
}
func (p *Page) Success(message string, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	p.PagingInfo(p.ActivePageIndex, p.noPages)
	pp := p.SENotice(message, bold(lang.TxtSuccess), msg...)
	disp.PrintAt(pp, inputColumn, p.footerBarMessage)
}

func (p *Page) SENotice(errText, promptTxt string, msg ...string) string {

	if len(msg) > 0 {
		// check for enough %v strings in the error
		// if not enough then add them on the end
		noVars := strings.Count(errText, "%v")
		if noVars < len(msg) {
			errText = errText + strings.Repeat(" %v", len(msg)-noVars)
		}
	}

	qq := errText
	for i := range msg {
		qq = strings.Replace(qq, "%v", fmt.Sprintf("%v", msg[i]), 1)
	}
	errText = ("" + promptTxt + "") + qq
	return errText
}

func (p *Page) Clearline(row int) {
	//disp.MoveCursor(startColumn, row)
	disp.ClearLine(row)
}

func (p *Page) ClearContent(row int) {
	disp.PrintAt(strings.Repeat(lang.Space, p.width-4), inputColumn, row)
}

func (p *Page) GetOptions(includeDefaults bool) string {

	xx := p.actions
	if !includeDefaults {
		xx = remove(xx, lang.SymActionQuit)
		xx = remove(xx, lang.SymActionForward)
		xx = remove(xx, lang.SymActionBack)
	}

	return qQuote(strings.Join(xx, ","))
}

func remove(s []string, r string) []string {
	var rtn []string
	for _, v := range s {
		if v != r {
			return append(rtn, v)
		}
	}
	return s
}

func (p *Page) Confirmation(msg string) (bool, error) {
	return p.DisplayConfirmation(msg)
}

func (p *Page) DisplayConfirmation(msg string) (bool, error) {

	if msg == "" {
		msg = "Proceed"
	}
	for {
		p.prompt = msg
		p.AddAction("Y")
		p.AddAction("N")
		p.actions = append(p.actions, lang.SymActionHelp)
		drawScreen(p)
		choice := p.Input(msg, "Y/N")
		switch {
		case choice == "Y":
			return true, nil
		case choice == "N":
			return false, nil
		case choice == lang.SymActionHelp:
			p.Help()
		default:
			p.Error(errs.ErrInvalidAction, choice)
		}
		// if choice == lang.SymActionHelp {
		// 	p.Help()
		// 	continue
		// } else {
		// 	if upcase(choice) == "Y" {
		// 		return true, nil
		// 	} else if upcase(choice) == "N" {
		// 		return false, nil
		// 	}
		// }
	}
	//return true, nil
}

func (p *Page) SetHelp(msg []string) {
	p.helpText = msg
}

func (p *Page) GetHelp() []string {
	// TODO: if helptext is null then generate some helptext based on options etc.
	if p.helpText == nil {
		var rtn []string
		rtn = append(rtn, lang.SymBlank)
		rtn = append(rtn, "Help for "+p.title)
		rtn = append(rtn, lang.SymBlank)
		rtn = append(rtn, "The following actions are supported:")
		rtn = append(rtn, lang.SymBlank)
		for _, v := range p.actions {
			rtn = append(rtn, "- "+upcase(v))
		}
		rtn = append(rtn, lang.SymBlank)
		rtn = append(rtn, "Autogenerated : "+time.Now().Format(time.RFC3339))
		p.SetHelp(rtn)
		return rtn
	}
	return p.helpText
}

func (p *Page) ResetSetHelp() {
	p.SetHelp(nil)
}

func (p *Page) Help() {
	help := p.viewPort.NewPage("Help")
	help.Clear()
	help.Header("Help for " + p.title)
	help.Body()
	help.Footer()
	help.AddParagraph(p.GetHelp())
	//help.SetPrompt("Press Y when done")
	for {
		ok, err := help.DisplayConfirmation("Press Y when done")
		if err != nil {
			p.Error(err)
		}
		if ok {
			help.ResetSetHelp()
			p.displayIt() // Re Display the originating page
			return
		}
	}
}
