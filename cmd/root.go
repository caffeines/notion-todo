package cmd

import (
	"os"

	"github.com/caffeines/notion-todo/cmd/processors"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "CLI for Todo with Notion database",
	Long:  `A modern command line interface for managing todos with Notion database integration.`,
	Run:   processors.Root,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags and configuration can be added here
}
