package commands

import (
	"github.com/ianchildress/branch-tools/tools"
	"github.com/ianchildress/dcli"
	"github.com/ianchildress/dcli/flags"
	"github.com/prometheus/common/log"
)

var Replace = &dcli.CommandNode{
	N:       "replace",
	D:       "replaces a string in a file with another string",
	RunFunc: replace,
}

func replace() {
	old := *flags.GetStringFlag("old").Value()
	new := *flags.GetStringFlag("new").Value()
	filename := *flags.GetStringFlag("file").Value()

	if err := tools.Replace(old, new, filename); err != nil {
		log.Fatal(err)
	}
}
