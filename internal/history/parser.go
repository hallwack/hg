// Package history provides functionality to parse shell history files.
package history

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Parser interface {
	Parse(r io.Reader) ([]Entry, error)
}

func NewParser() Parser {
	shell := os.Getenv("SHELL")
	switch {
	case strings.Contains(shell, "bash"):
		return &BashParser{}
	case strings.Contains(shell, "zsh"):
		return &ZshParser{}
	default:
		return &BashParser{} // Default to BashParser if shell is unknown
	}
}

func LoadHistory() ([]Entry, error) {
	shell := os.Getenv("SHELL")
	var path string

	switch {
	case strings.Contains(shell, "bash"):
		path = filepath.Join(os.Getenv("HOME"), ".bash_history")
	case strings.Contains(shell, "zsh"):
		path = filepath.Join(os.Getenv("HOME"), ".zsh_history")
	default:
		path = filepath.Join(os.Getenv("HOME"), ".bash_history")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open history file %s: %w", path, err)
	}
	defer file.Close()

	parser := NewParser()
	return parser.Parse(file)
}
