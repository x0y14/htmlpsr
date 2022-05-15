package htmlpsr

type Lexer struct {
	pos   int
	runes []rune
}

func NewLexer() *Lexer {
	return &Lexer{
		pos:   0,
		runes: nil,
	}
}

func (l *Lexer) Lex() {}
