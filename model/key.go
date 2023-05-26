package model

import (
	"encoding/json"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

// KeyView is the value of a key in a bucket
type KeyView struct {
	styles   *style.Styles
	value    []byte
	viewport viewport.Model
}

// NewKeyView creates a new key value model
func NewKeyView(value []byte, styles *style.Styles) *KeyView {
	return &KeyView{
		styles: styles,
		value:  value,
	}
}

// Init is a noop
func (kv *KeyView) Init() tea.Cmd {
	return nil
}

// Update is a noop
func (m *KeyView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// View renders the key value model
func (kv *KeyView) View() string {
	kv.viewport = viewport.New(98, 5000)

	obj := make(map[string]any)
	err := json.Unmarshal(kv.value, &obj)
	if err != nil {

		builder := strings.Builder{}
		// Try to separate by newlines and parse into an array
		newLines := strings.Split(string(kv.value), "\n")

		for _, line := range newLines {
			if line == "" {
				continue
			}

			// Try to parse into an object
			obj := make(map[string]any)
			err := json.Unmarshal([]byte(line), &obj)
			if err != nil {
				builder.WriteString(line)
				builder.WriteString("\n")
				continue
			} else {

				bytes, _ := json.MarshalIndent(obj, "", "  ")
				builder.WriteString(string(bytes))
				builder.WriteString("\n")
			}

		}
		return builder.String()
	}
	j, _ := json.MarshalIndent(obj, "", "  ")
	kv.viewport.SetContent(string(j))
	return kv.viewport.View()
}
