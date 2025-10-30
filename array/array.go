package array

import (
	"slices"

	"golang.org/x/exp/constraints"
)

// RemoveDuplicates removes duplicates from a sorted slice of any comparable type
func RemoveDuplicates[T comparable](nums []T) int {
	if len(nums) == 0 {
		return 0
	}

	L := 0
	for R := 1; R < len(nums); R++ {
		if nums[R] != nums[L] {
			L++
			nums[L] = nums[R]
		}
	}
	return L + 1
}

// Rotate rotates a generic slice to the right by k steps.
func Rotate[T any](items []T, k int) {
	n := len(items)
	if n == 0 {
		return
	}
	k = k % n
	if k == 0 {
		return
	}

	// Reverse the entire slice
	slices.Reverse(items)

	// Reverse the first k elements
	slices.Reverse(items[:k])

	// Reverse the remaining elements
	slices.Reverse(items[k:])
}

func SumDivisibleByK(nums []int, k int) int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}
	sum := 0
	for n, count := range hash {
		if count%k == 0 {
			sum += n * count
		}
	}
	return sum
}

func Merge[T constraints.Ordered](nums1 []T, m int, nums2 []T, n int) []T {
	L, R := 0, 0
	var merged []T
	for L < m && R < n {
		if nums1[L] < nums2[R] {
			merged = append(merged, nums1[L])
			L++
		} else {
			merged = append(merged, nums2[R])
			R++
		}
	}
	for L < m {
		merged = append(merged, nums1[L])
		L++
	}
	for R < n {
		merged = append(merged, nums2[R])
		R++
	}

	return merged
}
