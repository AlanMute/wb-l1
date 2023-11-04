//Реализовать собственную функцию sleep.

package main

import (
	"context"
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	<-time.After(t)
}

func sleepTwo(t time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	<-ctx.Done()
	cancel()
}

func main() {
	// 1 способ
	fmt.Println("Begin Sleep")
	sleep(time.Second * 2)
	fmt.Println("Hello")

	fmt.Println()

	// 2 способ
	fmt.Println("Begin Sleep")
	sleepTwo(time.Second * 2)
	fmt.Println("Hello")

}
