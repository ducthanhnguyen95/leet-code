package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))

}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxProfit := 0

	for index := 0; index < len(prices); index++ {
		if prices[index] < minPrice {
			minPrice = prices[index]
		} else if prices[index]-minPrice > maxProfit {
			maxProfit = prices[index] - minPrice
		}
	}
	return maxProfit
}
