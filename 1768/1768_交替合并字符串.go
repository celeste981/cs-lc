package mergeStringsAlternately

import (
	"strings"
)

// 2023-12-03 18:22:28
// leetcode submit region begin(Prohibit modification and deletion)
func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	i, j := 0, 0
	// 交替添加字符
	for i < len(word1) && j < len(word2) {
		res.WriteByte(word1[i])
		res.WriteByte(word2[j])
		i++
		j++
	}
	// 添加剩余字符
	for ; i < len(word1); i++ {
		res.WriteByte(word1[i])
	}
	for ; j < len(word2); j++ {
		res.WriteByte(word2[j])
	}

	return res.String()
}

//leetcode submit region end(Prohibit modification and deletion)
