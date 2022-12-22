package leetcode

import (
	"testing"
)

/**
Given the head of a linked list, remove the nᵗʰ node from the end of the list
and return its head.


 Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]


 Example 2:


Input: head = [1], n = 1
Output: []


 Example 3:


Input: head = [1,2], n = 1
Output: [1]



 Constraints:


 The number of nodes in the list is sz.
 1 <= sz <= 30
 0 <= Node.val <= 100
 1 <= n <= sz



 Follow up: Could you do this in one pass?

 Related Topics 链表 双指针 👍 2341 👎 0

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

	cnt := 0

	p1 := dummy

	for ; cnt != n; p1 = p1.Next {
		cnt++
	}

	p2 := dummy
	for ; p1.Next != nil; p1 = p1.Next {
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {

}
