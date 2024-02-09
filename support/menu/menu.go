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

type menu struct {
	title        string
	menuItems    []menuItem
	noItems      int
	prompt       string
	actions      []string
	actionMaxLen int
}

type menuItem struct {
	ID          int
	Number      string
	Title       string
	AlternateID string
	DateTime    string
}

func New(title string) *menu {
	// truncate title to 25 characters
	if len(title) > 25 {
		title = title[:25] + "..."
	}
	m := menu{title: title, menuItems: []menuItem{}, noItems: 0, prompt: promptString}
	return &m
}

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

func (m *menu) Display(crt *support.Crt) (nextAction string, selected menuItem) {
	crt.Clear()
	m.AddAction("Q") // Add Quit action
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
		crt.Print("\n")
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
