package leetcode

import (
	"testing"
)

/**
Given the head of a singly linked list, group all the nodes with odd indices
together followed by the nodes with even indices, and return the reordered list.

 The first node is considered odd, and the second node is even, and so on.

 Note that the relative order inside both the even and odd groups should remain
as it was in the input.

 You must solve the problem in O(1) extra space complexity and O(n) time
complexity.


 Example 1:


Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]


 Example 2:


Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]



 Constraints:


 The number of nodes in the linked list is in the range [0, 10⁴].
 -10⁶ <= Node.val <= 10⁶


 Related Topics 链表 👍 667 👎 0

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
func oddEvenList(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}

	for odd, evenHead, even := head, head.Next, head.Next; even != nil; {
		if even.Next == nil {
			break
		}

		odd.Next = even.Next
		even.Next = even.Next.Next
		odd.Next.Next = evenHead

		odd = odd.Next
		even = even.Next
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestOddEvenLinkedList(t *testing.T) {

}
