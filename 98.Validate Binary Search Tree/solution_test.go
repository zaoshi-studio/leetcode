package leetcode

import (
	"fmt"
	"math"
	"testing"
)

/**
Given the root of a binary tree, determine if it is a valid binary search tree (
BST).

 A valid BST is defined as follows:


 The left subtree of a node contains only nodes with keys less than the node's
key.
 The right subtree of a node contains only nodes with keys greater than the
node's key.
 Both the left and right subtrees must also be binary search trees.



 Example 1:


Input: root = [2,1,3]
Output: true


 Example 2:


Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 1843 👎 0

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

// 尝试进行自底向上
func isValidBST1(root *TreeNode) bool {
	ans := true

	var traverse func(node *TreeNode) (int, int)
	traverse = func(node *TreeNode) (int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt
		}

		//var leftMax, rightMin int

		_, leftMax := traverse(node.Left)

		rightMin, _ := traverse(node.Right)

		//fmt.Printf("node:%d , left:%d , right: %d\n", node.Val, leftMax, rightMin)

		//if leftMax == -1 && rightMin == -1 {
		//	return node.Val, node.Val
		//}
		//
		//if leftMax != -1 && rightMin == -1 && node.Val <= leftMax {
		//	ans = false
		//}
		//
		//if leftMax == -1 && rightMin != -1 && node.Val >= rightMin {
		//	ans = false
		//}

		if leftMax >= node.Val || node.Val >= rightMin {
			fmt.Printf("node:%d , left:%d , right: %d\n", node.Val, leftMax, rightMin)
			ans = false
		}

		// 右子树返回最小值
		// 左子树返回最大值

		//return max(node.Val, max(leftMax, rightMin)), min(node.Val, min(leftMax, rightMin))
		return max(node.Val, max(leftMax, rightMin)), min(node.Val, min(leftMax, rightMin))
	}

	traverse(root)

	return ans
}

func isValidBST(root *TreeNode) bool {
	// 自顶向下

	var traverse func(node *TreeNode, low, high int) bool
	traverse = func(node *TreeNode, low, high int) bool {
		if node == nil {
			return true
		}

		if node.Val <= low || node.Val >= high {
			return false
		}

		return traverse(node.Left, low, node.Val) && traverse(node.Right, node.Val, high)
	}

	return traverse(root, math.MinInt, math.MaxInt)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidateBinarySearchTree(t *testing.T) {

}
