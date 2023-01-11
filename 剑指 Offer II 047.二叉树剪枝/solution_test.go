package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 二叉树 👍 58 👎 0

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
func pruneTree(root *TreeNode) *TreeNode {

	dummy := &TreeNode{
		Left: root,
	}

	var traverse func(node *TreeNode) bool
	traverse = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		leftMatch := traverse(node.Left)
		rightMatch := traverse(node.Right)

		if leftMatch {
			node.Left = nil
		}

		if rightMatch {
			node.Right = nil
		}

		return node.Val == 0 && leftMatch && rightMatch
	}

	traverse(dummy)

	return dummy.Left
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPOCWxh(t *testing.T) {

}
