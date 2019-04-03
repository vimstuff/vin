// Copyright (c) 2019 The vin developers. All rights reserved.
// Project site: https://github.com/vimstuff/vin
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Vim plugins",
	Run: func(cmd *cobra.Command, args []string) {
		// Update Janus
		fmt.Println("Updating Janus (running rake)")
		janusDir, err := homedir.Expand("~/.vim")
		if err != nil {
			log.Fatal(err)
		}
		go rake(janusDir)

		// Determine the custom Janus plug-ins installed and update.
		pluginDir, err := homedir.Expand("~/.janus")
		if err != nil {
			log.Fatal(err)
		}
		plugins, err := ioutil.ReadDir(pluginDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, plugin := range plugins {
			fmt.Println(plugin.Name())
			p := path.Join(pluginDir, plugin.Name())
			go gitPull(p)
		}
	},
}

func rake(dir string) {
	// Update Janus
	c := exec.Command("rake")
	c.Dir = dir
	err := c.Run()
	if err != nil {
		log.Printf("rake had error: %s", err)
	}
}

func gitPull(dir string) {
	c := exec.Command("git", "pull")
	c.Dir = dir
	err := c.Run()
	if err != nil {
		log.Printf("git pull %s had error: %s", dir, err)
	}
}
