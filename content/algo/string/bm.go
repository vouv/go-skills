package main

import "fmt"

type Bm struct {
	bc     []int
	suffix []int  // p串的后缀子串在p串中靠后的idx
	prefix []bool // p串的后缀子串是否匹配模式串的前缀
}

// 目前只考虑ascii值
func (b *Bm) genBC(p string) {
	const size = int(1 << 8)
	b.bc = make([]int, size)
	for i := range b.bc {
		b.bc[i] = -1
	}
	for i, c := range []byte(p) {
		b.bc[c] = i
	}
}

// 仅仅用坏字符规则
func (b *Bm) find1(s, p string) int {
	b.genBC(p)
	n, m := len(s), len(p)
	var i = 0
	for i <= n-m {
		var j int
		// 坏字符
		for j = m - 1; j >= 0; j-- {
			if s[i+j] != p[j] {
				break
			}
		}
		// 找到了
		if j < 0 {
			return i
		}

		// 按照坏字符规则，移动到j-bc[s[i]]
		i = i + (j - b.bc[s[i+j]])
	}
	return -1
}

// todo 要理解！
// 巧妙的求得子串与公共后缀的匹配
func (b *Bm) genGS(p string) {
	m := len(p)
	b.suffix = make([]int, m)
	for i := range b.suffix {
		b.suffix[i] = -1
	}
	b.prefix = make([]bool, m)
	for i := 0; i < m-1; i++ {
		var j = i
		var k = 0
		for j >= 0 && p[j] == p[m-1-k] {
			j--
			k++
			b.suffix[k] = j + 1
		}
		if j == -1 {
			b.prefix[k] = true
		}
	}
}

// 仅仅用坏字符规则
func (b *Bm) Find(s, p string) int {
	b.genBC(p)
	b.genGS(p)
	n, m := len(s), len(p)
	var i = 0
	for i <= n-m {
		var j int
		// 坏字符
		for j = m - 1; j >= 0; j-- {
			if s[i+j] != p[j] {
				break
			}
		}
		// 找到了
		if j < 0 {
			return i
		}

		// 按照坏字符规则，移动到j-bc[s[i]]
		var x = j - b.bc[s[i+j]]
		var y = 0
		if j < m-1 {
			y = b.moveByGS(j, m)
		}
		max := func(a, b int) int {
			if a > b {
				return a
			} else {
				return b
			}
		}

		i = i + max(x, y)

	}
	return -1
}

func (b *Bm) moveByGS(j, m int) int {
	// 好后缀长度
	k := m - 1 - j
	if b.suffix[k] != -1 {
		return j - b.suffix[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if b.prefix[m-r] {
			return r
		}
	}
	return m
}

func main() {
	bm := Bm{}
	s := "abcacabdc"
	p := "abd"
	fmt.Println(bm.Find(s, p))
}
