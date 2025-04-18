package main

import "fmt"

func main() {
	fmt.Println(generate(5))

}

func generate(numRows int) [][]int {
	var output [][]int
	for i := 0; i < numRows; i++ {
		var num []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				num = append(num, 1)
			} else {
				left := output[i-1][j-1]
				right := output[i-1][j]
				num = append(num, left+right)
			}
		}
		output = append(output, num)
	}
	return output

}
