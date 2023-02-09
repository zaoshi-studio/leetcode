package leetcode

import (
	"testing"
)

/**
Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 8 👎 0

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
func expandBinaryTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	if root.Left != nil {
		newLeftNode := &TreeNode{
			Val: -1,
		}
		newLeftNode.Left = expandBinaryTree(root.Left)

		root.Left = newLeftNode
	}

	if root.Right != nil {
		newRightNode := &TreeNode{
			Val: -1,
		}
		newRightNode.Right = expandBinaryTree(root.Right)

		root.Right = newRightNode
	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKnLfVT(t *testing.T) {

}
