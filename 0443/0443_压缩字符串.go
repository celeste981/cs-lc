package stringCompression

import (
	"strconv"
)

// 2024-01-17 20:22:04
// leetcode submit region begin(Prohibit modification and deletion)
func compress(chars []byte) int {
	var ch, chLen, idx = chars[0], 1, 0
	for i := 1; i < len(chars); i++ {
		if ch == chars[i] {
			chLen++
		} else {
			chars[idx] = ch
			idx++
			ch = chars[i]
			if chLen != 1 {
				for _, num := range strconv.Itoa(chLen) {
					chars[idx] = byte(num)
					idx++
				}
				chLen = 1
			}
		}
	}
	chars[idx] = ch
	idx++
	if chLen != 1 {
		for _, num := range strconv.Itoa(chLen) {
			chars[idx] = byte(num)
			idx++
		}
	}
	return idx
}

//leetcode submit region end(Prohibit modification and deletion)
