package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ianchildress/branch-tools/branch"
	"github.com/pkg/errors"
)

var (
	rootDir, myBranch string
)

func locate() {
	var err error
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Missing required myBranch arg")
		fmt.Println("example: myBranch-locator master")
		return
	}
	myBranch = args[1]

	rootDir, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	// list directories
	directories, err := directories()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := range directories {
		if err := checkForBranch(directories[i]); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func directories() ([]os.FileInfo, error) {
	var output []os.FileInfo
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return output, errors.WithStack(err)
	}
	for i := range files {
		if files[i].IsDir() {
			output = append(output, files[i])
		}
	}
	return output, nil
}

func checkForBranch(dir os.FileInfo) error {
	defer os.Chdir(rootDir)
	if err := os.Chdir(dir.Name()); err != nil {
		return errors.WithStack(err)
	}
	b, err := branch.Exists(myBranch)
	if err != nil {
		return errors.WithStack(err)
	}
	if b {
		fmt.Println(dir.Name())
	}
	return nil
}
