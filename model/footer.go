package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

// KeyBinding is the binding of key
type KeyBinding struct {
	Key         string
	Description string
}

// Footer is a model displayed at the foot of the application
type Footer struct {
	bindings []KeyBinding
	styles   *style.Styles
}

// NewFooter creates a new footer model
func NewFooter(bindings []KeyBinding, styles *style.Styles) *Footer {
	return &Footer{
		bindings: bindings,
		styles:   styles,
	}
}

// Init is a noop
func (f *Footer) Init() tea.Cmd {
	return nil
}

// Update is a noop
func (f *Footer) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return f, nil
}

// View renders the footer model
func (f *Footer) View() string {
	s := strings.Builder{}
	for i, binding := range f.bindings {
		s.WriteString(f.styles.HelpKey.Render(binding.Key))
		s.WriteString(" ")
		s.WriteString(f.styles.HelpValue.Render(binding.Description))

		if i < len(f.bindings)-1 {
			s.WriteString(f.styles.HelpDivider.Render(f.styles.HelpDivider.Value()))
		}
	}
	return f.styles.Footer.Render(s.String())
}
