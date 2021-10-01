package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var group string

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "consume message from topic specified",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("function under development...")
	},
}

func init() {
	rootCmd.AddCommand(receiveCmd)
	receiveCmd.Flags().StringVarP(&brokers, "brokers", "b", "", "broker list")
	receiveCmd.Flags().StringVarP(&group, "group", "g", "", "gorup id")
	receiveCmd.Flags().StringVarP(&topic, "topic", "t", "", "topic")
}
