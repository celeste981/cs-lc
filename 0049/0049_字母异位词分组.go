package groupAnagrams

import (
	"sort"
	"strings"
)

// 2024-12-19 20:49:31
// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	strSortedMap := make(map[string][]string)
	for _, str := range strs { // n
		sortedStr := sortString(str)
		strSortedMap[sortedStr] = append(strSortedMap[sortedStr], str)
	}
	for _, s := range strSortedMap {
		res = append(res, s)
	}
	return res
}

// k*log k
func sortString(str string) string {
	strArray := strings.Split(str, "")
	sort.Strings(strArray)
	//strRune := []rune(str)
	//sort.Slice(strRune, func(i, j int) bool {
	//	return strRune[i] < strRune[j]
	//})
	return strings.Join(strArray, "")
}

//func groupAnagrams(strs []string) [][]string {
//	res := [][]string{}
//	strCountMap := make(map[[26]int][]string)
//
//	for _, str := range strs {// n
//		count := [26]int{}
//		for _, char := range str { // k
//			count[char-'a']++
//		}
//		strCountMap[count] = append(strCountMap[count], str)
//	} // n*k
//
//	for _, group := range strCountMap {
//		res = append(res, group)
//	} // n
//	return res
//	// n*k
//}

//leetcode submit region end(Prohibit modification and deletion)
