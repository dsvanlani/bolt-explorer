package model

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/style"
)

// Page is a model that displays the contents of the application
type Page struct {
	styles   *style.Styles
	content  tea.Model
	viewport viewport.Model
	height   int
}

// NewPage creates a new page model
func NewPage(content tea.Model, styles *style.Styles) *Page {
	return &Page{
		styles:  styles,
		content: content,
	}
}

// Init is a noop
func (p *Page) Init() tea.Cmd {
	return nil
}

// Update is a noop
func (p *Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.height = p.getDesiredHeight(msg.Height)
		p.viewport = viewport.New(p.styles.Page.GetWidth(), p.height)
	case MenuItem:
		p.content = msg.Content()
	}

	viewport, cmd := p.viewport.Update(msg)
	p.viewport = viewport

	return p, cmd
}

// View renders the page model
func (p *Page) View() string {
	p.viewport.SetContent(p.content.View())
	return p.styles.Page.Copy().Height(p.height).Render(p.viewport.View())
}

// getDesiredHeight returns the desiredHeight of the page component
func (p *Page) getDesiredHeight(windowHeight int) int {
	desiredHeight := windowHeight -
		p.styles.App.GetVerticalMargins() -
		p.styles.Page.GetVerticalMargins() -
		p.styles.Page.GetVerticalBorderSize() -
		p.styles.Footer.GetVerticalFrameSize() - 1
	if desiredHeight < p.styles.Menu.GetHeight() {
		desiredHeight = p.styles.Menu.GetHeight()
	}
	return desiredHeight
}
