//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
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

	timeout := time.After(time.Second * time.Duration(t)) // начинаем обратный отчет

	for {
		select {
		case <-timeout: // По истечению времени в канал будет отправлено текущее время и программа завершиться
			close(ch)
			w.Wait()
			fmt.Println("The End!")
			return
		default: // время еще есть? Отпарвляем в канал рандомные значения
			ch <- rand.Intn(100)
		}
	}

}
