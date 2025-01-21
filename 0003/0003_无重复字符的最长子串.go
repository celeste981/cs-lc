package longestSubstringWithoutRepeatingCharacters

// 2023-09-03 22:43:34
// leetcode submit region begin(Prohibit modification and deletion)

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max, l, r := 0, 0, 0
	windows := make(map[byte]int)
	for r < len(s) {
		c := s[r]
		windows[c]++
		r++

		// 保证窗口中无重复字符
		for windows[c] > 1 {
			windows[s[l]]--
			l++
		}

		// 更新窗口的值
		max = maxNum(max, r-l)
	}
	return max
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
