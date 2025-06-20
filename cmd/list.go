/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/caffeines/notion-todo/consts"
	"github.com/caffeines/notion-todo/service/config"
	"github.com/caffeines/notion-todo/service/files"
	"github.com/caffeines/notion-todo/service/notion"
	"github.com/spf13/cobra"
)

// Todo represents a todo item.
type Todo struct {
	ID      string
	Title   string
	Status  string
	DueDate *string
}

// Fetch todos command for async operations
func fetchTodosCmd(status string) tea.Cmd {
	return func() tea.Msg {
		// Try to fetch from Notion first
		credService := config.NewCredentialSvc(files.NewFileService(consts.ConfigFileName))
		notionSvc := notion.NewNotionImpl(credService)

		todos, err := notionSvc.QueryPages(status, "")
		if err != nil {
			return refreshMsg{
				success: false,
				todos:   nil,
				message: "Failed to load todos: " + err.Error(),
			}
		}

		// Convert Notion todos to local Todo format
		var result []Todo
		for _, todo := range todos {
			// Map Notion status to local status names
			localStatus := todo.Status
			switch todo.Status {
			case "Todo":
				localStatus = "Pending"
			case "In Progress":
				localStatus = "In Progress"
			case "Done":
				localStatus = "Done"
			}

			result = append(result, Todo{
				ID:      todo.ID,
				Title:   todo.Title,
				Status:  localStatus,
				DueDate: todo.DueDate,
			})
		}

		return refreshMsg{
			success: true,
			todos:   result,
			message: fmt.Sprintf("Loaded %d todos", len(result)),
		}
	}
}

// Status update message for async operations
type statusUpdateMsg struct {
	success   bool
	todoID    string
	newStatus string
	message   string
}

// Refresh message for async refresh operations
type refreshMsg struct {
	success bool
	todos   []Todo
	message string
}

// Bubble Tea model
type model struct {
	todos            []Todo
	cursor           int
	statusList       []string
	updating         bool
	refreshing       bool
	message          string
	messageTime      time.Time
	showConfirmation bool
	pendingTodoID    string
	pendingNewStatus string
	pendingOldStatus string
	statusFilter     string
	width            int
	height           int
	errorMsg         string
}

// Minimalistic and clean styling - responsive
func getTitleStyle(width int) lipgloss.Style {
	titleWidth := width - 4 // Account for container padding
	if titleWidth < 20 {
		titleWidth = 20
	}
	if titleWidth > 80 {
		titleWidth = 80
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{
			Light: "#1f2937", // Gray-800
			Dark:  "#f9fafb", // Gray-50
		}).
		Padding(0, 1).
		Bold(true).
		Width(titleWidth).
		Align(lipgloss.Center)
}

func getContainerStyle(width, height int) lipgloss.Style {
	containerWidth := width - 2 // Minimal margin
	if containerWidth < 30 {
		containerWidth = 30
	}

	return lipgloss.NewStyle().
		Padding(1, 2).
		Width(containerWidth).
		MaxHeight(height - 2)
}

func getConfirmationContainerStyle(width int) lipgloss.Style {
	confirmWidth := width / 2
	if confirmWidth < 30 {
		confirmWidth = 30
	}
	if confirmWidth > 60 {
		confirmWidth = 60
	}

	return lipgloss.NewStyle().
		Padding(1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.AdaptiveColor{
			Light: "#d1d5db", // Gray-300
			Dark:  "#4b5563", // Gray-600
		}).
		Width(confirmWidth).
		Align(lipgloss.Center)
}

var (
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{
			Light: "#1f2937", // Gray-800
			Dark:  "#f9fafb", // Gray-50
		}).
		Background(lipgloss.AdaptiveColor{
			Light: "#f3f4f6", // Gray-100
			Dark:  "#374151", // Gray-700
		}).
		Padding(0, 1)

	itemStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
			Light: "#374151", // Gray-700
			Dark:  "#d1d5db", // Gray-300
		})

	statusStyles = map[string]lipgloss.Style{
		"Pending": lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
				Light: "#92400e", // Amber-800
				Dark:  "#fbbf24", // Amber-400
			}),
		"In Progress": lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
				Light: "#1e40af", // Blue-800
				Dark:  "#60a5fa", // Blue-400
			}),
		"Done": lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
				Light: "#065f46", // Emerald-800
				Dark:  "#34d399", // Emerald-400
			}),
	}

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#6b7280", // Gray-500
			Dark:  "#9ca3af", // Gray-400
		}).
		Align(lipgloss.Center)

	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#065f46", // Emerald-800
			Dark:  "#34d399", // Emerald-400
		})

	updatingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#92400e", // Amber-800
			Dark:  "#fbbf24", // Amber-400
		})

	confirmationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{
			Light: "#dc2626", // Red-600
			Dark:  "#f87171", // Red-400
		}).
		Padding(0, 1)

	dueDateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#6b7280", // Gray-500
			Dark:  "#9ca3af", // Gray-400
		})

	emptyStateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#6b7280", // Gray-500
			Dark:  "#9ca3af", // Gray-400
		}).
		Align(lipgloss.Center)
)

func initialModel(statusFilter string) model {
	return model{
		todos:            []Todo{}, // Start with empty todos
		cursor:           0,
		statusList:       []string{"Pending", "In Progress", "Done"}, // Use actual status values that match the dummy data
		updating:         false,
		refreshing:       true, // Set to true to show loading state
		message:          "Loading todos...",
		messageTime:      time.Now(),
		showConfirmation: false,
		pendingTodoID:    "",
		pendingNewStatus: "",
		pendingOldStatus: "",
		statusFilter:     statusFilter,
		width:            80, // Default width
		height:           24, // Default height
	}
}

func (m model) Init() tea.Cmd {
	return fetchTodosCmd(m.statusFilter)
}

// Update status using Notion API
func updateStatusCmd(todoID, newStatus string) tea.Cmd {
	return func() tea.Msg {
		// Initialize services
		credService := config.NewCredentialSvc(files.NewFileService(consts.ConfigFileName))
		notionSvc := notion.NewNotionImpl(credService)

		// Map local status to Notion status names
		notionStatus := newStatus
		switch newStatus {
		case "Pending":
			notionStatus = "Todo"
		case "In Progress":
			notionStatus = "In Progress"
		case "Done":
			notionStatus = "Done"
		}

		// Call Notion API to update status
		err := notionSvc.UpdatePageStatus(todoID, notionStatus)
		if err != nil {
			return statusUpdateMsg{
				success:   false,
				todoID:    todoID,
				newStatus: newStatus,
				message:   fmt.Sprintf("Failed to update: %v", err),
			}
		}

		return statusUpdateMsg{
			success:   true,
			todoID:    todoID,
			newStatus: newStatus,
			message:   fmt.Sprintf("Updated to %s", newStatus),
		}
	}
}

// Simulate async refresh
func refreshTodosCmd(statusFilter string) tea.Cmd {
	return func() tea.Msg {
		// Fetch fresh data from Notion
		credService := config.NewCredentialSvc(files.NewFileService(consts.ConfigFileName))
		notionSvc := notion.NewNotionImpl(credService)

		todos, err := notionSvc.QueryPages(statusFilter, "")
		if err != nil {
			return refreshMsg{
				success: false,
				todos:   nil,
				message: "Failed to refresh: " + err.Error(),
			}
		}

		// Convert Notion todos to local Todo format
		var result []Todo
		for _, todo := range todos {
			// Map Notion status to local status names
			localStatus := todo.Status
			switch todo.Status {
			case "Todo":
				localStatus = "Pending"
			case "In Progress":
				localStatus = "In Progress"
			case "Done":
				localStatus = "Done"
			}

			result = append(result, Todo{
				ID:      todo.ID,
				Title:   todo.Title,
				Status:  localStatus,
				DueDate: todo.DueDate,
			})
		}

		return refreshMsg{
			success: true,
			todos:   result,
			message: fmt.Sprintf("Refreshed %d todos", len(result)),
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		if m.showConfirmation {
			switch msg.String() {
			case "y", "Y", "enter":
				// Confirm the update
				m.showConfirmation = false
				m.updating = true
				// Call API to update status (local update happens in statusUpdateMsg handler)
				return m, updateStatusCmd(m.pendingTodoID, m.pendingNewStatus)
			case "n", "N", "esc":
				// Cancel the update
				m.showConfirmation = false
				m.pendingTodoID = ""
				m.pendingNewStatus = ""
				m.pendingOldStatus = ""
			}
			return m, nil
		}

		// Prevent actions during update or refresh
		if m.updating || m.refreshing {
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			}
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.todos)-1 {
				m.cursor++
			}
		case "right", "l", "enter", " ":
			if len(m.todos) > 0 && !m.updating {
				// Cycle status forward
				todo := &m.todos[m.cursor]
				oldStatus := todo.Status
				newStatus := ""
				for i, s := range m.statusList {
					if s == todo.Status {
						newStatus = m.statusList[(i+1)%len(m.statusList)]
						break
					}
				}

				if newStatus != "" && newStatus != oldStatus {
					m.showConfirmation = true
					m.pendingTodoID = todo.ID
					m.pendingNewStatus = newStatus
					m.pendingOldStatus = oldStatus
				}
			}
		case "left", "h":
			if len(m.todos) > 0 && !m.updating {
				// Cycle status backward
				todo := &m.todos[m.cursor]
				oldStatus := todo.Status
				newStatus := ""
				for i, s := range m.statusList {
					if s == todo.Status {
						newStatus = m.statusList[(i-1+len(m.statusList))%len(m.statusList)]
						break
					}
				}

				if newStatus != "" && newStatus != oldStatus {
					m.showConfirmation = true
					m.pendingTodoID = todo.ID
					m.pendingNewStatus = newStatus
					m.pendingOldStatus = oldStatus
				}
			}
		case "r":
			m.refreshing = true
			m.message = "Syncing..."
			m.messageTime = time.Now()
			return m, refreshTodosCmd(m.statusFilter)

		}

	case statusUpdateMsg:
		m.updating = false
		if msg.success {
			// Update was successful - update local todo status and show success message
			for i := range m.todos {
				if m.todos[i].ID == msg.todoID {
					m.todos[i].Status = msg.newStatus
					break
				}
			}
			m.message = msg.message
		} else {
			// Update failed - show error message
			m.message = msg.message
		}
		m.messageTime = time.Now()

	case refreshMsg:
		m.refreshing = false
		if msg.success {
			m.todos = msg.todos
			// Reset cursor if it's out of bounds
			if m.cursor >= len(m.todos) {
				m.cursor = len(m.todos) - 1
			}
			if m.cursor < 0 {
				m.cursor = 0
			}
			m.message = msg.message
		} else {
			m.errorMsg = msg.message // Clear todos on refresh failure
			m.message = msg.message
		}
		m.messageTime = time.Now()
	}
	return m, nil
}

func (m model) View() string {
	containerStyle := getContainerStyle(m.width, m.height)
	titleStyle := getTitleStyle(m.width)

	if len(m.todos) == 0 {
		var emptyMessage string
		if m.refreshing {
			emptyMessage = "Loading todos..."
		} else {
			if m.errorMsg != "" {
				emptyMessage = m.errorMsg
			} else if m.showConfirmation {
				emptyMessage = "Confirm your action"
			} else {
				// Default empty message when no todos and not refreshing
				emptyMessage = "No todos available"
			}
		}

		emptyContent := titleStyle.Render("Todo") + "\n\n" +
			emptyStateStyle.Render(emptyMessage) + "\n\n" +
			helpStyle.Render("Press 'q' to quit")

		return containerStyle.Render(emptyContent)
	}

	// Show confirmation dialog if needed
	if m.showConfirmation {
		confirmationContainerStyle := getConfirmationContainerStyle(m.width)
		todo := ""
		for _, t := range m.todos {
			if t.ID == m.pendingTodoID {
				todo = t.Title
				break
			}
		}

		confirmationText := fmt.Sprintf("Change '%s' to '%s'?", todo, m.pendingNewStatus)
		confirmation := confirmationStyle.Render(confirmationText) + "\n\n" +
			helpStyle.Render("y: confirm • n: cancel")

		return confirmationContainerStyle.Render(confirmation)
	}

	// Simple header
	header := titleStyle.Render("Todo")

	// Todo list with minimal styling
	var todoItems []string
	maxTitleWidth := m.width - 25 // Reserve space for status and date
	if maxTitleWidth < 10 {
		maxTitleWidth = 10
	}

	for i, todo := range m.todos {
		cursor := " "
		style := itemStyle

		if m.cursor == i {
			cursor = ">"
			style = selectedItemStyle
		}

		// Get status style - minimal colors only
		statusStyle, exists := statusStyles[todo.Status]
		if !exists {
			statusStyle = itemStyle
		}

		// Status prefix indicators
		statusPrefix := ""
		switch todo.Status {
		case "Pending":
			statusPrefix = "[ ]"
		case "In Progress":
			statusPrefix = "[~]"
		case "Done":
			statusPrefix = "[✓]"
		}

		// Format the todo item with status prefix
		statusBadge := statusStyle.Render(statusPrefix)
		truncatedTitle := truncateText(todo.Title, maxTitleWidth)

		// Simple due date
		dueDateText := formatDueDate(todo.DueDate, m.width)

		// Create clean todo line with status prefix
		todoLine := fmt.Sprintf("%s %s %s%s", cursor, statusBadge, truncatedTitle, dueDateText)
		todoLine = style.Render(todoLine)

		todoItems = append(todoItems, todoLine)
	}

	todoList := strings.Join(todoItems, "\n")

	// Simple status message
	statusMsg := ""
	if m.updating {
		statusMsg = "\n" + updatingStyle.Render("Updating...")
	} else if m.refreshing {
		statusMsg = "\n" + updatingStyle.Render("Refreshing...")
	} else if m.message != "" && time.Since(m.messageTime) < 3*time.Second {
		statusMsg = "\n" + messageStyle.Render(m.message)
	}

	// Minimal help text
	helpText := "↑↓: navigate • ←→: status • r: refresh • q: quit"
	if m.width < 45 {
		helpText = "↑↓←→ r q"
	}
	help := helpStyle.Render(helpText)

	// Simple layout
	content := header + "\n\n" + todoList + statusMsg + "\n\n" + help

	return containerStyle.Render(content)
}

// Simple date formatting for minimalistic design
func formatDueDate(dateStr *string, screenWidth int) string {
	if dateStr == nil || *dateStr == "" {
		return ""
	}

	// Parse the ISO date (YYYY-MM-DD) from Notion
	parsedDate, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return ""
	}

	// Check if date is overdue, today, or upcoming
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	dueDate := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, parsedDate.Location())

	var dateColor string
	if dueDate.Before(today) {
		dateColor = "#dc2626" // Red - overdue
	} else if dueDate.Equal(today) {
		dateColor = "#d97706" // Orange - due today
	} else {
		dateColor = "#6b7280" // Gray - future
	}

	dateStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(dateColor))

	// Simple format based on screen width
	if screenWidth < 60 {
		return " " + dateStyle.Render(parsedDate.Format("1/2"))
	} else {
		if parsedDate.Year() == now.Year() {
			return " " + dateStyle.Render(parsedDate.Format("Jan 2"))
		} else {
			return " " + dateStyle.Render(parsedDate.Format("Jan 2, 06"))
		}
	}
}

// Helper function to truncate text for responsive design
func truncateText(text string, maxWidth int) string {
	if len(text) <= maxWidth {
		return text
	}
	if maxWidth <= 3 {
		return "..."
	}
	return text[:maxWidth-3] + "..."
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "ls"},
	Short:   "List all items in the Notion Todo database",
	Long: `list retrieves and displays all items from the Notion Todo database.
This command is useful for viewing all tasks, their statuses, and due dates in a structured format.`,
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetString("status")
		p := tea.NewProgram(
			initialModel(status),
			tea.WithAltScreen(),       // Use alternate screen buffer
			tea.WithMouseCellMotion(), // Enable mouse support
		)
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running Bubble Tea program:", err)
			os.Exit(1)
		}
	},
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
