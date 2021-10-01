package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show pangolin version information",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version: v0.1\nauthor:Logan Chen <ghouleztt@gmail.com>")
	},
}
