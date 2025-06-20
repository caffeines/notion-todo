package template

import (
	"fmt"
	"strings"

	"github.com/caffeines/notion-todo/cmd/steps"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type guideModel struct {
	currentStep  steps.Guide
	width        int
	height       int
	showNextHint bool
}

func InitialGuideModel() guideModel {
	return guideModel{
		currentStep:  steps.Welcome,
		width:        80,
		height:       24,
		showNextHint: true,
	}
}

func (m guideModel) Init() tea.Cmd {
	return nil
}

func (m guideModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter", " ", "n", "right", "l":
			if m.currentStep < steps.Complete {
				m.currentStep++
				return m, nil
			}
			return m, tea.Quit

		case "p", "left", "h":
			if m.currentStep > steps.Welcome {
				m.currentStep--
				return m, nil
			}

		case "r":
			// Reset to beginning
			m.currentStep = steps.Welcome
			return m, nil
		}
	}

	return m, nil
}

func (m guideModel) View() string {
	var content string

	switch m.currentStep {
	case steps.Welcome:
		content = m.renderWelcome()
	case steps.CreateDatabase:
		content = m.renderCreateDatabase()
	case steps.GetDatabaseID:
		content = m.renderGetDatabaseID()
	case steps.CreateIntegration:
		content = m.renderCreateIntegration()
	case steps.ConnectDatabase:
		content = m.renderGetToken()
	case steps.GetToken:
		content = m.renderConnectDatabase()
	case steps.TestConnection:
		content = m.renderTestConnection()
	case steps.Complete:
		content = m.renderComplete()
	}

	return content
}

func (m guideModel) renderWelcome() string {
	title := RenderTitle("🚀 Notion Todo Setup Guide", m.width)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		AccentStyle.Render("Welcome to the interactive setup guide!"),
		"",
		"This guide will help you:",
		"• Set up your todo database",
		"• Create a Notion integration",
		"• Configure the CLI tool",
		"",
		HelpStyle.Render("📝 You'll need a Notion account to get started"),
		"",
		m.renderNavigation(),
	)

	return RenderContainer(content, m.width, m.height)
}

func (m guideModel) renderCreateIntegration() string {
	title := RenderTitle("Step 3: Create Notion Integration", m.width)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		AccentStyle.Render("🔗 Create a new Notion integration:"),
		"",
		"1. Open your browser and go to:",
		"   "+InfoStyle.Render("https://www.notion.so/profile/integrations"),
		"",
		"2. Click "+AccentStyle.Render("\"New integration\""),
		"",
		"3. Fill in the form:",
		"   • Name: "+AccentStyle.Render("\"Todo CLI\"")+" (or any name you prefer)",
		"   • Associated workspace: Select your workspace",
		"   • Type: "+AccentStyle.Render("Internal"),
		"",
		"4. Click "+AccentStyle.Render("\"Save\""),
		"",
		WarningStyle.Render("💡 Keep this browser tab open - you'll need it in the next step!"),
		"",
		m.renderNavigation(),
	)

	return RenderContainer(content, m.width, m.height)
}

func (m guideModel) renderGetToken() string {
	title := RenderTitle("Step 4: Grant Database Access", m.width)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		AccentStyle.Render("🔐 Grant access to your database:"),
		"",
		"1. After creating the integration, a popup will open to configure it",
		"",
		"2. Click on "+AccentStyle.Render("\"Access\"")+" tab in the configuration page",
		"",
		"3. Select "+AccentStyle.Render("\"pages\"")+" where you will see your database listed",
		"",
		"4. Select your todo database from the list",
		"",
		"5. Click "+AccentStyle.Render("\"Update Access\""),
		"",
		"6. Go back to the "+AccentStyle.Render("\"Configuration\"")+" tab",
		"",
		SuccessStyle.Render("🎉 Perfect! Your integration can now access the database."),
		"",
		WarningStyle.Render("⚠️  This step is crucial - without it, the CLI won't work!"),
		"",
		m.renderNavigation(),
	)

	return RenderContainer(content, m.width, m.height)
}

func (m guideModel) renderCreateDatabase() string {
	title := RenderTitle("Step 1: Create Todo Database", m.width)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		AccentStyle.Render("📋 Create your todo database:"),
		"",
		AccentStyle.Render("Option 1: Use the template (Recommended)"),
		"1. Go to: "+InfoStyle.Render("https://www.notion.so/templates/cli-todo"),
		"2. Click "+AccentStyle.Render("\"Get template\"")+" and select your workspace",
		"",
		AccentStyle.Render("Option 2: Create manually"),
		"1. Open Notion and create a "+AccentStyle.Render("\"+ New page\""),
		"2. Choose "+AccentStyle.Render("\"Database\"")+" → "+AccentStyle.Render("\"Table\""),
		"3. Name your database: "+AccentStyle.Render("\"My Todos\"")+" (or any name)",
		"4. Add these properties:",
		"   • "+AccentStyle.Render("Title")+" (Title) - for the todo text",
		"   • "+AccentStyle.Render("Status")+" (Select) - with options: Todo, In progress, Done",
		"   • "+AccentStyle.Render("Due Date")+" (Date) - for due dates (Optional)",
		"   • "+AccentStyle.Render("Tags")+" (Multi-select) - for categorizing (Optional)",
		"",
		InfoStyle.Render("💡 The template already has the correct structure set up!"),
		"",
		m.renderNavigation(),
	)

	return RenderContainer(content, m.width, m.height)
}

func (m guideModel) renderGetDatabaseID() string {
	title := RenderTitle("Step 2: Get Database ID", m.width)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		AccentStyle.Render("🔍 Find your database ID:"),
		"",
		"1. Open your database in Notion",
		"",
		"2. Look at the URL in your browser address bar. It looks like:",
		"   "+HelpStyle.Render("https://www.notion.so/{workspace}/{database_id}?v=..."),
		"",
		"3. The database ID is the string after the last "+AccentStyle.Render("\"/\"")+" and before "+AccentStyle.Render("\"?\""),
		"",
		"4. Example URL breakdown:",
		"   "+HelpStyle.Render("https://www.notion.so/sadatos/217e31436430803999d6ecaabdf4e11f?v=217e3135963081c2b425000c383f3ac2"),
		"   Database ID: "+InputStyle.Render("217e31436430803999d6ecaabdf4e11f"),
		"",
		"5. Copy this ID - you'll need it for configuration",
		"",
		WarningStyle.Render("📝 The database ID is usually 32 characters long"),
		"",
		m.renderNavigation(),
	)

	return RenderContainer(content, m.width, m.height)
}

func (m guideModel) renderConnectDatabase() string {
	title := RenderTitle("Step 5: Copy Integration Token", m.width)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		AccentStyle.Render("🔑 Get your integration token:"),
		"",
		"1. Make sure you're still in the "+AccentStyle.Render("\"Configuration\"")+" tab",
		"",
		"2. Click "+AccentStyle.Render("\"Show\"")+" next to the \"Internal Integration Token\"",
		"",
		"3. Copy the token that appears",
		"",
		"4. Store it safely - it looks like:",
		"   "+InputStyle.Render("ntn_abcd1234..."),
		"",
		WarningStyle.Render("🔒 Keep this token secret and don't share it publicly!"),
		"",
		InfoStyle.Render("✅ Token copied? Great! You're almost done."),
		"",
		m.renderNavigation(),
	)

	return RenderContainer(content, m.width, m.height)
}

func (m guideModel) renderTestConnection() string {
	title := RenderTitle("Step 6: Configure CLI", m.width)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		AccentStyle.Render("⚙️  Now let's configure the CLI:"),
		"",
		"1. Open your terminal and run:",
		"   "+InputStyle.Render("todo config"),
		"",
		"2. When prompted, enter your:",
		"   • "+AccentStyle.Render("Integration Token")+" (from Step 2)",
		"   • "+AccentStyle.Render("Database ID")+" (from Step 4)",
		"",
		"3. Test your setup by adding a todo:",
		"   "+InputStyle.Render("todo add \"Test todo from CLI\""),
		"",
		"4. Check your Notion database - you should see the new todo!",
		"",
		InfoStyle.Render("🔧 If something doesn't work, go back and check the previous steps."),
		"",
		m.renderNavigation(),
	)

	return RenderContainer(content, m.width, m.height)
}

func (m guideModel) renderComplete() string {
	title := RenderTitle("🎉 Setup Complete!", m.width)

	content := lipgloss.JoinVertical(lipgloss.Left,
		title,
		"",
		SuccessStyle.Render("Congratulations! Your Notion Todo CLI is ready to use."),
		"",
		AccentStyle.Render("Quick commands to get started:"),
		"",
		"• "+InputStyle.Render("todo add \"Buy groceries\"")+" - Add a new todo",
		"• "+InputStyle.Render("todo add \"Meeting\" --date 25-06-2025")+" - Add with date",
		"• "+InputStyle.Render("todo list")+" - View and manage your todos",
		"• "+InputStyle.Render("todo --help")+" - See all available commands",
		"",
		HelpStyle.Render("Need help? Check the README.md or run 'todo guide' again anytime."),
		"",
		AccentStyle.Render("Press any key to exit..."),
	)

	return RenderContainer(content, m.width, m.height)
}

func (m guideModel) renderNavigation() string {
	var nav []string

	if m.currentStep > steps.Welcome {
		nav = append(nav, HelpStyle.Render("p/h/←")+" - Previous")
	}

	if m.currentStep < steps.Complete {
		nav = append(nav, HelpStyle.Render("Enter/Space/l/→")+" - Next")
	} else {
		nav = append(nav, HelpStyle.Render("Enter/Space")+" - Exit")
	}

	nav = append(nav, HelpStyle.Render("q")+" - Quit")

	if m.currentStep > steps.Welcome {
		nav = append(nav, HelpStyle.Render("r")+" - Restart")
	}

	progress := fmt.Sprintf("Step %d of %d", int(m.currentStep)+1, int(steps.Complete)+1)

	return lipgloss.JoinVertical(lipgloss.Center,
		"",
		InfoStyle.Render(progress),
		HelpStyle.Render(strings.Join(nav, " • ")),
	)
}
