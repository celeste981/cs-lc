package longestPalindromicSubstring

// 2023-09-04 11:02:56
// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	res, left, right := 0, 0, 0

	dp := make([][]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && (j-i+1 > 1 || dp[i-1][j+1])
			if dp[i][j] {
				if (j - i + 1) > res {
					res = j - i + 1
					left = i
					right = j
				}
			}
		}
	}
	return s[left : right+1]
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
