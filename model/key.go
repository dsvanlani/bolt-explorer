package model

import (
	"encoding/json"

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
	json.Unmarshal(kv.value, &obj)
	j, _ := json.MarshalIndent(obj, "", "  ")
	kv.viewport.SetContent(string(j))
	return kv.viewport.View()
}
