package lexer

import (
	"io"
	"zedgen/internal/lexer/scanner"
)

type Lexer struct {
	// Token is the lexical token
	Token scanner.Token

	// the lexical value
	Text string

	RawText []rune

	Pos scanner.Position

	scanner *scanner.Scanner

	scanErr error
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		scanner: scanner.NewScanner(r),
	}
}

func (l *Lexer) Next() {
	var err error
	l.Token, l.Text, l.Pos, err = l.scanner.Scan()
	l.RawText = l.scanner.LastScanRaw()
	if err != nil {
		l.scanErr = err
	}
}

func (l *Lexer) NextN(n int) {
	for i := 0; i < n; i++ {
		l.Next()
	}
}

func (l *Lexer) UnNext() {
	l.Pos = l.scanner.UnScan()
	l.Token = scanner.TILLEGAL
}

func (l *Lexer) IsEOF() bool {
	return l.Token == scanner.TEOF
}

func (l *Lexer) TryNextComment() bool {
	l.nextInMode(scanner.ScanComment)
	if l.Token == scanner.TCOMMENT {
		return true
	}
	l.UnNext()
	return false
}

func (l *Lexer) NextKeyword() {
	l.nextInMode(scanner.ScanKeyword)
}

func (l *Lexer) NextIdentifier() {
	l.nextInMode(scanner.ScanIdent)
}

func (lex *Lexer) NextKeywordOrIdentifier() {
	lex.nextInMode(scanner.ScanKeyword | scanner.ScanIdent)
}

func (l *Lexer) nextInMode(mode scanner.Mode) {
	oldMode := l.scanner.Mode
	defer func() {
		l.scanner.Mode = oldMode
	}()
	l.scanner.Mode = mode
	l.Next()
}

func (l *Lexer) ConsumeToken(t scanner.Token) {
	l.Next()
	if l.Token == t {
		return
	}
	l.UnNext()
}

func (l *Lexer) IsNextKeyword(t scanner.Token) bool {
	l.NextKeyword()
	defer l.UnNext()
	return l.Token == t
}

func (l *Lexer) TryInlineComment() string {
	if comment, err := l.scanner.TryInlineComment(); err == nil && comment != "" {
		return comment
	}
	return ""
}
