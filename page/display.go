package page

import (
	ansi "github.com/bit101/go-ansi"
)

// Replaces/Substitutions for the following gtrm functions:
// gtrm.MoveCursor(startColumn, 21)
// gtrm.Print(t.row())
// gtrm.Println(p.FormatRowOutput(p.pageRows[i].RowContent))
// gtrm.Flush()
// gtrm.Clear()

func PrintAt(content string, column, row int) {
	MoveCursor(column, row)
	ansi.Print(ansi.Green, content)
}

func Flush() {
	// Do nothing
}

func Clear() {
	// Do nothing
	ansi.ClearScreen()
}

func ClearLine(row int) {
	MoveCursor(0, row)
	ansi.ClearLine()
}

func MoveCursor(column, row int) {
	ansi.MoveTo(column, row)
}

func Println(content string) {
	ansi.Println(ansi.Green, content)
}

func Print(content string) {
	ansi.Print(ansi.Green, content)
}
