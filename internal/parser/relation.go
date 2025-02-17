package parser

import (
	"errors"
	"zedgen/internal/lexer/scanner"
	"zedgen/internal/parser/meta"
)

type Relation struct {
	// Name of the relation
	Name string
	// Definition of the relation
	Definition string
	// Comments associated with the relation
	Comments       []string
	InlineComments []string
	// meta information
	Meta meta.Meta
}

var relationDefDelimitters = map[scanner.Token]interface{}{
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
	definition := ""
	// get first identifier token
	p.lex.Next()
	if p.lex.Token == scanner.TIDENT {
		definition += p.lex.Text
	}
	// if there are options or relations after the first identifier, handle those
	for {
		p.lex.Next()
		if _, ok := relationDefDelimitters[p.lex.Token]; ok {
			definition += p.lex.Text
			p.lex.Next()
			if p.lex.Token != scanner.TIDENT && p.lex.Token != scanner.TWILDCARD {
				return nil, errors.New("expected relation definition")
			}
			definition += p.lex.Text
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
		Definition:     definition,
		InlineComments: inlineComments,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
