// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	w := 0            // стена
	p := len(arr) - 1 // опорный элемент будет последним

	for i := range arr {
		if arr[i] < arr[p] { // если текущий элемент меньше опорного перемещаем его слева от стены
			arr[i], arr[w] = arr[w], arr[i]
			w++
		}
	}
	arr[w], arr[p] = arr[p], arr[w] // Ставим опорный элемент на место стены. Теперь это его окончательное место

	quicksort(arr[:w]) // проделываем те же действия для получившихся частей
	quicksort(arr[w+1:])

	return arr
}

func main() {
	arr := []int{2, 4, 7, 0, 9, 1, 8}
	fmt.Println(quicksort(arr))
}
