// +build OMIT

package main

import "fmt"

func translateNum(num int) int {
	if num < 100 {
		if 10 <= num && num < 26 {
			return 2
		}
		return 1
	}
	v := num % 100
	if 10 <= v && v < 26 {
		return translateNum(num/10) + translateNum(num/100)
	}
	return translateNum(num / 10)
}

func main() {
	fmt.Println("expect 2 and get", translateNum(1068385902))
}
