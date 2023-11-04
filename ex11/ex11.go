//Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5, 4}
	a2 := []int{2, 4, 4, 10, 1, 1}

	fmt.Println(cross(a1, a2)) // [2 4 4 1]
}

func cross(a []int, b []int) []int {
	m := make(map[int]int)
	var ans []int

	for _, i := range a {
		m[i] += 1
	}

	for _, i := range b {
		if n, ok := m[i]; ok && n > 0 { // Если число было в первом множестве то это пересечение
			ans = append(ans, i)
			m[i]--
		}
	}

	return ans
}
