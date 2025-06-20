package processors

import (
	"fmt"

	tpl "github.com/caffeines/notion-todo/cmd/template"
	"github.com/caffeines/notion-todo/consts"
	"github.com/spf13/cobra"
)

func Version(cmd *cobra.Command, args []string) {
	width := 60
	height := 10

	content := tpl.RenderTitle("Version Information", width) + "\n\n" +
		tpl.AccentStyle.Render("📦 Notion Todo CLI") + "\n" +
		tpl.SubtitleStyle.Render("Version: ") + consts.GetVersion() + "\n\n" +
		tpl.HelpStyle.Render("A modern CLI tool for managing todos with Notion database integration")

	fmt.Println(tpl.RenderContainer(content, width, height))
}
