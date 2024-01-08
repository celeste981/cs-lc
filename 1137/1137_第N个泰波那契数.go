package nThTribonacciNumber

// 2023-12-21 20:57:30
// leetcode submit region begin(Prohibit modification and deletion)
func tribonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

//leetcode submit region end(Prohibit modification and deletion)
