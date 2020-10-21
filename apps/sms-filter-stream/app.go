package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"telenor.com/spam-filter-demo/sms-filter-stream/config"
	"telenor.com/spam-filter-demo/sms-filter-stream/services/kafka"
)

func main() {
	fmt.Println("Starting the LiveCache Triplet Stream Application...")
	cfg := config.New()

	// Define callbacks and processors
	callback := kafka.CallbackFilter(cfg)
	processor, err := kafka.CreateProcessor(cfg, callback)

	// Thread the stream process
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)

	// Trigger the triplet stream
	go func() {
		defer close(done)
		if err = processor.Run(ctx); err != nil {
			log.Fatalf("Error running processor: %v", err)
		} else {
			log.Printf("Processor shutdown cleanly")
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait   // wait for SIGINT/SIGTERM
	cancel() // gracefully stop processor
	<-done
}
