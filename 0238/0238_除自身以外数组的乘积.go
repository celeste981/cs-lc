package productOfArrayExceptSelf

// 2024-01-14 21:54:15
// leetcode submit region begin(Prohibit modification and deletion)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftProducts := make([]int, n)
	// 更新左侧乘积数组
	leftRes := 1
	for i, num := range nums {
		leftProducts[i] = leftRes
		leftRes *= num
	}
	// 更新右侧乘积结果并更新结果
	rightRes := 1
	for i := n - 1; i >= 0; i-- {
		leftProducts[i] *= rightRes
		rightRes *= nums[i]
	}
	return leftProducts
}

//leetcode submit region end(Prohibit modification and deletion)
