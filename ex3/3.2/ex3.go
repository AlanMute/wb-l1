// Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(2^2+3^2+4^2...) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Задача 3.2")

	arr := [5]int{2, 4, 6, 8, 10}

	var ans int // переменная результата

	m := sync.Mutex{} //Мьютекс для безопасного доступа

	ch := make(chan int, 5)
	defer close(ch)

	// отправка квадратов в горутине по каналам
	for _, i := range arr {
		go func(a int) {
			ch <- a * a
		}(i)
	}

	for range arr {
		m.Lock()
		ans += <-ch
		m.Unlock()
	}

	fmt.Print(ans)

}
