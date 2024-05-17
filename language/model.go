package language

import (
	"strings"
	"unicode"

	symb "github.com/mt1976/crt/strings/symbols"
)

type Text struct {
	// General
	content string
	len     int
}

type Paragraph struct {
	content []Text
	len     int
}

func New(message string) *Text {
	return &Text{
		content: message,
		len:     len(message),
	}
}

func NewParagraph(message []string) *Paragraph {
	para := &Paragraph{
		len: len(message),
	}
	for _, m := range message {
		para.content = append(para.content, *New(m))
	}
	return para
}

func (t *Text) Text() string {
	return t.content
}

func (t *Text) Len() int {
	return t.len
}

func isMessageInt(message string) bool {
	for _, c := range message {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func (p *Paragraph) Len() int {
	return p.len
}

func (p *Paragraph) String() []string {
	out := []string{}
	for _, t := range p.content {
		out = append(out, t.Text()+symb.Newline.Symbol())
	}
	return out
}

func (p *Paragraph) Add(message string) {
	p.content = append(p.content, *New(message))
	p.len++
}

func (p *Paragraph) AddBlankRow() {
	p.content = append(p.content, *New(symb.Newline.Symbol()))
	p.len++
}

func upcase(in string) string {
	return strings.ToUpper(in)
}

func (t *Text) IsEmpty() bool {
	return t.Len() == 0
}
