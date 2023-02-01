package leetcode

import (
	"testing"
)

/**
You are given two linked lists: list1 and list2 of sizes n and m respectively.

 Remove list1's nodes from the aᵗʰ node to the bᵗʰ node, and put list2 in their
place.

 The blue edges and nodes in the following figure indicate the result:

 Build the result list and return its head.


 Example 1:


Input: list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
Output: [0,1,2,1000000,1000001,1000002,5]
Explanation: We remove the nodes 3 and 4 and put the entire list2 in their
place. The blue edges and nodes in the above figure indicate the result.


 Example 2:


Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1
000003,1000004]
Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
Explanation: The blue edges and nodes in the above figure indicate the result.



 Constraints:


 3 <= list1.length <= 10⁴
 1 <= a <= b < list1.length - 1
 1 <= list2.length <= 10⁴


 Related Topics 链表 👍 71 👎 0

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
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Next: list1,
	}

	var (
		prea, bnode, list2_tail *ListNode
	)
	cnt := 0

	for p1 := dummy; p1 != nil; p1 = p1.Next {
		if a == cnt {
			prea = p1
		}
		if cnt-1 == b {
			bnode = p1
		}
		cnt++
	}

	for list2_tail = list2; list2_tail.Next != nil; list2_tail = list2_tail.Next {
	}

	prea.Next = list2
	list2_tail.Next = bnode.Next

	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeInBetweenLinkedLists(t *testing.T) {

}
