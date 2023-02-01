package leetcode

import (
	"testing"
)

/**
Given a binary tree, determine if it is height-balanced.


 Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: true


 Example 2:


Input: root = [1,2,2,3,3,null,null,4,4]
Output: false


 Example 3:


Input: root = []
Output: true



 Constraints:


 The number of nodes in the tree is in the range [0, 5000].
 -10⁴ <= Node.val <= 10⁴


 Related Topics 树 深度优先搜索 二叉树 👍 1232 👎 0

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

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := getHeight(node.Left)
	right := getHeight(node.Right)
	return max(left, right) + 1
}

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	// 高度差不小于 1
	if rightHeight-leftHeight >= -1 && rightHeight-leftHeight <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBalancedBinaryTree(t *testing.T) {

}
