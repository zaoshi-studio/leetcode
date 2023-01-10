package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 72 👎 0

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
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	var (
		ans, pre *TreeNode
	)

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {

		if nil == node {
			return
		}

		inorder(node.Right)

		if node.Val == p.Val {
			ans = pre
		}

		pre = node

		inorder(node.Left)

		return
	}

	inorder(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestP5rCT8(t *testing.T) {

}
