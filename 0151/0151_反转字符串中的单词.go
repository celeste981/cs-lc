package reverseWordsInAString

// 2024-01-08 22:10:43
// leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	// 清理空格
	var clean []rune
	for i, ch := range s {
		if ch != ' ' || (i > 0 && s[i-1] != ' ') {
			clean = append(clean, ch)
		}
	}

	// 去除末尾的空格
	if len(clean) > 0 && clean[len(clean)-1] == ' ' {
		clean = clean[:len(clean)-1]
	}

	// 反转字符串
	reverse(clean, 0, len(clean)-1)

	// 反转字符串里的每个单词
	idx := 0
	for i, ch := range clean {
		if ch != ' ' {
			continue
		}
		reverse(clean, idx, i-1)
		idx = i + 1
	}

	// 反转最后一个单词（如果需要）
	if idx < len(clean) {
		reverse(clean, idx, len(clean)-1)
	}

	return string(clean)
}

func reverse(data []rune, start, end int) {
	for start < end {
		data[start], data[end] = data[end], data[start]
		start++
		end--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
