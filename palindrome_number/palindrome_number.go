package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))

}

func isPalindrome(x int) bool {
	if x <= 9 && x >= -9 {
		if x >= 0 {
			return true
		}
		return false
	}
	r := strconv.Itoa(x)
	return r == Reverse(r)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
