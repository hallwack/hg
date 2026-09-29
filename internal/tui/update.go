package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/hallwack/hg/internal/clipboard"
)

type executeFinishedMsg struct {
	err error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.resize(msg.Width, msg.Height)

	case tea.KeyPressMsg:
		switch m.Mode {
		case NormalMode:
			return m.updateNormalMode(msg)

		case SearchMode:
			return m.updateSearchMode(msg)
		}

	}
	return m, nil
}

func (m Model) updateSearchMode(msg tea.KeyPressMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.Mode = NormalMode
		m.Input.Blur()
		return m, nil

	case "enter":
		m.Query = m.Input.Value()
		m.Mode = NormalMode
		m.Input.Blur()

		return m, nil

	case "ctrl+y":
		return m.yankSelected()

	case "ctrl+enter":
		return m.printSelected()
	}

	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)

	query := m.Input.Value()

	if query != m.Query {
		m.Query = query
		m.search()
	}

	return m, cmd
}

func (m Model) updateNormalMode(msg tea.KeyPressMsg) (Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit

	case "j", "down":
		m.moveCursor(1)

	case "k", "up":
		m.moveCursor(-1)

	case "ctrl+d":
		m.moveCursor(m.pageSize())

	case "ctrl+u":
		m.moveCursor(-m.pageSize())

	case "/":
		m.Mode = SearchMode
		m.Input.SetValue(m.Query)

		cmd := m.Input.Focus()
		return m, cmd

	case "y":
		return m.yankSelected()

	case "enter":
		return m.executeSelected()

	case "ctrl+enter":
		return m.printSelected()
	}

	return m, nil
}

func (m *Model) resize(width, height int) {
	viewportHeight := max(height-headerHeight-searchHeight-footerHeight, 1)

	m.Viewport.SetWidth(width)
	m.Viewport.SetHeight(viewportHeight)
}

func (m *Model) search() {
	m.Results = m.Finder.FindEntries(
		m.Input.Value(),
		m.Entries,
	)

	m.Cursor = 0
	m.Viewport.SetYOffset(0)

	m.updateViewport()
}

func (m *Model) ensureCursorVisible() {
	height := m.Viewport.Height()

	if height <= 0 {
		return
	}

	top := m.Viewport.YOffset()
	bottom := top + height - 1

	if m.Cursor < top {
		m.Viewport.SetYOffset(m.Cursor)
		return
	}

	if m.Cursor > bottom {
		m.Viewport.SetYOffset(m.Cursor - height + 1)
	}
}

func (m *Model) moveCursor(delta int) {
	if len(m.Results) == 0 {
		return
	}

	next := max(m.Cursor+delta, 0)

	if next >= len(m.Results) {
		next = len(m.Results) - 1
	}

	if next == m.Cursor {
		return
	}

	m.Cursor = next
	m.ensureCursorVisible()
	m.updateViewport()
}

func (m Model) pageSize() int {
	height := m.Viewport.Height()

	if height < 1 {
		return 1
	}

	return max(height/2, 1)
}

func (m Model) executeSelected() (Model, tea.Cmd) {
	if len(m.Results) == 0 {
		return m, nil
	}

	m.SelectedCommand = m.Results[m.Cursor].Entry.Command
	m.ExitAction = ExitExecute

	return m, tea.Quit
}

func (m Model) printSelected() (Model, tea.Cmd) {
	if len(m.Results) == 0 {
		return m, nil
	}

	m.SelectedCommand = m.Results[m.Cursor].Entry.Command
	m.ExitAction = ExitPrint

	return m, tea.Quit
}

func (m Model) yankSelected() (Model, tea.Cmd) {
	if len(m.Results) == 0 {
		return m, nil
	}

	command := m.Results[m.Cursor].Entry.Command

	if err := clipboard.Copy(command); err != nil {
		return m, tea.Printf("yank failed: %v", err)
	}

	return m, nil
}
