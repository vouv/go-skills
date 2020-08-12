// +build OMIT

package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Type of arr:", reflect.TypeOf(arr))
}
