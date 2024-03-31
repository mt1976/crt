package spinner

import (
	"log"
)

// New returns a new Spinner
// The New function returns a new instance of the Spinner type.
func New() *Spinner {
	// ...
	return new()
}

// Tick advances the spinner to the next state
// The `Tick()` function is a method of the `Spinner` type. It advances the spinner to the next state
// without displaying any message. It calls the `tick()` method of the `Spinner` instance with an empty
// string as the message parameter.
func (s *Spinner) Tick() {
	s.tick("")
}

// TickWithMessage advances the spinner to the next state and displays a message
// The `TickWithMessage` function is a method of the `Spinner` type. It advances the spinner to the
// next state and displays a message. It calls the `tick` method of the `Spinner` instance with the
// `message` parameter.
func (s *Spinner) TickWithMessage(message string) {
	s.tick(message)
}

// Style sets the style of the spinner
// The `Style` method is a function of the `Spinner` type. It sets the style of the spinner by calling
// the `setStyle` method of the `Spinner` instance with the `style` parameter. It then returns the
// `Spinner` instance.
func (s *Spinner) Style(style framesIndex) *Spinner {
	return s.setStyle(style)
}

//func (s *Spinner) SetLocation(row int, column int) *Spinner {
//	// ...
//	return s.setLocation(row, column)
//}

// Debug prints debug information to stdout
// The `Debug()` function is a method of the `Spinner` type. It is used to print debug information to
// the standard output.
func (s *Spinner) Debug() {
	log.Println("Debug")
	log.Println("style:", s.style)
	//log.Println("row:", s.row)
	//log.Println("column:", s.column)
	log.Println("frames:", s.frames)
	log.Println("speed:", s.slow)
	log.Println("Styles:", s.Styles)
	log.Println("cycle:", s.cycle)

}

// Delay sets the delay between frames
// The `Delay` function is a method of the `Spinner` type. It sets the delay between frames of the
// spinner animation. It takes a `seconds` parameter of type `float64` which represents the delay in
// seconds. It calls the `setDelay` method of the `Spinner` instance with the `seconds` parameter to
// set the delay. Finally, it returns the `Spinner` instance to allow for method chaining.
func (s *Spinner) Delay(seconds float64) *Spinner {
	s.setDelay(seconds)
	return s
}
