package template

import "github.com/charmbracelet/lipgloss"

// Shared minimalistic styling for consistent UI across all commands
var (
	// Title styling - minimal and clean
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#1f2937", // Gray-800
			Dark:  "#f9fafb", // Gray-50
		}).
		Padding(0, 1).
		Bold(true).
		Align(lipgloss.Center)

	// Success message styling - minimal
	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#065f46", // Emerald-800
			Dark:  "#34d399", // Emerald-400
		})

	// Error message styling - minimal
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#dc2626", // Red-600
			Dark:  "#f87171", // Red-400
		})

	// Warning message styling - minimal
	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#92400e", // Amber-800
			Dark:  "#fbbf24", // Amber-400
		})

	// Info message styling - minimal
	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#1e40af", // Blue-800
			Dark:  "#60a5fa", // Blue-400
		})

	// Help text styling - minimal
	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#6b7280", // Gray-500
			Dark:  "#9ca3af", // Gray-400
		}).
		Align(lipgloss.Center)

	// Container styling - minimal, no borders
	ContainerStyle = lipgloss.NewStyle().
			Padding(1, 2)

	// Input styling - minimal with subtle border
	InputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#374151", // Gray-700
			Dark:  "#d1d5db", // Gray-300
		}).
		Padding(0, 1).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.AdaptiveColor{
			Light: "#d1d5db", // Gray-300
			Dark:  "#4b5563", // Gray-600
		})

	// Subtitle styling - minimal
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#374151", // Gray-700
			Dark:  "#d1d5db", // Gray-300
		})

	// Accent styling - minimal
	AccentStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#1f2937", // Gray-800
			Dark:  "#f9fafb", // Gray-50
		})

	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{
			Light: "#1f2937", // Gray-800
			Dark:  "#f9fafb", // Gray-50
		}).
		Background(lipgloss.AdaptiveColor{
			Light: "#f3f4f6", // Gray-100
			Dark:  "#374151", // Gray-700
		}).
		Padding(0, 1)

	ItemStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(lipgloss.AdaptiveColor{
			Light: "#374151", // Gray-700
			Dark:  "#d1d5db", // Gray-300
		})

	StatusStyles = map[string]lipgloss.Style{
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

	MessageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#065f46", // Emerald-800
			Dark:  "#34d399", // Emerald-400
		})

	UpdatingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#92400e", // Amber-800
			Dark:  "#fbbf24", // Amber-400
		})

	ConfirmationStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{
			Light: "#dc2626", // Red-600
			Dark:  "#f87171", // Red-400
		}).
		Padding(0, 1)

	DueDateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#6b7280", // Gray-500
			Dark:  "#9ca3af", // Gray-400
		})

	EmptyStateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{
			Light: "#6b7280", // Gray-500
			Dark:  "#9ca3af", // Gray-400
		}).
		Align(lipgloss.Center)
)

// Helper functions for common UI patterns - simplified
func RenderTitle(text string, width int) string {
	titleWidth := width - 4
	if titleWidth < 20 {
		titleWidth = 20
	}
	if titleWidth > 80 {
		titleWidth = 80
	}

	return TitleStyle.Copy().Width(titleWidth).Render(text)
}

func RenderSuccess(text string) string {
	return SuccessStyle.Render(text)
}

func RenderError(text string) string {
	return ErrorStyle.Render(text)
}

func RenderWarning(text string) string {
	return WarningStyle.Render(text)
}

func RenderInfo(text string) string {
	return InfoStyle.Render(text)
}

func RenderContainer(content string, width, height int) string {
	containerWidth := width - 2
	if containerWidth < 30 {
		containerWidth = 30
	}

	return ContainerStyle.Copy().
		Width(containerWidth).
		MaxHeight(height - 2).
		Render(content)
}

func RenderHelp(text string) string {
	return HelpStyle.Render(text)
}
