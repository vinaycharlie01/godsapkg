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

func MaxProfit(prices []int) int {
	var total int
	if len(prices) == 0 {
		return 0
	}

	prev := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i]-prev > 0 {
			total += prices[i] - prev
		}
		prev = prices[i]

	}
	return total
}

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	// 7,1,5,3,6,4
	// [[0,-1] [0,]]

	for i := 1; i < n; i++ { 
		
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}
