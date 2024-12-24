package moveZeroes

// 2024-05-17 16:57:43
// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	nonZeroIndex := 0
	for _, num := range nums {
		if num != 0 {
			nums[nonZeroIndex] = num
			nonZeroIndex++
		}
	}
	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}

//leetcode submit region end(Prohibit modification and deletion)
