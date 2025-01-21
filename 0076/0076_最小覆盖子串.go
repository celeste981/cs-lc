package minimumWindowSubstring

// 2025-01-21 20:58:59
// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	min := ""
	need := make(map[byte]int)
	window := make(map[byte]int)
	valid := 0
	for _, v := range t {
		need[byte(v)]++
	}
	l, r := 0, 0
	for r < len(s) {
		window[s[r]]++
		if _, ok := need[s[r]]; ok && need[s[r]] == window[s[r]] {
			valid++
		}
		r++
		for valid == len(need) {
			if min == "" {
				min = s[l:r]
			}
			min = minStr(s[l:r], min)
			if _, ok := need[s[l]]; ok && need[s[l]] == window[s[l]] {
				valid--
			}
			window[s[l]]--
			l++
		}
	}
	return min
}

func minStr(a, b string) string {
	if len(a) < len(b) {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
