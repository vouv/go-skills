// +build OMIT

package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	var count = 0
	for _, c := range s {
		if c == ' ' {
			count++
		}
	}
	var result = make([]byte, len(s)+2*count)
	var k = 0
	for i := range s {
		if s[i] == ' ' {
			result[k] = '%'
			result[k+1] = '2'
			result[k+2] = '0'
			k += 3
		} else {
			result[k] = s[i]
			k++
		}
	}
	return string(result)
}

func main() {
	fmt.Println("expect: We%20are%20happy.")
	fmt.Println("get:", replaceSpace("We are happy."))
}
