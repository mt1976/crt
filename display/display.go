package displarow

import (
	"fmt"

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
	fmt.Print(content)
}

func Flush() {
	// Do nothing
}

func Clear() {
	// Do nothing
	ansi.ClearScreen()
}

func ClearLine(row int) {
	ansi.MoveTo(0, row)
	ansi.ClearLine()
}

func MoveCursor(column, row int) {
	ansi.MoveTo(column, row+2)
}

func Println(content string) {
	Print(content + "\n")
}

func Print(content string) {
	fmt.Print(content)
}
