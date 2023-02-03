package leetcode

import (
	"testing"
)

/**
Given the root of a binary search tree and the lowest and highest boundaries as
low and high, trim the tree so that all its elements lies in [low, high].
Trimming the tree should not change the relative structure of the elements that will
remain in the tree (i.e., any node's descendant should remain a descendant). It
can be proven that there is a unique answer.

 Return the root of the trimmed binary search tree. Note that the root may
change depending on the given bounds.


 Example 1:


Input: root = [1,0,2], low = 1, high = 2
Output: [1,null,2]


 Example 2:


Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
Output: [3,2,null,1]



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 0 <= Node.val <= 10⁴
 The value of each node in the tree is unique.
 root is guaranteed to be a valid binary search tree.
 0 <= low <= high <= 10⁴


 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 749 👎 0

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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	dummy := &TreeNode{
		Left: root,
		// 此处防止遍历时对 dummy 判断后不会更新 dummy 的左右子树
		Val: high,
	}

	var postorder func(node *TreeNode) *TreeNode
	postorder = func(node *TreeNode) *TreeNode {
		if nil == node {
			return nil
		}

		leftNode := postorder(node.Left)
		rightNode := postorder(node.Right)

		// 如果当前节点值小于下界，则该节点+左子树应该全部抛弃掉
		if node.Val < low {
			return rightNode
		}

		// 如果当前节点大于上界，则该节点+右子树应该全部抛弃掉
		if node.Val > high {
			return leftNode
		}

		node.Left = leftNode
		node.Right = rightNode

		return node
	}

	postorder(dummy)

	return dummy.Left
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTrimABinarySearchTree(t *testing.T) {

}
