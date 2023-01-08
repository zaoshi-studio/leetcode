package leetcode

import (
	"testing"
)

/**
Given the root of a binary search tree, and an integer k, return the kᵗʰ
smallest value (1-indexed) of all the values of the nodes in the tree.


 Example 1:


Input: root = [3,1,4,null,2], k = 1
Output: 1


 Example 2:


Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3



 Constraints:


 The number of nodes in the tree is n.
 1 <= k <= n <= 10⁴
 0 <= Node.val <= 10⁴



 Follow up: If the BST is modified often (i.e., we can do insert and delete
operations) and you need to find the kth smallest frequently, how would you
optimize?

 Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 692 👎 0

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
func kthSmallest(root *TreeNode, k int) int {

	var (
		ans  int
		rank int
	)

	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if nil == node {
			return
		}

		preorder(node.Left)

		rank++
		if k == rank {
			ans = node.Val
			return
		}

		preorder(node.Right)

	}

	preorder(root)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKthSmallestElementInABst(t *testing.T) {

}
