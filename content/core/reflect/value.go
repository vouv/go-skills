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

func (u *User) String() {
	fmt.Println("User", u.Id, u.Name, u.Age)
}

func Info(o interface{}) {
	v := reflect.ValueOf(o)
	t := v.Type()

	fmt.Println("Type:", t.Name())

	fmt.Println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		fid := t.Field(i)
		fmt.Print("\tField[", i, "]:", fid.Name, " ")
		val := v.Field(i).Interface()
		switch value := val.(type) {
		case int:
			fmt.Println(fid.Type, "=", value)
		case string:
			fmt.Println(fid.Type, "=", value)
		default:
			fmt.Println("value type not int or string")
		}
	}

}

func main() {
	u := User{
		Id:   666,
		Name: "vouv",
		Age:  18,
	}
	Info(u)
}
