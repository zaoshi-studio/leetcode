package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return its maximum depth.

 A binary tree's maximum depth is the number of nodes along the longest path
from the root node down to the farthest leaf node.


 Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 3


 Example 2:


Input: root = [1,null,2]
Output: 2



 Constraints:


 The number of nodes in the tree is in the range [0, 10⁴].
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1454 👎 0

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
func maxDepth(root *TreeNode) int {

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if nil == node {
			return 0
		}

		return max(dfs(node.Left), dfs(node.Right)) + 1
	}

	return dfs(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumDepthOfBinaryTree(t *testing.T) {

}
