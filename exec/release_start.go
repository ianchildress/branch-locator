package exec

import (
	"fmt"
	"log"
	"os"

	"github.com/ianchildress/branch-tools/branch"
	"github.com/pkg/errors"
)

// ReleaseStart will create a release branch for each of the provided directories
// todo: i would prefer to scan all directories for unsaved changes prior to beginning this procedure.
// to get this tool up and running we will go ahead and commit any unsaved changes without consideration for
// consequences
func ReleaseStart(release string, directories []string) error {
	home, err := os.Getwd()
	if err != nil {
		return errors.Wrap(errors.WithStack(err), "failed to get working directory")
	}

	// make sure each of the specified directories is a git enabled repository
	for i := range directories {
		if err := os.Chdir(home); err != nil {
			return errors.Wrapf(errors.WithStack(err), "failed to change directory %s", home)
		}

		if err := os.Chdir(directories[i]); err != nil {
			return errors.Wrapf(errors.WithStack(err), "failed to change directory %s", directories[i])
		}

		if !branch.Enabled() {
			return fmt.Errorf("unable to create branch for %s, it is not a git enabled directory", directories[i])
		}
	}

	// create the release branches
	for i := range directories {
		if err := os.Chdir(home); err != nil {
			return errors.Wrapf(errors.WithStack(err), "failed to change directory %s", home)
		}

		if err := os.Chdir(directories[i]); err != nil {
			return errors.Wrapf(errors.WithStack(err), "failed to change directory %s", directories[i])
		}

		// git checkout develop
		if err := branch.Checkout("develop"); err != nil {
			log.Fatal(err)
		}

		// git add
		if err := branch.Add(); err != nil {
			log.Fatal(err)
		}

		// git commit
		if err := branch.Commit(fmt.Sprintf("committing unsaved changes prior to release %s", release)); err != nil {
			log.Fatal(err)
		}

		// git pull
		if err := branch.Pull(); err != nil {
			log.Fatal(err)
		}

		// git flow release start
		if err := branch.GitFlowReleaseStart(release); err != nil {
			return errors.Wrapf(err, "unable to create release branch for %s, check if it is git flow enabled",
				directories[i])
		}

		// git flow publish
		if err := branch.GitFlowPublish(); err != nil {
			return errors.Wrapf(err, "unable to publish release branch for %s, check if it is git flow enabled",
				directories[i])
		}
	}
	return nil
}
