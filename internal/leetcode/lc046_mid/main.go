package lc046_mid

import "sort"

// 可能会超时
func findAnagrams(s string, p string) []int {
	ans := []int{}
	plen := len(p)
	bp := []byte(p)
	sort.Slice(bp, func(i, j int) bool {
		return p[i] < p[j]
	})

	p = string(bp)
	for left := 0; left <= len(s)-len(p); left++ {
		right := left + plen - 1
		sp := []byte(s[left : right+1])
		sort.Slice(sp, func(i, j int) bool {
			return sp[i] < sp[j]
		})

		if string(sp) == p {
			ans = append(ans, left)
		}
	}

	return ans
}

func findAnagrams2(s, p string) []int {
	var ans []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}
	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}

	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}

	return ans
}
