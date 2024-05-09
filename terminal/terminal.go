package terminal

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	gtrm "github.com/buger/goterm"
	beep "github.com/gen2brain/beeep"
	boxr "github.com/mt1976/crt/box"
	conf "github.com/mt1976/crt/config"
	dttm "github.com/mt1976/crt/datesTimes"
	errs "github.com/mt1976/crt/errors"
	hlpr "github.com/mt1976/crt/helpers"
	lang "github.com/mt1976/crt/language"
	strg "github.com/mt1976/crt/strings"
	symb "github.com/mt1976/crt/strings/symbols"

	"golang.org/x/term"
)

const StartColumn int = 1
const InputColumn int = StartColumn + 2

var config = conf.Configuration

// The "visibleContent" type represents a visibleContent with a map of rows and columns.
// @property row - The "row" property is a map that stores the values of each row in the visibleContent. The keys
// of the map are integers representing the row numbers, and the values are strings representing the
// content of each row.
// @property {int} cols - The "cols" property represents the number of columns in the visibleContent.
// @property {int} rows - The "rows" property represents the number of rows in the visibleContent.
type visibleContent struct {
	row    map[int]string
	cols   int
	rows   int
	prompt string
}

func (v *visibleContent) SetPrompt(prompt string) {
	v.prompt = prompt
}

func (v *visibleContent) GetPrompt() string {
	return v.prompt
}

// The ViewPort type represents a terminal screen with properties such as whether it is a terminal, its
// width and height, and whether it is the first row.
// @property {bool} isTerminal - A boolean value indicating whether the CRT (Cathode Ray Tube) is a
// terminal or not. If it is a terminal, it means that it is a device used for input and output of data
// to and from a computer. If it is not a terminal, it means that it is not used
// @property {int} width - The width property represents the number of characters that can be displayed
// horizontally on the terminal screen.
// @property {int} height - The `height` property represents the number of rows in the terminal or
// console window.
// @property {bool} firstRow - The `firstRow` property is a boolean value that indicates whether the
// current row is the first row of the terminal screen.
type ViewPort struct {
	isTerminal     bool             // true if running in terminal mode
	width          int              // the width of the terminal
	height         int              // the height of the terminal
	firstRow       bool             // true if the current row is the first row
	delay          int              // delay in milliseconds
	baudRate       int              // baud rate, which simulates the speed of a terminal
	currentRow     int              // the current row of the terminal
	currentCol     int              // the current column of the terminal
	visibleContent *visibleContent  // the current screen content
	Helpers        *hlpr.Helpers    // Helper functions
	Formatters     *hlpr.Formatters // Formatter functions
	Styles         *hlpr.Styles     // Colour functions
}

// The function `New` initializes a new `Crt` struct with information about the terminal size and
// whether it is a terminal or not.
func New() ViewPort {
	x := ViewPort{}
	x.isTerminal = true
	x.width = 0
	x.height = 0
	x.firstRow = true
	x.currentCol = 0
	x.currentRow = 0
	width, height, err := getTerminalSize()
	if err != nil {
		fmt.Println("ERROR: Unable to create new terminal")
		os.Exit(1)
	}

	x.width = width
	x.height = height

	x.SetTerminalSize(x.width, x.height)
	x.defaultDelay() // set delay to 0
	x.defaultBaud()  // set baud to 9600

	x.newPageContent(x.width, x.height)
	x.Helpers = hlpr.InitHelpers()
	x.Formatters = hlpr.InitFormatters()
	x.Styles = hlpr.InitStyles()
	return x
}

func getTerminalSize() (int, int, error) {

	if runtime.GOOS == "windows" {
		return 80, 25, nil
	} else {
		return term.GetSize(0)
	}

}

func NewWithSize(width, height int) ViewPort {
	xx := New()
	xx.SetTerminalSize(width, height)
	return xx
}

// The `row()` function is a method of the `Crt` struct. It is used to generate a formatted string that
// represents a row on the terminal.
func (t *ViewPort) row() string {
	displayChar := boxr.DividerLeft
	if t.firstRow {
		displayChar = boxr.StartLeft
		t.firstRow = false
	}
	return displayChar + strings.Repeat(boxr.Horizontal, t.width-3)
}

// The `Close()` function is a method of the `Crt` struct. It is used to print a closing line on the
// terminal. It calls the `row()` method of the `Crt` struct to get the formatted closing line string,
// and then it prints the string using `fmt.Println()`. This creates a visual separation between
// different sections or blocks of text on the terminal.
func (t *ViewPort) Close() {
	// t.PrintIt(t.row())
}

// The `SetDelayInMs` function is a method of the `Crt` struct. It takes an `int` parameter `delay` and
// sets the `delay` property of the `Crt` struct to the value of `delay`. This property represents the
// delay in milliseconds that should be applied before printing each character to the terminal.
func (t *ViewPort) SetDelayInMs(delayMs int) {
	t.delay = delayMs
}

// The `SetTerminalSize` function is a method of the `Crt` struct. It takes two parameters, `width` and
// `height`, which represent the desired width and height of the terminal screen.
func (t *ViewPort) SetTerminalSize(width, height int) {
	if !(width > 0 && height > 0) {
		t.Error(errs.ErrTerminalSize, strconv.Itoa(width), strconv.Itoa(height))
		os.Exit(1)
	}
	t.width = width
	t.height = height
}

// The `TerminalSize` function is a method of the `Crt` struct. It returns the width and height of the
// terminal screen. It retrieves the values of the `width` and `height` properties of the `Crt` struct
// and returns them as integers.
func (t *ViewPort) TerminalSize() (width, height int) {
	return t.width, t.height
}

// The `SetDelayInSec` function is a method of the `Crt` struct. It takes a parameter `delay` of type
// `interface{}`.
func (t *ViewPort) SetDelayInSec(delayMs float64) {
	t.delay = 0
	t.delay = int(delayMs * 1000)
}

// The `SetDelayInMin` function is a method of the `Crt` struct. It takes an `int` parameter `delay`
// and sets the `delay` property of the `Crt` struct to the value of `delay` multiplied by 60000. This
// function is used to set the delay in milliseconds that should be applied before printing each
// character to the terminal, but it takes the delay in minutes instead of milliseconds.
func (t *ViewPort) SetDelayInMin(delayMs int) {
	t.delay = delayMs * 60000
}

// The above code is defining a method called "ResetDelay" for a struct type "Crt". This method is a
// member of the "Crt" struct and has a receiver of type "*Crt". Inside the method, it calls another
// method called "defaultDelay" on the receiver "T".
func (t *ViewPort) ResetDelay() {
	t.defaultDelay()
}

// The above code is defining a method called "defaultDelay" for a struct type "Crt". This method sets
// the "delay" field of the struct to 0.
func (t *ViewPort) defaultDelay() {
	t.delay = 0
}

// The above code is defining a method called "DelayIt" for a struct type "Crt". This method takes no
// arguments and has no return value.
func (t *ViewPort) DelayIt() {
	if t.delay > 0 {
		time.Sleep(time.Duration(t.delay) * time.Millisecond)
	}
}

// Get Delay
// The above code is defining a method called "Delay" for a struct type "Crt". This method returns an
// integer value, which is the value of the "delay" field of the struct.
func (t *ViewPort) Delay() int {
	return t.delay
}

// Get Delay in seconds
// The above code is defining a method called "DelayInSec" for a struct type "Crt". This method returns
// the delay value of the "Crt" struct in seconds. The delay value is divided by 1000 to convert it
// from milliseconds to seconds and then returned as a float64.
func (t *ViewPort) DelayInSec() float64 {
	return float64(t.delay) / 1000
}

// The `Blank()` function is used to print a blank line on the terminal. It calls the `Format()` method
// of the `Crt` struct to format an empty string with the normal character (`chNormal`). Then, it
// prints the formatted string using `fmt.Println()`.
func (t *ViewPort) Blank() {
	t.Println(t.Format("", "") + symb.Newline.Symbol())
}

// The `Break()` function is used to print a line break on the terminal. It calls the `row()` method of
// the `Crt` struct to get the formatted line break string, and then it prints the string using
// `fmt.Println()`. This creates a visual separation between different sections or blocks of text on
// the terminal.
func (t *ViewPort) Break() {
	t.PrintIt(t.row() + symb.Newline.Symbol())
}

// The `Print` function is a method of the `Crt` struct. It takes a `msg` parameter of type string and
// prints it to the terminal. It uses the `Format` method of the `Crt` struct to format the message
// with the normal character (`chNormal`). Then, it prints the formatted string using `fmt.Println()`.
func (t *ViewPort) Print(msg string) {
	//log.Println(msg)x
	t.PrintIt(t.Format(msg, ""))
}

// Paragraph formats a list of strings as paragraphs, wrapping lines as needed to fit within the
// terminal width.
func (t *ViewPort) Paragraph(msg []string) {
	// make sure the lines are no longer than the screen width and wrap them if they are.
	out := []string{}
	for _, s := range msg {
		s = t.Formatters.TrimRepeatingCharacters(s, symb.Space.Symbol())
		if len(s) > t.Width() {
			out = append(out, s[:t.Width()])
			out = append(out, s[t.Width():])
		} else {
			out = append(out, s)
		}
	}

	for _, s := range out {
		t.Println(t.Format(s, ""))
	}
}

// The `Special` function is a method of the `Crt` struct. It takes a `msg` parameter of type string
// and prints it to the terminal using the `fmt.Println()` function. The message is formatted with the
// special character (`chSpecial`) using the `Format` method of the `Crt` struct. This function is used
// to print a special message or highlight certain text on the terminal.
func (t *ViewPort) Special(msg string) {
	t.Println(t.Format(msg, boxr.DividerLeft) + symb.Newline.Symbol())
}

// The `Input` function is a method of the `Crt` struct. It is used to display a prompt for the user for input on the
// terminal.
func (t *ViewPort) Input(msg string, options string) (output string) {
	gtrm.MoveCursor(StartColumn, 21)
	gtrm.Print(t.row())
	gtrm.MoveCursor(StartColumn, 22)
	mesg := msg
	//T.Format(msg, "")
	if options != "" {
		mesg = (t.Format(msg, "") + strg.PQuote(t.Styles.Bold(options)))
	}
	mesg = mesg + symb.PromptSymbol.Symbol()
	mesg = t.Format(mesg, "")
	//T.Print(mesg)
	gtrm.Print(mesg)
	gtrm.Flush()
	var out string
	fmt.Scan(&out)
	output = out
	return output
}

// The `InputError` function is a method of the `Crt` struct. It takes a `msg` parameter of type string and prints an error message to the terminal. It uses the `Format` method of the `Crt` struct to format the message with the bold red color and the special character (`chSpecial`). Then, it prints the formatted string using `fmt.Println()`.
func (t *ViewPort) InputError(err error, msg ...string) {
	gtrm.MoveCursor(StartColumn, 23)
	pp := t.SError(err, msg...)
	gtrm.Print(pp)
	gtrm.Flush()
	beep.Beep(config.DefaultBeepFrequency, config.DefaultBeepDuration)
	oldDelay := t.Delay()
	t.SetDelayInSec(config.DefaultErrorDelay)
	t.DelayIt()
	t.SetDelayInMs(oldDelay)
}

func (t *ViewPort) InfoMessage(msg string) {
	gtrm.MoveCursor(StartColumn, 23)
	//Print a line that clears the entire line
	blanks := strings.Repeat(symb.Space.Symbol(), t.width)
	gtrm.Print(t.Format(blanks, ""))
	gtrm.MoveCursor(StartColumn, 23)
	gtrm.Print(
		t.Format(t.Styles.Cyan(msg), ""))
	//T.Print(msg + t.SymNewline)
	gtrm.Flush()
	//beeep.Beep(defaultBeepFrequency, defaultBeepDuration)
	//oldDelay := T.Delay()
	//T.SetDelayInSec(errorDelay)
	//T.DelayIt()
	//T.SetDelayInMs(oldDelay)

}

// The `InputPagingInfo` function is a method of the `Crt` struct. It is used to print information about the current page and total number of pages to the terminal.
//
// Parameters:
// page: The current page number.
// ofPages: The total number of pages.
//
// Returns:
// None.
func (t *ViewPort) InputPagingInfo(page, ofPages int) {
	msg := fmt.Sprintf(lang.Paging.Text(), page, ofPages)
	lmsg := len(msg)
	gtrm.MoveCursor(t.width-lmsg-1, 22)
	//gT.MoveCursor(col, 23)
	gtrm.Print(
		t.Format(t.Styles.Yellow(msg), ""))
	//T.Print(msg + t.SymNewline)
	gtrm.Flush()
}

// lineBreakEnd returns a string that represents a line break with the end character.
func (t *ViewPort) lineBreakEnd() string {
	return t.lineBreakJunction(boxr.EndLeft)
}

// lineBreakJunction returns a string that represents a line break with the end character.
func (t *ViewPort) lineBreakJunction(displayChar string) string {
	return fmt.Sprintf(lang.LineConstructor.Text(), displayChar, strings.Repeat(boxr.Horizontal, t.width+1), boxr.Horizontal)
}

// The `Format` function is a method of the `Crt` struct. It takes two parameters: `in` of type string
// and `t` of type string.
func (t *ViewPort) Format(msg string, text string) string {
	char := boxr.Upright
	if text != "" {
		char = text
	}
	t.DelayIt()
	return fmt.Sprintf("%s %s", char, msg)
}

// clear the terminal screen
func (t *ViewPort) Clear() {

	t.firstRow = true
	t.currentRow = 0
	gtrm.Clear()
	gtrm.MoveCursor(StartColumn, 1)
	gtrm.Flush()
}

// The `Shout` function is a method of the `Crt` struct. It takes a `msg` parameter of type string and
// prints a formatted message to the terminal.
func (t *ViewPort) Shout(msg string) {
	t.PrintIt(t.row() + symb.Newline.Symbol())
	t.PrintIt(t.Format(t.Styles.Bold(msg), "") + symb.Newline.Symbol())
	t.PrintIt(t.lineBreakEnd() + symb.Newline.Symbol())
}

// The `Error` function is a method of the `Crt` struct. It takes two parameters: `msg` of type string
// and `err` of type error.
func (t *ViewPort) Error(err error, msg ...string) {
	t.Println(t.row())
	t.Println(t.SError(err, msg...))
	t.Println(t.row())
}

func (t *ViewPort) SError(err error, msg ...string) string {
	errText := err.Error()
	msgr := t.Styles.Red(lang.Error.Text())
	return t.fmtMessage(errText, msgr, "", msg...)
}

func (t *ViewPort) fmtMessage(errText, promptTxt, colour string, msg ...string) string {

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
	errText = (colour + promptTxt + t.Styles.Reset) + qq
	errText = t.Format(errText, "")
	return errText
}

// The `bold` method of the `Crt` struct is used to format a string with bold text. It takes a `msg`
// parameter of type string and returns a formatted string with the `msg` surrounded by the bold escape
// characters (`bold` and `reset`). The `fmt.Sprintf` function is used to concatenate the escape
// characters and the `msg` string.
// func (T *Crt) bold(msg string) string {
// 	return fmt.Sprintf(lang.TextLineConstructor, p.viewPort.Styles.Bold, msg, p.viewPort.Styles.Reset)
// }

// The `Underline` method of the `Crt` struct is used to format a string with an underline. It takes a
// `msg` parameter of type string and returns a formatted string with the `msg` surrounded by the
// underline escape characters (`underline` and `reset`). The `fmt.Sprintf` function is used to
// concatenate the escape characters and the `msg` string. This method is used to create an underlined
// text effect when printing to the terminal.
func (t *ViewPort) Underline(msg string) string {
	return t.Styles.Underline(msg)
}

// Spool prints the contents of a byte slice to the terminal.
//
// The byte slice is split into lines by the t.SymNewline character (\n). For each line, the function
// determines whether the line is empty. If the line is not empty, it is prepended with "  " (two
// spaces) and printed to the terminal.
//
// If the byte slice is empty, the function returns without printing anything.
//
// The function also prints a blank line after all lines have been printed.
func (t *ViewPort) Spool(msg []byte) {
	//output = []byte(strings.ReplaceAll(string(output), "\n", "\n"+T.Bold("  ")))
	//create an slice of strings, split by t.SymNewline
	lines := strings.Split(string(msg), symb.Newline.Symbol())
	// loop through the slice
	if len(msg) == 0 {
		return
	}
	t.Blank()
	for _, line := range lines {
		if line != "" {
			t.Print("  " + string(line))
		}
	}
	t.Blank()
}

// The `Banner` function is a method of the `Crt` struct. It is responsible for printing a banner
// message to the console.
func (t *ViewPort) Banner(msg string) {
	fmt.Println(t.row())
	//gtrm.Flush()
	for _, line := range lang.ApplicationHeader.String() {
		fmt.Println(t.Format(line, ""))
		//	gtrm.Flush()
	}
	fmt.Println(t.row())
	//gtrm.Flush()
	display := fmt.Sprintf(lang.ApplicationVersion.Text(), msg)
	fmt.Println(t.Format(display+symb.Newline.Symbol(), ""))
	//t.Break()
	//gtrm.Flush()
}

// The `Header` function is a method of the `Crt` struct. It is responsible for printing a banner
// message to the console.
func (t *ViewPort) Header(msg string) {
	// Print Header Line
	gtrm.MoveCursor(1, 1)
	gtrm.Println(t.row()) // + lang.SymNewline.String())
	gtrm.MoveCursor(StartColumn, 2)
	var line map[int]string = make(map[int]string)
	midway := (t.width - len(msg)) / 2
	for i := 0; i < len(lang.ApplicationName.Text()); i++ {
		line[i] = lang.ApplicationName.Text()[i : i+1]
	}
	for i := 0; i < len(msg); i++ {
		line[midway+i] = msg[i : i+1]
	}

	// Add DateTimeStamp to end of string
	for i := 0; i < len(dttm.DateTimeString()); i++ {
		line[t.width-len(dttm.DateTimeString())+i] = dttm.DateTimeString()[i : i+1]
	}

	//map to string
	var headerRowString string
	for i := 0; i < t.width; i++ {
		if line[i] == "" {
			line[i] = symb.Space.Symbol()
		}
		headerRowString = headerRowString + line[i]
	}

	gtrm.Print(t.Styles.Bold(headerRowString) + symb.Newline.Symbol())
	gtrm.Flush()
	t.Break()
}

// SetBaud sets the baud rate for the CRT.
//
// If the specified baud rate is not supported, an error is returned and the CRT's baud rate is reset to the default value.
func (t *ViewPort) SetBaud(baudRate int) {
	if sort.SearchInts(config.ValidBaudRates, baudRate) == -1 {
		t.Error(errs.ErrBaudRateError, strconv.Itoa(baudRate))
		t.defaultBaud()
		return
	}
	t.baudRate = baudRate
}

// Baud returns the current baud rate of the CRT.
func (t *ViewPort) Baud() int {
	return t.baudRate
}

// SetBaud sets the baud rate for the CRT.
//
// If the specified baud rate is not supported, an error is returned and the CRT's baud rate is reset to the default value.
func (t *ViewPort) defaultBaud() {
	t.baudRate = config.DefaultBaud
}

// PrintIt prints a message to the terminal.
//
// If the CRT's baud rate is set to col, the function prints the message without applying any delays or formatting.
// If the baud rate is non-zero, the function prints the message character by character, with a delay of 1000000 microseconds (1 millisecond) between each character.
// The function also prints the current row number at the end of the message.
//
// The function returns without printing a new line. To print a new line, use the Println method.
func (t *ViewPort) PrintIt(msg string) {
	t.currentRow++
	rowString := msg
	gtrm.MoveCursor(StartColumn, t.currentRow)
	//truncate rowString to length-1 and add a | character to the end
	//log.Printf("len(rowString): %v\n", len(rowString))
	//log.Printf("t.width: %v\n", t.width)
	//log.Printf("msg: %v\n", msg)
	//log.Printf("t.currentRow: %v\n", t.currentRow)
	if len(rowString) < t.width {
		rowString = rowString + strings.Repeat(".", t.width-(len(rowString)+1))
	} else {
		rowString = rowString[0:t.width]
	}
	//t.Print(rowString + msg
	rowString = rowString + boxr.Upright
	//log.Printf("rowString: [%v]\n", rowString)
	//log.Printf("len(rowString): %v\n", len(rowString))
	if t.NoBaudRate() {
		gtrm.Println(rowString)
		//fmt.Println(rowString)
		return
	} else {
		// print one character at a time
		for col, c := range msg {
			gtrm.MoveCursor(col, t.currentRow)
			gtrm.Print(string(c))
			//fmt.Print(string(c))
			time.Sleep(time.Duration(1000000/t.baudRate) * time.Microsecond)
		}
		//fmt.Print(lang.Space + rowString)
		//fmt.Println("")
	}
}

// Get the height of the terminal
func (t *ViewPort) Height() int {
	return t.height
}

// Println prints a message to the terminal and adds a new line.
//
// If the CRT's baud rate is set to col, the function prints the message without applying any delays or formatting.
// If the baud rate is non-zero, the function prints the message character by character, with a delay of 1000000 microseconds (1 millisecond) between each character.
// The function also prints the current row number at the end of the message.
//
// The function returns without printing a new line. To print a new line, use the Println method.
func (t *ViewPort) Println(msg string) {
	t.Print(msg + symb.Newline.Symbol())
}

// Get the width of the terminal
func (t *ViewPort) Width() int {
	return t.width
}

// Get the current row of the terminal
func (t *ViewPort) CurrentRow() int {
	return t.currentRow
}

// NoBaudRate returns true if the CRT's baud rate is set to col, false otherwise.
func (t *ViewPort) NoBaudRate() bool {
	return t.baudRate == 0
}

// ClearCurrentLine clears the current line in the terminal
func (t *ViewPort) ClearCurrentLine() {
	fmt.Print(t.Styles.ClearLine)
}

// newPageContent initializes a new page with the specified number of columns and rows.
func (t *ViewPort) newPageContent(cols, rows int) {
	v := visibleContent{}
	v.cols = cols
	v.rows = rows
	v.row = make(map[int]string)
	t.visibleContent = &v
}

func (t *ViewPort) Wait() {
	time.Sleep(time.Duration(t.delay) * time.Millisecond)
}
