package branch

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

// Checkout will check out the specified git branch
func Checkout(branch string) error {
	exists, err := Exists(branch)
	if err != nil {
		return errors.WithStack(err)
	}
	if !exists {
		return errors.WithStack(fmt.Errorf("branch %s doesn't exist", branch))
	}

	args := []string{"checkout", branch}
	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return errors.WithStack(err)
	}
	fmt.Println(string(output))
	return nil
}
