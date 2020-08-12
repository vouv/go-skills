// +build OMIT

package main

import "fmt"

func movingCount(m int, n int, k int) int {
	var visit = make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	return bfsSolver(visit, 0, 0, m, n, k)
}

func bfsSolver(visit [][]bool, i, j, m, n, k int) int {
	if check(visit, i, j, m, n, k) {
		visit[i][j] = true
		return 1 + bfsSolver(visit, i+1, j, m, n, k) +
			bfsSolver(visit, i-1, j, m, n, k) +
			bfsSolver(visit, i, j+1, m, n, k) +
			bfsSolver(visit, i, j-1, m, n, k)
	} else {
		return 0
	}
}

func check(visit [][]bool, i, j, m, n, k int) bool {
	return i >= 0 && j >= 0 && i < m && j < n && !visit[i][j] && getDigitSum(i)+getDigitSum(j) <= k
}

func getDigitSum(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10
	}
	return res
}

func main() {
	fmt.Println("expect 3 and get", movingCount(2, 3, 1))
}
