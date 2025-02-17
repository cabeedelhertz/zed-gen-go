package parser

import (
	"errors"
	"zedgen/internal/lexer/scanner"
	"zedgen/internal/parser/meta"
)

type Definition struct {
	// Name of the object
	Name string
	// the relations
	Relations []Relation
	// the permissions
	Permissions []Permission

	// Comments associated with the object
	Comments       []string
	InlineComments []string
	// meta information
	Meta meta.Meta
}

func (p *Parser) ParseDefinition() (*Definition, error) {
	p.lex.NextKeyword()
	if p.lex.Token != scanner.TDEFINITION {
		return nil, errors.New("expected definition keyword")
	}
	startPos := p.lex.Pos

	p.lex.ConsumeToken(scanner.TIDENT)
	if p.lex.Token != scanner.TIDENT {
		return nil, errors.New("expected definition name")
	}
	defName := p.lex.Text

	p.lex.ConsumeToken(scanner.TLEFTCURLY)
	if p.lex.Token != scanner.TLEFTCURLY {
		return nil, errors.New("expected left curly brace")
	}

	relations := make([]Relation, 0)
	permissions := make([]Permission, 0)
	for {
		var comments []string
		for p.lex.TryNextComment() {
			comments = append(comments, p.lex.Text)
		}

		p.lex.ConsumeToken(scanner.TRIGHTCURLY)
		if p.lex.Token == scanner.TRIGHTCURLY {
			break
		}

		if p.IsNextEOF() {
			break
		}

		if p.lex.IsNextKeyword(scanner.TRELATION) {
			relation, err := p.ParseRelation()
			if err != nil {
				return nil, err
			}
			if relation == nil {
				return nil, errors.New("expected relation")
			}
			relation.Comments = comments
			relations = append(relations, *relation)
			continue
		} else if p.lex.IsNextKeyword(scanner.TPERMISSION) {
			permission, err := p.ParsePermission()
			if err != nil {
				return nil, err
			}
			if permission == nil {
				return nil, errors.New("expected permission")
			}
			permission.Comments = comments
			permissions = append(permissions, *permission)
			continue
		} else {
			return nil, errors.New("expected relation or permission")
		}
	}

	return &Definition{
		Name:        defName,
		Relations:   relations,
		Permissions: permissions,
		Meta: meta.Meta{
			Pos:     startPos.Position,
			LastPos: p.lex.Pos.Position,
		},
	}, nil
}
