//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными
//параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func GetLength(a, b *Point) float64 {
	return math.Sqrt((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))
}

func main() {
	a := newPoint(1.1, 2.2)
	b := newPoint(3.3, 4.4)

	fmt.Println(GetLength(a, b))
}
