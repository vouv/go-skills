// +build OMIT

package main

import (
	"fmt"
	"sync"
)

// 正确方式
func main() {
	wg := sync.WaitGroup{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := range arr {
		wg.Add(1)
		go func(j int) {
			fmt.Println(j)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
