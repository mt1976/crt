package symbols

type Symbol struct {
	content string
	len     int
}

func New(content string) *Symbol {
	return &Symbol{
		content: content,
		len:     len(content),
	}
}

func (s *Symbol) Symbol() string {
	return s.content
}

func (s *Symbol) Len() int {
	return s.len
}

func (s *Symbol) String() string {
	return s.content
}

func (s *Symbol) Rune() []rune {
	return []rune(s.content)
}

func (s *Symbol) Equals(b string) bool {
	return s.content == b
}
