package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the postorder traversal of its nodes'
values.


 Example 1:


Input: root = [1,null,2,3]
Output: [3,2,1]


 Example 2:


Input: root = []
Output: []


 Example 3:


Input: root = [1]
Output: [1]



 Constraints:


 The number of the nodes in the tree is in the range [0, 100].
 -100 <= Node.val <= 100



Follow up: Recursive solution is trivial, could you do it iteratively?

 Related Topics 栈 树 深度优先搜索 二叉树 👍 965 👎 0

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
func postorderTraversal(root *TreeNode) []int {
	var ans []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		ans = append(ans, node.Val)
	}

	dfs(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreePostorderTraversal(t *testing.T) {

}
