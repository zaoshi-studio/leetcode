package leetcode

import (
	"testing"
)

/**
Given the root of a binary search tree (BST) with duplicates, return all the
mode(s) (i.e., the most frequently occurred element) in it.

 If the tree has more than one mode, return them in any order.

 Assume a BST is defined as follows:


 The left subtree of a node contains only nodes with keys less than or equal to
the node's key.
 The right subtree of a node contains only nodes with keys greater than or
equal to the node's key.
 Both the left and right subtrees must also be binary search trees.



 Example 1:


Input: root = [1,null,2,2]
Output: [2]


 Example 2:


Input: root = [0]
Output: [0]



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 -10⁵ <= Node.val <= 10⁵



Follow up: Could you do that without using any extra space? (Assume that the
implicit stack space incurred due to recursion does not count).

 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 576 👎 0

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
func findMode(root *TreeNode) []int {

	var ans []int
	var prev, cnt, maxCnt int
	// 此处 cnt 需要定义成全局变量
	// 在中序遍历中，可能左子树的最右侧结点先开始计数
	// 如果将 cnt 放到每一个递归函数，则可能当左子树遍历完成回到根节点时，cnt 会重新计数

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		if prev == node.Val {
			cnt++
		} else {
			cnt = 1
			prev = node.Val
		}

		if cnt == maxCnt {

			ans = append(ans, prev)
		} else if cnt > maxCnt {
			ans = []int{prev}
			maxCnt = cnt
		}

		inorder(node.Right)
	}

	inorder(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindModeInBinarySearchTree(t *testing.T) {

}
