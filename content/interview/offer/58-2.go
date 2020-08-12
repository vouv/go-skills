// +build OMIT

package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	ln := len(s)
	// defensive
	if n+1 > ln {
		return ""
	}
	bs := []byte(s)

	reverse(bs)
	reverse(bs[ln-n:])
	reverse(bs[:ln-n])
	return string(bs)
}

// 反转
func reverse(arr []byte) {
	n := len(arr)
	lo, hi := 0, n-1
	for lo < hi {
		arr[lo], arr[hi] = arr[hi], arr[lo]
		lo++
		hi--
	}
}

func main() {
	fmt.Println("expect cdefgab and get", reverseLeftWords("abcdefg", 2))
}
