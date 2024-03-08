package support

import "strings"

const (
	BoxCharacterNormal      string = "┃"
	BoxCharacterBreak       string = "┣"
	BoxCharacterStart       string = "┏"
	BoxCharacterBar         string = "━"
	BoxCharacterBarBreak    string = "┗"
	TableCharacterUnderline string = "-"
	bold                    string = "\033[1m"
	reset                   string = "\033[0m"
	underline               string = "\033[4m"
	red                     string = "\033[31m"
	clearline               string = "\033[2K"
	lineSymbol              string = "%s%s%s"

	// const chEnd = "┛"
	// const chJunction = "┣"
	// const chEndFirst = "┓"
	// const chClose = "┗"
)

// var smHeader string
var header []string = []string{
	"███████ ████████  █████  ██████  ████████ ███████ ██████  ███    ███ ",
	"██         ██    ██   ██ ██   ██    ██    ██      ██   ██ ████  ████ ",
	"███████    ██    ███████ ██████     ██    █████   ██████  ██ ████ ██ ",
	"     ██    ██    ██   ██ ██   ██    ██    ██      ██   ██ ██  ██  ██ ",
	"███████    ██    ██   ██ ██   ██    ██    ███████ ██   ██ ██      ██ ",
}

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

func Bold(s string) string {
	return bold + s + reset
}

func SQuote(s string) string {
	return "'" + s + "'"
}

func PQuote(s string) string {
	return "(" + s + ")"
}
