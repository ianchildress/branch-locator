package tools

import (
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

// Replace will open the specified file and replace the old string with the new string
func Replace(old, new, filename string) error {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.WithStack(err)
	}

	newContents := strings.Replace(string(input), old, new, -1)

	err = ioutil.WriteFile(filename, []byte(newContents), 0644)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
