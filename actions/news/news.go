package menu

import (
	"fmt"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/mt1976/admin_me/support"
)

func Run(crt *support.Crt) {
	//crt.Clear()
	//crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Menu")
	//crt.Print(fmt.Sprintf("Delay in seconds: %v", crt.GetDelayInSec()))

	if err := ui.Init(); err != nil {
		//log.Fatalf("failed to initialize termui: %v", err)
		crt.Error("failed to initialize termui", err)
	}
	defer ui.Close()

	l := widgets.NewList()
	l.Title = "StarTerm [News]"
	l.Rows = []string{
		menuItem(0, "News 1"),
		menuItem(1, "News 2"),
		menuItem(2, "github.com/mt1976/admin_me/support"),
		menuItem(3, "github.com/mt1976/admin_me/support/crt"),
		menuItem(4, "[你好，世界](fg:blue)"),
		menuItem(5, "[こんにちは世界](fg:red)"),
		menuItem(6, "[color](fg:white,bg:green) output"),
		menuItem(7, "random_out.go"),
		menuItem(8, "dashboard.go"),
		menuItem(9, "foo"),
		menuItem(10, "bar"),
		menuItem(11, "baz"),
		menuItem(12, "qux"),
	}

	l.TextStyle = ui.NewStyle(ui.ColorGreen)
	l.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorGreen)
	l.WrapText = false
	l.SetRect(0, crt.CurrentRow(), crt.Width(), crt.Height()-crt.CurrentRow())
	l.TitleStyle.Fg = ui.ColorGreen + ui.Color(ui.ModifierBold)
	l.BorderStyle.Fg = ui.ColorGreen

	ui.Render(l)
	x := widgets.NewParagraph()
	x.Title = "StarTerm [Actions]"
	x.Text = "Select [0-%v](mod:bold) or [q](mod:bold) to quit"
	x.Text = fmt.Sprintf(x.Text, len(l.Rows)-1)
	x.SetRect(0, 0, crt.Width(), crt.CurrentRow())
	x.BorderStyle.Fg = ui.ColorGreen
	x.TitleStyle.Fg = ui.ColorGreen + ui.Color(ui.ModifierBold)
	x.TextStyle = ui.NewStyle(ui.ColorGreen)
	ui.Render(x)
	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			l.ScrollDown()
		case "k", "<Up>":
			l.ScrollUp()
		case "<C-d>":
			l.ScrollHalfPageDown()
		case "<C-u>":
			l.ScrollHalfPageUp()
		case "<C-f>":
			l.ScrollPageDown()
		case "<C-b>":
			l.ScrollPageUp()
		case "g":
			if previousKey == "g" {
				l.ScrollTop()
			}
		case "<Home>":
			l.ScrollTop()
		case "G", "<End>":
			l.ScrollBottom()
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}

}

func menuItem(pos int, title string) string {
	return fmt.Sprintf("%2v│ %v", pos, title)
}
