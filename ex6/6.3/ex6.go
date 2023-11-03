package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := false
	var w sync.WaitGroup
	w.Add(1)

	go func() {
		defer w.Done()
		for {
			if done {
				fmt.Println("The end")
				return
			}
			fmt.Println("Work")
		}

	}()

	time.Sleep(time.Second)
	done = true
	w.Wait()
}
