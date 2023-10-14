package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	reader *kafka.Reader
)

func readKafka(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "test",
		CommitInterval: 1 * time.Second,
		GroupID:        "group1",
		StartOffset:    kafka.FirstOffset,
	})

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("读失败：%v\n", err)
		} else {
			log.Fatalf("success: topic=%s, partition=%d, offset=%d, key=%s, value=%s\n", message.Topic, message.Partition, message.Offset, message.Key, message.Value)
		}
	}
}

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Fatalf("接受到信号 %s", sig.String())
	if reader != nil {
		reader.Close()
	}
	os.Exit(0)
}
