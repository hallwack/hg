package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hallwack/hg/internal/clipboard"
	"github.com/hallwack/hg/internal/history"
	"github.com/hallwack/hg/internal/search"
)

func main() {
	var copyFlag bool
	flag.BoolVar(&copyFlag, "c", false, "Copy the selected command to the clipboard")
	flag.BoolVar(&copyFlag, "copy", false, "Copy the selected command to the clipboard")

	flag.Parse()

	if clipboard.MaybeRunHolder(os.Args[1:]) {
		return
	}

	entries, err := history.LoadHistory()

	if err != nil {
		fmt.Fprintf(os.Stderr, "hg: %v\n", err)
		os.Exit(1)
	}

	query := strings.Join(flag.Args(), " ")

	finder := &search.FuzzyFinder{}
	results := finder.FindEntries(query, entries)

	if copyFlag {
		if len(results) == 0 {
			fmt.Fprintln(os.Stderr, "hg: tidak ada hasil yang cocok untuk disalin")
			os.Exit(1)
		}

		textToCopy := results[0].Entry.Command
		if err := clipboard.Copy(textToCopy); err != nil {
			fmt.Fprintf(os.Stderr, "hg: gagal menyalin ke clipboard: %v\n", err)
			os.Exit(1)
		}

		// Beri pemberitahuan singkat ke stderr agar tidak mengganggu stdout
		fmt.Fprintf(os.Stderr, "Copied to clipboard: %s\n", textToCopy)
	}

	for _, command := range results {
		fmt.Println(command.Entry.Command)
	}
}
