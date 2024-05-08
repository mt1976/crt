package language

type Text struct {
	// General
	content string
	len     int
}

type Paragraph struct {
	content []Text
	len     int
}

func (t *Text) String() string {
	return t.content
}

func (t *Text) Len() int {
	return t.len
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

func (p *Paragraph) Len() int {
	return p.len
}

func (p *Paragraph) String() []string {
	out := []string{}
	for _, t := range p.content {
		out = append(out, t.String()+SymNewline)
	}
	return out
}

func (p *Paragraph) Add(message string) {
	p.content = append(p.content, *New(message))
	p.len++
}

func (p *Paragraph) AddBlankRow() {
	p.content = append(p.content, *New(SymNewline))
	p.len++
}
