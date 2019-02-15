package branch

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// Locate will search for a branch in directories 1 level deep and return all directories that are git enabled
// and have the specified branch
func Locate(branch string) ([]string, error) {
	var output []string

	// get working directory
	wd, err := os.Getwd()
	if err != nil {
		return output, errors.WithStack(err)
	}

	// get a list of subdirectories
	sd, err := subdirectories()
	if err != nil {
		return output, errors.WithStack(err)
	}

	// iterate over the subdirectories and find the directories that have the specified branch
	for i := range sd {
		cd := fmt.Sprintf("%s/%s", wd, sd[i].Name())
		if err := os.Chdir(cd); err != nil {
			return output, errors.WithStack(err)
		}

		// check if the directory is git enabled
		if !Enabled() {
			continue
		}

		// check if the branch exists locally in this repo
		exists, err := Exists(branch)
		if err != nil {
			return output, errors.WithStack(err)
		}
		if exists {
			output = append(output, sd[i].Name())
		}
	}
	return output, nil
}

func subdirectories() ([]os.FileInfo, error) {
	var output []os.FileInfo
	fi, err := ioutil.ReadDir(".")
	if err != nil {
		return output, errors.WithStack(err)
	}
	for i := range fi {
		if fi[i].IsDir() {
			output = append(output, fi[i])
		}
	}
	return output, nil
}
