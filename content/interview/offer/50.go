// +build OMIT

package main

import "fmt"

func firstUniqChar(s string) byte {
	var mp = make([]int, 26)
	for i := 0; i < len(s); i++ {
		mp[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if mp[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}

func main() {
	fmt.Println(firstUniqChar("lovevouvgoskills"))

}
