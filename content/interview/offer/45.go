// +build OMIT

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	ns := make([]string, len(nums))
	for i, v := range nums {
		ns[i] = strconv.Itoa(v)
	}
	sort.Slice(ns, func(a, b int) bool {
		return ns[a]+ns[b] < ns[b]+ns[a]
	})
	return strings.Join(ns, "")
}

func main() {
	fmt.Println(minNumber([]int{1, 2, 3, 4, 5, 0}))
}
