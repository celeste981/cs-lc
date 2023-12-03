package greatestCommonDivisorOfStrings

// 2023-12-03 22:20:35
// leetcode submit region begin(Prohibit modification and deletion)
func gcdOfStrings(str1 string, str2 string) string {
	// 最大公约数
	if str1+str2 != str2+str1 {
		return ""
	}

	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
