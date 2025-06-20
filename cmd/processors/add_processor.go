package processors

import (
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	tpl "github.com/caffeines/notion-todo/cmd/template"
	"github.com/caffeines/notion-todo/consts"
	"github.com/caffeines/notion-todo/service/config"
	"github.com/caffeines/notion-todo/service/files"
	"github.com/caffeines/notion-todo/service/notion"
	"github.com/caffeines/notion-todo/service/utility"
	"github.com/spf13/cobra"
)

func Add(cmd *cobra.Command, args []string) {
	todoItem := strings.Join(args, " ")
	date, err := cmd.Flags().GetString("date")

	// Validation
	if todoItem == "" {
		fmt.Println(tpl.RenderContainer(
			tpl.RenderTitle("Add Todo", 80)+"\n\n"+
				tpl.RenderError("Please provide a todo item to add.")+"\n\n"+
				tpl.RenderHelp("Usage: todo add \"Your task description\" [--date DD-MM-YYYY]"),
			80, 24,
		))
		return
	}

	if err != nil {
		fmt.Println(tpl.RenderContainer(
			tpl.RenderTitle("Add Todo", 80)+"\n\n"+
				tpl.RenderError("Error getting date flag: "+err.Error()),
			80, 24,
		))
		return
	}

	// Check date format DD-MM-YYYY only if date is provided
	if date != "" && !utility.IsValidDateFormat(date) {
		fmt.Println(tpl.RenderContainer(
			tpl.RenderTitle("Add Todo", 80)+"\n\n"+
				tpl.RenderError("Invalid date format. Please use DD-MM-YYYY.")+"\n\n"+
				tpl.RenderHelp("Example: 25-12-2024"),
			80, 24,
		))
		return
	}

	// Create and start spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " Creating todo..."
	s.Color("cyan")
	s.Start()

	// Create todo
	credService := config.NewCredentialSvc(
		files.NewFileService(consts.ConfigFileName),
	)
	notionSvc := notion.NewNotionImpl(credService)

	// Convert date if provided, otherwise pass empty string
	dateForAPI := ""
	if date != "" {
		dateForAPI = utility.ConvertDateToYYYYMMDD(date)
	}

	err = notionSvc.AddPage(todoItem, dateForAPI)

	// Stop spinner
	s.Stop()
	if err != nil {
		fmt.Println(tpl.RenderContainer(
			tpl.RenderTitle("Add Todo", 80)+"\n\n"+
				tpl.RenderError("Todo creation failed: "+err.Error())+"\n\n"+
				tpl.RenderHelp("Check configuration: todo config"),
			80, 24,
		))
		return
	}

	// Success message with minimal styling
	successContent := tpl.RenderTitle("Add Todo", 80) + "\n\n" +
		tpl.RenderSuccess("Todo added successfully!") + "\n\n" +
		"Task: " + todoItem + "\n"

	if date != "" {
		successContent += "Date: " + date + "\n"
	} else {
		successContent += "Date: No due date\n"
	}

	successContent += "\n" + tpl.RenderHelp("Use 'todo list' to view all todos")

	fmt.Println(tpl.RenderContainer(successContent, 80, 24))
}
