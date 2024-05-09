package page

import (
	"fmt"
	"strings"

	symb "github.com/mt1976/crt/strings/symbols"
)

// cleanContent removes unwanted characters from the rowContent string
func cleanContent(msg string) string {
	// replace \n, \r, \t, and " with empty strings
	msg = strings.Replace(msg, symb.Newline.Symbol(), "", -1)
	msg = strings.Replace(msg, symb.CarridgeReturn.Symbol(), "", -1)
	msg = strings.Replace(msg, symb.Tab.Symbol(), "", -1)
	msg = strings.Replace(msg, symb.DoubleQuote.Symbol(), symb.Space.Symbol(), -1)
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

// The function "format" takes a pointer to a Crt object and a menuItem object, and returns a
// formatted string containing the menu item's ID, title, and date.
func (p *Page) formatOption(row pageRow) string {
	miNumber := fmt.Sprintf(bold("%3v"), row.ID)

	//add Date to end of row
	miTitle := row.Title
	//padd out to 70 characters
	width := p.width - 7
	pad := width - (len(miTitle) + len(row.DateTime))
	if pad > 0 {
		miTitle = miTitle + strings.Repeat(symb.Space.Symbol(), pad)
	} else {
		miTitle = miTitle[:width-(len(row.DateTime)+1)] + " | " + row.DateTime
	}

	miString := fmt.Sprintf(miNumber + ") " + miTitle)
	return miString
}
