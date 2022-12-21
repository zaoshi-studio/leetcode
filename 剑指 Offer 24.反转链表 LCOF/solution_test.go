package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 递归 链表 👍 517 👎 0

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

func TestFanZhuanLianBiaoLcof(t *testing.T) {

}
