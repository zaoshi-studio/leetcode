package leetcode

import (
	"testing"
)

/**
Given the head of a linked list and an integer val, remove all the nodes of the
linked list that has Node.val == val, and return the new head.


 Example 1:


Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]


 Example 2:


Input: head = [], val = 1
Output: []


 Example 3:


Input: head = [7,7,7,7], val = 7
Output: []



 Constraints:


 The number of nodes in the list is in the range [0, 10⁴].
 1 <= Node.val <= 50
 0 <= val <= 50


 Related Topics 递归 链表 👍 1149 👎 0

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
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}

	for p := dummy; p != nil && p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveLinkedListElements(t *testing.T) {

}
