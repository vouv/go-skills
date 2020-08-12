// +build OMIT

package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}

	sm := m - 1
	sn := 0
	for sm >= 0 && sn < n {
		v := matrix[sn][sm]
		if v == target {
			return true
		} else if v > target {
			sm--
			continue
		} else {
			sn++
		}
	}
	return false
}

func main() {
	var matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	fmt.Println("expect true, get", findNumberIn2DArray(matrix, 5))
	fmt.Println("expect false, get", findNumberIn2DArray(matrix, 20))

}
