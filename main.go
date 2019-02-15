package main

import (
	"github.com/ianchildress/dcli"
)

const (
	service     = "branch-tools"
	description = "a series of tools i use to start and finish git flow releases across multiple repositories"
)

func main() {
	// create the top level menu node
	top := dcli.New(service, description)

	// Start
	dcli.Start(top)
}
