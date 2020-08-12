// +build OMIT

package main

import (
	"fmt"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	// 是否出现过数字、点、e
	numFlag := false
	dotFlag := false
	eFlag := false

	for i := range s {
		c := s[i]
		if (c == '+' || c == '-') && (i == 0 || s[i-1] == 'e') {
			numFlag = false
			continue
		} else if c >= '0' && c <= '9' {
			numFlag = true
		} else if c == '.' && !dotFlag && !eFlag {
			dotFlag = true
		} else if c == 'e' && !eFlag && numFlag {
			eFlag = true
			// e后必须有num
			numFlag = false
		} else {
			return false
		}
	}
	return numFlag
}

func main() {
	var num = " 2e0"
	fmt.Println("isNumber(\""+num+"\"):", isNumber(num))
}
