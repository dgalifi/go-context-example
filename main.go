package main

import (
	"context"
	"fmt"
	"time"
)

var cancelTimeout time.Duration = 5 * time.Second
var longRunningOperationDuration time.Duration = 7 * time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout)
	defer cancel()

	connectToExternalResource(ctx)
}

func connectToExternalResource(ctx context.Context) {
	fmt.Println("Connecting to external resource...")

	done := make(chan bool)
	go longRunningOperation(done)

	select {
	case <-done:
		fmt.Println("Connected to external resource.")
	case <-ctx.Done():
		fmt.Println("Connection to external resource canceled. Context is Done")
	}
}

func longRunningOperation(done chan bool) {
	time.Sleep(longRunningOperationDuration)
	done <- true
}
