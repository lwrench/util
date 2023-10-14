package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	// writeKafka(ctx)

	go listenSignal()
	readKafka(ctx)
}
