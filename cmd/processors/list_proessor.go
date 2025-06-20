package processors

import (
	"fmt"
	"os"
	"strings"
	"time"

	tpl "github.com/caffeines/notion-todo/cmd/template"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/caffeines/notion-todo/consts"
	"github.com/caffeines/notion-todo/service/config"
	"github.com/caffeines/notion-todo/service/files"
	"github.com/caffeines/notion-todo/service/notion"
	"github.com/charmbracelet/lipgloss"
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

// Delete message for async delete operations
type deleteMsg struct {
	success bool
	todoID  string
	message string
}

// Refresh message for async refresh operations
type refreshMsg struct {
	success bool
	todos   []Todo
	message string
}

// Bubble Tea model
type model struct {
	todos              []Todo
	cursor             int
	statusList         []string
	updating           bool
	refreshing         bool
	message            string
	messageTime        time.Time
	showConfirmation   bool
	showDeleteConfirm  bool
	pendingTodoID      string
	pendingNewStatus   string
	pendingOldStatus   string
	pendingDeleteTitle string
	statusFilter       string
	width              int
	height             int
	errorMsg           string
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

func initialModel(statusFilter string) model {
	return model{
		todos:              []Todo{}, // Start with empty todos
		cursor:             0,
		statusList:         []string{"Pending", "In Progress", "Done"}, // Use actual status values that match the dummy data
		updating:           false,
		refreshing:         true, // Set to true to show loading state
		message:            "Loading todos...",
		messageTime:        time.Now(),
		showConfirmation:   false,
		showDeleteConfirm:  false,
		pendingTodoID:      "",
		pendingNewStatus:   "",
		pendingOldStatus:   "",
		pendingDeleteTitle: "",
		statusFilter:       statusFilter,
		width:              80, // Default width
		height:             24, // Default height
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

		if m.showDeleteConfirm {
			switch msg.String() {
			case "y", "Y", "enter":
				// Confirm the delete
				m.showDeleteConfirm = false
				m.updating = true
				// Call API to delete todo
				return m, deleteTodoCmd(m.pendingTodoID)
			case "n", "N", "esc":
				// Cancel the delete
				m.showDeleteConfirm = false
				m.pendingTodoID = ""
				m.pendingDeleteTitle = ""
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
		case "d", "D":
			if len(m.todos) > 0 && !m.updating {
				// Show delete confirmation
				todo := &m.todos[m.cursor]
				m.showDeleteConfirm = true
				m.pendingTodoID = todo.ID
				m.pendingDeleteTitle = todo.Title
			}

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

	case deleteMsg:
		m.updating = false
		if msg.success {
			// Delete was successful - remove todo from local list and show success message
			for i := range m.todos {
				if m.todos[i].ID == msg.todoID {
					// Remove the todo from the slice
					m.todos = append(m.todos[:i], m.todos[i+1:]...)
					// Adjust cursor if necessary
					if m.cursor >= len(m.todos) && len(m.todos) > 0 {
						m.cursor = len(m.todos) - 1
					}
					if len(m.todos) == 0 {
						m.cursor = 0
					}
					break
				}
			}
			m.message = msg.message
		} else {
			// Delete failed - show error message
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
			} else if m.showConfirmation || m.showDeleteConfirm {
				emptyMessage = "Confirm your action"
			} else {
				// Default empty message when no todos and not refreshing
				emptyMessage = "No todos available"
			}
		}

		emptyContent := titleStyle.Render("Todo") + "\n\n" +
			tpl.EmptyStateStyle.Render(emptyMessage) + "\n\n" +
			tpl.HelpStyle.Render("Press 'q' to quit")

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
		confirmation := tpl.ConfirmationStyle.Render(confirmationText) + "\n\n" +
			tpl.HelpStyle.Render("y: confirm • n: cancel")

		return confirmationContainerStyle.Render(confirmation)
	}

	// Show delete confirmation dialog if needed
	if m.showDeleteConfirm {
		confirmationContainerStyle := getConfirmationContainerStyle(m.width)

		confirmationText := fmt.Sprintf("Delete '%s'?", m.pendingDeleteTitle)
		confirmation := tpl.ConfirmationStyle.Render(confirmationText) + "\n\n" +
			tpl.HelpStyle.Render("y: confirm • n: cancel")

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
		style := tpl.ItemStyle

		if m.cursor == i {
			cursor = ">"
			style = tpl.SelectedItemStyle
		}

		// Get status style - minimal colors only
		statusStyle, exists := tpl.StatusStyles[todo.Status]
		if !exists {
			statusStyle = tpl.ItemStyle
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
		statusMsg = "\n" + tpl.UpdatingStyle.Render("Updating...")
	} else if m.refreshing {
		statusMsg = "\n" + tpl.UpdatingStyle.Render("Refreshing...")
	} else if m.message != "" && time.Since(m.messageTime) < 3*time.Second {
		statusMsg = "\n" + tpl.MessageStyle.Render(m.message)
	}

	// Minimal help text
	helpText := "↑↓: navigate • ←→: status • d: delete • r: refresh • q: quit"
	if m.width < 60 {
		helpText = "↑↓←→ d r q"
	}
	help := tpl.HelpStyle.Render(helpText)

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

func List(cmd *cobra.Command, args []string) {
	status, _ := cmd.Flags().GetString("status")

	// Validate status filter
	if status != "" && !consts.IsValidStatus(status) {
		fmt.Printf("Invalid status filter: '%s'. Valid statuses are: %s\n", status, consts.GetAllStatuses())
		os.Exit(1)
	}

	status = cases.Title(language.English).String(status) // Normalize status to title case

	p := tea.NewProgram(
		initialModel(status),
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running Bubble Tea program:", err)
		os.Exit(1)
	}
}

// Delete todo using Notion API
func deleteTodoCmd(todoID string) tea.Cmd {
	return func() tea.Msg {
		// Initialize services
		credService := config.NewCredentialSvc(files.NewFileService(consts.ConfigFileName))
		notionSvc := notion.NewNotionImpl(credService)

		// Call Notion API to delete todo
		err := notionSvc.DeletePage(todoID)
		if err != nil {
			return deleteMsg{
				success: false,
				todoID:  todoID,
				message: fmt.Sprintf("Failed to delete: %v", err),
			}
		}

		return deleteMsg{
			success: true,
			todoID:  todoID,
			message: "Todo deleted successfully",
		}
	}
}
