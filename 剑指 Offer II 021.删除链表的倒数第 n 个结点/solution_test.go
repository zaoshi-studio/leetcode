package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 链表 双指针 👍 63 👎 0

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	p1 := dummy
	cnt := 0
	for ; cnt != n; p1 = p1.Next {
		cnt++
	}

	p2 := dummy
	for p1.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}

	p2.Next = p2.Next.Next

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSLwz0R(t *testing.T) {

}
