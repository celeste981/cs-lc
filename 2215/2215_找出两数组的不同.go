package findTheDifferenceOfTwoArrays

// 2024-11-05 18:02:15
// leetcode submit region begin(Prohibit modification and deletion)
func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)

	// 创建两个哈希表记录nums1和nums2中的元素
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)
	for _, n1 := range nums1 {
		m1[n1] = true
	}
	for _, n2 := range nums2 {
		m2[n2] = true
	}

	// 找出nums1中不在nums2的元素
	for num := range m1 {
		if !m2[num] {
			res[0] = append(res[0], num)
		}
	}
	// 找出nums2中不在nums1的元素
	for num := range m2 {
		if !m1[num] {
			res[1] = append(res[1], num)
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
