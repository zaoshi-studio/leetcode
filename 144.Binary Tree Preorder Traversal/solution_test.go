package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the preorder traversal of its nodes'
values.


 Example 1:


Input: root = [1,null,2,3]
Output: [1,2,3]


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

 Related Topics 栈 树 深度优先搜索 二叉树 👍 970 👎 0

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
func preorderTraversal(root *TreeNode) []int {
	var ans []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		ans = append(ans, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreePreorderTraversal(t *testing.T) {

}
