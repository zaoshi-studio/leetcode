package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 栈 链表 数学 👍 73 👎 0

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)

	head := &ListNode{}
	cur := head

	extra := 0
	p1, p2 := l1, l2
	for p1 != nil || p2 != nil {
		sum := extra

		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}

		remain := sum % 10
		extra = sum / 10

		cur.Next = &ListNode{
			Next: cur.Next,
			Val:  remain,
		}

		cur = cur.Next
	}

	if extra != 0 {
		cur.Next = &ListNode{
			Val: extra,
		}
	}

	return reverseList(head.Next)
}

func reverseList(head *ListNode) *ListNode {
	dummy := &ListNode{Next: nil}

	for p := head; p != nil; {
		holder := p.Next
		p.Next = dummy.Next
		dummy.Next = p

		p = holder
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLMSNwu(t *testing.T) {

}
