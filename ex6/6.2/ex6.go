package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	w.Add(1)
	go foo(ctx)

	time.Sleep(time.Second)
	cancel()
	w.Wait()
}

func foo(ctx context.Context) {
	defer w.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gorutine stop")
			return
		default:
			fmt.Println("Go!")
		}
	}
}
