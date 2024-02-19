package lc041_mid

func maxArea(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	for left < right {
		temp := (right - left) * min(height[left], height[right])
		res = max(temp, res)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
