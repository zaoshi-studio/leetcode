package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 哈希表 双指针 二叉树 👍 48 👎 0

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
func findTarget(root *TreeNode, k int) bool {

	var ans bool
	record := map[int]struct{}{}

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		if _, ok := record[k-node.Val]; ok {
			ans = true
			return
		}

		record[node.Val] = struct{}{}

		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestOpLdQZ(t *testing.T) {

}
