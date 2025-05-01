package main

import "fmt"

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3, 1}))

}

func getConcatenation(nums []int) []int {
	n := len(nums)
	arr := make([]int, 2*n)

	for i := 0; i < n; i++ {
		arr[i] = nums[i]
		arr[i+n] = nums[i]
	}

	return arr
}
