// +build OMIT

package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	n := len(matrix)
	if n == 0 {
		return []int{}
	}
	m := len(matrix[0])
	steps := n * m

	var res []int
	direct := 0

	var i, j = 0, 0
	var up, right, down, left = -1, m, n, -1
	for ; steps > 0; steps-- {
		res = append(res, matrix[i][j])
		switch direct {
		case 0:
			j++
			if j == right {
				j, i, up, direct = j-1, i+1, up+1, 1
			}
		case 1:
			i++
			if i == down {
				j, i, right, direct = j-1, i-1, right-1, 2
			}
		case 2:
			j--
			if j == left {
				j, i, down, direct = j+1, i-1, down-1, 3
			}
		case 3:
			i--
			if i == up {
				j, i, left, direct = j+1, i+1, left+1, 0
			}
		}
	}

	return res
}

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(matrix))
}
