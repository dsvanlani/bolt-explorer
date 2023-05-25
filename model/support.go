package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

// Support is a model displayed for support content
type Support struct {
	name   string
	styles *style.Styles
}

// NewSupport creates a new support model
func NewSupport(styles *style.Styles) *Support {
	return &Support{
		styles: styles,
	}
}

// Init is a noop
func (s *Support) Init() tea.Cmd {
	return nil
}

// Update is a noop
func (s *Support) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

// View renders the home model
func (s *Support) View() string {
	return "support"
}
