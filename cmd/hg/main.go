package main

import (
	"fmt"
	"os"

	"github.com/hallwack/hg/internal/history"
)

func main() {
	entries, err := history.LoadHistory()

	if err != nil {
		fmt.Fprintf(os.Stderr, "hg: %v\\n", err)
		os.Exit(1)
	}

	for _, entry := range entries {
		fmt.Println(entry.Command)
	}
}
