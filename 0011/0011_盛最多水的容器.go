package containerWithMostWater

// 2023-09-02 15:35:34
// leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0
	for i < j {
		tmp := (j - i) * minNum(height[i], height[j])
		max = maxNum(max, tmp)
		// 指针移动条件
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func minNum(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func maxNum(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

//leetcode submit region end(Prohibit modification and deletion)
