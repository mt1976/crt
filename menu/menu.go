package menu

import (
	"fmt"
	"log"
	"strings"

	"github.com/mt1976/admin_me/support"
)

const maxMenuItems = 15

type menu struct {
	title        string
	menuItems    []menuItem
	noItems      int
	prompt       string
	actions      []string
	actionMaxLen int
}

type menuItem struct {
	menuItemNumber       int
	menuItemNumberString string
	menuItemTitle        string
}

func NewMenu(title string) *menu {
	m := menu{title: title, menuItems: []menuItem{}, noItems: 0, prompt: promptString}
	return &m
}

func (m *menu) AddMenuItem(menuItemNumber int, menuItemTitle string) {
	if m.noItems >= maxMenuItems {
		log.Fatal(m.title + " - Max menu items reached")
		return
	}
	mi := menuItem{menuItemNumber, fmt.Sprintf("%2v", menuItemNumber), menuItemTitle}
	m.menuItems = append(m.menuItems, mi)
	m.noItems++
}

func (m *menu) AddAction(validAction string) {
	if validAction == "" {
		log.Fatal("Invalid action")
		return
	}
	validAction = strings.ReplaceAll(validAction, " ", "")
	m.actions = append(m.actions, validAction)
	if len(validAction) > m.actionMaxLen {
		m.actionMaxLen = len(validAction)
	}
}

func Run(crt *support.Crt) {
	crt.Clear()
	crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Main Menu")
	m := NewMenu(mainMenuTitle)
	for i := range 11 {
		m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
	}
	action := DisplayMenu(m, crt)
	y := NewMenu("Sub Menu")
	for i := range 14 {
		y.AddMenuItem(i, fmt.Sprintf("Sub Menu Item %v", action))
	}
	action = DisplayMenu(y, crt)
	crt.Println("Final Action: " + action)
}

func DisplayMenu(m *menu, crt *support.Crt) (nextAction string) {
	crt.Clear()
	m.AddAction("Q") // Add Quit action
	crt.Header(m.title)
	for i := range m.menuItems {
		crt.Println(printmenuItem(crt, m.menuItems[i].menuItemNumber, m.menuItems[i].menuItemTitle))
		m.AddAction(m.menuItems[i].menuItemNumberString) // Add action for each menu item
	}
	extraRows := (maxMenuItems - m.noItems) + 1
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
			crt.InputError("Invalid action '" + nextAction + "'")
			//crt.Shout("Invalid action '" + crt.Bold(nextAction) + "'")
			continue
		}

		for i := range m.actions {
			if upcase(nextAction) == upcase(m.actions[i]) {
				ok = true
				break
			}
		}
		if !ok {
			//crt.Shout("Invalid action '" + crt.Bold(nextAction) + "'")
			crt.InputError("Invalid action '" + nextAction + "'")

		}
	}
	//spew.Dump(m)
	return upcase(nextAction)
}

func upcase(s string) string {
	return strings.ToUpper(s)
}

func printmenuItem(crt *support.Crt, pos int, title string) string {
	miNumber := fmt.Sprintf(crt.Bold("%2v"), pos)
	miString := fmt.Sprintf(miNumber + ") " + title)
	return miString
}
