package palindromeNumber

// 2023-09-02 15:14:49
// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// change to other num
	reverseX := 0
	originX := x
	for x > 0 {
		reverseX = reverseX*10 + x%10
		x /= 10
	}

	return reverseX == originX
}

//leetcode submit region end(Prohibit modification and deletion)
