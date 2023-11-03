//Реализовать все возможные способы остановки выполнения горутины

package main

import (
	"fmt"
	"sync"
	"time"
)

var w sync.WaitGroup

func main() {
	done := make(chan bool)

	w.Add(1)
	go foo(done)

	time.Sleep(time.Second)
	done <- true
	w.Wait()
}

func foo(done chan bool) {
	defer w.Done()
	for {
		select {
		case <-done:
			fmt.Println("Gorutine stop")
			return
		default:
			fmt.Println("Go!")
		}
	}
}
