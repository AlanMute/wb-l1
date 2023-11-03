//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type Contest struct {
	mp map[int]int
	mu sync.Mutex
}

func main() {
	var m Contest
	m.mp = make(map[int]int)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(i int) {

			m.mu.Lock()
			m.mp[i] += i * i
			m.mu.Unlock()
			wg.Done()

		}(i)
	}

	wg.Wait()

	for i, a := range m.mp {
		fmt.Println(i, a)
	}

}
