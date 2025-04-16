package main

import "fmt"

func FindMaxInArray(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return max(FindMaxInArray(array))
}
func main() {
	array := [5]int{12, 4, 56, 7, 8}
	fmt.Println(FindMaxInArray(array[:]))
}
