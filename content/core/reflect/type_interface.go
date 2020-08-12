// +build OMIT

package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"1111" b:"2222"`
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")
	if ok {
		fmt.Println("field Name's Tag:", fieldName.Tag)
	}
	fieldAge, ok := rt.FieldByName("Age")
	if ok {
		fmt.Println("field Age's Tag[a]:", fieldAge.Tag.Get("a"))
		fmt.Println("field Age's Tag[b]:", fieldAge.Tag.Get("b"))
	}

	fmt.Println("rt.Name:", rt.Name())
	fmt.Println("rt.NumField:", rt.NumField())
	fmt.Println("rt.PkgPath:", rt.PkgPath())
	fmt.Println("rt.String:", rt.String())

	fmt.Println("rt.Kind:", rt.Kind())
	fmt.Println("rt.Kind.String:", rt.Kind().String())

	for i := 0; i < rt.NumField(); i++ {
		fmt.Println("rt.Field[", i, "]:", rt.Field(i).Name, rt.Field(i).Type.Name())
	}
}
