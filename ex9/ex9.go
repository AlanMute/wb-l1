//Разработать конвейер чисел. Даны два канала:
//в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout

package main

import "fmt"

func generator(ch chan int) {
	for x := 1; x <= 5; x++ {
		ch <- x
	}
	close(ch)
}

func pow(ch chan int, pw chan int) {
	for a := range ch {
		pw <- a * a
	}
	close(pw)
}

func main() {
	ch := make(chan int, 5)
	pw := make(chan int, 5)

	go generator(ch)
	go pow(ch, pw)

	for a := range pw {
		fmt.Println(a)
	}
}
