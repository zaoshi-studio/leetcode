package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 二叉树 👍 506 👎 0

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var ans *TreeNode

	var traverse func(node *TreeNode) bool
	traverse = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		leftHas := traverse(node.Left)
		rightHas := traverse(node.Right)

		if leftHas && rightHas {
			ans = node
		}

		if (leftHas || rightHas) && (node == p || node == q) {
			ans = node
		}

		return leftHas || rightHas || node == p || node == q
	}

	traverse(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestErChaShuDeZuiJinGongGongZuXianLcof(t *testing.T) {

}
