package commands

import (
	"fmt"
	"log"

	"github.com/ianchildress/branch-tools/branch"
	"github.com/ianchildress/dcli"
	"github.com/ianchildress/dcli/flags"
)

var Locate = &dcli.CommandNode{
	N:       "locate",
	D:       "returns all subdirectories that have the specified git branch",
	RunFunc: locate,
}

func locate() {
	checkoutBranch := *flags.GetStringFlag("branch").Value()
	sd, err := branch.Locate(checkoutBranch)
	if err != nil {
		log.Fatal(err)
	}
	for i := range sd {
		fmt.Println(sd[i])
	}
}
