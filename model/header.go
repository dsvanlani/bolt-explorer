package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

// Header is a model displayed at the head of the application
type Header struct {
	name   string
	styles *style.Styles
}

// NewHeader creates a new header model
func NewHeader(name string, styles *style.Styles) *Header {
	return &Header{
		name:   name,
		styles: styles,
	}
}

// Init is a noop
func (h *Header) Init() tea.Cmd {
	return nil
}

// Update is a noop
func (h *Header) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return h, nil
}

// View renders the header model
func (h *Header) View() string {
	return h.styles.Header.Render(h.name)
}
