package main

import "fmt"

func reverseString(s string) string {
	s1 := []byte(s)
	left, right := 0, len(s1)-1
	for left < right {
		s1[left], s1[right] = s[right], s[left]
		left++
		right--
	}

	return string(s1)
}

func main() {
	fmt.Println(reverseString("hello"))
}
