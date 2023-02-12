package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the same tree where every subtree (of
the given tree) not containing a 1 has been removed.

 A subtree of a node node is node plus every node that is a descendant of node.



 Example 1:


Input: root = [1,null,0,0,1]
Output: [1,null,0,null,1]
Explanation:
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.


 Example 2:


Input: root = [1,0,1,0,0,0,1]
Output: [1,null,1,null,1]


 Example 3:


Input: root = [1,1,0,1,1,0,1,0]
Output: [1,1,0,1,1,null,1]



 Constraints:


 The number of nodes in the tree is in the range [1, 200].
 Node.val is either 0 or 1.


 Related Topics 树 深度优先搜索 二叉树 👍 318 👎 0

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
func pruneTree(root *TreeNode) *TreeNode {

	dummy := &TreeNode{Left: root}

	var postorder func(node *TreeNode) bool
	postorder = func(node *TreeNode) bool {
		if nil == node {
			return false
		}

		leftHas := postorder(node.Left)
		rightHas := postorder(node.Right)

		if !leftHas {
			node.Left = nil
		}

		if !rightHas {
			node.Right = nil
		}

		return leftHas || rightHas || node.Val == 1
	}

	postorder(dummy)

	return dummy.Left
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreePruning(t *testing.T) {

}
