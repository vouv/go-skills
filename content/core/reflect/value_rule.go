// +build OMIT

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{
		Id:   1,
		Name: "vouv",
		Age:  18,
	}
	va := reflect.ValueOf(u)
	vb := reflect.ValueOf(&u)

	fmt.Println("User:", va.CanSet(), va.FieldByName("Name").CanSet())

	//vb.FieldByName("Name") // panic

	fmt.Println("&User:", vb.CanSet(), vb.Elem().FieldByName("Name").CanSet())
}
