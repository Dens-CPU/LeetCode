// Работа программы основана на принципе чисел Фибоначи
package main

import "fmt"

func climbStairs(n int) int {
	firstLastNumber := 1
	secondLastNumber := 1

	for i := 2; i <= n; i++ {
		secondLastNumber, firstLastNumber = secondLastNumber+firstLastNumber, secondLastNumber
	}
	return secondLastNumber
}

func main() {
	fmt.Println(climbStairs(5))
}
