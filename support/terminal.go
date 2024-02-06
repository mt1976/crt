package support

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const chNormal = "┃"
const chSpecial = "┣"
const chStart = "┏"

// const chClose = "┗"
const bold = "\033[1m"
const reset = "\033[0m"
const underline = "\033[4m"
const red = "\033[31m"

var header []string
var smHeader []string

func init() {
	header = []string{
		"███████ ████████  █████  ██████  ████████ ███████ ██████  ███    ███ ",
		"██         ██    ██   ██ ██   ██    ██    ██      ██   ██ ████  ████ ",
		"███████    ██    ███████ ██████     ██    █████   ██████  ██ ████ ██ ",
		"     ██    ██    ██   ██ ██   ██    ██    ██      ██   ██ ██  ██  ██ ",
		"███████    ██    ██   ██ ██   ██    ██    ███████ ██   ██ ██      ██ ",
	}

	smHeader = []string{
		"StartTerm",
	}
}

var baudRates = []int{0, 300, 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200}
var defaultBaud = 9600
var baudRateError = "Invalid Baud Rate"

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
	T.Println(T.row())
}

func (T *Crt) SetDelayInMs(delay int) {
	T.delay = delay
}

func (T *Crt) SetDelayInSec(delay interface{}) {
	if delay.(float64) > 0 {
		T.delay = int(delay.(float64) * 1000)
	} else {
		T.delay = delay.(int) * 1000
	}
}

func (T *Crt) SetDelayInMin(delay int) {
	T.delay = delay * 60000
}

func (T *Crt) ResetDelay() {
	T.defaultDelay()
}

func (T *Crt) defaultDelay() {
	T.delay = 0
}

func (T *Crt) Delay() {
	if T.delay > 0 {
		time.Sleep(time.Duration(T.delay) * time.Millisecond)
	}
}

// Get Delay
func (T *Crt) GetDelay() int {
	return T.delay
}

// Get Delay in seconds
func (T *Crt) GetDelayInSec() float64 {
	return float64(T.delay) / 1000
}

// The `Blank()` function is used to print a blank line on the terminal. It calls the `Format()` method
// of the `Crt` struct to format an empty string with the normal character (`chNormal`). Then, it
// prints the formatted string using `fmt.Println()`.
func (T *Crt) Blank() {
	T.Println(T.Format("", ""))
}

// The `Break()` function is used to print a line break on the terminal. It calls the `row()` method of
// the `Crt` struct to get the formatted line break string, and then it prints the string using
// `fmt.Println()`. This creates a visual separation between different sections or blocks of text on
// the terminal.
func (T *Crt) Break() {
	T.Println(T.row())
}

// The `Print` function is a method of the `Crt` struct. It takes a `msg` parameter of type string and
// prints it to the terminal. It uses the `Format` method of the `Crt` struct to format the message
// with the normal character (`chNormal`). Then, it prints the formatted string using `fmt.Println()`.
func (T *Crt) Print(msg string) {
	T.Println(T.Format(msg, ""))
}

// The `Special` function is a method of the `Crt` struct. It takes a `msg` parameter of type string
// and prints it to the terminal using the `fmt.Println()` function. The message is formatted with the
// special character (`chSpecial`) using the `Format` method of the `Crt` struct. This function is used
// to print a special message or highlight certain text on the terminal.
func (T *Crt) Special(msg string) {
	T.Println(T.Format(msg, chSpecial))
}

// The `Input` function is a method of the `Crt` struct. It is used to display a prompt for the user for input on the
// terminal.
func (T *Crt) Input(msg string, ops string) {
	mesg := T.Format(msg, "")
	if ops != "" {
		mesg = (T.Format(msg, "") + " (" + T.Bold(ops) + ")")
	}
	mesg = mesg + "? :"
	fmt.Print(mesg)
}

func (T *Crt) lineBreakEnd() string {
	return T.lineBreakJunction("┗")
}

func (T *Crt) lineBreakJunction(displayChar string) string {
	return fmt.Sprintf("%s%s", displayChar, strings.Repeat("━", T.width-2))
}

// The `Format` function is a method of the `Crt` struct. It takes two parameters: `in` of type string
// and `t` of type string.
func (T *Crt) Format(in string, t string) string {
	char := chNormal
	if t != "" {
		char = t
	}
	T.Delay()
	return fmt.Sprintf("%s %s", char, in)
}

// clear the terminal screen
func (T *Crt) Clear() {
	T.Println("\033[H\033[2J")
	T.firstRow = true
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

// The function `NewCrt` initializes a new `Crt` struct with information about the terminal size and
// whether it is a terminal or not.
func NewCrt() Crt {
	x := Crt{}
	x.isTerminal = true
	x.width = 0
	x.height = 0
	x.firstRow = true

	x.width = 80
	x.height = 25
	x.defaultDelay() // set delay to 0
	x.defaultBaud()  // set baud to 9600

	return x
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
	T.Print(display)
	T.Break()
}

func (T *Crt) SmBanner(msg string) {
	T.Println(T.row())
	for _, line := range smHeader {
		T.Print(line + " - " + msg)
	}
	T.Break()
}

func (T *Crt) SetBaud(baud int) {
	if sort.SearchInts(baudRates, baud) == -1 {
		T.Error(baudRateError, nil)
		T.defaultBaud()
		return
	}
	T.baud = baud
}

func (T *Crt) GetBaud() int {
	return T.baud
}

func (T *Crt) defaultBaud() {
	T.baud = defaultBaud
}

func (T *Crt) Println(msg string) {
	fmt.Println(msg)
}
