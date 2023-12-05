package kidsWithTheGreatestNumberOfCandies

// 2023-12-05 20:21:42
// leetcode submit region begin(Prohibit modification and deletion)
// 确定每个孩子是否可以通过额外的糖果成为拥有最多糖果的孩子
func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, len(candies))
	maxCandy := findMax(candies)

	for i, candy := range candies {
		res[i] = candy+extraCandies >= maxCandy
	}
	return res
}

// 找到数组中的最大值
func findMax(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
