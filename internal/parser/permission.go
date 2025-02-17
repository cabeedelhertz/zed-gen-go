package parser

import (
	"errors"
	"zedgen/internal/lexer/scanner"
	"zedgen/internal/parser/meta"
)

type Permission struct {
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

var permissionDefDelimitters = map[scanner.Token]interface{}{
	scanner.TRIGHTARROW: nil,
	scanner.TPLUS:       nil,
	scanner.TMINUS:      nil,
}

func (p *Parser) ParsePermission() (*Permission, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner.TPERMISSION {
		return nil, nil
	}
	startPos := p.lex.Pos

	p.lex.NextIdentifier()
	if p.lex.Token != scanner.TIDENT {
		return nil, nil
	}
	permName := p.lex.Text

	p.lex.Next()
	if p.lex.Token != scanner.TEQUALS {
		return nil, nil
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
		if _, ok := permissionDefDelimitters[p.lex.Token]; ok {
			definition += p.lex.Text
			p.lex.ConsumeToken(scanner.TIDENT)
			if p.lex.Token != scanner.TIDENT {
				return nil, errors.New("expected permission definition")
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

	return &Permission{
		Name:           permName,
		Definition:     definition,
		InlineComments: inlineComments,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
