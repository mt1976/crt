package actions

import (
	"strings"
	"unicode"

	lang "github.com/mt1976/crt/language"
)

type Action struct {
	content string
	len     int
	isNum   bool
}

func NewAction(message string) *Action {
	action := &Action{}
	action.content = strings.ReplaceAll(message, lang.Space.Symbol(), "")
	action.len = len(message)
	action.isNum = isMessageInt(message)
	return action
}

func (a *Action) Action() string {
	return a.content
}

func (a *Action) Len() int {
	return a.len
}

func (a *Action) Equals(b string) bool {
	sideA := strings.ToUpper(a.content)
	sideB := strings.ToUpper(b)
	return sideA == sideB
}

func (a *Action) Is(b *Action) bool {

	sideA := strings.ToUpper(a.content)
	sideB := strings.ToUpper(b.content)
	if sideA != sideB {
		return false
	}

	if a.len != b.len {
		return false
	}
	return true
}

func isMessageInt(message string) bool {
	for _, c := range message {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func (a *Action) IsInt() bool {
	return a.isNum
}

// isActionIn determines if the input string contains any of the specified actions.
// It is case-insensitive.
//
// Parameters:
// in: The input string to search.
// check: The list of actions to check for.
//
// Returns:
// A boolean indicating whether the input string contains any of the specified actions.
func IsActionIn(in string, check ...*Action) bool {
	for i := 0; i < len(check); i++ {
		if strings.Contains(upcase(in), check[i].Action()) {
			return true
		}
	}
	return false
}

func upcase(in string) string {
	return strings.ToUpper(in)
}
