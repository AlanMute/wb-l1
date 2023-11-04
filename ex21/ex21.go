//Реализовать паттерн «адаптер» на любом примере.

package main

import "fmt"

type day struct {
	temp int
}

func (d *day) GetDay() {
	fmt.Println("Здесь идут какие-то очень важные вычисления")
}

type Human interface {
	PrintName()
	TakeDay() // Допустим тут должно происходить то же что и в методе GetDay
}

type AdapterHuman struct { //адаптер, позволяющий использовать метод GetDay как TakeDay.
	day
	name string
}

func (a *AdapterHuman) PrintName() {
	fmt.Println(a.name)
}

func (a *AdapterHuman) TakeDay() { // производим "обертку" метода GetDay
	a.GetDay()
}

func someProg(h Human) {
	h.PrintName()
	h.TakeDay()
}

func main() {
	a := &AdapterHuman{
		day: day{
			temp: 8,
		},
		name: "Alan",
	}

	someProg(a)

}
