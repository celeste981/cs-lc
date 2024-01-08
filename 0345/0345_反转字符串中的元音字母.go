package reverseVowelsOfAString

// 2024-01-08 21:15:06
// leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	chars := []rune(s)
	for l < r {
		for l < r && !vowels[chars[l]] {
			l++
		}
		for l < r && !vowels[chars[r]] {
			r--
		}
		chars[l], chars[r] = chars[r], chars[l]
		l++
		r--
	}
	return string(chars)
}

//leetcode submit region end(Prohibit modification and deletion)
