// +build OMIT

package main

import (
	"fmt"
	"strconv"
)

func singleNumbers(nums []int) []int {
	var tmp int
	for i, v := range nums {
		if i == 0 {
			tmp = v
		} else {
			tmp ^= v
		}
	}
	var idx = findFirstBit1(tmp)
	var r1, r2 int
	for _, v := range nums {
		if isBit1(v, idx) {
			r1 ^= v
		} else {
			r2 ^= v
		}
	}

	return []int{r1, r2}
}

func findFirstBit1(or int) int {
	var idx = 0
	for idx < strconv.IntSize && (or>>idx)&1 == 0 {
		idx++
	}
	return idx
}

func isBit1(num, idx int) bool {
	return (num>>idx)&1 == 1
}

func main() {
	arr := []int{4, 1, 4, 6}
	fmt.Println("expect [1 6] and get", singleNumbers(arr))
}
