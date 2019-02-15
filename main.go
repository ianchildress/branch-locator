package main

import (
	"github.com/ianchildress/branch-tools/commands"
	"github.com/ianchildress/dcli"
)

const (
	service     = "myBranch-tools"
	description = "a series of tools i use to start and finish git flow releases across multiple repositories"
)

func main() {
	// create the top level menu node
	top := dcli.New(service, description)

	// branch checkout
	top.AddSubCommand(commands.Checkout)
	commands.Checkout.NewStringFlag("branch", "name of the branch to checkout", true)

	// branch locator
	top.AddSubCommand(commands.Locate)
	commands.Locate.NewStringFlag("branch", "name of the branch to locate", true)

	// replace tool
	top.AddSubCommand(commands.Replace)
	commands.Replace.NewStringFlag("old", "old string", true)
	commands.Replace.NewStringFlag("new", "new string", true)
	commands.Replace.NewStringFlag("file", "file location", true)

	// Start
	dcli.Start(top)
}
