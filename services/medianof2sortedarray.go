package services

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, i2 int
	var arr []int
	for len(nums1) > i || len(nums2) > i2 {
		var one, two int
		if len(nums1) > i {
			one = nums1[i]
		} else {
			arr = append(arr, nums2[i2])
			i2++
			continue
		}
		if len(nums2) > i2 {
			two = nums2[i2]
		} else {
			arr = append(arr, nums1[i])
			i++
			continue
		}
		if one <= two {
			arr = append(arr, one)
			i++
		} else { // one > two
			arr = append(arr, two)
			i2++
		}
	}
	midpoint := len(arr) / 2 //4/2 = 2
	if len(arr)%2 == 0 {     //arr[2],arr[3]
		return float64(arr[midpoint]+arr[midpoint-1]) / 2
	} else {
		a := math.Round(float64(len(arr) / 2))
		return float64(arr[int(a)])
	}
}

func Min(nums1, nums2 int) int {
	if nums1 < nums2 {
		return nums1
	} else {
		return nums2
	}
}
