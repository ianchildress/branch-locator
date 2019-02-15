package branch

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

// Exists will check if the current working directory has a git branch of the specified name
func Exists(branch string) (bool, error) {
	args := []string{"branch"}
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return false, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := strings.TrimPrefix(scanner.Text(), "*")

		if strings.TrimSpace(line) == branch {
			return true, nil
		}
	}
	return false, nil
}
