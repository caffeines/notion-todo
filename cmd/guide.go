package cmd

import (
	"github.com/caffeines/notion-todo/cmd/processors"
	"github.com/spf13/cobra"
)

// guideCmd represents the guide command
var guideCmd = &cobra.Command{
	Use:     "guide",
	Aliases: []string{"g"},
	Short:   "Interactive guide to setup Notion database and integration token",
	Long:    `An interactive, step-by-step guide to help you set up your Notion integration and database for the todo CLI.`,
	Run:     processors.Guide,
}

func init() {
	rootCmd.AddCommand(guideCmd)
}
