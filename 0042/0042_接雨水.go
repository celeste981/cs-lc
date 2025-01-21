package trappingRainWater

// 2025-01-15 18:18:51
// leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	lMax, rMax := height[l], height[r]
	// 指针最终会停在整个height数组的最高处
	for l < r {
		lMax = maxNum(height[l], lMax)
		rMax = maxNum(height[r], rMax)
		// 比较指针两侧的最大值，用最小的最大值来计算雨水量
		if lMax < rMax {
			// 计算雨水量
			if lMax > height[l] {
				res += lMax - height[l]
			}
			// 移动指针
			l++
		} else {
			if rMax > height[r] {
				res += rMax - height[r]
			}
			r--
		}
	}

	return res
}

func maxNum(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

//leetcode submit region end(Prohibit modification and deletion)
