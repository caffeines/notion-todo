/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/caffeines/notion-todo/cmd/processors"
	"github.com/caffeines/notion-todo/consts"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "List all items in the Notion Todo database",
	Long: `list retrieves and displays all items from the Notion Todo database.
This command is useful for viewing all tasks, their statuses, and due dates in a structured format.`,
	Run: processors.List,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("status", "s", "", fmt.Sprintf("Filter items by status (e.g., %s)", consts.GetAllStatuses()))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
