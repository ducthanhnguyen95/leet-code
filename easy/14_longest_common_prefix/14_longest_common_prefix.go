package main

import "fmt"

func main() {

	test := []string{"flower", "flow", "flight"}
	//fmt.Println("Vertical:", longestCommonPrefixVertical(test))
	fmt.Println("Vertical:", longestCommonPrefixHorizontal(test))

}

// 1. Vertical Scanning
//func longestCommonPrefixVertical(strs []string) string {
//	if len(strs) == 0 {
//		return ""
//	}
//	for i := 0; i < len(strs[0]); i++ {
//		c := strs[0][i]
//		for j := 1; j < len(strs); j++ {
//			if i >= len(strs[j]) || strs[j][i] != c {
//				return strs[0][:i]
//			}
//		}
//	}
//	return strs[0]
//}

func longestCommonPrefixHorizontal(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && strs[i][:min(len(prefix), len(strs[i]))] != prefix {
			// shrink prefix by one
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
