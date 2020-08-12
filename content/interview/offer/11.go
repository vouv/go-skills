// +build OMIT

package main

import "fmt"

func minArray(numbers []int) int {
	n := len(numbers)

	lo, hi := 0, n-1

	for lo < hi {
		mid := lo + (hi-lo)>>1
		if numbers[mid] < numbers[hi] {
			hi = mid
		} else if numbers[mid] > numbers[hi] {
			lo = mid + 1
		} else {
			// mid 与右边相等，不能确定mid在左还是在右，但是hi能排除
			hi--
		}
	}

	return numbers[lo]
}

func main() {
	fmt.Println("expect 1 and get", minArray([]int{3, 3, 3, 3, 3, 1, 3, 3, 3, 3, 3, 3, 3}))
}
