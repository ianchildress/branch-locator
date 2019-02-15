package branch

import (
	"os/exec"
	"strings"
)

// Enabled checks if the current directory is git enabled
func Enabled() bool {
	args := []string{"status"}
	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		return false
	}

	if strings.Contains(string(output), "On branch") {
		return true
	}

	return false
}
