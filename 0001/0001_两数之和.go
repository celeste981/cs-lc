package twoSum

// 2023-08-29 20:28:02
// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for idx, num := range nums {
		if idx2, ok := hashMap[target-num]; ok {
			return []int{idx2, idx}
		}
		hashMap[num] = idx
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
