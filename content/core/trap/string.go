// +build OMIT

package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "a"

	// str[i] byte
	// c rune
	for i, c := range str {
		fmt.Println("c in range:", c, reflect.TypeOf(c).Name())
		fmt.Println("str[i]:", str[i], reflect.TypeOf(str[i]).Name())
	}

	str = "我爱中国"
	fmt.Println("len(我爱中国) =", len(str))
	for i, c := range str {
		fmt.Println(i, c, string(c))
	}

}
