// +build OMIT

package main

import "fmt"

func reversePairs(nums []int) int {
	return mergeSort(nums)
}

func mergeSort(arr []int) int {
	if len(arr) < 2 {
		return 0
	}
	mid := len(arr) / 2
	mergeSort(arr[:mid])
	mergeSort(arr[mid:])
	var sum int
	for i, v := range merge(arr[:mid], arr[mid:], &sum) {
		arr[i] = v
	}
	return sum
}

func merge(arr1, arr2 []int, sum *int) []int {
	var result = make([]int, len(arr1)+len(arr2))
	var i, j int
	for k := range result {
		fmt.Println(i, len(arr1))
		if i == len(arr1) || arr1[i] > arr2[j] {
			result[k] = arr2[j]
			*sum += len(arr2) - j
			j++
		} else {
			result[k] = arr1[i]
			*sum += len(arr1) - i
			i++
		}
	}

	return result
}

func main() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}
