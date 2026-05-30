package ui

import "github.com/charmbracelet/lipgloss"

var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15"))
	subStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	highStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("1"))
	medStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("3"))
	lowStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	fileStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	flameStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15"))
)

func GetRisk(flame int) (string, lipgloss.Style) {
	switch {
	case flame >= 10:
		return "HIGH", highStyle
	case flame >= 4:
		return "MED ", medStyle
	default:
		return "LOW ", lowStyle
	}
}
