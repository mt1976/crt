package actions

var (
	// Actions
	Yes         *Action = New("Y")
	No          *Action = New("N")
	Quit        *Action = New("Q")
	Forward     *Action = New("F")
	Back        *Action = New("B")
	Exit        *Action = New("EX")
	Help        *Action = New("?") // Help
	Up          *Action = New("U")
	UpDoubleDot *Action = New("..")
	UpArrow     *Action = New("^")
	Go          *Action = New("G")
	Select      *Action = New("S")
)
