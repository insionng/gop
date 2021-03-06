// Copyright 2017 The Gop Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

const (
	defaultMainFile = `package main

func main() {
		
}
`

	defaultYaml = `targets:
- name: %s
  dir: main
  assets:
  - templates
  - public
`
)

// CmdInit represents
var CmdInit = cli.Command{
	Name:        "init",
	Usage:       "Init a new project",
	Description: `Init a new project`,
	Action:      runInit,
}

func runInit(ctx *cli.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Join(wd, "src"), os.ModePerm)
	os.MkdirAll(filepath.Join(wd, "src", "vendor"), os.ModePerm)
	os.MkdirAll(filepath.Join(wd, "src", "main"), os.ModePerm)
	os.MkdirAll(filepath.Join(wd, "bin"), os.ModePerm)

	ymlPath := filepath.Join(wd, "gop.yml")
	_, err = os.Stat(ymlPath)
	if err != nil {
		if os.IsNotExist(err) {
			y, err := os.Create(ymlPath)
			if err != nil {
				return err
			}
			defer y.Close()

			_, err = y.Write([]byte(fmt.Sprintf(defaultYaml, filepath.Base(wd))))
			if err != nil {
				return err
			}
		}
	}

	mainFile := filepath.Join(wd, "src", "main", "main.go")
	_, err = os.Stat(mainFile)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(mainFile)
			if err != nil {
				return fmt.Errorf("os.Create: %v", err)
			}
			defer f.Close()
			_, err = f.Write([]byte(defaultMainFile))
			if err != nil {
				return fmt.Errorf("create main file failed: %v", err)
			}
			return nil
		}
		return fmt.Errorf("os.State: %v", err)
	}

	return nil
}
