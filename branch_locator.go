package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

var (
	rootDir, branch string
)

func locate() {
	var err error
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Missing required branch arg")
		fmt.Println("example: branch-locator master")
		return
	}
	branch = args[1]

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
	b, err := hasBranch(branch)
	if err != nil {
		return errors.WithStack(err)
	}
	if b {
		fmt.Println(dir.Name())
	}
	return nil
}

func hasBranch(branch string) (bool, error) {
	args := []string{"branch"}
	out, err := exec.Command("git", args...).Output()
	if err != nil {
		return false, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := strings.TrimPrefix(scanner.Text(), "*")

		if strings.TrimSpace(line) == branch {
			return true, nil
		}
	}
	return false, nil
}
