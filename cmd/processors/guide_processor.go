package processors

import (
	"fmt"
	"os"

	tpl "github.com/caffeines/notion-todo/cmd/template"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Guide(cmd *cobra.Command, args []string) {
	p := tea.NewProgram(tpl.InitialGuideModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running guide: %v", err)
		os.Exit(1)
	}
}
