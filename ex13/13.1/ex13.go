//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func swap(a *int, b *int) {
	*a = *a * *b
	*b = *a / *b
	*a = *a / *b
}

func main() {
	a := 1
	b := 2

	fmt.Println(a, b)
	fmt.Println("Swap!")

	swap(&a, &b)

	fmt.Println(a, b)
}
