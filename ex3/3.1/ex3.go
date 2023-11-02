// Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(2^2+3^2+4^2...) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Задача 3.1")

	arr := [5]int{2, 4, 6, 8, 10}

	var ans int // переменная результата

	var w sync.WaitGroup
	m := sync.Mutex{} //Мьютекс для безопасного доступа

	for _, i := range arr {
		w.Add(1)

		go func(a int) {
			//Блокируем доступ для других горутин
			m.Lock()
			ans += a * a
			m.Unlock() //Возвращаем доступ для других горутин

			w.Done()
		}(i)
	}

	w.Wait()

	fmt.Print(ans)

}
