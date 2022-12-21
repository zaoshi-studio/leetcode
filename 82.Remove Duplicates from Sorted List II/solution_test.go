package leetcode

import (
	"testing"
)

/**
Given the head of a sorted linked list, delete all nodes that have duplicate
numbers, leaving only distinct numbers from the original list. Return the linked
list sorted as well.


 Example 1:


Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]


 Example 2:


Input: head = [1,1,1,2,3]
Output: [2,3]



 Constraints:


 The number of nodes in the list is in the range [0, 300].
 -100 <= Node.val <= 100
 The list is guaranteed to be sorted in ascending order.


 Related Topics 链表 双指针 👍 1034 👎 0

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

	dummy := &ListNode{Next: head}

	for pre, cur := dummy, head; cur != nil && cur.Next != nil; {

		p := cur.Next

		// cur 指向当前需要判断的结点
		for ; p != nil && p.Val == cur.Val; p = p.Next {
		}

		// p 指向与 cur 不相等的第一个节点
		if cur.Next == p {
			pre = cur
		} else {
			pre.Next = p
		}

		cur = p
	}

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRemoveDuplicatesFromSortedListIi(t *testing.T) {

}
