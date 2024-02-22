package support

import "strings"

const BoxCharacterNormal = "┃"
const BoxCharacterBreak = "┣"
const BoxCharacterStart = "┏"

// const chEnd = "┛"
// const chJunction = "┣"
// const chEndFirst = "┓"
const BoxCharacterBar = "━"
const TableCharacterUnderline = "-"

// const chClose = "┗"
const bold = "\033[1m"
const reset = "\033[0m"
const underline = "\033[4m"
const red = "\033[31m"
const clearline = "\033[2K"

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

const newline string = "\n"

var versionText = "StarTerm - Utilities 1.0 %s"

// The Upcase function in Go converts a string to uppercase.
// Upcase converts a string to uppercase.
func Upcase(s string) string {
	return strings.ToUpper(s)
}

// The function `TrimRepeatingCharacters` takes a string `s` and a character `c` as input, and returns
// a new string with all consecutive occurrences of `c` trimmed down to a single occurrence.
func TrimRepeatingCharacters(s string, c string) string {

	result := ""
	lenS := len(s)

	for i := 0; i < lenS; i++ {
		if i == 0 {
			result = string(s[i])
		} else {
			if string(s[i]) != c || string(s[i-1]) != c {
				result = result + string(s[i])
			}
		}
	}
	return result
}
