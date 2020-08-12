// +build OMIT

package main

import "fmt"

func strToInt(str string) int {
	bs := []byte(str)
	var n = len(bs)
	if n == 0 {
		return 0
	}
	var i = 0
	// 去掉开头空格
	for i < n && bs[i] == ' ' {
		i++
	}
	if i == n || !IsValid(bs[i]) {
		return 0
	}

	var minus bool
	if bs[i] == '-' || bs[i] == '+' {
		minus = bs[i] == '-'
		i++
	}

	var sum = 0
	for i < n && IsDigit(bs[i]) {
		sum = sum*10 + int(bs[i]-'0')
		if sum > math.MaxInt32 {
			break
		}
		i++
	}

	if minus {
		sum = -sum
	}

	if sum > math.MaxInt32 {
		sum = math.MaxInt32
	} else if sum < math.MinInt32 {
		sum = math.MinInt32
	}
	return sum
}

func IsValid(c byte) bool {
	return '0' <= c && c <= '9' || c == '+' || c == '-'
}

func IsDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func main() {
	fmt.Println("expect -91283472332 and get", strToInt("91283472332"))
}
