package scanner

import (
	"unicode/utf8"
	"zedgen/internal/parser/meta"
)

// Position represents the position of a token in the source code.
type Position struct {
	meta.Position
	// line -> column
	columns map[int]int
}

func NewPosition() *Position {
	return &Position{
		Position: meta.Position{
			Offset: 0,
			Line:   1,
			Column: 1,
		},
		columns: make(map[int]int),
	}
}

func (pos *Position) Advance(r rune) {
	length := utf8.RuneLen(r)
	pos.Offset += length
	if r == '\n' {
		pos.columns[pos.Line] = pos.Column
		pos.Line++
		pos.Column = 1
	} else {
		pos.Column++
	}
}

func (pos *Position) Revert(r rune) {
	length := utf8.RuneLen(r)
	pos.Offset -= length

	if r == '\n' {
		pos.Line--
		pos.Column = pos.columns[pos.Line]
	} else {
		pos.Column--
	}
}
