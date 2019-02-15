package commands

import (
	"github.com/ianchildress/dcli"
	"github.com/ianchildress/dcli/flags"
)

var ReleaseStart = &dcli.CommandNode{
	N:       "release-start",
	D:       "creates a git flow release branch for all repositories specified in the config file",
	RunFunc: releaseStart,
}

func releaseStart() {
	release := *flags.GetStringFlag("release").Value()
	config := *flags.GetStringFlag("config").Value()

}
