// Написать программу, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import "fmt"

func powtwo(a int, c chan int) {
	c <- a * a
}

func main() {
	fmt.Println("Задача 2.1")

	arr := [5]int{2, 4, 6, 8, 10}

	ch := make(chan int)
	defer close(ch)

	for _, i := range arr {
		go powtwo(i, ch)
	}

	for range arr {
		fmt.Println(<-ch)
	}

}
