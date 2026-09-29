package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/hallwack/hg/internal/history"
	"github.com/hallwack/hg/internal/tui"
)

func main() {
	entries, err := history.LoadHistory()
	if err != nil {
		fmt.Printf("Error loading history: %v\n", err)
		os.Exit(1)
	}

	model := tui.InitModel(entries)
	finalModel, err := tea.NewProgram(model).Run()
	if err != nil {
		log.Fatalf("Error running program: %v", err)
	}

	m := finalModel.(tui.Model)

	switch m.ExitAction {
	case tui.ExitExecute:
		fmt.Println(m.SelectedCommand)

	case tui.ExitPrint:
		fmt.Println(m.SelectedCommand)
	}
}
