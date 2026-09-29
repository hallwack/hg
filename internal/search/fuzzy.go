package search

import (
	"github.com/hallwack/hg/internal/history"
	"github.com/sahilm/fuzzy"
)

type Result struct {
	Entry          history.Entry
	Index          int
	Score          int
	MatchedIndexes []int
}

type FuzzyFinder struct{}

func (f *FuzzyFinder) FindEntries(query string, entries []history.Entry) []Result {
	if query == "" {
		result := make([]Result, len(entries))

		for i, entry := range entries {
			result[i] = Result{
				Entry: entry,
			}
		}

		return result
	}

	commands := make([]string, len(entries))
	for i, entry := range entries {
		commands[i] = entry.Command
	}

	matches := fuzzy.Find(query, commands)

	results := make([]Result, 0, len(matches))

	for _, result := range matches {
		results = append(results, Result{
			Entry:          entries[result.Index],
			Index:          result.Index,
			Score:          result.Score,
			MatchedIndexes: result.MatchedIndexes,
		})
	}

	return results
}
