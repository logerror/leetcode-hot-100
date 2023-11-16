package lc032_easy

import "sort"

// 排序 之后取中间值一定为众数 因为目标数的个数一定大于n/2
func majorityElement(nums []int) int {
	sort.Ints(nums)
	index := len(nums) / 2
	return nums[index]
}

//投票
func majorityElement2(nums []int) int {
	var count, target int
	for _, num := range nums {
		if count == 0 {
			target = num
		}
		if num == target {
			count++
		} else {
			count--
		}
	}
	return target
}
