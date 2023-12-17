package findTheHighestAltitude

// 2023-12-17 10:50:27
// leetcode submit region begin(Prohibit modification and deletion)
func largestAltitude(gain []int) int {
	max := 0
	cur := 0
	for _, g := range gain {
		cur += g
		if cur > max {
			max = cur
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
