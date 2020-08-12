// +build OMIT

package main

import (
	"fmt"
	"strings"
)

var dict = [...]int{
	1: 1e1, 2: 1e2, 3: 1e3, 4: 1e4, 5: 1e5, 6: 1e6, 7: 1e7,
	9: 1e8, 10: 1e9, 11: 1e10, 12: 1e11, 13: 1e12, 14: 1e13,
}

func printNumbers(n int) []string {
	var buf = make([]byte, n)
	var result = make([]string, dict[n])
	var k = 0
	helper(buf, 0, &result, &k)
	return result[1:]
}

func helper(buf []byte, cur int, result *[]string, k *int) {
	n := len(buf)
	if cur == n-1 {
		for i := 0; i < 10; i++ {
			buf[cur] = byte(i + '0')
			(*result)[*k] = strings.TrimLeft(string(buf), "0")
			*k++
		}
	} else {
		for i := 0; i < 10; i++ {
			buf[cur] = byte(i + '0')
			helper(buf, cur+1, result, k)
		}
	}
}

func main() {
	fmt.Println(printNumbers(1))
}
