package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the inorder traversal of its nodes'
values.


 Example 1:


Input: root = [1,null,2,3]
Output: [1,3,2]


 Example 2:


Input: root = []
Output: []


 Example 3:


Input: root = [1]
Output: [1]



 Constraints:


 The number of nodes in the tree is in the range [0, 100].
 -100 <= Node.val <= 100



Follow up: Recursive solution is trivial, could you do it iteratively?

 Related Topics栈 | 树 | 深度优先搜索 | 二叉树

 👍 1645, 👎 0bug 反馈 | 使用指南 | 更多配套插件

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
func inorderTraversal(root *TreeNode) []int {

	var ans []int

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if nil == node {
			return
		}

		traverse(node.Left)
		ans = append(ans, node.Val)
		traverse(node.Right)

	}

	traverse(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeInorderTraversal(t *testing.T) {

}
