package leetcode

import (
	"math"
	"testing"
)

/**
Given a binary tree, find its minimum depth.

 The minimum depth is the number of nodes along the shortest path from the root
node down to the nearest leaf node.

 Note: A leaf is a node with no children.


 Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 2


 Example 2:


Input: root = [2,null,3,null,4,null,5,null,6]
Output: 5



 Constraints:


 The number of nodes in the tree is in the range [0, 10⁵].
 -1000 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 901 👎 0

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
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 && right != 0 {
		return right + 1
	}

	//if root.Left != nil && root.Right == nil {
	//	return 1 + minDepth1(root.Left)
	//}

	if left != 0 && right == 0 {
		return left + 1
	}

	//if root.Left == nil && root.Right != nil {
	//	return 1 + minDepth1(root.Right)
	//}

	return min(left, right) + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	return minD + 1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumDepthOfBinaryTree(t *testing.T) {

}
