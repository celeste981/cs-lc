package decodeString

import (
	"unicode"
)

// 2024-11-20 20:50:07
// leetcode submit region begin(Prohibit modification and deletion)
func decodeString(s string) string {
	var numStack []int
	var strStack []string
	curNum := 0
	curStr := ""
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			curNum = curNum*10 + int(ch-'0')
		} else if unicode.IsLetter(ch) {
			curStr += string(ch)
		} else if ch == '[' {
			numStack = append(numStack, curNum)
			strStack = append(strStack, curStr)
			curNum = 0
			curStr = ""
		} else {
			repeatTimes := numStack[len(numStack)-1]
			prevStr := strStack[len(strStack)-1]
			curStr = prevStr + getRepeatStr(repeatTimes, curStr)
			numStack = numStack[:len(numStack)-1]
			strStack = strStack[:len(strStack)-1]
		}
	}
	return curStr
}

func getRepeatStr(times int, str string) string {
	res := ""
	for times > 0 {
		res += str
		times--
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
