package support

import "strings"

const chNormal = "┃"
const chSpecial = "┣"
const chStart = "┏"
const chEnd = "┛"
const chJunction = "┣"
const chEndFirst = "┓"
const chBar = "━"

// const chClose = "┗"
const bold = "\033[1m"
const reset = "\033[0m"
const underline = "\033[4m"
const red = "\033[31m"
const clearline = "\033[2K"

var header []string
var smHeader string

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

func Upcase(s string) string {
	return strings.ToUpper(s)
}
