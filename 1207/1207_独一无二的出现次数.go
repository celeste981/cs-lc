package uniqueNumberOfOccurrences

// 2024-11-05 20:13:06
// leetcode submit region begin(Prohibit modification and deletion)
func uniqueOccurrences(arr []int) bool {
	// 统计所有num出现的次数
	cntMap := make(map[int]int)
	for _, num := range arr {
		cntMap[num]++
	}

	// 用集合再次统计是否出现
	existMap := make(map[int]struct{})
	for _, cnt := range cntMap {
		if _, ok := existMap[cnt]; ok {
			return false
		}
		existMap[cnt] = struct{}{}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
