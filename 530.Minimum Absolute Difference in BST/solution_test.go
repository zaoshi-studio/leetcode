package leetcode

import (
	"math"
	"testing"
)

/**
Given the root of a Binary Search Tree (BST), return the minimum absolute
difference between the values of any two different nodes in the tree.


 Example 1:


Input: root = [4,2,6,1,3]
Output: 1


 Example 2:


Input: root = [1,0,48,null,null,12,49]
Output: 1



 Constraints:


 The number of nodes in the tree is in the range [2, 10⁴].
 0 <= Node.val <= 10⁵



 Note: This question is the same as 783: https://leetcode.com/problems/minimum-
distance-between-bst-nodes/

 Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 二叉树 👍 422 👎 0

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
func getMinimumDifference(root *TreeNode) int {

	ans := math.MaxInt32

	var prev *TreeNode

	// 二叉搜索树保证中序遍历有序
	// 仅需要比较当前节点与前一个节点的差值
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if nil == node {
			return
		}

		inorder(node.Left)

		if prev != nil {
			ans = min(ans, node.Val-prev.Val)
		}

		prev = node

		inorder(node.Right)
	}

	inorder(root)

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumAbsoluteDifferenceInBst(t *testing.T) {

}
