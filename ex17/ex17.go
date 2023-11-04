// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
)

func binsearch(arr []int, s int) int {
	p := len(arr) / 2 // опорный элемент
	f := 0            // первый
	l := len(arr) - 1 // последний

	for {
		if arr[p] == s { // если нашли - вернем идекс
			return p
		} else if p == l { // если не нашли и при этом дошли до конца - элемента нет
			return -1
		}
		if s < arr[p] {
			l = p // индекс последнего элемента сместился
		} else {
			f = p // индекс первого элемента сместился
		}
		p = (f + l + 1) / 2
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 6, 10, 11} // массив должен быть упорядочен изначально
	fmt.Println(arr)
	fmt.Println(binsearch(arr, 6))
}
