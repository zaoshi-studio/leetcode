package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, check whether it is a mirror of itself (i.e.,
symmetric around its center).


 Example 1:


Input: root = [1,2,2,3,4,4,3]
Output: true


 Example 2:


Input: root = [1,2,2,null,3,null,3]
Output: false



 Constraints:


 The number of nodes in the tree is in the range [1, 1000].
 -100 <= Node.val <= 100



Follow up: Could you solve it both recursively and iteratively?

 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2225 👎 0

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
func isSymmetric(root *TreeNode) bool {

	var compare func(nodeA, nodeB *TreeNode) bool
	compare = func(nodeA, nodeB *TreeNode) bool {
		if nodeA == nil && nodeB == nil {
			return true
		}

		if nodeA != nil && nodeB == nil ||
			nodeA == nil && nodeB != nil {
			return false
		}

		if nodeA.Val != nodeB.Val {
			return false
		}

		return compare(nodeA.Left, nodeB.Right) && compare(nodeA.Right, nodeB.Left)
	}

	return compare(root, root)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSymmetricTree(t *testing.T) {

}
