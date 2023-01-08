package leetcode

import (
	"testing"
)

/**
Given the root of a Binary Search Tree (BST), convert it to a Greater Tree such
that every key of the original BST is changed to the original key plus the sum
of all keys greater than the original key in BST.

 As a reminder, a binary search tree is a tree that satisfies these constraints:



 The left subtree of a node contains only nodes with keys less than the node's
key.
 The right subtree of a node contains only nodes with keys greater than the
node's key.
 Both the left and right subtrees must also be binary search trees.



 Example 1:


Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]


 Example 2:


Input: root = [0,null,1]
Output: [1,null,1]



 Constraints:


 The number of nodes in the tree is in the range [1, 100].
 0 <= Node.val <= 100
 All the values in the tree are unique.



 Note: This question is the same as 538: https://leetcode.com/problems/convert-
bst-to-greater-tree/

 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 191 👎 0

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
func bstToGst(root *TreeNode) *TreeNode {
	// 逆前缀和
	// 使用中序遍历(右中左), 等同于逆序遍历二叉搜索树
	// 在逆序过程中用前缀和计算当前节点的值
	preSum := 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Right)
		preSum += node.Val
		node.Val = preSum
		inorder(node.Left)
	}

	inorder(root)

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinarySearchTreeToGreaterSumTree(t *testing.T) {

}
