package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 48 👎 0

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
func convertBST(root *TreeNode) *TreeNode {

	var inorder func(node *TreeNode, preSum int) int
	inorder = func(node *TreeNode, preSum int) int {
		if node == nil {
			return preSum
		}

		if node.Left == nil && node.Right == nil {
			node.Val += preSum
			return node.Val
		}

		leftSum := inorder(node.Right, preSum)
		node.Val += leftSum

		return inorder(node.Left, node.Val)
	}

	inorder(root, 0)

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestW6cpku(t *testing.T) {

}
