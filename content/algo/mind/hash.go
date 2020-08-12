package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var raw = []byte("hello world")
var encode = hex.EncodeToString

func main() {
	MD5, SHA1, SHA256 := md5.New(), sha1.New(), sha256.New()
	MD5.Write(raw)
	SHA1.Write(raw)
	SHA256.Write(raw)

	fmt.Println(`MD5("hello world") =`, encode(MD5.Sum(nil)))
	fmt.Println(`SHA1("hello world") =`, encode(SHA1.Sum(raw)))
	fmt.Println(`SHA256("hello world") =`, encode(SHA256.Sum(raw)))
}
