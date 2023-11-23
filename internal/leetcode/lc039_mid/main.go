//package lc039_mid
package main

func longestConsecutive(nums []int) int {
	res := 0
	var nmap = map[int]bool{}
	for _, num := range nums {
		nmap[num] = true
	}

	for k, _ := range nmap {
		if !nmap[k-1] {
			temp := 1
			n := k
			for nmap[n+1] {
				temp++
				n++
			}
			res = max(temp, res)
		}

	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	longestConsecutive([]int{100, 4, 200, 1, 3, 2})
}
