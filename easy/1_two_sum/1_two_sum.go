package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[j] == target-nums[i] {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	numIndexes := make(map[int]int) // num -> index

	for i, num := range nums {
		diff := target - num
		if idx, ok := numIndexes[diff]; ok {
			return []int{i, idx}
		}
		numIndexes[num] = i
	}
	return nil
}
