package canPlaceFlowers

// 2023-12-05 20:51:03
// leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	// 遍历数组，检查当前位置是否为0（空间最大化）
	// 如果当前位置为0，检查左右两边是否为0，或者左右两边为边界（有效性检查）
	// 可以种花，修改当前位置为1，n减1
	// 遍历完成判断n是否小于等于0

	for i, flower := range flowerbed {
		if flower == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
			if n <= 0 {
				return true
			}
		}
	}

	return n <= 0
}

//
//leetcode submit region end(Prohibit modification and deletion)
