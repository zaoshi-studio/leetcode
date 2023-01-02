package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, invert the tree, and return its root.


 Example 1:


Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]


 Example 2:


Input: root = [2,1,3]
Output: [2,3,1]


 Example 3:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 100].
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1468 👎 0

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
func invertTree(root *TreeNode) *TreeNode {

	var invert func(node *TreeNode)
	invert = func(node *TreeNode) {
		if nil == node {
			return
		}

		// 先在左右子树内进行交换, 再交换左右子树
		invert(node.Left)
		invert(node.Right)

		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp

		return
	}

	invert(root)

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestInvertBinaryTree(t *testing.T) {

}
