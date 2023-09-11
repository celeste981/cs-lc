package mergeSortedArray

// 2023-09-11 20:56:20
// leetcode submit region begin(Prohibit modification and deletion)
//func merge(nums1 []int, m int, nums2 []int, n int) {
//	sorted := make([]int, 0, m+n)
//	x, y := 0, 0
//	for {
//		if x == m {
//			sorted = append(sorted, nums2[y:]...)
//			break
//		}
//		if y == n {
//			sorted = append(sorted, nums1[x:]...)
//			break
//		}
//		if nums1[x] > nums2[y] {
//			sorted = append(sorted, nums2[y])
//			y++
//		} else {
//			sorted = append(sorted, nums1[x])
//			x++
//		}
//	}
//	copy(nums1, sorted)
//}

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	x, y := m-1, n-1
//	idx := m+n-1
//	for x >= 0 {
//		if y >= 0 && nums2[y] >= nums1[x] {
//			nums1[idx] = nums2[y]
//			y--
//		} else {
//			nums1[idx] = nums1[x]
//			x--
//		}
//		idx--
//	}
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	x, y := m-1, n-1
	idx := m+n-1
	for y >= 0 {
		if x >= 0 && nums1[x] > nums2[y] {
			nums1[idx] = nums1[x]
			x--
		} else {
			nums1[idx] = nums2[y]
			y--
		}
		idx--
	}
}




//leetcode submit region end(Prohibit modification and deletion)
