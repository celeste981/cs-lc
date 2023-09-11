package zigzagConversion

import (
	"bytes"
)

// 2023-09-11 20:23:02
// leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}
	mat := make([][]byte, numRows)
	circle := 2*numRows - 2
	line := 0
	for i := 0; i < len(s); i++ {
		mat[line] = append(mat[line], s[i])
		if i%circle < numRows-1 {
			line++
		} else {
			line--
		}
	}
	return string(bytes.Join(mat, nil))
}

//leetcode submit region end(Prohibit modification and deletion)
