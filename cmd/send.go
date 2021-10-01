package cmd

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var brokers string
var topic string
var file bool

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send message to specified topic",
	Run: func(cmd *cobra.Command, args []string) {
		if !file {
			syncProducer(brokers, topic, strings.Join(args[:], ""))
			return
		}

		log.Printf("start to read file%v\n", args[0])

		// open file
		f, err := os.Open(args[0])
		defer func(f *os.File) {
			_ = f.Close()
		}(f)
		if err != nil {
			log.Fatalf("failed to open file %v :%v", args[0], err)
			return
		}

		// read file
		content, _ := ioutil.ReadAll(f)
		if err != nil {
			log.Fatalf("failed to read file %v:%v", args[0], err)
			return
		}
		syncProducer(brokers, topic, string(content))
	},
}

func syncProducer(broker, topic, message string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	p, err := sarama.NewSyncProducer(strings.Split(brokers, ","), config)
	defer func(p sarama.SyncProducer) {
		_ = p.Close()
	}(p)
	if err != nil {
		log.Fatalf("failed to connect to kafka:%v", err)
		return
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}
	if _, _, err := p.SendMessage(msg); err != nil {
		log.Fatalf("failed to send message:%v", err)
		return
	}
	log.Printf("success\n")
}

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&brokers, "brokers", "b", "", "broker list")
	sendCmd.Flags().StringVarP(&topic, "topic", "t", "", "topic")
	sendCmd.Flags().BoolVarP(&file, "file", "f", false, "message file")
}
