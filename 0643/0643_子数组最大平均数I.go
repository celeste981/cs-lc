package maximumAverageSubarrayI

import (
	"math"
)

// 2025-01-21 20:48:53
// leetcode submit region begin(Prohibit modification and deletion)
func findMaxAverage(nums []int, k int) float64 {
	max, sum := math.MinInt, 0
	l, r := 0, 0
	for r < len(nums) {
		// 扩展窗口
		sum += nums[r]
		r++
		// 收缩窗口
		for r-l > k {
			sum -= nums[l]
			l++
		}
		// 窗口满足条件时，更新结果
		if r-l == k && max < sum {
			max = sum
		}
	}

	return float64(max) / float64(k)
}

//leetcode submit region end(Prohibit modification and deletion)
