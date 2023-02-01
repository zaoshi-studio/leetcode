package leetcode

import (
	"testing"
)

/**
Given the root of a complete binary tree, return the number of the nodes in the
tree.

 According to Wikipedia, every level, except possibly the last, is completely
filled in a complete binary tree, and all nodes in the last level are as far left
as possible. It can have between 1 and 2ʰ nodes inclusive at the last level h.

 Design an algorithm that runs in less than O(n) time complexity.


 Example 1:


Input: root = [1,2,3,4,5,6]
Output: 6


 Example 2:


Input: root = []
Output: 0


 Example 3:


Input: root = [1]
Output: 1



 Constraints:


 The number of nodes in the tree is in the range [0, 5 * 10⁴].
 0 <= Node.val <= 5 * 10⁴
 The tree is guaranteed to be complete.


 Related Topics 树 深度优先搜索 二分查找 二叉树 👍 855 👎 0

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
func countNodes1(root *TreeNode) int {
	var cnt int

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if nil == node {
			return
		}
		cnt++
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)

	return cnt
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 利用完全二叉树特性的递归解法
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0

	// 检查左右子树的高度是否相等来验证当前二叉树是否是man二叉树
	leftNode := root.Left
	rightNode := root.Right
	for leftNode != nil {
		leftNode = leftNode.Left
		leftH++
	}
	for rightNode != nil {
		rightNode = rightNode.Right
		rightH++
	}
	if leftH == rightH {
		return (2 << leftH) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountCompleteTreeNodes(t *testing.T) {

}
