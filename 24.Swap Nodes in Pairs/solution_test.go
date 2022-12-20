package leetcode

import (
	"testing"
)

/**
Given a linked list, swap every two adjacent nodes and return its head. You
must solve the problem without modifying the values in the list's nodes (i.e., only
nodes themselves may be changed.)


 Example 1:


Input: head = [1,2,3,4]
Output: [2,1,4,3]


 Example 2:


Input: head = []
Output: []


 Example 3:


Input: head = [1]
Output: [1]



 Constraints:


 The number of nodes in the list is in the range [0, 100].
 0 <= Node.val <= 100


 Related Topics 递归 链表 👍 1646 👎 0

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
func swapPairs(head *ListNode) *ListNode {

	dummy := &ListNode{Next: head}

	for cur, pre := head, dummy; cur != nil && cur.Next != nil; {

		holder := cur.Next
		cur.Next = holder.Next
		holder.Next = cur
		pre.Next = holder

		pre = cur
		cur = cur.Next
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSwapNodesInPairs(t *testing.T) {

}
