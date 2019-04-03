// Copyright (c) 2019 The vin developers. All rights reserved.
// Project site: https://github.com/vimstuff/vin
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list Vim plug-ins",
	Run: func(cmd *cobra.Command, args []string) {
		// FIXME(mdr): List the disabled Janus plugins.

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
		}
	},
}
