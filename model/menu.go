package model

import (
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/observiq/bolt-explorer/logger"
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
	scrollOffset  int
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
		case "up":
			if m.selectedIndex > 0 {
				m.selectedIndex--
				return m.sendMenuChangeEvent()
			}

			if m.selectedIndex == 0 && m.scrollOffset > 0 {
				m.scrollOffset--
				return m, nil
			}

		case "down":
			if m.selectedIndex < len(m.GetItems())-1 && m.selectedIndex < m.maxItems()-1 {
				m.selectedIndex++
				return m.sendMenuChangeEvent()
			}

			if m.selectedIndex == m.maxItems()-1 {
				m.scrollOffset++
				return m, nil
			}

		case "left":
			if m.active {

				m.router.GoUpOneLevel()

				m.items = []MenuItem{}
				m.selectedIndex = 0
				m.scrollOffset = 0
				for _, item := range m.router.GetPathsForLocation() {
					m.items = append(m.items, NewMenuItem(item, m.router, m.styles))
				}
				return m.sendMenuChangeEvent()
			}
		case "enter":
			if m.active {
				if !m.router.Paths[m.items[m.selectedIndex].key].IsBucket() {
					return m, nil
				}

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
	if m.items == nil || len(m.items) == 0 {
		return m, nil
	}
	cmd := func() tea.Msg {
		return m.GetItems()[m.selectedIndex]
	}

	return m, cmd
}

// View renders the menu model
func (m *Menu) View() string {
	items := m.GetItems()

	s := strings.Builder{}
	viewport := items[m.scrollOffset:]

	if len(viewport) > m.maxItems() {
		viewport = viewport[:m.maxItems()]
	}

	for i, item := range viewport {
		if m.scrollOffset > 0 && i == 0 {
			s.WriteString(m.styles.MenuItem.Render("..."))
			s.WriteString("\n")
		}
		switch {
		case i == m.selectedIndex:

			s.WriteString(m.styles.MenuCursor.String())

			// split the runcated string by the m.router.SearchValue
			truncated := truncateString(item.key, 80)
			builder := strings.Builder{}
			if !m.router.SearchMode {
				s.WriteString(m.styles.SelectedMenuItem.Render(truncated))
				s.WriteString("\n")
				continue
			}

			if !strings.Contains(truncated, m.router.SearchValue) {
				s.WriteString(m.styles.SelectedMenuItem.Render(truncated))
			}

			splits := strings.Split(truncated, m.router.SearchValue)

			if len(splits) > 1 {
				for i, split := range splits {
					builder.WriteString(split)
					if i < len(splits)-1 {
						builder.WriteString(m.styles.MenuItemTextHighlight.Render(m.router.SearchValue))
					}
				}
			}

			s.WriteString(m.styles.SelectedMenuItem.Render(builder.String()))
		default:
			// split the runcated string by the m.router.SearchValue
			truncated := truncateString(item.key, 80)

			builder := strings.Builder{}

			if !m.router.SearchMode {
				s.WriteString(m.styles.MenuItem.Render(truncated))
				s.WriteString("\n")
				continue
			}

			if !strings.Contains(truncated, m.router.SearchValue) {
				s.WriteString(m.styles.MenuItem.Render(truncated))
			}

			splits := strings.Split(truncated, m.router.SearchValue)
			logger.Logger().Debug("splits: ", splits)
			logger.Logger().Debug("truncated: ", truncated)
			if len(splits) > 1 {
				for i, split := range splits {
					builder.WriteString(split)
					if i < len(splits)-1 {
						builder.WriteString(m.styles.MenuItemTextHighlight.Render(m.router.SearchValue))
					}
				}
			}

			s.WriteString(m.styles.MenuItem.Render(builder.String()))
		}
		s.WriteRune('\n')

		if i == m.maxItems()-1 {
			s.WriteString(m.styles.MenuItem.Render("..."))
			break
		}
	}

	return m.styles.Menu.Copy().Height(m.desiredHeight()).Render(s.String())
}

func (m *Menu) GetItems() []MenuItem {
	items := []MenuItem{}
	if m.router.SearchMode {
		if m.router.SearchValue != "" {
			for path := range m.router.Paths {
				if strings.Contains(strings.ToLower(path), strings.ToLower(m.router.SearchValue)) {
					items = append(items, NewMenuItem(path, m.router, m.styles))
				}
			}
		}

	} else {
		items = m.items
	}
	// sort the items
	sort.Slice(items, func(i, j int) bool {
		return items[i].key < items[j].key
	})
	return items
}

func (m *Menu) desiredHeight() int {
	return m.windowHeight -
		m.styles.App.GetVerticalMargins() -
		m.styles.Header.GetVerticalFrameSize() -
		m.styles.Header.GetVerticalBorderSize() -
		m.styles.Menu.GetVerticalMargins() -
		m.styles.Menu.GetVerticalBorderSize() -
		m.styles.Footer.GetVerticalFrameSize()
}

func (m *Menu) maxItems() int {
	max := m.desiredHeight() - 4

	if max < 0 {
		return 0
	}

	return max
}

func truncateString(str string, length int) string {
	runeStr := []rune(str)

	if len(runeStr) > length {
		return string(runeStr[:length])
	}

	return str
}
