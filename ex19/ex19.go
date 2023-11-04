// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s *string) {
	l := utf8.RuneCountInString(*s)
	r := []rune(*s)
	for i := 0; i <= l/2; i++ {
		r[i], r[l-1] = r[l-1], r[i]
	}
	*s = string(r)
}

func main() {
	var s string
	if _, err := fmt.Scan(&s); err != nil {
		fmt.Println(err)
		return
	}

	reverse(&s)

	fmt.Println(s)
}
