package leetcode

import (
	"testing"
)

/**
You are given the root of a binary search tree (BST) and an integer val.

 Find the node in the BST that the node's value equals val and return the
subtree rooted with that node. If such a node does not exist, return null.


 Example 1:


Input: root = [4,2,7,1,3], val = 2
Output: [2,1,3]


 Example 2:


Input: root = [4,2,7,1,3], val = 5
Output: []



 Constraints:


 The number of nodes in the tree is in the range [1, 5000].
 1 <= Node.val <= 10⁷
 root is a binary search tree.
 1 <= val <= 10⁷


 Related Topics 树 二叉搜索树 二叉树 👍 350 👎 0

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
func searchBST(root *TreeNode, val int) *TreeNode {
	var ans *TreeNode
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if nil == node {
			return
		}

		if node.Val == val {
			ans = node
			return
		}

		if node.Val > val {
			traverse(node.Left)
			return
		}

		if node.Val < val {
			traverse(node.Right)
			return
		}
	}

	traverse(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSearchInABinarySearchTree(t *testing.T) {

}
