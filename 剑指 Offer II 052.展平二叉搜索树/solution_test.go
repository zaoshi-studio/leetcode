package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 栈 树 深度优先搜索 二叉搜索树 二叉树 👍 47 👎 0

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
func increasingBST(root *TreeNode) *TreeNode {

	dummy := &TreeNode{}
	cursor := dummy

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		// cursor 指向中序顺序中的前一个节点
		cursor.Right = node
		node.Left = nil

		// 该步是以一个中序的顺序让游标进行遍历
		cursor = node

		inorder(node.Right)
	}

	inorder(root)

	return dummy.Right
}

//leetcode submit region end(Prohibit modification and deletion)

func TestNYBBNL(t *testing.T) {

}
