// +build OMIT

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[:3]

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	s2 = append(s2, 6)
	fmt.Println("s2 append 6:", s2)
	fmt.Println("s1", s1)

	var s3 []int
	var s4 = make([]int, 0)

	fmt.Println("s3 == nil:", s3 == nil)
	fmt.Println("s4 == nil:", s4 == nil)
}
