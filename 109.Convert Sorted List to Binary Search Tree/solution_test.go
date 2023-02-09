package leetcode

import (
	"testing"
)

/**
Given the head of a singly linked list where elements are sorted in ascending
order, convert it to a height-balanced binary search tree.


 Example 1:


Input: head = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the
shown height balanced BST.


 Example 2:


Input: head = []
Output: []



 Constraints:


 The number of nodes in head is in the range [0, 2 * 10⁴].
 -10⁵ <= Node.val <= 10⁵


 Related Topics 树 二叉搜索树 链表 分治 二叉树 👍 798 👎 0

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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}

	mid := binaryList(left, right)
	node := &TreeNode{
		Val: mid.Val,
	}

	node.Left = buildTree(left, mid)
	node.Right = buildTree(mid.Next, right)
	return node
}

// 查找链表的中间节点
func binaryList(left, right *ListNode) *ListNode {

	fast, slow := left, left

	for fast != right && fast.Next != right {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConvertSortedListToBinarySearchTree(t *testing.T) {

}
