package style

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles that define the app
type Styles struct {
	ActiveBorderColor   lipgloss.Color
	InactiveBorderColor lipgloss.Color

	App lipgloss.Style

	HeaderBorder lipgloss.Border
	Header       lipgloss.Style

	PageBorder lipgloss.Border
	Page       lipgloss.Style

	Menu             lipgloss.Style
	MenuInactive     lipgloss.Style
	MenuCursor       lipgloss.Style
	MenuItem         lipgloss.Style
	MenuBorder       lipgloss.Border
	SelectedMenuItem lipgloss.Style

	Bucket lipgloss.Style
	Key    lipgloss.Style

	Footer      lipgloss.Style
	HelpKey     lipgloss.Style
	HelpValue   lipgloss.Style
	HelpDivider lipgloss.Style
}

// DefaultStyles returns default styles for the app
func DefaultStyles() *Styles {
	s := new(Styles)

	s.ActiveBorderColor = lipgloss.Color("62")
	s.InactiveBorderColor = lipgloss.Color("236")

	s.App = lipgloss.NewStyle().
		Margin(1, 2)

	s.HeaderBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "┬",
		BottomLeft:  "├",
		BottomRight: "┤",
	}

	s.Header = lipgloss.NewStyle().
		BorderStyle(s.HeaderBorder).
		BorderForeground(s.ActiveBorderColor).
		Foreground(lipgloss.Color("39")).
		PaddingLeft(1).
		MarginLeft(2).
		Width(55).
		Bold(true)

	s.PageBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "",
		Right:       "│",
		TopLeft:     "─",
		TopRight:    "╮",
		BottomLeft:  "─",
		BottomRight: "╯",
	}

	s.Page = lipgloss.NewStyle().
		BorderStyle(s.PageBorder).
		BorderForeground(s.ActiveBorderColor).
		Width(99).
		Height(11)

	s.MenuBorder = lipgloss.Border{
		Top:         "",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "│",
		TopRight:    "│",
		BottomLeft:  "╰",
		BottomRight: "┴",
	}

	s.Menu = lipgloss.NewStyle().
		BorderStyle(s.MenuBorder).
		BorderForeground(s.ActiveBorderColor).
		PaddingLeft(2).
		MarginLeft(2).
		Width(55).
		Height(12)

	s.MenuInactive = lipgloss.NewStyle().
		BorderStyle(s.MenuBorder).
		BorderForeground(s.ActiveBorderColor).
		Foreground(lipgloss.Color("236")).
		PaddingLeft(2).
		MarginLeft(2).
		Width(20).
		Height(12)

	s.MenuCursor = lipgloss.NewStyle().
		Foreground(lipgloss.Color("213")).
		SetString(">")

	s.MenuItem = lipgloss.NewStyle().
		PaddingLeft(2)

	s.SelectedMenuItem = lipgloss.NewStyle().
		Foreground(lipgloss.Color("207")).
		PaddingLeft(1)

	s.Bucket = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	s.Key = lipgloss.NewStyle()

	s.Footer = lipgloss.NewStyle().
		MarginLeft(3)

	s.HelpKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Bold(true)

	s.HelpValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color("239"))

	s.HelpDivider = lipgloss.NewStyle().
		Foreground(lipgloss.Color("237")).
		SetString(" • ")

	return s
}
