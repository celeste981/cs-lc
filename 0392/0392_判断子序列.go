package isSubsequence

import (
	"sort"
)

// 2024-05-17 18:52:44
// leetcode submit region begin(Prohibit modification and deletion)
//func isSubsequence(s string, t string) bool {
//	i, j := 0, 0
//	for i < len(s) && j < len(t) {
//		if s[i] == t[j] {
//			i++
//		}
//		j++
//	}
//	return i == len(s)
//}

// 进阶
// 对字符串t进行预处理
func sortStringT(t string) map[rune][]int {
	res := make(map[rune][]int)
	for i, char := range t {
		res[char] = append(res[char], i)
	}
	return res
}

func isSubsequence(s string, t string) bool {
	curPos := -1
	charIndex := sortStringT(t)
	for _, char := range s {
		if indices, found := charIndex[char]; found {
			index := sort.Search(len(indices), func(i int) bool {
				return indices[i] > curPos
			})
			if index == len(indices) {
				return false
			}
			curPos = indices[index]
		} else {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
