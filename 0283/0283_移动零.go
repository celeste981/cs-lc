package moveZeroes

// 2024-05-17 16:57:43
// leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int) {
	zeroCount := 0
	for i, num := range nums {
		if num == 0 {
			zeroCount++
			continue
		}
		if zeroCount > 0 {
			tmp := num
			nums[i-zeroCount] = tmp
			nums[i] = 0
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
