package leetcode

import (
	"testing"
)

/**
Given the head of a singly linked list and two integers left and right where
left <= right, reverse the nodes of the list from position left to position right,
and return the reversed list.


 Example 1:


Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]


 Example 2:


Input: head = [5], left = 1, right = 1
Output: [5]



 Constraints:


 The number of nodes in the list is n.
 1 <= n <= 500
 -500 <= Node.val <= 500
 1 <= left <= right <= n



Follow up: Could you do it in one pass?

 Related Topics 链表 👍 1483 👎 0

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

var static *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		static = head.Next
		return head
	}

	nextNode := reverseN(head.Next, n-1)

	head.Next.Next = head

	head.Next = static

	return nextNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseLinkedListIi(t *testing.T) {

}
