// +build OMIT

package main

import "fmt"

func hammingWeight(num uint32) int {
	var count = 0
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}

// 学自Java - Integer.bitCount
func hammingWeight2(i uint32) int {
	// HD, Figure 5-2
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	return int(i & 0x3f)
}

func main() {
	fmt.Println("expect 17 and get", hammingWeight(87654321))
	fmt.Println("expect 17 and get", hammingWeight2(87654321))
}
