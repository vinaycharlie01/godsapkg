package array

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

func Merge(nums1 []int, m int, nums2 []int, n int) {
	L, R := 0, 0
	var merged []int
	
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
	copy(nums1, merged)
}
