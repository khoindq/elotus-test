package main

import (
	"fmt"
)

func findLength(nums1 []int, nums2 []int) int {
	length_nums1 := len(nums1)
	length_nums2 := len(nums2)
	d1 := make([]int, length_nums2+1)
	res := 0
	for i := length_nums1 - 1; i >= 0; i-- {
		for j := 0; j < length_nums2; j++ {
			if nums1[i] == nums2[j] {
				d1[j] = d1[j+1] + 1
				if d1[j] > res {
					res = d1[j]
				}
			} else {
				d1[j] = 0
			}
		}
	}
	return res
}

func grayCode(n int) []int {
	//edge case
	if n <= 0 {
		return []int{0}
	}

	if n == 1 {
		return []int{0, 1}
	}
	//end of edge case

	s1 := grayCode(n - 1)

	s2 := make([]int, len(s1)*2)
	copy(s2, s1)
	for i := 0; i < len(s1); i++ {
		s2[len(s1)+i] = 1<<uint(n-1) + s1[len(s1)-1-i]
	}
	return s2
}

func main() {
	fmt.Println("type go test")

}
