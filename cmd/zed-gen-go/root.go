package main

import (
	"context"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

func InitializeRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "zgen",
		Short: "Code generator for Authzed client code",
		Long:  "Zedgen is a code generator for Authzed client code. Defining strict types and structures based on your Authzed Schema.",
	}
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Println(err)
		cmd.Println(cmd.UsageString())
		return errors.New("flag error")
	})

	// register generate command
	registerGenerateCmd(rootCmd)
	return rootCmd
}

func main() {
	rootCmd := InitializeRootCmd()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
