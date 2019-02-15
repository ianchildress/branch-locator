package commands

import (
	"github.com/ianchildress/branch-tools/branch"
	"github.com/ianchildress/dcli"
	"github.com/ianchildress/dcli/flags"
	"github.com/prometheus/common/log"
)

var Checkout = &dcli.CommandNode{
	N:       "checkout",
	D:       "checks out a git repository",
	RunFunc: checkout,
}

func checkout() {
	checkoutBranch := *flags.GetStringFlag("branch").Value()
	if err := branch.Checkout(checkoutBranch); err != nil {
		log.Fatal(err)
	}
}
