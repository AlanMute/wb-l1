// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от
// родительской структуры Human (аналог наследования)

package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) hello() {
	fmt.Printf("Hello! I`m %s, and I see u\n", h.name)
}

func (h *Human) goodbye() {
	fmt.Println("See u!")
}

type Action struct {
	Human //Встраивание родительской структуры Human
}

func main() {
	fmt.Println("Задача 1")

	h := Human{
		name: "Alan",
		age:  19,
	}

	a := Action{
		Human: Human{name: "Ivan"},
	}

	h.hello()
	a.hello()   // Вызов встроенного метода
	a.goodbye() // Вызов встроенного метода

	// Result:
	// 		Hello! I`m Alan, and I see u
	// 		Hello! I`m Ivan, and I see u
	//		See u!
}
