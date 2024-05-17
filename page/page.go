package page

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"time"

	//	disp "github.com/buger/goterm"
	spew "github.com/davecgh/go-spew/spew"
	beep "github.com/gen2brain/beeep"
	boxr "github.com/mt1976/crt/box"
	conf "github.com/mt1976/crt/config"
	dttm "github.com/mt1976/crt/datesTimes"
	errs "github.com/mt1976/crt/errors"
	lang "github.com/mt1976/crt/language"
	numb "github.com/mt1976/crt/numbers"
	actn "github.com/mt1976/crt/page/actions"
	strg "github.com/mt1976/crt/strings"
	symb "github.com/mt1976/crt/strings/symbols"
	term "github.com/mt1976/crt/terminal"
)

var config = conf.Configuration

const (
	first = iota
	middle
	last
	lineBreak
)

func (p *Page) ViewPort() term.ViewPort {
	return *p.viewPort
}

//func NewPage() *Page {
//	return &Page{}
//}

// The NewPage function creates a new page with a truncated title and initializes other properties.
func NewPage(t *term.ViewPort, title string) *Page {
	// truncate title to 25 characters
	if len(title) > config.TitleLength {
		title = title[:config.TitleLength] + symb.Truncate.Symbol()
	}
	p := Page{title: title, pageRows: []pageRow{}, noRows: 0, prompt: lang.TxtPagingPrompt, actions: []*actn.Action{}, actionLen: 0, noPages: 0, ActivePageIndex: 0, counter: 0}
	p.viewPort = t
	// Now for the more complex setup
	p.SetTitle(title)
	p.AddAction(actn.Quit)    // Add Quit action
	p.AddAction(actn.Forward) // Add Next action
	p.AddAction(actn.Back)    // Add Previous action
	p.showOptions = false
	p.pageRowCounter = 0

	// Setup viewport page info
	p.height = t.Height()
	p.width = t.Width()
	p.headerBarTop = 1
	p.headerBarContent = 2
	p.headerBarBotton = 3
	p.textAreaStart = 4
	p.textAreaEnd = t.Height() - 4
	p.footerBarTop = t.Height() - 3
	p.footerBarInput = t.Height() - 2
	p.footerBarMessage = t.Height() - 1
	p.footerBarBottom = t.Height()
	p.maxContentRows = (t.Height() - 4)     // Remove the number of rows used for the footer
	p.maxContentRows = p.maxContentRows - 3 // Remove the number of rows used for the header
	p.blockedActions = []string{}           // No Blocked Actions
	p.ResetSetHelp()
	p.Clear()

	return &p
}

func (p *Page) SetTitle(title string) {
	p.title = p.viewPort.Styles.Bold(title)
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (p *Page) Add(rowContent string, altID string, dateTime string) {
	//lets clean the rowContent
	rowContent = strg.CleanContent(rowContent)

	if rowContent == "" {
		return
	}

	if strings.Trim(rowContent, symb.Space.Symbol()) == "" {
		return
	}

	if rowContent == symb.Blank.Symbol() {
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
		p.AddAction(actn.Forward) // Add Next action
		p.AddAction(actn.Back)    // Add Previous action
	}
	if remainder != "" {
		p.Add(remainder, altID, dateTime)
	}
}

// AddAction takes a validAction string as a parameter. The function adds the validAction to the list of available actions on the page.
func (p *Page) AddAction(validAction *actn.Action) {

	if validAction.Equals("?") {
		p.Error(errs.ErrInvalidAction, validAction.Action())
		return
	}

	//validAction = strings.ReplaceAll(validAction, lang.Space.Symbol(), "")

	// if validAction == "" {
	// 	log.Fatal(errs.ErrNoActionSpecified)
	// 	return
	// }
	//If the validAction is already in the list of actions, return
	if slices.Contains(p.actions, validAction) {
		//do nothing
		return
	}
	p.actions = append(p.actions, validAction)
	if validAction.Len() > p.actionLen {
		p.actionLen = validAction.Len()
	}
}

// AddIntAction adds an action to the page with the given integer value
func (p *Page) AddIntAction(num int) {
	p.AddAction(actn.New(fmt.Sprintf("%v", num)))
}

func (p *Page) GetActions() []*actn.Action {
	return p.actions
}

func (p *Page) BlockAction(action string) {
	p.blockedActions = append(p.blockedActions, action)
}

func (p *Page) BlockIntAction(action int) {
	p.BlockAction(fmt.Sprintf("%v", action))
}

func (p *Page) UnblockAction(action string) {
	newList := []string{}
	for _, v := range p.blockedActions {
		if v != action {
			newList = append(newList, v)
		}
	}
	p.blockedActions = newList
}

func (p *Page) GetBlockedActions() []string {
	return p.blockedActions
}

func (p *Page) IsBlockedAction(action string) bool {
	return slices.Contains(p.blockedActions, action)
}

// The `Add` function is used to add a new row of data to a page. It takes four parameters:
// `pageRowNumber`, `rowContent`, `altID`, and `dateTime`.
func (p *Page) AddMenuOption(id int, rowContent string, altID string, dateTime string) {
	// lets clean the rowContent
	rowContent = strg.CleanContent(rowContent)

	if rowContent == "" {
		return
	}

	if strings.Trim(rowContent, symb.Space.Symbol()) == "" {
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
	p.AddIntAction(id)
	p.pageRows = append(p.pageRows, mi)
	p.noRows++
}

func (p *Page) formatNumberedOptionText(row pageRow) string {
	si := strconv.Itoa(row.ID)
	if len(si) < 4 {
		si = si + strings.Repeat(symb.Space.Symbol(), 4-len(si))
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
func (p *Page) AddFieldValuePair(key any, value string) {

	keyString, err := translate(key)
	if err != nil {
		p.Error(err, keyString)
		return
	}

	// format the field value pair
	format := "%-25s : %s"
	keyString = bold(keyString)
	//+ Printewline
	p.Add(fmt.Sprintf(format, keyString, value), "", "")
}

func translate(key any) (string, error) {
	var keyString string

	switch t := key.(type) {
	case string:
		keyString = key.(string)
	case lang.Text:
		keyText := key.(lang.Text)
		keyString = keyText.Text()
	default:
		errTxt := fmt.Sprintf("invalid object type [%v]", reflect.TypeOf(t).String())
		err := errors.New(errTxt)
		return "", err
	}
	return keyString, nil
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
func (p *Page) addColumns(isBold bool, columns ...string) {

	noColumns := len(columns)
	maxCols := 12
	colSize := p.calcColSize(len(columns))

	// Check if the colsize will be wide enough
	if colSize < 5 {
		p.Error(errs.ErrAddColumns, strconv.Itoa(noColumns))
		os.Exit(1)
	}

	// Check the number of columns
	if noColumns > maxCols {
		p.Error(errs.ErrAddColumns, strconv.Itoa(len(columns)), strconv.Itoa(maxCols))
		os.Exit(1)
	}

	// Loop through each column
	var output []string
	for i := 0; i < noColumns; i++ {
		// Get the current column
		op := columns[i]

		// Check if the column is longer than the column width
		if len(op) > colSize {
			// Truncate the column to the column width
			op = op[0 : colSize-1]
		} else {
			// Calculate the number of spaces to add
			noToAdd := colSize - (len(op) + 1)

			// Add the spaces to the column
			if noToAdd > 0 {
				op = op + strings.Repeat(symb.Space.Symbol(), noToAdd)
			}
		}

		// Add the column to the output slice
		output = append(output, op)
	}

	dsp := strings.Join(output, symb.Space.Symbol())
	//if isBold {
	//	dsp = p.viewPort.Styles.Bold(dsp)
	//}

	// Join the output slice into a single string and add it to the page
	p.Add(dsp, "", "")
}

func (p *Page) AddColumns(columns ...string) {

	p.addColumns(false, columns...)

}

func (p *Page) calcColSize(nocols int) int {
	// Calculate the column width
	colSize := ((p.width - 4) / nocols)
	//spew.Dump(".............", "colsize", colSize, "width", p.width, "nocols", nocols, ".............")
	return colSize
}

// AddColumnsTitle adds a ruler to the page, separating the columns
func (p *Page) AddColumnsTitle(columns ...string) {

	p.addColumns(true, columns...)

	var output []string
	noCols := len(columns)

	colSize := p.calcColSize(noCols)

	for i := 0; i < noCols; i++ {

		// noChars := len(op)
		op := strings.Repeat(lang.Underline.Text(), colSize-1)

		output = append(output, op)
	}

	// turn string array into sigle string
	rtn := strings.Join(output, symb.Space.Symbol())
	//rtn = p.viewPort.Styles.Bold(rtn)
	p.Add(rtn, "", "")
}

// AddBlankRow adds a blank row to the page
func (p *Page) AddBlankRow() {
	p.Add(symb.Blank.Symbol(), "", "")
}

func (p *Page) AddParagraph(msg []string) {
	// make sure the lines are no longer than the screen width and wrap them if they are.
	for _, s := range msg {
		s = strg.TrimRepeatingCharacters(s, symb.Space.Symbol())
		p.Add(s, "", "")
	}
}

func (p *Page) AddParagraphString(msg string) {
	// make sure the lines are no longer than the screen width and wrap them if they are.
	var msgSlice []string
	msgSlice = append(msgSlice, msg)
	p.AddParagraph(msgSlice)
}

func (p *Page) Display_Actions() (nextAction *actn.Action) {
	//t := p.viewPort.Formatters.Upcase
	Clear()
	exit := false
	for !exit {
		nextAction, _ := p.displayIt()
		switch {
		case nextAction.Is(actn.Help):
			p.Help()
		case nextAction.Is(actn.Quit):
			exit = true
			return actn.Quit
		case nextAction.Is(actn.Forward):
			p.Forward()
		case nextAction.Is(actn.Back):
			p.Back()
		case actn.IsInActions(&nextAction, p.actions):
			// upcase the action
			exit = true
			// if isInt(nextAction) {
			// 	return nextAction
			// }
			return &nextAction
		default:
			p.Error(errs.ErrInvalidAction, nextAction.Action())
		}
	}
	return &actn.Action{}
}

func (p *Page) Clear() {
	Clear()
	p.Header(p.title)
	p.Body()
	p.Footer()
}

func (p *Page) Display_Input(minLen, maxLen int) (nextAction string, selected pageRow) {
	if p.prompt.Text() == "" {
		p.Error(errs.ErrNoPromptSpecified, lang.SetPrompt.Text())
		os.Exit(1)
	}
	if minLen > 0 || maxLen > 0 {
		//	p.Hint(lang.TxtMinMaxLength, strconv.Itoa(minLen), strconv.Itoa(maxLen))
		p.Add(symb.Blank.Symbol(), "", "")
		p.Add(lang.HelpHint.Text(), "", "")
		p.Add(p.minMaxHint(minLen, maxLen), "", "")
	}
	drawScreen(p)

	for {

		p.PagingInfo(p.ActivePageIndex+1, p.noPages+1)

		out := p.Input(p.prompt, "")
		if actn.IsActionIn(out, actn.Quit) {
			return actn.Quit.Action(), pageRow{}
		}

		if actn.IsActionIn(out, actn.Exit) {
			os.Exit(0)
		}

		if actn.IsActionIn(out, actn.Help) {
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

	Clear()
	p.Header(p.title)
	p.Body()

	for i := range p.pageRows {
		if p.ActivePageIndex == p.pageRows[i].PageIndex {
			rowsDisplayed++
			lineNumber := (p.textAreaStart + rowsDisplayed) - 1
			if p.pageRows[i].RowContent == "" || symb.Blank.Equals(p.pageRows[i].RowContent) {
				continue
			}
			PrintAt(p.pageRows[i].RowContent, term.InputColumn, lineNumber)
		}
	}
	p.Footer()
	p.PagingInfo(p.ActivePageIndex+1, p.noPages+1)
}

// The `Header` function is a method of the `Crt` struct. It is responsible for printing a banner
// message to the console.
func (p *Page) Header(msg string) {
	// Print Header Line
	PrintAt(p.boxPartDraw(first), term.StartColumn, p.headerBarTop)
	width := p.width
	PrintAt(p.boxPartDraw(99), term.StartColumn, p.headerBarContent)
	PrintAt(lang.ApplicationName.Text(), term.InputColumn, p.headerBarContent)
	midway := (width - len(msg)) / 2
	PrintAt(msg, midway, p.headerBarContent)
	PrintAt(dttm.DateTimeString(), width-(len(dttm.DateTimeString())+1), p.headerBarContent)
	PrintAt(p.boxPartDraw(middle), term.StartColumn, p.headerBarBotton)
}
func (p *Page) Body() {
	for x := 4; x < p.footerBarMessage; x++ {
		PrintAt(p.FormatRowOutput(""), 0, x)
	}
}

func (p *Page) Footer() {
	PrintAt(p.boxPartDraw(middle), term.StartColumn, p.footerBarTop)
	PrintAt(p.boxPartDraw(99), term.StartColumn, p.footerBarInput)
	PrintAt(p.FormatRowOutput(p.prompt.Text()), term.StartColumn, p.footerBarMessage)
	PrintAt(p.boxPartDraw(last), term.StartColumn, p.footerBarBottom)
}

// Display displays the page content to the user and handles user input.
func (p *Page) displayIt() (actn.Action, pageRow) {

	drawScreen(p)

	inputAction := ""
	ok := false
	for !ok {
		inputAction = p.Input(p.prompt, "")

		if len(inputAction) > p.actionLen {
			p.Error(errs.ErrInvalidActionLen, inputAction, strconv.Itoa(len(inputAction)), strconv.Itoa(p.actionLen))
			continue
		}

		if inputAction == actn.Help.Action() {
			p.Help()
			continue
		}

		ok = actn.IsActionIn(strg.Upcase(inputAction), p.actions...)
		if !ok {
			p.Error(errs.ErrInvalidAction, inputAction)
		}
	}
	// if nextAction is a numnber, find the menu item
	if numb.IsInt(inputAction) {
		pos, _ := strconv.Atoi(inputAction)
		rtnAction := actn.New(inputAction)
		return *rtnAction, p.pageRows[pos-1]
	}

	if actn.Exit.Equals(inputAction) {
		os.Exit(0)
	}
	return *actn.New(inputAction), pageRow{}
}

// The `Input` function is a method of the `Crt` struct. It is used to display a prompt for the user for input on the
// terminal.
func (p *Page) Input(msg *lang.Text, options string) string {
	mesg := msg.Text() + symb.PromptSymbol.Symbol() + symb.Space.Symbol()
	if p.showOptions {
		mesg = msg.Text() + symb.Space.Symbol() + strg.Italic(p.GetOptions(true))
		p.showOptions = false
	}

	PrintAt(mesg, term.InputColumn, p.footerBarMessage)
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
	MoveCursor(term.InputColumn, p.footerBarInput)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", errs.ErrInputScannerFailure
	}
	var input string
	fmt.Sscanf(scanner.Text(), "%s", &input)
	return input, nil
}

func (p *Page) Dump(in ...string) {

	// Only proceed if page dumping is active in the config file
	if !config.PageDumpActive {
		return
	}
	// OK Proceed - Sleep for a second to stop dumping multiple files with same timestamp
	time.Sleep(1 * time.Second)

	seconds := strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "")

	filename := fmt.Sprintf("dump_%v.txt", seconds)

	thisPath, _ := os.Getwd()
	currentpath := filepath.Join(thisPath, filename)

	if config.PageDumpPath != "" {
		//thisPath = thisPath + config.PageDumpPath
		currentpath = filepath.Join(thisPath, config.PageDumpPath, filename)
	}

	f, err := os.Create(currentpath)
	if err != nil {
		fmt.Printf("FAILED to create file %v", currentpath)
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
	space := strings.Repeat(symb.Space.Symbol(), p.width-2)
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
	MoveCursor(term.StartColumn, row)
	Print(p.boxPartDraw(middle))
}

func (p *Page) AddBreakRow() {
	line := strings.Repeat("-", p.width-4)
	p.Add(line, "", "")
}

func (p *Page) PagingInfo(page, ofPages int) {
	msg := fmt.Sprintf(lang.Paging.Text(), page+1, ofPages+1)
	lmsg := len(msg)
	if ofPages == 0 {
		msg = strings.Repeat(" ", lmsg)
	}
	msg = p.viewPort.Styles.Yellow(msg)
	PrintAt(msg, p.width-lmsg-1, p.footerBarMessage)
}

func (p *Page) InputHintInfo(msg *lang.Text) {
	//lmsg := msg.Len()
	PrintAt(msg.Text(), p.width-msg.Len()-1, p.footerBarMessage)
}

func (p *Page) minMaxHint(min, max int) string {
	if min <= 0 && max <= 0 {
		return ""
	}
	msg := fmt.Sprintf(lang.MinMax.Text(), min, max)
	return msg
}

// Forward moves to the next page.
// If the current page is the last page, it returns an error.
func (p *Page) Forward() {
	if p.ActivePageIndex == p.noPages {
		p.Error(errs.ErrNoMorePages)
		return
	}
	p.ActivePageIndex++
}

// Back moves to the previous page.
// If the current page is the first page, it returns an error.
func (p *Page) Back() {
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
func (p *Page) SetPrompt(prompt *lang.Text) {
	p.prompt = prompt
}

// SetPrompt sets the prompt for the page
func (p *Page) setPromptString(prompt string) {
	p.prompt = lang.New(prompt)
}

// ResetPrompt resets the prompt to the default value
func (p *Page) ResetPrompt() {
	p.prompt = lang.TxtPagingPrompt
}

func (p *Page) Error(err error, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	pp := p.formatMessage(err.Error(), p.viewPort.Styles.Red(lang.Warning.Text()), msg...)
	PrintAt(pp, term.InputColumn, p.footerBarMessage)
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := p.viewPort.Delay()
	p.viewPort.SetDelayInSec(config.DefaultErrorDelay)
	p.viewPort.DelayIt()
	p.viewPort.SetDelayInMs(oldDelay)
	p.ClearContent(p.footerBarInput)
	p.ClearContent(p.footerBarMessage)
}

func (p *Page) Info(info *lang.Text, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	p.PagingInfo(p.ActivePageIndex, p.noPages)
	pp := p.formatMessage(info.Text(), p.viewPort.Styles.White(lang.Info.Text()), msg...)
	PrintAt(pp, term.InputColumn, p.footerBarMessage)
}

func (p *Page) Hint(info *lang.Text, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	p.PagingInfo(p.ActivePageIndex, p.noPages)
	pp := p.formatMessage(info.Text(), p.viewPort.Styles.Cyan(lang.Hint.Text()), msg...)
	PrintAt(pp, term.InputColumn, p.footerBarMessage)
}

func (p *Page) Warning(warning lang.Text, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	pp := p.formatMessage(warning.Text(), p.viewPort.Styles.Yellow(lang.Warning.Text()), msg...)
	PrintAt(pp, term.InputColumn, p.footerBarMessage)
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := p.viewPort.Delay()
	p.viewPort.SetDelayInSec(config.DefaultErrorDelay)
	p.viewPort.DelayIt()
	p.viewPort.SetDelayInMs(oldDelay)
	p.ClearContent(p.footerBarInput)
	p.ClearContent(p.footerBarMessage)
}
func (p *Page) Success(message *lang.Text, msg ...string) {
	p.ClearContent(p.footerBarMessage)
	p.PagingInfo(p.ActivePageIndex, p.noPages)
	pp := p.formatMessage(message.Text(), bold(lang.Success.Text()), msg...)
	PrintAt(pp, term.InputColumn, p.footerBarMessage)
}

func (p *Page) formatMessage(errText, promptTxt string, msg ...string) string {

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
	//MoveCursor(term.StartColumn, row)
	ClearLine(row)
}

func (p *Page) ClearContent(row int) {
	PrintAt(strings.Repeat(symb.Space.Symbol(), p.width-4), term.InputColumn, row)
}

func (p *Page) GetOptions(includeDefaults bool) string {

	var xx []string
	for _, option := range p.actions {
		switch option {
		case actn.Help:
			continue
		case actn.Quit:
			continue
		case actn.Forward:
			continue
		case actn.Back:
			continue
		case actn.Yes:
			continue
		case actn.No:
			continue
		default:
			xx = append(xx, option.Action())
		}
	}
	// xx := p.actions
	// if !includeDefaults {
	// 	xx = removeOption(xx, lang.Quit.Action())
	// 	xx = removeOption(xx, lang.Forward.Action())
	// 	xx = removeOption(xx, lang.Back.Action())
	// }

	return strg.QQuote(strings.Join(xx, ","))
}

func removeOption(s []string, r string) []string {
	var rtn []string
	for _, v := range s {
		if v != r {
			return append(rtn, v)
		}
	}
	return s
}

func (p *Page) Display_Confirmation(msg *lang.Text) (bool, error) {

	if msg.IsEmpty() {
		msg = lang.Proceed
	}
	for {
		//	p.prompt = msg
		p.SetPrompt(msg)
		p.AddAction(actn.Yes)
		p.AddAction(actn.No)
		p.actions = append(p.actions, actn.Help)
		drawScreen(p)
		choice := p.Input(msg, actn.Yes.Action()+actn.No.Action())
		switch {
		case actn.Yes.Equals(choice):
			return true, nil
		case actn.No.Equals(choice):
			return false, nil
		case actn.Forward.Equals(choice) && actn.IsInActions(actn.Forward, p.actions):
			p.Forward()
		case actn.Back.Equals(choice) && actn.IsInActions(actn.Back, p.actions):
			p.Back()
		case actn.Help.Equals(choice):
			if !p.IsBlockedAction(actn.Help.Action()) {
				p.Help()
				continue
			}
			fallthrough
		default:
			p.Error(errs.ErrInvalidAction, choice)
		}
	}
}

func (p *Page) SetHelp(msg []string) {
	p.helpText = msg
}

func (p *Page) GetHelp() []string {
	if p.helpText == nil {
		var rtn []string
		rtn = append(rtn, symb.Blank.Symbol())
		rtn = append(rtn, lang.HelpFor.Text()+p.title)
		rtn = append(rtn, symb.Blank.Symbol())
		rtn = append(rtn, lang.HelpSupportedActions.Text())
		rtn = append(rtn, symb.Blank.Symbol())
		for _, v := range p.actions {
			rtn = append(rtn, symb.Bullet.Symbol()+strg.Upcase(v.Action()))
		}
		rtn = append(rtn, symb.Blank.Symbol())
		rtn = append(rtn, lang.HelpAutoGenerated.Text()+time.Now().Format(time.RFC822))
		p.SetHelp(rtn)
		return rtn
	}
	return p.helpText
}

func (p *Page) ResetSetHelp() {
	p.SetHelp(nil)
}

func (p *Page) Help() {
	help := NewPage(p.viewPort, lang.HelpPageTitle.Text())
	help.Clear()
	help.Header(lang.HelpFor.Text() + p.title)
	help.Body()
	help.Footer()
	help.AddParagraph(p.GetHelp())
	help.AddAction(actn.Yes)
	help.BlockAction(actn.Help.Action())
	//help.SetPrompt("Press Y when done")
	prompt := lang.HelpPromptSinglePage

	if len(p.actions) > 10 {
		prompt = lang.HelpPromptMultiPage
		help.AddAction(actn.Back)
		help.AddAction(actn.Forward)
	}

	for {
		ok, err := help.Display_Confirmation(prompt)
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

func bold(s string) string {
	return s
}
