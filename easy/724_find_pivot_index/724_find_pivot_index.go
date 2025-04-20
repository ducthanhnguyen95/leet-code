package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}

//func findPivotIndex(nums []int) int {
//
//	for i := 0; i < len(nums); i++ {
//		var leftRightSum int
//		var rightLeftSum int
//		for j := i + 1; j < len(nums); j++ {
//			leftRightSum += nums[j]
//		}
//		for k := i - 1; k >= 0; k-- {
//			rightLeftSum += nums[k]
//		}
//		if i == 0 && leftRightSum == 0 {
//			return i
//		}
//		if i == len(nums)-1 && rightLeftSum == 0 {
//			return i
//		}
//		if leftRightSum == rightLeftSum {
//			return i
//		}
//
//	}
//	return -1
//}

func pivotIndex(nums []int) int {

	total := 0
	for _, value := range nums {
		total += value
	}

	left := 0

	for i, value := range nums {
		if left == total-left-value {
			return i
		}
		left += value
	}
	return -1
}
