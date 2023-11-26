package lc047_mid

//package main

import "fmt"

// 耗时较长的解法
func subarraySum(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			ans++
		}
		right := i + 1
		sum := nums[i]
		for right < len(nums) {
			sum = sum + nums[right]
			if sum == k {
				ans++
			}
			right++
		}
	}

	return ans
}

func subarraySum2(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end > 0; end-- {
			sum = sum + nums[end]
			if sum == k {
				count++
			}
		}

	}
	return count
}

//前缀和 + 哈希表优化
func subarraySum3(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		pre = pre + nums[i]
		if _, ok := m[pre-k]; ok {
			count = count + m[pre-k]
		}
		m[pre] = m[pre] + 1
	}

	return count
}

func main() {
	fmt.Println(subarraySum([]int{0, 0}, 0))
}
