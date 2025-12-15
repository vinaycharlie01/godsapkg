package array

import (
	"slices"
	"sort"

	"github.com/vinaycharlie01/godsapkg/array/model"
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

// SingleElement returns the single non-duplicate element from a sorted slice.
// Example: [1,1,2,2,3] -> 3
func SingleElement[T constraints.Ordered](items []T) T {
	if len(items) == 0 {
		var zero T
		return zero // return zero value if empty slice
	}

	slices.Sort(items)

	for i := 0; i < len(items)-1; i += 2 {
		if items[i] != items[i+1] {
			return items[i]
		}
	}

	// If all pairs matched, last one is the single element
	return items[len(items)-1]
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

func MaxProfit[T model.Number](prices []T) T {
	if len(prices) == 0 {
		return 0
	}
	prevNumber := prices[0]
	var maxProfit T
	for i := 1; i < len(prices); i++ {
		if prices[i]-prevNumber > 0 {
			maxProfit += prices[i] - prevNumber
		}
		prevNumber = prices[i]
	}
	return maxProfit
}

func Intersect[T constraints.Ordered](element1 []T, element2 []T) []T {
	var result []T
	hash := map[T]int{}
	for _, v := range element1 {
		hash[v]++
	}
	for _, v := range element2 {
		if hash[v] > 0 {
			result = append(result, v)
			hash[v]--
		}
	}
	return result
}

func Insert(intervals [][]int, newInterval []int) [][]int {
	finalSlice := [][]int{}
	intervals = append(intervals, newInterval)
	sortSliceUsingStart(intervals)

	start, end := intervals[0][0], intervals[0][1]
	for idx := 1; idx < len(intervals); idx++ {
		currentStart, currentEnd := intervals[idx][0], intervals[idx][1]
		if currentStart <= end {
			if currentEnd > end {
				end = currentEnd
			}
		} else {
			finalSlice = append(finalSlice, []int{start, end})
			start, end = currentStart, currentEnd
		}
	}

	finalSlice = append(finalSlice, []int{start, end})
	return finalSlice
}

func sortSliceUsingStart(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	return intervals
}
