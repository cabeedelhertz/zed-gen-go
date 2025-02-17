package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"zedgen/internal/lexer"
	"zedgen/internal/parser"
)

type GenerateInput struct {
	SchemaFile string // args[0]
	OutputDir  string // -o
	Verbose    bool   // -v
}

func Run(ctx context.Context, in *GenerateInput) error {
	if in == nil {
		return fmt.Errorf("invalid input")
	}

	file, err := os.Open(in.SchemaFile)
	if err != nil {
		return err
	}
	defer file.Close()

	lex := lexer.NewLexer(file)
	parser := parser.NewParser(lex)

	// Parse the spiceDB schema
	schema, err := parser.ParseSchema()
	if err != nil {
		return err
	}
	if in.Verbose {
		gotJSON, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to marshal, err %v\n", err)
		}
		fmt.Print(string(gotJSON))
	}

	// Generate the file
	err = generateFile(in, schema)
	if err != nil {
		return err
	}

	return nil
}
