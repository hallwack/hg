package tui

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"github.com/hallwack/hg/internal/history"
	"github.com/hallwack/hg/internal/search"
)

type Mode int

type ExitAction int

const (
	NormalMode Mode = iota
	SearchMode
)

const (
	ExitNone ExitAction = iota
	ExitExecute
	ExitPrint
)

const (
	headerHeight = 8
	searchHeight = 1
	footerHeight = 8
)

type Model struct {
	Entries []history.Entry
	Results []search.Result

	Finder *search.FuzzyFinder

	Query  string
	Cursor int

	PreviousQuery string

	Mode Mode

	ExitAction      ExitAction
	SelectedCommand string

	Input    textinput.Model
	Viewport viewport.Model
}

func InitModel(entries []history.Entry) Model {
	finder := &search.FuzzyFinder{}

	input := textinput.New()
	input.Prompt = "/ "
	input.Placeholder = "Search..."
	input.CharLimit = 256

	result := finder.FindEntries("", entries)

	vp := viewport.New(
		viewport.WithWidth(80),
		viewport.WithHeight(10),
	)

	model := Model{
		Entries:  entries,
		Results:  result,
		Finder:   finder,
		Input:    input,
		Viewport: vp,
		Mode:     NormalMode,
	}

	model.updateViewport()

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) updateViewport() {
	var s strings.Builder

	for i, result := range m.Results {
		cursor := " "

		if i == m.Cursor {
			cursor = ">"
		}

		fmt.Fprintf(
			&s,
			"%s %s\n",
			cursor,
			result.Entry.Command,
		)
	}

	m.Viewport.SetContent(s.String())
}
