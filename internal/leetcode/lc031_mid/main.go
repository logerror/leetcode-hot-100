package lc031_mid

func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	result := ""
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	// length 2表示aba类型；1表示aa类型；0表示a类型
	for length := 0; length < n; length++ {
		for i := 0; i+length < n; i++ {
			j := i + length
			if length == 0 {
				dp[i][j] = true
			} else if length == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
			}
			if dp[i][j] && length+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}

	return result
}
