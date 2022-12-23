package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 401 👎 0

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
func isSymmetric(root *TreeNode) bool {
	var compare func(nodeA, nodeB *TreeNode) bool
	compare = func(nodeA, nodeB *TreeNode) bool {

		if nodeA == nil && nodeB == nil {
			return true
		}

		if nodeA != nil && nodeB == nil || nodeA == nil && nodeB != nil || nodeA.Val != nodeB.Val {
			return false
		}

		return compare(nodeA.Left, nodeB.Right) && compare(nodeA.Right, nodeB.Left)
	}
	return compare(root, root)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDuiChengDeErChaShuLcof(t *testing.T) {

}
