package lc045_mid

//package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	ans := 0
	right := -1
	sLen := len(s)
	for i := 0; i < sLen; i++ {

		for right < sLen-1 && m[s[right+1]] == 0 {
			m[s[right+1]]++
			right++
		}

		ans = max(ans, right-i+1)

		m[s[i]]--
	}

	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
