// +build OMIT

package main

import "fmt"

// 堆排序
// 假设从0开始存储
func HeapSort(arr []int) {
	n := len(arr)

	// 构建大顶堆
	for i := (n - 1) / 2; i >= 0; i-- {
		siftDown(arr[i:n])
	}

	// 不断弹出堆顶元素
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		siftDown(arr[0:i])
	}
}

// 自上往下堆化
func siftDown(arr []int) {
	root := 0
	n := len(arr)
	for {
		child := 2 * root + 1 // 左子节点
		if child >= n {
			break
		}
		if child+1 < n && arr[child] < arr[child+1] { // 选取左右子节点大的节点
			child++ // 右子节点
		}
		if arr[root] >= arr[child] { // 根节点比左右子节点都大
			return
		}
		arr[root], arr[child] = arr[child], arr[root]
		root = child
	}
}

func main() {
	arr := []int{7,8,9,5,6,3,4,2,1}
	fmt.Println("before", arr)
	HeapSort(arr)
	fmt.Println("after", arr)
}
