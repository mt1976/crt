package support

import "strings"

const chNormal = "┃"
const chSpecial = "┣"
const chStart = "┏"

// const chEnd = "┛"
// const chJunction = "┣"
// const chEndFirst = "┓"
const chBar = "━"

// const chClose = "┗"
const bold = "\033[1m"
const reset = "\033[0m"
const underline = "\033[4m"
const red = "\033[31m"
const clearline = "\033[2K"

const errorDelay = 7.0

var header []string
var smHeader string

// The init function initializes the header and smHeader variables with ASCII art and a string value,
// respectively.
func init() {
	header = []string{
		"███████ ████████  █████  ██████  ████████ ███████ ██████  ███    ███ ",
		"██         ██    ██   ██ ██   ██    ██    ██      ██   ██ ████  ████ ",
		"███████    ██    ███████ ██████     ██    █████   ██████  ██ ████ ██ ",
		"     ██    ██    ██   ██ ██   ██    ██    ██      ██   ██ ██  ██  ██ ",
		"███████    ██    ██   ██ ██   ██    ██    ███████ ██   ██ ██      ██ ",
	}

	smHeader = "StarTerm"

}

var newline = "\n"

// The Upcase function in Go converts a string to uppercase.
// Upcase converts a string to uppercase.
func Upcase(s string) string {
	return strings.ToUpper(s)
}
