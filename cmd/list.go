// Copyright (c) 2019 The vin developers. All rights reserved.
// Project site: https://github.com/vimstuff/vin
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List Vim plug-ins",
	Run: func(cmd *cobra.Command, args []string) {
		// Disabled Janus included plug-ins.
		vimrcBefore, err := homedir.Expand("~/.vimrc.before")
		if err != nil {
			log.Fatal(err)
		}
		disabled := findDisabled(vimrcBefore)
		if len(disabled) > 0 {
			fmt.Println("The following Janus included plug-ins are disabled:")
			for _, d := range disabled {
				fmt.Printf("  %s\n", d)
			}
		} else {
			fmt.Println("None of the included Janus plug-ins are disabled.")
			fmt.Println("")
		}

		// Determine the custom Janus plug-ins installed and update.
		pluginDir, err := homedir.Expand("~/.janus")
		if err != nil {
			log.Fatal(err)
		}
		plugins, err := ioutil.ReadDir(pluginDir)
		if err != nil {
			log.Fatal(err)
		}
		if len(plugins) > 0 {
			fmt.Println("The following additional plug-ins are installed:")
			for _, plugin := range plugins {
				fmt.Printf("  %s\n", plugin.Name())
			}
		} else {
			fmt.Println("There are no additional plug-ins installed.")
		}
	},
}

func findDisabled(file string) []string {
	submatches := []string{}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return submatches
	}
	re := regexp.MustCompile(`[^#][^ ]call janus#disable_plugin\('(.*)'\)`)
	matches := re.FindAllSubmatch(b, -1)
	for _, match := range matches {
		submatches = append(submatches, string(match[1]))
	}
	return submatches
}
