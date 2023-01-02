package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 350 👎 0

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
func kthLargest(root *TreeNode, k int) int {

	var ans int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if nil == node {
			return
		}

		// 在右节点中尝试找有没有比当前节点大的
		dfs(node.Right)

		// -- 是当前节点占据一个位置
		k--
		if k == 0 {
			ans = node.Val
			return
		}

		dfs(node.Left)
	}

	dfs(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestErChaSouSuoShuDeDiKdaJieDianLcof(t *testing.T) {

}
