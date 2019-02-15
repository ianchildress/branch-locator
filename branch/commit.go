package branch

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func Commit(message string) error {
	args := []string{"commit", "-m", message}
	out, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		if strings.Contains(string(out), "nothing to commit") {
			fmt.Println(string(out))
			return nil
		}
		return errors.Wrap(errors.WithStack(err), string(out))
	}
	return nil
}
