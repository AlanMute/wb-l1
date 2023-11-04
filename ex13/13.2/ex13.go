//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 1
	b := 2

	fmt.Println(a, b)
	fmt.Println("Swap!")

	a, b = b, a

	fmt.Println(a, b)
}
