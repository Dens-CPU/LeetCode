package main

import (
	"fmt"
	"strings"
)

// func strStr(haystack string, needle string) int {
// 	var k int
// 	last := len(needle) - 1
// 	for k < len(haystack)-last {
// 		p := 0
// 		for p <= last && haystack[k+p] == needle[p] {
// 			p++
// 		}
// 		if p == len(needle) {
// 			return k
// 		}
// 		k++
// 	}
// 	return -1

// }

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func main() {
	haystack := "sado"
	needle := "sad"
	fmt.Println(strStr(haystack, needle))
}
