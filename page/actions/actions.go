package actions

var (
	// Actions
	Yes         *Action = NewAction("Y")
	No          *Action = NewAction("N")
	Quit        *Action = NewAction("Q")
	Forward     *Action = NewAction("F")
	Back        *Action = NewAction("B")
	Exit        *Action = NewAction("EX")
	Help        *Action = NewAction("?") // Help
	Up          *Action = NewAction("U")
	UpDoubleDot *Action = NewAction("..")
	UpArrow     *Action = NewAction("^")
	Go          *Action = NewAction("G")
	Select      *Action = NewAction("S")
)
