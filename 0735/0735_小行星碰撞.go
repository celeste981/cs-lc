package asteroidCollision

// 2024-11-18 20:42:42
// leetcode submit region begin(Prohibit modification and deletion)
func asteroidCollision(asteroids []int) []int {
	var res []int

	for _, asteroid := range asteroids {
		for len(res) > 0 && res[len(res)-1] > 0 && asteroid < 0 {
			prev := abs(res[len(res)-1])
			cur := abs(asteroid)
			if prev == cur {
				res = res[:len(res)-1]
				asteroid = 0
				break
			} else if prev > cur {
				asteroid = 0
				break
			} else {
				res = res[:len(res)-1]
			}
		}
		if asteroid != 0 {
			res = append(res, asteroid)
		}
	}

	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

//leetcode submit region end(Prohibit modification and deletion)
