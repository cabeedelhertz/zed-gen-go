package parser

import (
	"errors"
	"zedgen/internal/lexer/scanner"
	"zedgen/internal/parser/meta"
)

type Relation struct {
	// Name of the relation
	Name string
	// Computation definition of the relation
	Computation string
	// Comments associated with the relation
	Comments       []string
	InlineComments []string
	// meta information
	Meta meta.Meta
}

var relationCompDelimitters = map[scanner.Token]interface{}{
	scanner.THASH:  nil,
	scanner.TPIPE:  nil,
	scanner.TCOLON: nil,
}

func (p *Parser) ParseRelation() (*Relation, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner.TRELATION {
		return nil, errors.New("expected relation keyword")
	}
	startPos := p.lex.Pos

	p.lex.NextIdentifier()
	if p.lex.Token != scanner.TIDENT {
		return nil, errors.New("expected relation name")
	}
	relName := p.lex.Text

	p.lex.Next()
	if p.lex.Token != scanner.TCOLON {
		return nil, errors.New("expected colon")
	}
	computation := ""
	// get first identifier token
	p.lex.Next()
	if p.lex.Token == scanner.TIDENT {
		computation += p.lex.Text
	}
	// if there are options or relations after the first identifier, handle those
	for {
		p.lex.Next()
		if _, ok := relationCompDelimitters[p.lex.Token]; ok {
			computation += p.lex.Text
			p.lex.Next()
			if p.lex.Token != scanner.TIDENT && p.lex.Token != scanner.TWILDCARD {
				return nil, errors.New("expected relation computation")
			}
			computation += p.lex.Text
		} else {
			p.lex.UnNext()
			break
		}
	}

	var inlineComments []string
	if il := p.lex.TryInlineComment(); il != "" {
		inlineComments = append(inlineComments, il)
	}

	return &Relation{
		Name:           relName,
		Computation:    computation,
		InlineComments: inlineComments,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
