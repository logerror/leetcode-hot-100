package lc034_easy

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > res {
			res = prices[i] - minPrice
		}
	}
	return res
}
