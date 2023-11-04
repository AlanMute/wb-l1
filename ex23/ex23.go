// //Удалить i-ый элемент из слайса.

package main

import "fmt"

func delOfIndex(arr []int, idx int) []int {
	if idx < 0 || idx >= len(arr) {
		return arr // Если индекс недопустим, возвращаем исходный срез.
	}

	arr = append(arr[:idx], arr[idx+1:]...)
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(delOfIndex(arr, 3))
}
