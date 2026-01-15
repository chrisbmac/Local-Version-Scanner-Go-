package runner

import (
	"os/exec"
	"strings"
)

type Result struct {
	Name    string
	Version string
	Error   string
}

func Run(name string, cmd []string) Result {
	c := exec.Command(cmd[0], cmd[1:]...)

	out, error := c.CombinedOutput()

	if error != nil {
		return Result{
			Name:  name,
			Error: error.Error(),
		}
	}

	text := strings.TrimSpace(string(out))

	firstLine := text

	if idx := strings.Index(text, "\n"); idx != -1 {
		firstLine = text[:idx]
	}

	return Result{
		Name:    name,
		Version: firstLine,
	}
}
