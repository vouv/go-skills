// +build OMIT

package main

import (
	"fmt"
	"unsafe"
)

type Demo struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}

func main() {
	info(Demo{})
}

func info(demo Demo) {
	fmt.Printf("bool size: %d, align: %d, offset: %d\n",
		unsafe.Sizeof(demo.a),
		unsafe.Alignof(demo.a),
		unsafe.Offsetof(demo.a))
	fmt.Printf("int32 size: %d, align: %d, offset: %d\n",
		unsafe.Sizeof(demo.b), unsafe.Alignof(demo.b), unsafe.Offsetof(demo.b))
	fmt.Printf("int8 size: %d, align: %d, offset: %d\n",
		unsafe.Sizeof(demo.c), unsafe.Alignof(demo.c), unsafe.Offsetof(demo.c))
	fmt.Printf("int64 size: %d, align: %d, offset: %d\n",
		unsafe.Sizeof(demo.d), unsafe.Alignof(demo.d), unsafe.Offsetof(demo.d))
	fmt.Printf("byte size: %d, align: %d, offset: %d\n",
		unsafe.Sizeof(demo.e), unsafe.Alignof(demo.e), unsafe.Offsetof(demo.e))
	fmt.Printf("demo size: %d, align: %d\n",
		unsafe.Sizeof(demo), unsafe.Alignof(demo))

	fmt.Printf("struct size: %d, align: %d\n",
		unsafe.Sizeof(demo), unsafe.Alignof(demo))
}
