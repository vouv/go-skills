// +build OMIT

package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	var heap = make([]int, k)
	for i, v := range arr {
		if i < k {
			heap[i] = v
		}
		if i == k-1 {
			buildHeap(heap)
		} else if v < heap[0] {
			heap[0] = v
			heapify(heap, 0)
		}
	}
	return heap
}

func buildHeap(arr []int) {
	for i := len(arr) / 2; i > -1; i-- {
		heapify(arr, i)
	}
}

func heapify(arr []int, i int) {
	for {
		maxi := i
		if i*2+1 < len(arr) && arr[i*2+1] > arr[maxi] {
			maxi = i*2 + 1
		}
		if i*2+2 < len(arr) && arr[i*2+2] > arr[maxi] {
			maxi = i*2 + 2
		}
		if maxi == i {
			break
		}
		arr[maxi], arr[i] = arr[i], arr[maxi]
		i = maxi
	}
}

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}
