package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hallwack/hg/internal/history"
	"github.com/hallwack/hg/internal/search"
)

func main() {
	entries, err := history.LoadHistory()

	if err != nil {
		fmt.Fprintf(os.Stderr, "hg: %v\n", err)
		os.Exit(1)
	}

	var query string
	if len(os.Args) > 1 {
		query = strings.Join(os.Args[1:], " ")
	}

	commands := make([]history.Entry, 0, len(entries))
	commands = append(commands, entries...)

	finder := &search.FuzzyFinder{}
	results := finder.Find(query, commands)

	for _, command := range results {
		fmt.Println(command)
	}
}
