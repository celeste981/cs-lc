package containerWithMostWater

// 2023-09-02 15:35:34
// leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}

	max, i, j := 0, 0, len(height)-1

	for i < j {
		cur := minNum(height[i], height[j]) * (j - i)
		// 如果较短的柱子不动，面积一定不会再增大，所以需要移动较短的柱子以寻求最优解
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		if cur > max {
			max = cur
		}
	}

	return max
}

func minNum(i int, j int) int {
	if i < j {
		return i
	}
	return j
}

//leetcode submit region end(Prohibit modification and deletion)
