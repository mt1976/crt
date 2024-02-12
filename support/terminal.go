package support

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	gT "github.com/buger/goterm"
	"github.com/gen2brain/beeep"
)

var baudRates = []int{0, 300, 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200}
var defaultBaud = 0

var defaultBeepFrequency float64 = beeep.DefaultFreq
var defaultBeepDuration int = beeep.DefaultDuration

// The Crt type represents a terminal screen with properties such as whether it is a terminal, its
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
type Crt struct {
	isTerminal bool
	width      int
	height     int
	firstRow   bool
	delay      int // delay in milliseconds
	baud       int
	currentRow int
	currentCol int
	scr        page
}

// The "page" type represents a page with a map of rows and columns.
// @property row - The "row" property is a map that stores the values of each row in the page. The keys
// of the map are integers representing the row numbers, and the values are strings representing the
// content of each row.
// @property {int} cols - The "cols" property represents the number of columns in the page.
// @property {int} rows - The "rows" property represents the number of rows in the page.
type page struct {
	row  map[int]string
	cols int
	rows int
}

// The `row()` function is a method of the `Crt` struct. It is used to generate a formatted string that
// represents a row on the terminal.
func (T *Crt) row() string {
	displayChar := chSpecial
	if T.firstRow {
		displayChar = chStart
		T.firstRow = false
	}
	return T.lineBreakJunction(displayChar)
}

// The `Close()` function is a method of the `Crt` struct. It is used to print a closing line on the
// terminal. It calls the `row()` method of the `Crt` struct to get the formatted closing line string,
// and then it prints the string using `fmt.Println()`. This creates a visual separation between
// different sections or blocks of text on the terminal.
func (T *Crt) Close() {
	T.PrintIt(T.row())
}

// The `SetDelayInMs` function is a method of the `Crt` struct. It takes an `int` parameter `delay` and
// sets the `delay` property of the `Crt` struct to the value of `delay`. This property represents the
// delay in milliseconds that should be applied before printing each character to the terminal.
func (T *Crt) SetDelayInMs(delay int) {
	T.delay = delay
}

// The `SetTerminalSize` function is a method of the `Crt` struct. It takes two parameters, `width` and
// `height`, which represent the desired width and height of the terminal screen.
func (T *Crt) SetTerminalSize(width, height int) {
	if !(width > 0 && height > 0) {
		T.Error(TerminalSizeError, nil)
		os.Exit(1)
	}
	T.width = width
	T.height = height
}

// The `TerminalSize` function is a method of the `Crt` struct. It returns the width and height of the
// terminal screen. It retrieves the values of the `width` and `height` properties of the `Crt` struct
// and returns them as integers.
func (T *Crt) TerminalSize() (width int, height int) {
	return T.width, T.height
}

// The `SetDelayInSec` function is a method of the `Crt` struct. It takes a parameter `delay` of type
// `interface{}`.
func (T *Crt) SetDelayInSec(delay interface{}) {
	T.delay = 0

	if delay.(float64) > 0 {
		T.delay = int(delay.(float64) * 1000)
	} else {
		T.delay = delay.(int) * 1000
	}
}

// The `SetDelayInMin` function is a method of the `Crt` struct. It takes an `int` parameter `delay`
// and sets the `delay` property of the `Crt` struct to the value of `delay` multiplied by 60000. This
// function is used to set the delay in milliseconds that should be applied before printing each
// character to the terminal, but it takes the delay in minutes instead of milliseconds.
func (T *Crt) SetDelayInMin(delay int) {
	T.delay = delay * 60000
}

// The above code is defining a method called "ResetDelay" for a struct type "Crt". This method is a
// member of the "Crt" struct and has a receiver of type "*Crt". Inside the method, it calls another
// method called "defaultDelay" on the receiver "T".
func (T *Crt) ResetDelay() {
	T.defaultDelay()
}

// The above code is defining a method called "defaultDelay" for a struct type "Crt". This method sets
// the "delay" field of the struct to 0.
func (T *Crt) defaultDelay() {
	T.delay = 0
}

// The above code is defining a method called "DelayIt" for a struct type "Crt". This method takes no
// arguments and has no return value.
func (T *Crt) DelayIt() {
	if T.delay > 0 {
		time.Sleep(time.Duration(T.delay) * time.Millisecond)
	}
}

// Get Delay
// The above code is defining a method called "Delay" for a struct type "Crt". This method returns an
// integer value, which is the value of the "delay" field of the struct.
func (T *Crt) Delay() int {
	return T.delay
}

// Get Delay in seconds
// The above code is defining a method called "DelayInSec" for a struct type "Crt". This method returns
// the delay value of the "Crt" struct in seconds. The delay value is divided by 1000 to convert it
// from milliseconds to seconds and then returned as a float64.
func (T *Crt) DelayInSec() float64 {
	return float64(T.delay) / 1000
}

// The `Blank()` function is used to print a blank line on the terminal. It calls the `Format()` method
// of the `Crt` struct to format an empty string with the normal character (`chNormal`). Then, it
// prints the formatted string using `fmt.Println()`.
func (T *Crt) Blank() {
	T.Println(T.Format("", "") + newline)
}

// The `Break()` function is used to print a line break on the terminal. It calls the `row()` method of
// the `Crt` struct to get the formatted line break string, and then it prints the string using
// `fmt.Println()`. This creates a visual separation between different sections or blocks of text on
// the terminal.
func (T *Crt) Break() {
	T.PrintIt(T.row() + newline)
}

// The `Print` function is a method of the `Crt` struct. It takes a `msg` parameter of type string and
// prints it to the terminal. It uses the `Format` method of the `Crt` struct to format the message
// with the normal character (`chNormal`). Then, it prints the formatted string using `fmt.Println()`.
func (T *Crt) Print(msg string) {
	T.PrintIt(T.Format(msg, ""))
}

// The `Special` function is a method of the `Crt` struct. It takes a `msg` parameter of type string
// and prints it to the terminal using the `fmt.Println()` function. The message is formatted with the
// special character (`chSpecial`) using the `Format` method of the `Crt` struct. This function is used
// to print a special message or highlight certain text on the terminal.
func (T *Crt) Special(msg string) {
	T.Println(T.Format(msg, chSpecial) + newline)
}

// The `Input` function is a method of the `Crt` struct. It is used to display a prompt for the user for input on the
// terminal.
func (T *Crt) Input(msg string, ops string) (output string) {
	gT.MoveCursor(2, 21)
	gT.Print(T.row())
	gT.MoveCursor(2, 22)
	mesg := msg
	//T.Format(msg, "")
	if ops != "" {
		mesg = (T.Format(msg, "") + " (" + T.Bold(ops) + ")")
	}
	mesg = mesg + "? "
	mesg = T.Format(mesg, "")
	//T.Print(mesg)
	gT.Print(mesg)
	gT.Flush()
	var out string
	fmt.Scan(&out)
	output = out
	return output
}

// The `InputError` function is a method of the `Crt` struct. It takes a `msg` parameter of type string and prints an error message to the terminal. It uses the `Format` method of the `Crt` struct to format the message with the bold red color and the special character (`chSpecial`). Then, it prints the formatted string using `fmt.Println()`.
func (T *Crt) InputError(msg string) {
	gT.MoveCursor(2, 23)
	gT.Print(
		T.Format(gT.Color(gT.Bold("ERROR : "), gT.RED)+msg, ""))
	//T.Print(msg + newline)
	gT.Flush()
	beeep.Beep(defaultBeepFrequency, defaultBeepDuration)
	oldDelay := T.Delay()
	T.SetDelayInSec(errorDelay)
	T.DelayIt()
	T.SetDelayInMs(oldDelay)

}

// The `InputPageInfo` function is a method of the `Crt` struct. It is used to print information about the current page and total number of pages to the terminal.
//
// Parameters:
// page: The current page number.
// ofPages: The total number of pages.
//
// Returns:
// None.
func (T *Crt) InputPageInfo(page, ofPages int) {
	msg := fmt.Sprintf("Page %v of %v", page, ofPages)
	lmsg := len(msg)
	gT.MoveCursor(T.width-lmsg-1, 22)
	//gT.MoveCursor(2, 23)
	gT.Print(
		T.Format(gT.Color(msg, gT.YELLOW), ""))
	//T.Print(msg + newline)
	gT.Flush()
}

// lineBreakEnd returns a string that represents a line break with the end character.
func (T *Crt) lineBreakEnd() string {
	return T.lineBreakJunction("┗")
}

// lineBreakJunction returns a string that represents a line break with the end character.
func (T *Crt) lineBreakJunction(displayChar string) string {
	//endChar := chEnd
	//if T.currentRow == 0 {
	//	endChar = chEndFirst
	//}
	return fmt.Sprintf("%s%s%s", displayChar, strings.Repeat(chBar, T.width+1), chBar)
}

// The `Format` function is a method of the `Crt` struct. It takes two parameters: `in` of type string
// and `t` of type string.
func (T *Crt) Format(in string, t string) string {
	char := chNormal
	if t != "" {
		char = t
	}
	T.DelayIt()
	return fmt.Sprintf("%s %s", char, in)
}

// clear the terminal screen
func (T *Crt) Clear() {
	//goTerminal.Clean()
	//goTerminal.CursorLineColumn(0, 0)
	//fmt.Println("Bold")
	//T.Println("\033[H\033[2J")
	T.firstRow = true
	T.currentRow = 0
	gT.Clear()
	gT.MoveCursor(2, 1)
	gT.Flush()
}

// The `Shout` function is a method of the `Crt` struct. It takes a `msg` parameter of type string and
// prints a formatted message to the terminal.
func (T *Crt) Shout(msg string) {
	T.Println(T.row())
	T.Println(T.Format(bold+"MESSAGE: "+reset+msg, ""))
	T.Println(T.lineBreakEnd())
}

// The `Error` function is a method of the `Crt` struct. It takes two parameters: `msg` of type string
// and `err` of type error.
func (T *Crt) Error(msg string, err error) {
	T.Println(T.row())
	T.Println(T.Format(T.Bold(red+"ERROR: ")+msg+fmt.Sprintf(" [%v]", err), ""))
	T.Println(T.row())
}

// The function `New` initializes a new `Crt` struct with information about the terminal size and
// whether it is a terminal or not.
func New() Crt {
	x := Crt{}
	x.isTerminal = true
	x.width = 0
	x.height = 0
	x.firstRow = true
	x.currentCol = 0
	x.currentRow = 0

	x.width = 80
	x.height = 25
	x.defaultDelay() // set delay to 0
	x.defaultBaud()  // set baud to 9600

	x.scr = NewPage(x.width, x.height)

	return x
}

func NewWithSize(width, height int) Crt {
	xx := New()
	xx.SetTerminalSize(width, height)
	return xx
}

// NewPage initializes a new page with the specified number of columns and rows.
func NewPage(cols, rows int) page {
	p := page{}
	p.cols = cols
	p.rows = rows
	p.row = make(map[int]string)
	return p
}

// The `Bold` method of the `Crt` struct is used to format a string with bold text. It takes a `msg`
// parameter of type string and returns a formatted string with the `msg` surrounded by the bold escape
// characters (`bold` and `reset`). The `fmt.Sprintf` function is used to concatenate the escape
// characters and the `msg` string.
func (T *Crt) Bold(msg string) string {
	return fmt.Sprintf("%s%s%s", bold, msg, reset)
}

// The `Underline` method of the `Crt` struct is used to format a string with an underline. It takes a
// `msg` parameter of type string and returns a formatted string with the `msg` surrounded by the
// underline escape characters (`underline` and `reset`). The `fmt.Sprintf` function is used to
// concatenate the escape characters and the `msg` string. This method is used to create an underlined
// text effect when printing to the terminal.
func (T *Crt) Underline(msg string) string {
	return fmt.Sprintf("%s%s%s", underline, msg, reset)
}

// Spool prints the contents of a byte slice to the terminal.
//
// The byte slice is split into lines by the newline character (\n). For each line, the function
// determines whether the line is empty. If the line is not empty, it is prepended with "  " (two
// spaces) and printed to the terminal.
//
// If the byte slice is empty, the function returns without printing anything.
//
// The function also prints a blank line after all lines have been printed.
func (T *Crt) Spool(msg []byte) {
	//output = []byte(strings.ReplaceAll(string(output), "\n", "\n"+T.Bold("  ")))
	//create an slice of strings, split by newline
	lines := strings.Split(string(msg), "\n")
	// loop through the slice
	if len(msg) == 0 {
		return
	}
	T.Blank()
	for _, line := range lines {
		if line != "" {
			T.Print("  " + string(line))
		}
	}
	T.Blank()
}

// The `Banner` function is a method of the `Crt` struct. It is responsible for printing a banner
// message to the console.
func (T *Crt) Banner(msg string) {
	T.Println(T.row())
	for _, line := range header {
		T.Print(line)
	}
	T.Blank()
	display := fmt.Sprintf("StarTerm - Utilities 1.0 %s", msg)
	T.Println(display)
	T.Break()
}

// The `Header` function is a method of the `Crt` struct. It is responsible for printing a banner
// message to the console.
func (T *Crt) Header(msg string) {
	T.PrintIt(T.row() + newline)
	var line map[int]string = make(map[int]string)
	midway := (T.width - len(msg)) / 2
	for i := 0; i < len(smHeader); i++ {
		line[i] = smHeader[i : i+1]
	}
	for i := 0; i < len(msg); i++ {
		line[midway+i] = msg[i : i+1]
	}

	// Add DateTimeStamp to end of string
	for i := 0; i < len(DateTimeString()); i++ {
		line[T.width-len(DateTimeString())+i] = DateTimeString()[i : i+1]
	}

	//map to string
	var headerRowString string
	for i := 0; i < T.width; i++ {
		if line[i] == "" {
			line[i] = " "
		}
		headerRowString = headerRowString + line[i]
	}

	T.Print(T.Bold(headerRowString) + newline)

	// TODO Print Date/Time
	T.Break()
}

// SetBaud sets the baud rate for the CRT.
//
// If the specified baud rate is not supported, an error is returned and the CRT's baud rate is reset to the default value.
func (T *Crt) SetBaud(baud int) {
	if sort.SearchInts(baudRates, baud) == -1 {
		T.Error(BaudRateError, nil)
		T.defaultBaud()
		return
	}
	T.baud = baud
}

// Baud returns the current baud rate of the CRT.
func (T *Crt) Baud() int {
	return T.baud
}

// SetBaud sets the baud rate for the CRT.
//
// If the specified baud rate is not supported, an error is returned and the CRT's baud rate is reset to the default value.
func (T *Crt) defaultBaud() {
	T.baud = defaultBaud
}

// PrintIt prints a message to the terminal.
//
// If the CRT's baud rate is set to 0, the function prints the message without applying any delays or formatting.
// If the baud rate is non-zero, the function prints the message character by character, with a delay of 1000000 microseconds (1 millisecond) between each character.
// The function also prints the current row number at the end of the message.
//
// The function returns without printing a new line. To print a new line, use the Println method.
func (T *Crt) PrintIt(msg string) {
	T.currentRow++
	rowString := fmt.Sprintf("%v", T.currentRow-1)
	if T.NoBaudRate() {
		fmt.Print(msg + " ")
		return
	} else {
		// print one character at a time
		for _, c := range msg {
			fmt.Print(string(c))
			time.Sleep(time.Duration(1000000/T.baud) * time.Microsecond)
		}
		fmt.Print(" " + rowString)
		//fmt.Println("")
	}
}

// Get the height of the terminal
func (T *Crt) Height() int {
	return T.height
}

// Println prints a message to the terminal and adds a new line.
//
// If the CRT's baud rate is set to 0, the function prints the message without applying any delays or formatting.
// If the baud rate is non-zero, the function prints the message character by character, with a delay of 1000000 microseconds (1 millisecond) between each character.
// The function also prints the current row number at the end of the message.
//
// The function returns without printing a new line. To print a new line, use the Println method.
func (T *Crt) Println(msg string) {
	T.Print(msg + "\n")
}

// Get the width of the terminal
func (T *Crt) Width() int {
	return T.width
}

// Get the current row of the terminal
func (T *Crt) CurrentRow() int {
	return T.currentRow
}

// NoBaudRate returns true if the CRT's baud rate is set to 0, false otherwise.
func (T *Crt) NoBaudRate() bool {
	return T.baud == 0
}

// ClearCurrentLine clears the current line in the terminal
func (T *Crt) ClearCurrentLine() {
	fmt.Print(clearline)
}
