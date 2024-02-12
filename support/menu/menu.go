package menu

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/mt1976/admin_me/support"
	"github.com/xeonx/timeago"
)

const MaxMenuItems = 15

// The "menu" type represents a menu with a title, menu items, a prompt, and associated actions.
// @property {string} title - The title property represents the title of the menu.
// @property {[]menuItem} menuItems - The `menuItems` property is an array of `menuItem` objects. Each
// `menuItem` object represents an item in the menu and contains information such as the item's title,
// description, and action.
// @property {int} noItems - The `noItems` property represents the number of menu items in the menu.
// @property {string} prompt - The prompt is a string that represents the message or question displayed
// to the user to prompt them for input or action.
// @property {[]string} actions - The "actions" property is a slice of strings that represents the
// available actions or options for the menu. Each string in the slice represents an action that the
// user can choose from when interacting with the menu.
// @property {int} actionMaxLen - The property "actionMaxLen" represents the maximum length of the
// actions in the menu. It is used to determine the width of the menu display, ensuring that all
// actions are properly aligned.
type menu struct {
	title        string
	menuItems    []menuItem
	noItems      int
	prompt       string
	actions      []string
	actionMaxLen int
}

// The above type represents a menu item with various properties such as ID, number, title, alternate
// ID, and date/time.
// @property {int} ID - An integer representing the unique identifier of the menu item.
// @property {string} Number - The "Number" property of the menuItem struct is a string that represents
// the number associated with the menu item.
// @property {string} Title - The "Title" property of the menuItem struct represents the title or name
// of the menu item. It is a string type property.
// @property {string} AlternateID - The AlternateID property is a string that represents an alternate
// identifier for the menuItem. It can be used as an additional way to identify the menuItem, apart
// from the ID property.
// @property {string} DateTime - The DateTime property in the menuItem struct represents the date and
// time associated with the menu item. It is of type string.
type menuItem struct {
	ID          int
	Number      string
	Title       string
	AlternateID string
	DateTime    string
}

// The New function creates a new menu with a truncated title and initializes its properties.
func New(title string) *menu {
	// truncate title to 25 characters
	if len(title) > 25 {
		title = title[:25] + "..."
	}
	m := menu{title: title, menuItems: []menuItem{}, noItems: 0, prompt: promptString}
	return &m
}

// The `Add` function is a method of the `menu` struct. It is used to add a new menu item to the menu.
func (m *menu) Add(menuItemNumber int, menuItemTitle string, altID string, dateTime string) {
	if m.noItems >= MaxMenuItems {
		log.Fatal(m.title + " " + maxMenuItemsError)
		return
	}
	if len(menuItemTitle) > 50 {
		menuItemTitle = menuItemTitle[:50] + "..."
	}
	if menuItemTitle != "" {
		m.AddAction(fmt.Sprintf("%v", menuItemNumber))
	}
	mi := menuItem{menuItemNumber, fmt.Sprintf("%2v", menuItemNumber), menuItemTitle, altID, dateTime}
	m.menuItems = append(m.menuItems, mi)
	m.noItems++
}

// The `AddAction` method is used to add a valid action to the menu. It takes a `validAction` string as
// a parameter.
func (m *menu) AddAction(validAction string) {
	if validAction == "" {
		log.Fatal(invalidActionError)
		return
	}
	validAction = strings.ReplaceAll(validAction, " ", "")
	m.actions = append(m.actions, validAction)
	if len(validAction) > m.actionMaxLen {
		m.actionMaxLen = len(validAction)
	}
}

// The `Display` method of the `menu` struct is responsible for displaying the menu to the user and
// handling user input. Here is a breakdown of what it does:
func (m *menu) Display(crt *support.Crt) (nextAction string, selected menuItem) {
	crt.Clear()
	m.AddAction(Quit) // Add Quit action
	crt.Header(m.title)
	for i := range m.menuItems {
		if m.menuItems[i].Title == "" {
			crt.Println("")
			continue
		}
		crt.Println(format(crt, m.menuItems[i]))
		m.AddAction(m.menuItems[i].Number) // Add action for each menu item
	}
	extraRows := (MaxMenuItems - m.noItems) + 1
	//log.Println("Extra Rows: ", extraRows)
	for i := 0; i <= extraRows; i++ {
		crt.Print(newline)
	}
	crt.Break()
	//crt.Print(m.prompt)
	ok := false
	for !ok {
		nextAction = crt.Input(m.prompt, "")
		if len(nextAction) > m.actionMaxLen {
			crt.InputError(invalidActionError + "'" + nextAction + "'")
			//crt.Shout("Invalid action '" + crt.Bold(nextAction) + "'")
			continue
		}

		for i := range m.actions {
			if support.Upcase(nextAction) == support.Upcase(m.actions[i]) {
				ok = true
				break
			}
		}
		if !ok {
			//crt.Shout("Invalid action '" + crt.Bold(nextAction) + "'")
			crt.InputError(invalidActionError + " '" + nextAction + "'")

		}
	}
	// if nextAction is a numnber, find the menu item
	if support.IsInt(nextAction) {
		pos, _ := strconv.Atoi(nextAction)
		return support.Upcase(nextAction), m.menuItems[pos]
	}
	//spew.Dump(m)
	return support.Upcase(nextAction), menuItem{}
}

// The function "format" takes a pointer to a support.Crt object and a menuItem object, and returns a
// formatted string containing the menu item's ID, title, and date.
func format(crt *support.Crt, m menuItem) string {
	miNumber := fmt.Sprintf(crt.Bold("%2v"), m.ID)
	//spew.Dump(m)
	// Example time Thu, 25 Jan 2024 09:56:00 +0000
	// Setup a time format and parse the time
	if m.DateTime != "" {
		mdt, _ := time.Parse(time.RFC1123Z, m.DateTime)
		m.DateTime = timeago.English.Format(mdt)
	}
	//add Date to end of row
	miTitle := m.Title
	//padd out to 70 characters
	pad := 74 - (len(miTitle) + len(m.DateTime))

	miTitle = miTitle + strings.Repeat(" ", pad)
	miDate := m.DateTime
	miString := fmt.Sprintf(miNumber + ") " + miTitle + " " + miDate)
	return miString
}
