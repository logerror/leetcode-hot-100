package lc040_easy

// left左边不为0， right右边都是0
func moveZeroes(nums []int) {
	left, right := 0, 0
	for right = 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
