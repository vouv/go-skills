// +build OMIT

package main

import (
	"fmt"
)

type Sorts interface {
	BubbleSort(arr []int)
	SelectionSort(arr []int)
	InsertionSort(arr []int)

	MergeSort(arr []int)
	QuickSort(arr []int)
	HeapSort(arr []int)
}

func main() {
	fmt.Println("vouv's doc for gopher")
}
