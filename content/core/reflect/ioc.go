package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

// 传入一个int参数类型的函数，使用随机输入函数对其进行调用
func Invoke(f interface{}) {
	ft := reflect.TypeOf(f)

	params := make([]reflect.Value, ft.NumIn())

	for i := 0; i < ft.NumIn(); i++ {
		paramType := ft.In(i)
		pv := reflect.New(paramType).Elem()
		pv.Set(reflect.ValueOf(rand.Intn(20)))
		params[i] = pv
	}

	// 调用
	reflect.ValueOf(f).Call(params)
}

func main() {
	show := func(a, b int) {
		fmt.Println("a=", a, "b=", b) // a=1 b=7
	}
	Invoke(show)
}
