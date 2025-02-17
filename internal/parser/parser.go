package parser

import (
	"errors"
	"zedgen/internal/lexer"
)

type Parser struct {
	lex *lexer.Lexer
}

func NewParser(lex *lexer.Lexer) *Parser {
	return &Parser{
		lex: lex,
	}
}

// IsNextEOF checks if the next token is EOF
func (p *Parser) IsNextEOF() bool {
	p.lex.Next()
	defer p.lex.UnNext()
	return p.lex.IsEOF()
}

// ParseSchema parses the entire schema file and returns a list of definitions
func (p *Parser) ParseSchema() ([]Definition, error) {
	var definitions []Definition
	for {
		var comments []string
		for p.lex.TryNextComment() {
			comments = append(comments, p.lex.Text)
		}

		if p.IsNextEOF() {
			break
		}

		def, err := p.ParseDefinition()
		if err != nil {
			return nil, err
		}
		if def == nil {
			return nil, errors.New("expected definition")
		}

		def.Comments = comments
		definitions = append(definitions, *def)

	}
	return definitions, nil
}
