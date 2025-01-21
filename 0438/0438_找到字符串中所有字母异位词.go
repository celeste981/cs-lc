package findAllAnagramsInAString

// 2025-01-21 19:35:26
// leetcode submit region begin(Prohibit modification and deletion)
// -- 会超时
//func findAnagrams(s string, p string) []int {
//	if len(s) < len(p) {
//		return nil
//	}
//	ans := make([]int, 0)
//	// 固定长度为p.len的窗口
//	l, r := 0, len(p)-1
//	for r < len(s) {
//		curStr := s[l : r+1]
//		if sortStr(curStr) == sortStr(p) {
//			ans = append(ans, l)
//		}
//		l++
//		r++
//	}
//
//	return ans
//}
//
//func sortStr(s string) string {
//	bytes := []byte(s)
//	sort.Slice(bytes, func(i, j int) bool {
//		return bytes[i] < bytes[j]
//	})
//	return string(bytes)
//}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	ans := make([]int, 0)

	// 定义频率窗口
	need := make([]int, 26)
	tmp := make([]int, 26)

	// 初始化need窗口
	for _, v := range p {
		need[int(v-'a')]++
	}

	l, r := 0, 0
	// 滑动窗口来搜集index
	for r < len(s) {
		tmp[int(s[r]-'a')]++
		r++

		for r-l > len(p) {
			tmp[int(s[l]-'a')]--
			l++
		}

		if isEqual(need, tmp) {
			ans = append(ans, l)
		}

	}

	return ans
}

// 检查两个map是否一致
func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
