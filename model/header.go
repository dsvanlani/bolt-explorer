package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

// Header is a model displayed at the head of the application
type Header struct {
	name         string
	styles       *style.Styles
	searchMode   bool
	searchString string
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

func (h *Header) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch h.searchMode {
		case true:
			switch msg.String() {
			case "backspace":
				if len(h.searchString) > 0 {
					h.searchString = h.searchString[:len(h.searchString)-1]
				}
			case "esc":
				h.searchMode = false
			default:
				if len(msg.String()) > 1 {
					// do nothing
				} else {
					h.searchString += msg.String()
				}
			}
		case false:
			switch msg.String() {

			case "f":
				h.searchMode = true
				return h, nil
			case "esc":
				h.searchMode = false
				return h, nil
			}
		}

	}
	return h, nil
}

// View renders the header model
func (h *Header) View() string {
	str := strings.Builder{}
	str.WriteString(h.name)
	if h.searchMode {
		str.WriteString("   >  ")
		str.WriteString(h.searchString)
	}

	return h.styles.Header.Render(str.String())
}
