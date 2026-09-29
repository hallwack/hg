package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	var s strings.Builder

	s.WriteString("History Finder\n")

	if m.Mode == SearchMode {
		fmt.Fprintf(&s, "/ %s", m.Query)
	} else {
		fmt.Fprintf(&s, "  %s", m.Query)
	}

	s.WriteString("\n")

	s.WriteString(m.Viewport.View())

	s.WriteString("\n")

	if m.Mode == SearchMode {
		s.WriteString("SEARCH • Enter confirm • Esc cancel")
	} else {
		fmt.Fprintf(
			&s,
			"NORMAL • %d results • / search • q quit",
			len(m.Results),
		)
	}

	v := tea.NewView(s.String())
	v.AltScreen = true
	v.KeyboardEnhancements.ReportAllKeysAsEscapeCodes = true
	return v
}
