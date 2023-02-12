package leetcode

import (
	"testing"
)

/**
Design an algorithm and write code to find the first common ancestor of two
nodes in a binary tree. Avoid storing additional nodes in a data structure. NOTE:
This is not necessarily a binary search tree.

 For example, Given the following tree: root = [3,5,1,6,2,0,8,null,null,7,4]


    3
   / \
  5   1
 / \ / \
6  2 0  8
  / \
 7   4


 Example 1:


Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Input: 3
Explanation: The first common ancestor of node 5 and node 1 is node 3.

 Example 2:


Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
Output: 5
Explanation: The first common ancestor of node 5 and node 4 is node 5.

 Notes:


 All node values are pairwise distinct.
 p, q are different node and both can be found in the given tree.


 Related Topics 树 深度优先搜索 二叉树 👍 89 👎 0

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
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

	var ans *TreeNode

	var postorder func(node *TreeNode) bool
	postorder = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		left := postorder(node.Left)
		right := postorder(node.Right)

		if (left || right) && (node.Val == q.Val || node.Val == p.Val) {
			ans = node
		}

		if left && right {
			ans = node
		}

		return left || right || node.Val == q.Val || node.Val == p.Val
	}

	postorder(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFirstCommonAncestorLcci(t *testing.T) {

}
