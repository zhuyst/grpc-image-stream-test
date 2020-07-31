package main

import (
	"grpc-image-stream-test/client"
	"grpc-image-stream-test/server"
	"log"
	"os"
)

func main() {
	port := 50052
	go func() {
		if err := server.Run(port); err != nil {
			log.Fatalf("server.Run err: %v", err)
		}
	}()
	client.Request(port)
	os.Exit(0)
}
