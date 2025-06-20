package cmd

import (
	"github.com/caffeines/notion-todo/cmd/processors"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo item",
	Long:  `Add a todo item to the Notion database with an optional due date.`,
	Run:   processors.Add,
	Args:  cobra.MinimumNArgs(1),
	Example: `todo add "Buy groceries" --date 15-03-25
todo add "Finish project report"`,
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("date", "d", "", "Due date for the todo item (optional, format: DD-MM-YYYY)")
}
