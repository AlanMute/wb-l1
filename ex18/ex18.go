//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type count struct {
	num int
	m   sync.Mutex
}

func main() {
	var wg sync.WaitGroup

	var c count
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 6; i++ {
				c.m.Lock()
				c.num++
				c.m.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(c.num)
}
