package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the length of the diameter of the tree.


 The diameter of a binary tree is the length of the longest path between any
two nodes in a tree. This path may or may not pass through the root.

 The length of a path between two nodes is represented by the number of edges
between them.


 Example 1:


Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].


 Example 2:


Input: root = [1,2]
Output: 1



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 二叉树 👍 1208 👎 0

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
func diameterOfBinaryTree(root *TreeNode) int {

	var ans int

	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 不能直接使用 max(root.Left + root.Right)
		// 因为可能最长的直径在某个子树中
		// 某根节点下的最长直径 = 左子树最大深度 + 右子树最大深度
		left := height(node.Left)
		right := height(node.Right)

		ans = max(ans, left+right)
		return max(left, right) + 1
	}

	height(root)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDiameterOfBinaryTree(t *testing.T) {

}
