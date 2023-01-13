package leetcode

import (
	"math"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 动态规划 二叉树 👍 64 👎 0

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
func maxPathSum(root *TreeNode) int {

	ans := math.MinInt32

	// 任意两节点之间的路径只有两种情况
	// 从上到右子树
	// 从上到左子树
	// 总左子树到右子树
	var preorder func(node *TreeNode) int
	preorder = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 从上到左子树
		leftMax := max(preorder(node.Left), 0)
		// 从上到右子树
		rightMax := max(preorder(node.Right), 0)
		// 从左子树到右子树
		leftToRight := leftMax + node.Val + rightMax

		ans = max(ans, leftToRight)

		return node.Val + max(leftMax, rightMax)
	}

	preorder(root)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestJC7MId(t *testing.T) {

}
