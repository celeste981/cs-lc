package removingStarsFromAString

// 2024-11-18 20:25:52
// leetcode submit region begin(Prohibit modification and deletion)
func removeStars(s string) string {
	var res []rune
	for _, ch := range s {
		if ch != '*' {
			res = append(res, ch)
		} else {
			res = res[:len(res)-1]
		}
	}

	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)
