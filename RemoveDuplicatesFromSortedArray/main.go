package main

import "fmt"

func removeDuplicates(nums []int) int {
	var k int
	if len(nums) == 0 {
		return 0
	}
	for j := k + 1; j < len(nums); j++ {
		if nums[k] == nums[j] {
			continue
		}
		k++
		nums[k] = nums[j]
	}
	return k + 1
}
func main() {
	nums := []int{1, 2, 2, 2, 3, 3, 4} // Входной массив
	expectedNums := []int{1, 2, 3, 4}  // Ожидаемый ответ с правильной длиной

	var k = removeDuplicates(nums) // Вызывает вашу реализацию

	if k == len(expectedNums) {
		for i := 0; i < k; i++ {
			if nums[i] == expectedNums[i] {
				fmt.Println("Равно")
			}
		}
	} else {
		fmt.Println("Длины не равны")
	}
}
