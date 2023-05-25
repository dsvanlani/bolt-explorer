package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

type BucketView struct {
	styles   *style.Styles
	children []string
}

func NewBucketView(children []string, styles *style.Styles) *BucketView {
	return &BucketView{
		styles:   styles,
		children: children,
	}
}

func (b *BucketView) Init() tea.Cmd {
	return nil
}

func (b *BucketView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return b, nil
}

func (b *BucketView) View() string {
	s := strings.Builder{}
	for _, child := range b.children {
		s.WriteString(child)
		s.WriteString("\n")
	}

	return b.styles.Bucket.Render(s.String())
}
