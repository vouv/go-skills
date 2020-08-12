// +build OMIT

package main

import "fmt"

func isMatch(s string, p string) bool {
	if len(s) == 0 {
		return len(p) == 0 || len(p) > 1 && p[1] == '*' && isMatch(s, p[2:])
	}

	if len(p) == 0 {
		return false
	} else if len(p) > 1 && p[1] == '*' {
		if p[0] == s[0] || p[0] == '.' {
			return isMatch(s[1:], p[2:]) || isMatch(s[1:], p) || isMatch(s, p[2:])
		} else {
			return isMatch(s, p[2:])
		}
	}
	if s[0] == p[0] || p[0] == '.' {
		return isMatch(s[1:], p[1:])
	}
	return false
}

func main() {
	var s = "aaaaaaaaaaaaab"
	var p = "a*a*a*a*a*a*a*a*a*a*b"
	fmt.Println("match:")
	fmt.Println("\t", s, "\n\t", p, ":", isMatch(s, p))
}
