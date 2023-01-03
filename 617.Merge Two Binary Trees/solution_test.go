package leetcode

import (
	"testing"
)

/**
You are given two binary trees root1 and root2.

 Imagine that when you put one of them to cover the other, some nodes of the
two trees are overlapped while the others are not. You need to merge the two trees
into a new binary tree. The merge rule is that if two nodes overlap, then sum
node values up as the new value of the merged node. Otherwise, the NOT null node
will be used as the node of the new tree.

 Return the merged tree.

 Note: The merging process must start from the root nodes of both trees.


 Example 1:


Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
Output: [3,4,5,5,4,null,7]


 Example 2:


Input: root1 = [1], root2 = [1,2]
Output: [2,2]



 Constraints:


 The number of nodes in both trees is in the range [0, 2000].
 -10⁴ <= Node.val <= 10⁴


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1132 👎 0

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
func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil && root2 != nil {
		return root2
	}

	if root2 == nil && root1 != nil {
		return root1
	}

	if root1 == nil && root2 == nil {
		return nil
	}

	var merge func(nodeA, nodeB *TreeNode)
	merge = func(nodeA, nodeB *TreeNode) {
		// nodeA 与  nodeB 一定不为 nil
		// 将 nodeB 拷贝到 nodeA 上
		nodeA.Val += nodeB.Val

		if nodeA.Left != nil && nodeB.Left != nil {
			merge(nodeA.Left, nodeB.Left)
		}

		if nodeA.Left == nil && nodeB.Left != nil {
			nodeA.Left = nodeB.Left
		}

		if nodeA.Right != nil && nodeB.Right != nil {
			merge(nodeA.Right, nodeB.Right)
		}

		if nodeA.Right == nil && nodeB.Right != nil {
			nodeA.Right = nodeB.Right
		}
	}

	merge(root1, root2)

	return root1
}

func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	var merge func(nodeA, nodeB *TreeNode) *TreeNode
	merge = func(nodeA, nodeB *TreeNode) *TreeNode {
		// nodeA 与  nodeB 一定不为 nil
		// 将 nodeB 拷贝到 nodeA 上

		if nodeA == nil {
			return nodeB
		}

		if nodeB == nil {
			return nodeA
		}

		nodeA.Val += nodeB.Val

		nodeA.Left = merge(nodeA.Left, nodeB.Left)
		nodeA.Right = merge(nodeA.Right, nodeB.Right)

		return nodeA
	}

	return merge(root1, root2)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMergeTwoBinaryTrees(t *testing.T) {

}
