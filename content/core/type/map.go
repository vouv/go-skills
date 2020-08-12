// +build OMIT

package main

import "fmt"

var dict = map[rune]rune{
	'a': 'n', 'b': 'o', 'c': 'p', 'd': 'q', 'e': 'r', 'f': 's', 'g': 't',
	'h': 'u', 'i': 'v', 'j': 'w', 'k': 'x', 'l': 'y', 'm': 'z', 'n': 'a',
	'o': 'b', 'p': 'c', 'q': 'd', 'r': 'e', 's': 'f', 't': 'g', ' ': ' ',
	'u': 'h', 'v': 'i', 'w': 'j', 'x': 'k', 'y': 'l', 'z': 'm', '!': '!',
}

func main() {
	for _, c := range "uryyb jbeyq!" {
		fmt.Printf("%c", dict[c])
	}
	fmt.Println()
}
