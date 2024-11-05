package determineIfTwoStringsAreClose

// 2024-11-05 20:31:28
// leetcode submit region begin(Prohibit modification and deletion)
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	// 确定两个字符串中的 字符种类、种类数 和每个种类下的个数一致就是"接近"

	// 获取字符种类统计
	uniqueCharsInWord1 := make(map[rune]int)
	for _, ch := range word1 {
		uniqueCharsInWord1[ch]++
	}

	uniqueCharsInWord2 := make(map[rune]int)
	for _, ch := range word2 {
		uniqueCharsInWord2[ch]++
	}

	// 种类数比较
	if len(uniqueCharsInWord1) != len(uniqueCharsInWord2) {
		return false
	}

	// 种类比较
	for k := range uniqueCharsInWord1 {
		if _, exist := uniqueCharsInWord2[k]; !exist {
			return false
		}
	}

	// 统计每个种类出现的次数是否完全一致
	cateCnt1 := make(map[int]int)
	for _, cnt := range uniqueCharsInWord1 {
		cateCnt1[cnt]++
	}

	cateCnt2 := make(map[int]int)
	for _, cnt := range uniqueCharsInWord2 {
		cateCnt2[cnt]++
	}

	for key, val := range cateCnt1 {
		if _, exist := cateCnt2[key]; !exist {
			return false
		}
		if val != cateCnt2[key] {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
