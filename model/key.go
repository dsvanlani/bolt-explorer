package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

// KeyView is the value of a key in a bucket
type KeyView struct {
	styles *style.Styles
	value  []byte
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
func (kv *KeyView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return kv, nil
}

// View renders the key value model
func (kv *KeyView) View() string {
	return kv.styles.Key.Render(string(kv.value))
}
