package lc048_mid

import "awesomeProject1/pkg/util"

// dp
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = util.Max(res, dp[i])
	}

	return res
}

func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}

		if res < 0 {
			res = 0
		}
		p++
	}

	return maxSum
}
