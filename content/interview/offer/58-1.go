// +build OMIT

package main

import (
	"fmt"
)

func reverseWords(s string) string {
	bs := clearSpace([]byte(s))
	n := len(bs)
	reverse(bs)
	i, j := 0, 0
	for i < n {
		if j == n || bs[j] == ' ' {
			reverse(bs[i:j])
			i = j + 1
			j = i
		} else {
			j++
		}
	}
	return string(bs)
}

// 空格处理
func clearSpace(bs []byte) []byte {
	n := len(bs)
	var i, flag = 0, true
	for j := 0; j < n; j++ {
		if bs[j] != ' ' || !flag {
			bs[i] = bs[j]
			flag = bs[j] == ' '
			i++
		}
	}
	// 串尾空格
	if i > 0 && flag {
		i -= 1
	}
	return bs[:i]
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
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println(reverseWords("       a good     example     "))

}
