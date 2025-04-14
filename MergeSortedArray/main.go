package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[0]
		nums2 = nums2[1:]
	}
	sort.Ints(nums1)
	fmt.Println(nums1)
}

func main() {
	var n, m int = 3, 3
	nums1 := make([]int, n+m)
	nums2 := make([]int, n)

	for i := 0; i < len(nums1)-n; i++ {
		n := rand.Intn(10)
		if i != 0 && n > nums1[i-1] {
			nums1[i] = n
		} else if i != 0 {
			nums1[i] = n
			nums1[i], nums1[i-1] = nums1[i-1], nums1[i]

		}
	}

	for i := 0; i < len(nums2); i++ {
		n := rand.Intn(10)
		if i != 0 && n > nums2[i-1] {
			nums2[i] = n
		}
	}
	fmt.Println(nums1, nums2)
	merge(nums1, m, nums2, n)

}
