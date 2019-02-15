package branch

import (
	"os/exec"

	"github.com/pkg/errors"
)

func Add() error {
	args := []string{"add", "."}
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return errors.Wrap(errors.WithStack(err), string(out))
	}
	return nil
}
