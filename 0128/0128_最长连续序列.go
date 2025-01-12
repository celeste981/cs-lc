package longestConsecutiveSequence

// 2025-01-12 17:42:18
// leetcode submit region begin(Prohibit modification and deletion)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 1. 构建 HashSet
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	// 2. 遍历 numSet，找到最长的序列
	maxLength := 0

	for num := range numSet {
		// 如果 num-1 不存在，说明是一个新序列的起点
		if _, ok := numSet[num-1]; !ok {
			currentLength := 1
			currentNum := num

			// 查找 num 的连续序列
			for {
				if _, found := numSet[currentNum+1]; found {
					currentNum++
					currentLength++
				} else {
					break
				}
			}

			// 更新最长序列长度
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

// leetcode submit region end(Prohibit modification and deletion)
