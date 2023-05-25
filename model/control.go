package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

// Control is a model displayed for control content
type Control struct {
	name   string
	styles *style.Styles
}

// NewControl creates a new control model
func NewControl(styles *style.Styles) *Control {
	return &Control{
		styles: styles,
	}
}

// Init is a noop
func (c *Control) Init() tea.Cmd {
	return nil
}

// Update is a noop
func (c *Control) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

// View renders the home model
func (c *Control) View() string {
	return "control"
}
