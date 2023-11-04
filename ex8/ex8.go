//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"log"
	"strconv"
)

func IntToBits(num int64) string {
	return strconv.FormatInt(int64(num), 2)
}

func setBit(a *int64, pos uint, b uint) { // метод уставновки нужного бита в нужную позицию
	pos = 63 - pos
	fmt.Println(IntToBits(*a)) // Вывод исходного числа в двоичном формате
	if b == 1 {                // уставновка бита в единицу путем наложения маски и битовой операции ИЛИ
		*a |= int64(1) << pos // Сдвиг до нужной позиции
	} else if b == 0 { // уставновка бита в 0 путем наложения маски и битовой операции И НЕ
		*a &^= int64(1) << pos
	} else {
		log.Fatal("Бит только 0 или 1")

	}
	fmt.Println(IntToBits(*a)) // Вывод полученного числа в двоичном формате
	fmt.Println(*a)
}

func main() {
	var a int64 = 12
	setBit(&a, 41, 1)
}
