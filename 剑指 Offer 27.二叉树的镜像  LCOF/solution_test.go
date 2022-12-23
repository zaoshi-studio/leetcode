package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 324 👎 0

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
func mirrorTree(root *TreeNode) *TreeNode {

	if nil == root {
		return nil
	}

	node := &TreeNode{
		Val: root.Val,
	}

	var mirror func(nodeA, nodeB *TreeNode)
	mirror = func(nodeA, nodeB *TreeNode) {

		if nodeB == nil {
			return
		}

		if nodeB.Left != nil {
			nodeA.Right = &TreeNode{
				Val: nodeB.Left.Val,
			}
		}

		if nodeB.Right != nil {
			nodeA.Left = &TreeNode{
				Val: nodeB.Right.Val,
			}
		}

		mirror(nodeA.Left, nodeB.Right)
		mirror(nodeA.Right, nodeB.Left)
	}

	mirror(node, root)

	return node
}

func mirrorTreeOfficial(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := mirrorTreeOfficial(root.Left)
	right := mirrorTreeOfficial(root.Right)
	root.Left = right
	root.Right = left
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestErChaShuDeJingXiangLcof(t *testing.T) {

}
