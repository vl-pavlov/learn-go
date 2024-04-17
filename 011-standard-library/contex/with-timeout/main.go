package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Do some work
	select {
	case <-ctx.Done():
		fmt.Println("Work completed")
	case <-time.After(10 * time.Second):
		fmt.Println("Work took longer than 10 seconds")
	}
}
