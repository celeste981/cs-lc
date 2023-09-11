package longestSubstringWithoutRepeatingCharacters

// 2023-09-03 22:43:34
// leetcode submit region begin(Prohibit modification and deletion)
//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	max := 1
//	i := 0
//	j := 0
//	// map: character -> idx on the right
//	hashMap := make(map[byte]bool)
//	for j < len(s) {
//		if flag, _ := hashMap[s[j]]; !flag {
//			hashMap[s[j]] = true
//			j++
//			max = maxNum(max, j-i)
//		} else {
//			hashMap[s[i]] = false
//			i++
//		}
//	}
//	return max
//}
//
//func maxNum(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 1

	i, j := 0, 0
	location := make([]int, 128)
	for j < len(s) {
		// 1. 保证区间内无重复
		i = maxNum(i, location[s[j]])
		// 2. 更新遍历过的字母的下标
		location[s[j]] = j + 1
		// 3. 更新子串长度
		res = maxNum(res, j-i+1)
		// 4. 右指针右移
		j++
	}
	return res
}

func maxNum(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//leetcode submit region end(Prohibit modification and deletion)
