package leetcode

import (
	"testing"
)

/**
Given a binary search tree (BST), find the lowest common ancestor (LCA) node of
two given nodes in the BST.

 According to the definition of LCA on Wikipedia: “The lowest common ancestor
is defined between two nodes p and q as the lowest node in T that has both p and
q as descendants (where we allow a node to be a descendant of itself).”


 Example 1:


Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.


 Example 2:


Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of
itself according to the LCA definition.


 Example 3:


Input: root = [2,1], p = 2, q = 1
Output: 2



 Constraints:


 The number of nodes in the tree is in the range [2, 10⁵].
 -10⁹ <= Node.val <= 10⁹
 All Node.val are unique.
 p != q
 p and q will exist in the BST.


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1008 👎 0

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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode

	var postorder func(node *TreeNode) bool
	postorder = func(node *TreeNode) bool {
		if nil == node {
			return false
		}

		leftHas := postorder(node.Left)
		rightHas := postorder(node.Right)

		if leftHas && rightHas {
			ans = node
		}

		if (leftHas || rightHas) && (node.Val == p.Val || node.Val == q.Val) {
			ans = node
		}

		return (leftHas || rightHas) || node.Val == p.Val || node.Val == q.Val
	}

	postorder(root)

	return ans
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if nil == root {
		return nil
	}

	var smallNode, largeNode *TreeNode

	if p.Val > q.Val {
		largeNode = p
		smallNode = q
	} else {
		largeNode = q
		smallNode = p
	}

	if root.Val < smallNode.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	if root.Val > largeNode.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLowestCommonAncestorOfABinarySearchTree(t *testing.T) {

}
