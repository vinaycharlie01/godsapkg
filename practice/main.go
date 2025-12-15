package main

import "fmt"

func RemoveDuplicates(nums []int) int {
	var count int
	hash := map[int]int{}
	var arr []int
	for _, v := range nums {
		if val, ok := hash[v]; ok && val > 1 {
			continue
		} else {
			count++
			hash[v]++
			arr = append(arr, v)
		}
	}
	fmt.Println(hash)
	copy(nums, arr)
	return count
}

func main() {
	arr := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(RemoveDuplicates(arr))
	fmt.Println(arr)

}
