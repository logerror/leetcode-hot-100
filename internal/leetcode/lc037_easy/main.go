package lc037_easy

func twoSum(nums []int, target int) []int {
	vmap := make(map[int]int)
	res := []int{}
	for i, v := range nums {
		if n, ok := vmap[target-v]; ok {
			return []int{i, n}
		} else {
			vmap[v] = i
		}
	}

	return res
}
