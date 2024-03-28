package display

import "fmt"

// gtrm.MoveCursor(startColumn, 21)
// gtrm.Print(t.row())
// gtrm.Println(p.FormatRowOutput(p.pageRows[i].RowContent))
// gtrm.Flush()
// gtrm.Clear()

func PrintAt(content string, x, y int) {
	fmt.Print(content)
}

func Flush() {
	// Do nothing
}

func Clear() {
	// Do nothing
}

func MoveCursor(x, y int) {
	// Do nothing
}

func Println(content string) {
	// Do nothing
	Print(content + "\n")
}

func Print(content string) {
	// Do nothing
}
