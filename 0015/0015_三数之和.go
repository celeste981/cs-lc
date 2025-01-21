package threeSum

import (
	"sort"
)

// 2025-01-13 19:53:27
// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sum := 0
	// 排序便于指针处理
	sort.Ints(nums)
	for i, num := range nums {
		// 去除重复的
		if i > 0 && nums[i-1] == num {
			continue
		}
		// 计算目标值
		target := sum - num
		// 移动双指针，将所有符合条件的数据加进结果中
		l, r := i+1, len(nums)-1
		for l < r {
			if target == (nums[l] + nums[r]) {
				res = append(res, []int{num, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if target < (nums[l] + nums[r]) {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
