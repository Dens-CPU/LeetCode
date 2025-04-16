package main

import "fmt"

func removeElement(nums []int, val int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{3, 2, 2, 3, 0, 4, 3, 3, 6, 7, 9}
	val := 3
	fmt.Println(removeElement(nums, val))

}
