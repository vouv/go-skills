// +build OMIT

package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[i] {
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(
		validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}),
	)
	fmt.Println(
		validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}),
	)
}
