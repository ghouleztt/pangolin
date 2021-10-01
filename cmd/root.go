package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pangolin",
	Short: "Pangolin is a tool for sending messages to Kafka.",
	Long: "Pangolin is a tool for sending messages to Kafka.\n" +
		"The reason for writing is that the terminal tool does not support the input of super-long text, " +
		"so we cannot send super-long text to Kafka directly. This tool can send long texts to Kafka by reading files",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please use this tool via sub command. lookup -h for help")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		return 
	}
}
