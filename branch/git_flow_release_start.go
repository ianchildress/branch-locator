package branch

import (
	"os/exec"

	"github.com/pkg/errors"
)

func GitFlowReleaseStart(release string) error {
	args := []string{"flow", "release", "start", release}
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return errors.Wrap(errors.WithStack(err), string(out))
	}
	return nil
}
