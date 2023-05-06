package elotusteam

import (
	"math"
	"strconv"
)

func findLength(nums1 []int, nums2 []int) int {
	maxLength := 0
	for i1, num1 := range nums1 {
		tempMaxLength := 0
		for i2, num2 := range nums2 {
			if num1 == num2 {
				for u, v := i1, i2; u < len(nums1) && v < len(nums2); {
					if nums1[u] == nums2[v] {
						tempMaxLength++
						u++
						v++
					} else {
						break
					}
				}
				if maxLength < tempMaxLength {
					maxLength = tempMaxLength
				}
			}
			tempMaxLength = 0
		}
	}

	return maxLength
}

type GrayCode struct {
	Binary string
	Index  int64
}

// DiffOneBit check next integer differs by exactly one bit
func (g *GrayCode) DiffOneBit(o *GrayCode) bool {
	numDiff := 0
	for i := 0; i < len(g.Binary); i++ {
		if g.Binary[i] != o.Binary[i] {
			numDiff++
		}
		if numDiff > 1 {
			return false
		}
	}

	if numDiff == 1 {
		return true
	}

	return false
}

func (g *GrayCode) FullFillBinary(n int) {
	for i := 0; i < n-len(g.Binary); i++ {
		g.Binary = g.Binary + "0"
	}
}

func grayCode(n int) []int {
	// generate array
	grayCodeArr := make([]*GrayCode, 0, n)
	for i := int64(0); i <= int64(math.Pow(2, float64(n))-1); i++ {
		grayCodeArr = append(grayCodeArr, &GrayCode{
			Binary: strconv.FormatInt(i, 2),
			Index:  i,
		})
	}

	// normalize
	maxLength := len(grayCodeArr[len(grayCodeArr)-1].Binary)
	for i := range grayCodeArr {
		grayCodeArr[i].FullFillBinary(maxLength)
	}

	for i := 0; i < len(grayCodeArr); i++ {
		if findNextInteger(grayCodeArr, i) {
			break
		}
	}

	res := make([]int, 0, len(grayCodeArr))
	for _, v := range grayCodeArr {
		res = append(res, int(v.Index))
	}
	return res
}

func findNextInteger(grayCodeArr []*GrayCode, i int) bool {
	if i == len(grayCodeArr)-1 {
		// check the first and last integers
		if grayCodeArr[0].DiffOneBit(grayCodeArr[i]) {
			return true
		} else {
			return false
		}
	}

	for j := i + 1; j < len(grayCodeArr); j++ {
		if grayCodeArr[i].DiffOneBit(grayCodeArr[j]) {
			// swap
			grayCodeArr[i+1], grayCodeArr[j] = grayCodeArr[j], grayCodeArr[i+1]
			if findNextInteger(grayCodeArr, i+1) {
				return true
			}
			// roll back
			grayCodeArr[j], grayCodeArr[i+1] = grayCodeArr[i+1], grayCodeArr[j]
		}
	}

	return false
}
