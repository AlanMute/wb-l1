// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog bro crew sun"

	sr := strings.Fields(s) // Массив из слов разделенные пробелом
	l := len(sr)
	for i := 0; i <= l/2; i++ {
		sr[i], sr[l-1-i] = sr[l-1-i], sr[i]
	}

	res := strings.Join(sr, " ")

	fmt.Print(res)
}
