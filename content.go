package crt

import (
	"fmt"
	"strings"

	lang "github.com/mt1976/crt/language"
)

// The "pageContent" type represents a pageContent with a map of rows and columns.
// @property row - The "row" property is a map that stores the values of each row in the pageContent. The keys
// of the map are integers representing the row numbers, and the values are strings representing the
// content of each row.
// @property {int} cols - The "cols" property represents the number of columns in the pageContent.
// @property {int} rows - The "rows" property represents the number of rows in the pageContent.
type pageContent struct {
	row  map[int]string
	cols int
	rows int
}

// cleanContent removes unwanted characters from the rowContent string
func cleanContent(msg string) string {
	// replace \n, \r, \t, and " with empty strings
	msg = strings.Replace(msg, lang.SymNewline, "", -1)
	msg = strings.Replace(msg, lang.SymCarridgeReturn, "", -1)
	msg = strings.Replace(msg, lang.SymTab, "", -1)
	msg = strings.Replace(msg, lang.SymDoubleQuote, lang.Space, -1)
	return msg
}

// isInList determines if the given action is in the list of available actions
func isInList(value string, list []string) bool {
	// loop through each action in the list
	for i := range list {
		// if the given action matches an action in the list, return true
		if value == list[i] {
			return true
		}
	}
	// if no match was found, return false
	return false
}

// The format function returns the first 50 characters of the content in a pageRow object.
// format returns the first n characters of the content in a pageRow object.
func format(t *Crt, m pageRow) string {
	return m.Content
}

// The function "format" takes a pointer to a Crt object and a menuItem object, and returns a
// formatted string containing the menu item's ID, title, and date.
func formatOption(row pageRow) string {
	miNumber := fmt.Sprintf(bold("%3v"), row.ID)

	//add Date to end of row
	miTitle := row.Title
	//padd out to 70 characters
	width := config.TerminalWidth - 7
	pad := width - (len(miTitle) + len(row.DateTime))
	if pad > 0 {
		miTitle = miTitle + strings.Repeat(lang.Space, pad)
	} else {
		miTitle = miTitle[:width-(len(row.DateTime)+1)] + " | " + row.DateTime
	}

	miString := fmt.Sprintf(miNumber + ") " + miTitle)
	return miString
}
