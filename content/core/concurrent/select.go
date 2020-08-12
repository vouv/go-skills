// +build OMIT

package main

import (
	"fmt"
	"math/rand"
)

func GenInt(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- rand.Intn(99):
			case <-done:
				close(ch)
				return
			}
		}
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenInt(done)

	fmt.Println("guess is 23 and get", <-ch) // always 23

	done <- struct{}{}
	for v := range ch {
		fmt.Println(v) // never reached
	}
}
