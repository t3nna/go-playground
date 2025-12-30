package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func doWork(ctx context.Context, wg *sync.WaitGroup) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work Done")
	case <-ctx.Done():
		fmt.Println("Canceled: ", ctx.Err())
	}
	wg.Done()
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go doWork(ctx, wg)

	wg.Wait()
}
