package twoSum

// 2023-08-29 20:28:02
// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		leftPart := target - num
		if idx, ok := hashMap[leftPart]; ok {
			return []int{idx, i}
		}
		hashMap[num] = i
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
