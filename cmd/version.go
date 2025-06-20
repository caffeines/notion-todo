package cmd

import (
	"github.com/caffeines/notion-todo/cmd/processors"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version of the application",
	Long: `Print the current version of the Notion Todo application.
This command displays version information and build details.`,
	Run:               processors.Version,
	Args:              cobra.NoArgs,
	Example:           `todo version, todo v`,
	ValidArgsFunction: cobra.NoFileCompletions,
	Annotations: map[string]string{
		"help:args": "No arguments required",
	},
	DisableFlagsInUseLine: true,
	DisableAutoGenTag:     true,
	DisableSuggestions:    true,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
