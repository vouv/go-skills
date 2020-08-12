// +build OMIT

package main

import "fmt"

func permutation(s string) []string {
	var res []string
	dfs([]byte(s), 0, &res)
	return res
}

func dfs(s []byte, cur int, res *[]string) {
	if cur == len(s) {
		*res = append(*res, string(s))
		return
	}
	used := map[byte]bool{}
	for i := cur; i < len(s); i++ {
		if used[s[i]] {
			continue
		}
		s[i], s[cur] = s[cur], s[i]
		dfs(s, cur+1, res)
		s[i], s[cur] = s[cur], s[i]
		used[s[i]] = true
	}
}

func main() {
	fmt.Println(permutation("abc"))
}
