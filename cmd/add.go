/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/caffeines/notion-todo/notion"
	"github.com/caffeines/notion-todo/service"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a todo item",
	Long:  `add a todo item to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		todoItem := strings.Join(args, " ")
		todoSvc := service.NewTodoImpl()
		credService := service.NewCredentialSvc(
			service.NewFile("config.json"),
		)
		notionSvc := notion.NewNotionImpl(todoSvc, credService)
		err := notionSvc.AddPage("✼ " + todoItem)
		if err != nil {
			fmt.Println("Todo creation failed with error: ", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
