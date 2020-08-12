// +build OMIT

package main

import "fmt"

func twoSum(n int) []float64 {
	if n < 1 {
		return nil
	}
	var dp = make([][]int, 12)
	for i := range dp {
		dp[i] = make([]int, 6*n+1)
	}
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	// n 个色子
	for i := 2; i <= n; i++ {
		// 可能的点数
		for j := i; j <= 6*n; j++ {
			// 最后一个骰子的点数
			for k := 1; k <= 6; k++ {
				if j-k < 1 {
					break
				}
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}
	sum := pow(6, n)
	var result = make([]float64, 5*n+1)
	for i := n; i <= 6*n; i++ {
		result[i-n] = float64(dp[n][i]) / float64(sum)
	}
	return result
}

func pow(a, n int) int {
	var res = 1
	for i := 0; i < n; i++ {
		res *= a
	}
	return res
}

func main() {
	fmt.Println(twoSum(5))
}
