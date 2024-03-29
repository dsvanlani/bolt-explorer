package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/observiq/bolt-explorer/router"
	"github.com/observiq/bolt-explorer/style"
)

var keyBindings = []KeyBinding{
	{Key: "w/s", Description: "navigate"},
	{Key: "enter", Description: "select"},
	{Key: "a", Description: "up one level"},
	{Key: "q", Description: "quit"},
	{Key: "f", Description: "toggle search"},
}

// App is the top level model of the application
type App struct {
	header *Header
	menu   *Menu
	page   *Page
	footer *Footer
	styles *style.Styles
	router *router.Router
}

// NewApp creates a new app with the supplied styles
func NewApp(router *router.Router, styles *style.Styles) *App {
	menuItems := []MenuItem{}
	initItems := router.GetPathsForLocation()
	for _, item := range initItems {
		menuItems = append(menuItems, NewMenuItem(item, router, styles))
	}

	return &App{
		styles: styles,
		header: NewHeader("bolt-explorer", styles, router),
		menu:   NewMenu(menuItems, styles, router),
		footer: NewFooter(keyBindings, styles),
		page:   NewPage(menuItems[0].Content(), styles),
		router: router,
	}
}

// Init is a noop
func (a *App) Init() tea.Cmd {
	return nil
}

// Update listens for update messages and sends them to child components
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	responses := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":

			return a, tea.Quit
		case "q":
			if !a.router.SearchMode {
				return a, tea.Quit
			}
		}
	}

	_, response := a.page.Update(msg)
	responses = append(responses, response)

	_, response = a.menu.Update(msg)
	responses = append(responses, response)

	_, response = a.header.Update(msg)
	responses = append(responses, response)

	_, response = a.footer.Update(msg)
	responses = append(responses, response)

	for _, item := range a.menu.items {
		_, response = item.Content().Update(msg)
		responses = append(responses, response)
	}

	return a, tea.Batch(responses...)
}

// View renders the app model
func (a *App) View() string {
	s := strings.Builder{}

	sideBar := lipgloss.JoinVertical(lipgloss.Left, a.header.View(), a.menu.View())
	app := lipgloss.JoinHorizontal(lipgloss.Top, sideBar, a.page.View())
	s.WriteString(app)
	s.WriteRune('\n')
	s.WriteString(a.footer.View())
	return a.styles.App.Render(s.String())
}
