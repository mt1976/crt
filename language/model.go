package language

import (
	"strings"
	"unicode"
)

type Text struct {
	// General
	content string
	len     int
}

type Symbol struct {
	content string
	len     int
}

type Action struct {
	content string
	len     int
	isNum   bool
}

type Paragraph struct {
	content []Text
	len     int
}

func NewText(message string) *Text {
	return &Text{
		content: message,
		len:     len(message),
	}
}

func NewSymbol(content string) *Symbol {
	return &Symbol{
		content: content,
		len:     len(content),
	}
}

func NewAction(message string) *Action {
	action := &Action{}
	action.content = strings.ReplaceAll(message, Space.Symbol(), "")
	action.len = len(message)
	action.isNum = isMessageInt(message)
	return action
}

func NewParagraph(message []string) *Paragraph {
	para := &Paragraph{
		len: len(message),
	}
	for _, m := range message {
		para.content = append(para.content, *NewText(m))
	}
	return para
}

func (t *Text) Text() string {
	return t.content
}

func (t *Text) Len() int {
	return t.len
}

func (s *Symbol) Symbol() string {
	return s.content
}

func (s *Symbol) Len() int {
	return s.len
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

func (p *Paragraph) Len() int {
	return p.len
}

func (p *Paragraph) String() []string {
	out := []string{}
	for _, t := range p.content {
		out = append(out, t.Text()+Newline.Symbol())
	}
	return out
}

func (p *Paragraph) Add(message string) {
	p.content = append(p.content, *NewText(message))
	p.len++
}

func (p *Paragraph) AddBlankRow() {
	p.content = append(p.content, *NewText(Newline.Symbol()))
	p.len++
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
