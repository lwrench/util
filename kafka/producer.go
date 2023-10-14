package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func writeKafka(ctx context.Context) {
	writer := kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "test",
		Balancer:               &kafka.Hash{},
		WriteTimeout:           1 * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()

	for i := 0; i < 3; i++ {
		if err := writer.WriteMessages(ctx, kafka.Message{Key: []byte("1"), Value: []byte("test")}); err != nil {
			if err == kafka.LeaderNotAvailable {
				time.Sleep(500 * time.Millisecond)
				continue
			} else {
				log.Fatalf("写入失败:%v/n", err)
			}
		} else {
			break
		}
	}
}
