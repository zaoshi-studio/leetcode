package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the sum of all left leaves.

 A leaf is a node with no children. A left leaf is a leaf that is the left
child of another node.


 Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15
respectively.


 Example 2:


Input: root = [1]
Output: 0



 Constraints:


 The number of nodes in the tree is in the range [1, 1000].
 -1000 <= Node.val <= 1000


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 541 👎 0

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
func sumOfLeftLeaves(root *TreeNode) int {

	var ans int

	var traverse func(node *TreeNode, isLeft bool)
	traverse = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}

		// 叶节点返回值
		if node.Left == nil && node.Right == nil && isLeft {
			ans += node.Val
		}

		traverse(node.Left, true)
		traverse(node.Right, false)
	}

	traverse(root.Left, true)
	traverse(root.Right, false)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSumOfLeftLeaves(t *testing.T) {

}
