// Copyright (c) 2019 The vin developers. All rights reserved.
// Project site: https://github.com/vimstuff/vin
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vin",
	Short: "vin provides Vim Installation tools",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
