package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 链表 👍 269 👎 0

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
func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	for p := dummy; p.Next != nil; p = p.Next {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			break
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestShanChuLianBiaoDeJieDianLcof(t *testing.T) {

}
