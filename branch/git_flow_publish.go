package branch

import (
	"os/exec"

	"github.com/pkg/errors"
)

func GitFlowPublish() error {
	args := []string{"flow", "publish"}
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return errors.Wrap(errors.WithStack(err), string(out))
	}
	return nil
}
