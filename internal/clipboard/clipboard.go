package clipboard

import (
	"fmt"
	"os/exec"
	"strings"

	goclip "github.com/tpyle/goclip"
)

func Copy(text string) error {
	err := goclip.Write(text)
	if err == nil {
		return nil
	}

	// Fallback
	if path, lookErr := exec.LookPath("wl-copy"); lookErr == nil {
		cmd := exec.Command(path)
		cmd.Stdin = strings.NewReader(text)
		if runErr := cmd.Run(); runErr == nil {
			return nil
		}
	}

	return fmt.Errorf("clipboard: failed to copy text to clipboard: %w", err)
}

func MaybeRunHolder(args []string) bool {
	return goclip.MaybeRunHolder(args)
}
