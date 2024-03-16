// SPDX-FileCopyrightText: 2019 Kent Gibson <warthog618@gmail.com>
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "undefined"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version",
	Long:  `All software has versions. This is gpiocdev's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s (gpiocdev) %s\n", os.Args[0], version)
	},
}
