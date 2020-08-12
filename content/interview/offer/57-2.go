// +build OMIT

package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	var result = [][]int{}
	if target < 2 {
		return result
	}
	n := (target + 1) / 2

	lo := 1
	hi := 2
	sum := lo + hi

	for hi < n {
		if sum == target {
			arr := make([]int, hi-lo+1)
			for i := range arr {
				arr[i] = i + lo
			}
			result = append(result, arr)
			hi++
			sum += hi
		} else if sum > target {
			sum -= lo
			lo++
		} else {
			hi++
			sum += hi
		}
	}
	return result
}

func main() {
	fmt.Println(findContinuousSequence(9))
}
