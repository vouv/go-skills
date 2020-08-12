// +build OMIT

package main

func merge(A []int, m int, B []int, n int) {
	var i, j, k = m - 1, n - 1, m + n - 1

	for j >= 0 {
		if i >= 0 && A[i] > B[j] {
			A[k] = A[i]
			i--
		} else {
			A[k] = B[j]
			j--
		}
		k--
	}
}

func main() {

}
