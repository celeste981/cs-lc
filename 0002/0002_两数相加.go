package addTwoNumbers

// 2023-09-02 19:27:48
//leetcode submit region begin(Prohibit modification and deletion)
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	tail := res
	carry := 0
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + carry
		carry = tmp / 10
		left := tmp % 10
		res.Next = &ListNode{
			Val:  left,
			Next: nil,
		}
		l1 = l1.Next
		l2 = l2.Next
		res = res.Next
	}
	if l2 != nil {
		l1 = l2
	}
	for l1 != nil {
		tmp := l1.Val + carry
		carry = tmp / 10
		left := tmp % 10
		res.Next = &ListNode{
			Val:  left,
			Next: nil,
		}
		l1 = l1.Next
		res = res.Next
	}

	// 如果还有进位
	if carry != 0 {
		res.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return tail.Next
}

//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//leetcode submit region end(Prohibit modification and deletion)
