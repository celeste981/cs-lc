package minCostClimbingStairs

// 2023-12-17 10:57:41
// leetcode submit region begin(Prohibit modification and deletion)
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < len(cost)+1; i++ {
		dp[i] = min(cost[i-2]+dp[i-2], cost[i-1]+dp[i-1])
	}

	return dp[len(cost)]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//leetcode submit region end(Prohibit modification and deletion)
