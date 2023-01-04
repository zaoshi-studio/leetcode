package leetcode

import (
	"testing"
)

/**
Given the roots of two binary trees p and q, write a function to check if they
are the same or not.

 Two binary trees are considered the same if they are structurally identical,
and the nodes have the same value.


 Example 1:


Input: p = [1,2,3], q = [1,2,3]
Output: true


 Example 2:


Input: p = [1,2], q = [1,null,2]
Output: false


 Example 3:


Input: p = [1,2,1], q = [1,1,2]
Output: false



 Constraints:


 The number of nodes in both trees is in the range [0, 100].
 -10⁴ <= Node.val <= 10⁴


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 946 👎 0

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
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

	var compare func(nodeA, nodeB *TreeNode) bool
	compare = func(nodeA, nodeB *TreeNode) bool {
		if nodeA == nil && nodeB == nil {
			return true
		}

		if (nodeA == nil && nodeB != nil) || (nodeB == nil && nodeA != nil) {
			return false
		}

		if nodeA.Val != nodeB.Val {
			return false
		}

		return compare(nodeA.Left, nodeB.Left) && compare(nodeA.Right, nodeB.Right)
	}

	return compare(p, q)

}

//leetcode submit region end(Prohibit modification and deletion)

func TestSameTree(t *testing.T) {

}
