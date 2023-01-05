package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 二叉树 👍 44 👎 0

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
func sumNumbers(root *TreeNode) int {

	var traverse func(node *TreeNode, val int) int
	traverse = func(node *TreeNode, val int) int {
		if node == nil {
			return 0
		}

		val = val*10 + node.Val

		if node.Left == nil && node.Right == nil {
			return val
		}
		return traverse(node.Left, val) + traverse(node.Right, val)
	}

	return traverse(root, 0)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestThreeEtpl5(t *testing.T) {

}
