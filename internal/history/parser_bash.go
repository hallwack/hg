package history

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
	"time"
)

type BashParser struct {
	Path string
}

func (p *BashParser) Parse(r io.Reader) ([]Entry, error) {
	scanner := bufio.NewScanner(r)

	var entries []Entry
	var currentCmd strings.Builder
	var currentTimestamp time.Time

	for scanner.Scan() {
		line := scanner.Text()

		if after, ok := strings.CutPrefix(line, "#"); ok {
			if timeStamp, err := strconv.ParseInt(after, 10, 64); err == nil {
				currentTimestamp = time.Unix(timeStamp, 0)
				continue
			}
		}

		if before, ok := strings.CutSuffix(line, "\\"); ok {
			currentCmd.WriteString(before)
			currentCmd.WriteString("\n")
			continue
		}

		if currentCmd.Len() > 0 {
			currentCmd.WriteString(line)
			entries = append(entries, Entry{
				Command: currentCmd.String(),
			})
			currentCmd.Reset()
			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		entries = append(entries, Entry{
			Command:   line,
			Timestamp: currentTimestamp,
		})

		currentTimestamp = time.Time{}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	slices.Reverse(entries)
	return entries, nil
}
