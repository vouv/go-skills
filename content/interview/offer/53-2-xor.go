// +build OMIT

package main

func missingNumber(nums []int) int {
	var tmp = len(nums)
	for i, v := range nums {
		tmp ^= v
		tmp ^= i
	}
	return tmp
}

func main() {

}
