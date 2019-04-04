// Copyright (c) 2019 The vin developers. All rights reserved.
// Project site: https://github.com/vimstuff/vin
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"log"
	"os/exec"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"in", "i"},
	Short:   "Install a Vim plug-in",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pluginDir, err := homedir.Expand("~/.janus")
		if err != nil {
			log.Fatal(err)
		}
		gitClone(pluginDir, args[0])
	},
}

func gitClone(dir, repo string) {
	c := exec.Command("git", "clone", repo)
	c.Dir = dir
	err := c.Run()
	if err != nil {
		log.Printf("error on git clone %s: %s", repo, err)
	}
}
