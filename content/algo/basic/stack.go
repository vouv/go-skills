// +build OMIT

package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Pop() int {
	n := len(s.data)
	if n == 0 {
		panic("pop from empty stack")
	}
	defer func() {
		s.data = s.data[:n-1]
	}()
	return s.data[n-1]
}

func (s *Stack) Push(val ...int) {
	s.data = append(s.data, val...)
	return
}

func main() {
	stack := Stack{}
	stack.Push(1, 2, 3)
	fmt.Println("stack pop:", stack.Pop())
	fmt.Println("stack pop:", stack.Pop())
	fmt.Println("stack len:", stack.Len())
}
