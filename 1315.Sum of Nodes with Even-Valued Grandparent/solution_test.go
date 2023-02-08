package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the sum of values of nodes with an even-
valued grandparent. If there are no nodes with an even-valued grandparent,
return 0.

 A grandparent of a node is the parent of its parent if it exists.


 Example 1:


Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 18
Explanation: The red nodes are the nodes with even-value grandparent while the
blue nodes are the even-value grandparents.


 Example 2:


Input: root = [1]
Output: 0



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 1 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 82 👎 0

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
func sumEvenGrandparent(root *TreeNode) int {

	var ans int

	var dfs func(node *TreeNode, parent, gParent bool)
	dfs = func(node *TreeNode, parent, gParent bool) {
		if nil == node {
			return
		}

		if gParent {
			ans += node.Val
		}

		dfs(node.Left, node.Val%2 == 0, parent)
		dfs(node.Right, node.Val%2 == 0, parent)
	}

	dfs(root, false, false)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSumOfNodesWithEvenValuedGrandparent(t *testing.T) {

}
