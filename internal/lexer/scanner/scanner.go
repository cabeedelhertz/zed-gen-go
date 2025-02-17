package scanner

import (
	"bufio"
	"io"
	"unicode"
)

var eof = rune(0)

type Text struct {
	Lit string
	Pos Position
}

type Scanner struct {
	r              *bufio.Reader
	lastReadBuffer []rune
	lastScanRaw    []rune

	// current source position
	pos *Position

	Mode Mode

	comments []Text
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		r:   bufio.NewReader(r),
		pos: NewPosition(),
	}
}

func (s *Scanner) read() (r rune) {
	defer func() {
		if r == eof {
			return
		}
		s.lastScanRaw = append(s.lastScanRaw, r)

		s.pos.Advance(r)
	}()

	// if we already read the rune (from a peek -> unread), then just return it. Otherwise read it from the reader.
	if 0 < len(s.lastReadBuffer) {
		var ch rune
		ch, s.lastReadBuffer = s.lastReadBuffer[len(s.lastReadBuffer)-1], s.lastReadBuffer[:len(s.lastReadBuffer)-1]
		return ch
	}
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread(ch rune) {
	s.lastReadBuffer = append(s.lastReadBuffer, ch)

	s.pos.Revert(ch)
}

func (s *Scanner) peek() rune {
	ch := s.read()
	if ch != eof {
		s.lastScanRaw = s.lastScanRaw[0 : len(s.lastScanRaw)-1]
		s.unread(ch)
	}
	return ch
}

func (s *Scanner) isEOF() bool {
	ch := s.peek()
	return ch == eof
}

func isLetter(r rune) bool {
	if r < 'A' {
		return false
	}

	if r > 'z' {
		return false
	}

	if r > 'Z' && r < 'a' {
		return false
	}

	return true
}

func isDecimalDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func (s *Scanner) Comments() []Text {
	return s.comments
}

func (s *Scanner) LastScanRaw() []rune {
	r := make([]rune, len(s.lastScanRaw))
	copy(r, s.lastScanRaw)
	return r
}

func (s *Scanner) Scan() (Token, string, Position, error) {
	s.lastScanRaw = s.lastScanRaw[:0]
	return s.scan()
}

func (s *Scanner) UnScan() Position {
	var reversed []rune
	for _, ch := range s.lastScanRaw {
		reversed = append([]rune{ch}, reversed...)
	}
	for _, ch := range reversed {
		s.unread(ch)
	}
	return *s.pos
}

func (s *Scanner) scan() (Token, string, Position, error) {
	ch := s.peek()

	startPos := *s.pos

	switch {
	case unicode.IsSpace(ch):
		s.read()
		return s.scan()
	case s.isEOF():
		return TEOF, "", startPos, nil
	case isLetter(ch) || ch == '_':
		ident := s.scanIdent()
		if s.Mode&ScanKeyword != 0 && asKeywordToken(ident) != TIDENT {
			return asKeywordToken(ident), ident, startPos, nil
		}
		return TIDENT, ident, startPos, nil
	case ch == '-':
		// if the next is '>' then its a right arrow, otherwise its a minus
		s.read()
		if s.peek() == '>' {
			s.read()
			return TRIGHTARROW, "->", startPos, nil
		}
		return TMINUS, "-", startPos, nil

	case ch == '/':
		lit, err := s.scanComment()
		if err != nil {
			return TILLEGAL, "", startPos, err
		}
		return TCOMMENT, lit, startPos, nil

	default:
		return asMiscToken(ch), string(s.read()), startPos, nil
	}

}
