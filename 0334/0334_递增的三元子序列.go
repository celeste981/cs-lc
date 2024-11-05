package increasingTripletSubsequence

import (
	"math"
)

// 2024-01-16 14:30:14
// leetcode submit region begin(Prohibit modification and deletion)
func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
