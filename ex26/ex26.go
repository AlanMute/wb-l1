//Разработать программу, которая проверяет, что все символы в строке уникальные
//(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func isUniq(s string) bool {
	s = strings.ToLower(s) //строку в нижний регистр

	m := make(map[rune]int)

	for _, i := range s {
		m[i]++
		if m[i] > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isUniq("asdqwezxc"))
	fmt.Println(isUniq("asdqAezxc"))
	fmt.Println(isUniq("фывйцуяс"))
}
