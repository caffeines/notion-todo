/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/caffeines/notion-todo/service"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the app",
	Long:  `Configure the app by setting the token and database id.`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenValidate := func(input string) error {
			if len(input) == 0 {
				return errors.New("Token cannot be empty")
			}
			return nil
		}

		databaseIDValidate := func(input string) error {
			if len(input) == 0 {
				return errors.New("Database ID cannot be empty")
			}
			return nil
		}

		tokenPrompt := promptui.Prompt{
			Label:    "Token",
			Validate: tokenValidate,
			Mask:     '*',
		}

		databaseIDPrompt := promptui.Prompt{
			Label:    "Database ID",
			Validate: databaseIDValidate,
			Mask:     '*',
		}

		token, err := tokenPrompt.Run()
		if err != nil {
			return
		}
		databaseId, err := databaseIDPrompt.Run()
		if err != nil {
			return
		}

		file := service.NewFile("config.json")
		credService := service.NewCredentialSvc(file)
		err = credService.SetConfig(token, databaseId)
		if err != nil {
			fmt.Println("❌ Error setting config:", err)
			return
		}

		fmt.Printf("\n✅ Application configured successfully")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
