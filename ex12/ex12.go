//Имеется последовательность строк - (cat, cat, dog, cat, tree)
//создать для нее собственное множество.

package main

import "fmt"

func getSet(arr []string) []string {
	m := make(map[string]int)
	var ans []string
	for _, i := range arr {
		if _, ok := m[i]; !ok {
			ans = append(ans, i)
			m[i] = 1
		}
	}
	return ans
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(getSet(arr))
}
