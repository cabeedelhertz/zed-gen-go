package parser

import (
	"errors"
	"zedgen/internal/lexer/scanner"
	"zedgen/internal/parser/meta"
)

type Permission struct {
	// Name of the permission
	Name string
	// Computation definition of the permission
	Computation string
	// Comments associated with the permission
	Comments       []string
	InlineComments []string
	// meta information
	Meta meta.Meta
}

var permissionCompDelimitters = map[scanner.Token]interface{}{
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
	computation := ""
	// get first identifier token
	p.lex.Next()
	if p.lex.Token == scanner.TIDENT {
		computation += p.lex.Text
	}
	// TODO: parse permission computation logic
	for {
		p.lex.Next()
		if _, ok := permissionCompDelimitters[p.lex.Token]; ok {
			computation += p.lex.Text
			p.lex.ConsumeToken(scanner.TIDENT)
			if p.lex.Token != scanner.TIDENT {
				return nil, errors.New("expected permission computation")
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

	return &Permission{
		Name:           permName,
		Computation:    computation,
		InlineComments: inlineComments,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
