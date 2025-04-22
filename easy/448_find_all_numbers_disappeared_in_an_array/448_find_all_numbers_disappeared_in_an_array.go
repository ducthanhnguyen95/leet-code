package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	nums2 := []int{1, 1, 1}
	fmt.Println(findDisappearedNumbers(nums))  // Output: [5 6]
	fmt.Println(findDisappearedNumbers(nums2)) // Output: [2 3]
}

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}

	var result []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
