//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("Input quantity of workers:")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}

	ch := make(chan int)          // канал по которому будут передаваться значения
	br := make(chan os.Signal, 1) // буфферизированный канал нужен для того, что бы сигнал не потерялся

	signal.Notify(br, os.Interrupt, syscall.SIGTERM) // Отправка в канал br сигнала при нажатии на Ctrl + C

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(number int) {
			for c := range ch {
				fmt.Printf("Worker %d: %d\n", number, c)
			}
			fmt.Printf("Worker %d is dead :(\n", number)
			wg.Done()
		}(i)
	}

	for {
		select {
		case <-br: // Ожидания сигнала
			close(ch)
			wg.Wait()
			return
		default: // нет сигнала: отправляем в канал числа
			ch <- rand.Intn(100)
		}
	}

}
