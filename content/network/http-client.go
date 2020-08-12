// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://baidu.com")
	fmt.Println(resp.Status)
}
