// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

func main() {
	var ch interface{}

	ch = make(chan int)

	switch t := ch.(type) {
	case int:
		fmt.Println("This is int")
	case string:
		fmt.Println("This is string")
	case bool:
		fmt.Println("This is bool")
	default:
		st := fmt.Sprintf("%T", t)
		if st[:4] == "chan" {
			fmt.Println("This is channel")
		} else {
			fmt.Println("Idk what`s it: ", st)
		}
	}
}
