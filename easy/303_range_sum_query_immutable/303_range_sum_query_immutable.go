package main

import "fmt"

func main() {

	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)

	fmt.Printf("SumRange(0, 2) = %d  (expected 1)\n", obj.SumRange(0, 2))
	fmt.Printf("SumRange(2, 5) = %d  (expected -1)\n", obj.SumRange(2, 5))
	fmt.Printf("SumRange(0, 5) = %d  (expected -3)\n", obj.SumRange(0, 5))
}

//type NumArray struct {
//	nums  []int
//	cache []int
//}
//
//func Constructor(nums []int) NumArray {
//	sum := 0
//	cache := make([]int, len(nums))
//
//	for i := 0; i < len(nums); i++ {
//		sum += nums[i]
//		cache[i] = sum
//	}
//
//	return NumArray{
//		nums:  nums,
//		cache: cache,
//	}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	if left == 0 {
//		return this.cache[right]
//	}
//
//	return this.cache[right] - this.cache[left] + this.nums[left]
//}

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}
	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}
