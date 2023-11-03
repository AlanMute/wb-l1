package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var w sync.WaitGroup

func read(ch chan int) {
	for a := range ch {
		fmt.Println(a)
	}
	w.Done()
}

func main() {
	ch := make(chan int)

	w.Add(1)

	fmt.Println("Input the time duration:")
	var t int
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Fatal(err)
	}

	go read(ch)
	// контекст для отслеживания времени
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(t))
	defer cancel()

	for {
		select {
		case <-ctx.Done(): // По истечению времени программа завершиться
			close(ch)
			w.Wait()
			fmt.Println("The End!")
			return
		default: // время еще есть? Отпарвляем в канал рандомные значения
			ch <- rand.Intn(100)
		}
	}

}
