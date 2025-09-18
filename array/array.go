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
