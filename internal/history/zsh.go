// Package history provides functionality to parse shell history files.
package history

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type ZshParser struct {
}

func (p *ZshParser) Parse(r io.Reader) ([]Entry, error) {
	scanner := bufio.NewScanner(r)

	var entries []Entry
	var currentCmd strings.Builder
	var currentTimestamp time.Time

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, ": ") {
			if currentCmd.Len() > 0 {
				entries = append(entries, Entry{
					Command:   strings.TrimSpace(currentCmd.String()),
					Timestamp: currentTimestamp,
				})
				currentCmd.Reset()
			}

			timeStamp, command, err := parseZshLine(line)
			if err != nil {
				continue // Skip lines that don't match the expected format
			}

			currentTimestamp = timeStamp
			currentCmd.WriteString(command)
			continue
		}

		if currentCmd.Len() > 0 {
			currentCmd.WriteString("\n")
			currentCmd.WriteString(line)
		}
	}

	if currentCmd.Len() > 0 {
		entries = append(entries, Entry{
			Command:   strings.TrimSpace(currentCmd.String()),
			Timestamp: currentTimestamp,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func parseZshLine(line string) (time.Time, string, error) {
	line = strings.TrimPrefix(line, ": ")

	parts := strings.SplitN(line, ";", 2)
	if len(parts) != 2 {
		return time.Time{}, "", fmt.Errorf("zsh format invalid")
	}

	meta := parts[0]
	command := parts[1]

	metaParts := strings.SplitN(meta, ":", 2)
	if len(metaParts) < 1 {
		return time.Time{}, "", fmt.Errorf("zsh metadata invalid")
	}

	timeStampUnix, err := strconv.ParseInt(metaParts[0], 10, 64)
	if err != nil {
		return time.Time{}, command, nil
	}

	return time.Unix(timeStampUnix, 0), command, nil
}
