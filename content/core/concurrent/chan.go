// +build OMIT

package main

import "fmt"

func genID() chan int {
	id := 0
	ch := make(chan int)
	go func() {
		for {
			ch <- id
			id++
		}
	}()
	return ch
}

func main() {
	ch := genID()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
