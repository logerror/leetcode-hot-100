package lc042_mid

//package main

import "sort"

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	sort.Ints(nums)
	res := [][]int{}
	for first := 0; first < numsLen; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := numsLen - 1
		target := -nums[first]
		for second := first + 1; second < numsLen; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			sum := nums[first] + nums[second]

			for second < third && nums[third]+nums[second] > target {
				third--
			}

			if second == third {
				continue
			}

			if nums[third] == -sum {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}

		}
	}

	return res
}

func main() {
	threeSum([]int{-1, -1, 0, 1})
}
