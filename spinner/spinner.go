package spinner

import (
	"fmt"
	"time"
)

type framesIndex int

// The Spinner type represents a spinning animation with various styles and settings.
// @property {framesIndex} style - The style property is of type framesIndex, which is used to
// determine the style of the spinner. It is an index that points to a specific style in the
// spinnerStyles struct.
// @property {int} row - The `row` property represents the current row position of the spinner. It is
// used to determine where the spinner should be displayed on the screen or in a terminal window.
// @property {int} column - The `column` property represents the column position of the spinner in a
// grid or table layout. It determines where the spinner will be displayed horizontally.
// @property {[]string} frames - The `frames` property is a slice of strings that represents the
// different frames of the spinner animation. Each string in the slice represents a single frame of the
// animation.
// @property {int} cycle - The `cycle` property represents the number of times the spinner animation
// should repeat before stopping.
// @property {int} sequence - The `sequence` property is an integer that represents the current
// position in the sequence of frames for the spinner. It is used to determine which frame to display
// next when the spinner is animated.
// @property slow - The `slow` property is of type `time.Duration` and represents the duration of each
// frame in the spinner animation. It determines how fast or slow the spinner rotates.
// @property Styles - The `Styles` property is a pointer to a `spinnerStyles` struct.
type Spinner struct {
	// ...
	style    framesIndex
	row      int
	column   int
	frames   []string
	cycle    int
	sequence int
	slow     time.Duration
	Styles   *spinnerStyles
}

type spinnerStyles struct {
	// ...
	Default        framesIndex
	Plus           framesIndex
	Directions     framesIndex
	Dots           framesIndex
	Ball           framesIndex
	SquareClock    framesIndex
	Clock          framesIndex
	Snake          framesIndex
	ChasingDots    framesIndex
	Arrows         framesIndex
	Grow           framesIndex
	Cross          framesIndex
	Flip           framesIndex
	Cylon          framesIndex
	DirectionsSlow framesIndex
}

// new returns a new Spinner, with defaults
func new() *Spinner {
	sp := &Spinner{row: 0, column: 0}
	sp.sequence = 0
	sp.slow = 0
	sp.Styles = initialiseStyles()
	sp.style = sp.Styles.Default
	sp.frames = sp.getFrames(sp.style)
	sp.cycle = len(sp.frames)
	return sp
}

// The `tick` function is responsible for advancing the spinner to the next state and displaying the
// current frame of the spinner animation.
func (s *Spinner) tick(msg string) {
	// ...
	//	log.Println("tick")
	s.sequence = (s.sequence + 1)
	if s.sequence >= s.cycle {
		s.sequence = 0
	}
	//	log.Println("sequence:", s.sequence)
	fmt.Print("\033[u\033[K[" + s.frames[s.sequence] + "] " + msg)
	if s.slow > 0 {
		//		log.Println("sleeping")
		time.Sleep(s.slow)
	}
}

// setStyle sets the style of the spinner
// The `setStyle` function is a method of the `Spinner` struct. It takes a `style` parameter of type
// `framesIndex` and sets the `style` property of the `Spinner` to the provided value. It then updates
// the `frames` property of the `Spinner` by calling the `getFrames` method with the new style. The
// `cycle` property is updated to the length of the new frames, and the `sequence` property is reset to
// 0. Finally, it returns a pointer to the updated `Spinner` object.
func (s *Spinner) setStyle(style framesIndex) *Spinner {
	// ...
	s.style = style
	s.frames = s.getFrames(style)
	s.cycle = len(s.frames)
	s.sequence = 0
	return s
}

// setLocation sets the location of the spinner
// The `setLocation` function is a method of the `Spinner` struct. It takes two parameters, `row` and
// `column`, which represent the new row and column positions of the spinner.
func (s *Spinner) setLocation(row int, column int) *Spinner {
	// ...
	s.row = row
	s.column = column
	return s
}

// getFrames returns the characters for a given style
// The `getFrames` function is a method of the `Spinner` struct. It takes a `style` parameter of type
// `framesIndex` and returns a slice of strings representing the frames of the spinner animation for
// the given style.
func (s *Spinner) getFrames(style framesIndex) []string {

	rtn := []string{"-", "\\", "|", "/"}

	switch style {
	case s.Styles.Default:
	//	rtn = []string{"-", "\\", "|", "/"}
	case s.Styles.Plus:
		rtn = []string{"+", "x"}
	case s.Styles.Directions:
		rtn = []string{"v", "<", "^", ">"}
	case s.Styles.Dots:
		rtn = []string{".   ", " .  ", "  . ", "   ."}
	case s.Styles.Ball:
		rtn = []string{"◐", "◓", "◑", "◒"}
	case s.Styles.SquareClock:
		rtn = []string{"◰", "◳", "◲", "◱"}
	case s.Styles.Clock:
		rtn = []string{"◴", "◷", "◶", "◵"}
	case s.Styles.Snake:
		rtn = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	case s.Styles.ChasingDots:
		rtn = []string{".  ", ".. ", "...", " ..", "  .", "   "}
	case s.Styles.Arrows:
		rtn = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	case s.Styles.Grow:
		rtn = []string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"}
	case s.Styles.Cross:
		rtn = []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}
	case s.Styles.Flip:
		rtn = []string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"}
	case s.Styles.Cylon:
		rtn = []string{"( ●    )",
			"(  ●   )",
			"(   ●  )",
			"(    ● )",
			"(     ●)",
			"(    ● )",
			"(   ●  )",
			"(  ●   )",
			"( ●    )",
			"(●     )"}
	case s.Styles.DirectionsSlow:
		rtn = []string{"<", "<", "∧", "∧", ">", ">", "v", "v"}
	default:
		//rtn = []string{"-", "\\", "|", "/"}
	}
	// ...
	return rtn
}

// initialiseStyles sets the default styles
// The function `initialiseStyles()` initializes and returns a pointer to a `spinnerStyles` struct with
// predefined values for different spinner styles.
func initialiseStyles() *spinnerStyles {
	// ...
	s := &spinnerStyles{}
	s.Default = 1
	s.Plus = 2
	s.Directions = 3
	s.Dots = 4
	s.Ball = 5
	s.SquareClock = 6
	s.Clock = 7
	s.Snake = 8
	s.ChasingDots = 9
	s.Arrows = 10
	s.Grow = 11
	s.Cross = 12
	s.Flip = 13
	s.Cylon = 14
	s.DirectionsSlow = 15
	return s
}

// Delay sets the delay between frames
// The `setDelay` function is a method of the `Spinner` struct. It takes a `seconds` parameter of type
// `float64` and sets the `slow` property of the `Spinner` to the duration specified by `seconds`.
func (s *Spinner) setDelay(seconds float64) *Spinner {
	nanos := time.Second.Nanoseconds()
	seconds = float64(nanos) * seconds
	s.slow = time.Duration(seconds)
	return s
}
