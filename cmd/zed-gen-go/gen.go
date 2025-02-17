package main

import (
	"zedgen/internal/generate"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "gen <schema_file>",
	Short: "Generates Go code from a Zed schema file",
	Args:  cobra.MinimumNArgs(1),
	RunE:  generateCmdFunc,
}

func registerGenerateCmd(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("output", "o", "out", "Output directory for generated code")
	cmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")

	cmd.AddCommand(generateCmd)
}

func generateCmdFunc(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return cmd.Help()
	}

	if len(args) != 1 {
		return cmd.Help()
	}

	schemaFile := args[0]
	if schemaFile == "" {
		return cmd.Help()
	}
	outputDir, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return err
	}

	input := generate.GenerateInput{
		SchemaFile: schemaFile,
		OutputDir:  outputDir,
		Verbose:    verbose,
	}
	return generate.Run(cmd.Context(), &input)
}
