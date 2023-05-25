package model

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/router"
	"github.com/observiq/bolt-explorer/style"
)

// MenuItem is a an item in the menu tied to a model
type MenuItem struct {
	key    string
	router *router.Router
	styles *style.Styles
}

// NewMenuItem creates a new menu item
func NewMenuItem(key string, router *router.Router, styles *style.Styles) MenuItem {
	return MenuItem{
		key:    key,
		router: router,
		styles: styles,
	}
}

// Menu is a model used to select items
type Menu struct {
	items         []MenuItem
	active        bool
	styles        *style.Styles
	selectedIndex int
	windowHeight  int
	router        *router.Router
}

// NewMenu creates a new menu model
func NewMenu(items []MenuItem, styles *style.Styles, router *router.Router) *Menu {
	return &Menu{
		items:  items,
		styles: styles,
		active: true,
		router: router,
	}
}

func (mi *MenuItem) Content() tea.Model {
	route := mi.router.Paths[mi.key]
	if route.IsBucket() {
		return NewBucketView(route.Children, mi.styles)
	}

	return NewKeyView(route.Value, mi.styles)
}

// Init is a noop
func (m *Menu) Init() tea.Cmd {
	return nil
}

// Update listens for key messages and updates the menu selection
func (m *Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
	case tea.KeyMsg:

		switch msg.String() {
		case "k", "up":
			if m.active && m.selectedIndex > 0 {
				m.selectedIndex--
				return m.sendMenuChangeEvent()
			}
		case "j", "down":
			if m.active && m.selectedIndex < len(m.items)-1 {
				m.selectedIndex++
				return m.sendMenuChangeEvent()
			}
		case "esc":
			if !m.active {
				m.active = true
			}
		case "enter":
			if m.active {
				m.router.SetLocation(m.items[m.selectedIndex].key)

				m.items = []MenuItem{}
				m.selectedIndex = 0
				for _, item := range m.router.GetPathsForLocation() {
					m.items = append(m.items, NewMenuItem(item, m.router, m.styles))
				}

				return m.sendMenuChangeEvent()
			}
		}
	}
	return m, nil
}

// sendMenuChangeEvent sends an event indicating the menu selection has changed.
func (m *Menu) sendMenuChangeEvent() (tea.Model, tea.Cmd) {
	cmd := func() tea.Msg {
		return m.items[m.selectedIndex]
	}

	return m, cmd
}

// View renders the menu model
func (m *Menu) View() string {
	s := strings.Builder{}
	for i, item := range m.items {
		switch {
		case i == m.selectedIndex:
			s.WriteString(m.styles.MenuCursor.String())
			s.WriteString(m.styles.SelectedMenuItem.Render(item.key))
		default:
			s.WriteString(m.styles.MenuItem.Render(item.key))
		}
		s.WriteRune('\n')
		s.WriteRune('\n')
	}

	desiredHeight := m.windowHeight -
		m.styles.App.GetVerticalMargins() -
		m.styles.Header.GetVerticalFrameSize() -
		m.styles.Header.GetVerticalBorderSize() -
		m.styles.Menu.GetVerticalMargins() -
		m.styles.Menu.GetVerticalBorderSize() -
		m.styles.Footer.GetVerticalFrameSize()
	if desiredHeight < m.styles.Menu.GetHeight() {
		desiredHeight = m.styles.Menu.GetHeight()
	}

	if !m.active {
		m.styles.MenuInactive.Copy().Height(desiredHeight).Render(s.String())
	}

	return m.styles.Menu.Copy().Height(desiredHeight).Render(s.String())
}
