package lc038_easy

import "sort"

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	vmap := make(map[string][]string)
	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		sortStr := string(chars)
		if temp, ok := vmap[sortStr]; ok {
			temp = append(temp, str)
			vmap[sortStr] = temp
		} else {
			vmap[sortStr] = []string{str}
		}
	}

	for _, v := range vmap {
		res = append(res, v)
	}
	return res
}

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}
