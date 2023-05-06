package elotusteam

import (
	"fmt"
	"testing"
)

func TestFindLength(t *testing.T) {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	result := findLength(nums1, nums2)
	fmt.Println(result) // Output: 3

	nums1 = []int{1, 2, 3, 2, 1}
	nums2 = []int{3, 2, 1, 4, 7}
	result = findLength(nums1, nums2)
	fmt.Println(result) // Output: 3

	nums1 = []int{0, 0, 0, 0, 0}
	nums2 = []int{0, 0, 0, 0, 0}
	result = findLength(nums1, nums2)
	fmt.Println(result) // Output: 5
}

func TestFindNextInteger(t *testing.T) {
	fmt.Println(grayCode(2)) // Output: [0,1,3,2]

	fmt.Println(grayCode(1)) // Output: [0,1]
}
