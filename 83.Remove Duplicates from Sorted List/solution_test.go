package leetcode

import (
	"testing"
)

/**
Given the head of a sorted linked list, delete all duplicates such that each
element appears only once. Return the linked list sorted as well.


 Example 1:


Input: head = [1,1,2]
Output: [1,2]


 Example 2:


Input: head = [1,1,2,3,3]
Output: [1,2,3]



 Constraints:


 The number of nodes in the list is in the range [0, 300].
 -100 <= Node.val <= 100
 The list is guaranteed to be sorted in ascending order.


 Related Topics 链表 👍 906 👎 0

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
func deleteDuplicates(head *ListNode) *ListNode {
	for p := head; p != nil && p.Next != nil; {
		if p.Next.Val == p.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedList(t *testing.T) {

}
