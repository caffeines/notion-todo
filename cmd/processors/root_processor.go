package processors

import (
	"fmt"

	tpl "github.com/caffeines/notion-todo/cmd/template"
	"github.com/spf13/cobra"
)

func Root(cmd *cobra.Command, args []string) {
	// Display welcome screen when no subcommand is provided
	width := 80
	height := 20

	content := tpl.RenderTitle("Notion Todo CLI", width) + "\n\n" +
		tpl.AccentStyle.Render("📝 Welcome to Notion Todo!") + "\n\n" +
		tpl.SubtitleStyle.Render("Available Commands:") + "\n" +
		"• " + tpl.AccentStyle.Render("guide") + "   - Interactive setup guide\n" +
		"• " + tpl.AccentStyle.Render("config") + "  - Configure Notion integration\n" +
		"• " + tpl.AccentStyle.Render("add") + "     - Add a new todo item\n" +
		"• " + tpl.AccentStyle.Render("list") + "    - View and manage todos\n" +
		"• " + tpl.AccentStyle.Render("version") + " - Show version information\n\n" +
		tpl.HelpStyle.Render("Get started:\n1. Run 'todo guide' for interactive setup\n2. Or run 'todo config' to set up manually\n3. Use 'todo add \"task\" --date DD-MM-YYYY' to create todos\n4. Use 'todo list' to view and update your todos\n\nFor detailed help: todo --help")

	fmt.Println(tpl.RenderContainer(content, width, height))
}
