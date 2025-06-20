package cmd

import (
	"errors"
	"fmt"

	"github.com/caffeines/notion-todo/consts"
	"github.com/caffeines/notion-todo/service/config"
	"github.com/caffeines/notion-todo/service/files"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"c"},
	Short:   "Configure the app",
	Long:    `Configure the app by setting the token and database id.`,
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

		file := files.NewFileService(consts.ConfigFileName)
		credService := config.NewCredentialSvc(file)
		err = credService.SetConfig(token, databaseId)
		if err != nil {
			fmt.Println("Error setting config: " + err.Error())
			return
		}

		fmt.Printf("\n✅ Application configured successfully\n")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
