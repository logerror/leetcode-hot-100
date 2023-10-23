package lc005_mid

func productExceptSelf(nums []int) []int {
	length := len(nums)
	// 初始化两个空数组 L 和 R。对于给定索引 i，L[i] 代表的是 i 左侧所有数字的乘积，R[i] 代表的是 i 右侧所有数字的乘积。
	L, R, ans := make([]int, length), make([]int, length), make([]int, length)
	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}
	for i := 0; i < length; i++ {
		ans[i] = L[i] * R[i]
	}

	return ans
}
