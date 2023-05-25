package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/router"
	"github.com/observiq/bolt-explorer/style"
)

// Header is a model displayed at the head of the application
type Header struct {
	name   string
	styles *style.Styles
	router *router.Router
}

// NewHeader creates a new header model
func NewHeader(name string, styles *style.Styles, router *router.Router) *Header {
	return &Header{
		name:   name,
		styles: styles,
		router: router,
	}
}

// Init is a noop
func (h *Header) Init() tea.Cmd {
	return nil
}

func (h *Header) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch h.router.SearchMode {
		case true:
			switch msg.String() {
			case "esc":
				h.router.SearchMode = false
			case "backspace":
				if len(h.router.SearchValue) > 0 {
					h.router.SearchValue = h.router.SearchValue[:len(h.router.SearchValue)-1]
				}
			default:
				if len(msg.String()) > 1 {
					// do nothing
				} else {
					h.router.SearchValue += msg.String()
				}
			}
		case false:
			switch msg.String() {
			case "f":
				h.router.SearchMode = true
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
	if h.router.SearchMode {
		str.WriteString("   >  ")
		str.WriteString(h.router.SearchValue)
	}

	return h.styles.Header.Render(str.String())
}
