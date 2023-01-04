package leetcode

import (
	"testing"
)

/**
Given two integer arrays preorder and inorder where preorder is the preorder
traversal of a binary tree and inorder is the inorder traversal of the same tree,
construct and return the binary tree.


 Example 1:


Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]


 Example 2:


Input: preorder = [-1], inorder = [-1]
Output: [-1]



 Constraints:


 1 <= preorder.length <= 3000
 inorder.length == preorder.length
 -3000 <= preorder[i], inorder[i] <= 3000
 preorder and inorder consist of unique values.
 Each value of inorder also appears in preorder.
 preorder is guaranteed to be the preorder traversal of the tree.
 inorder is guaranteed to be the inorder traversal of the tree.


 Related Topics 树 数组 哈希表 分治 二叉树 👍 1821 👎 0

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
func buildTree1(preorder []int, inorder []int) *TreeNode {

	var build func(preorder, inorder []int) *TreeNode
	build = func(preorder, inorder []int) *TreeNode {

		if len(preorder) == 0 {
			return nil
		}

		var idx int

		for ; inorder[idx] != preorder[0]; idx++ {
		}

		left := build(preorder[1:len(inorder[0:idx])+1], inorder[0:idx])
		right := build(preorder[len(inorder[0:idx])+1:], inorder[idx+1:])

		cur := &TreeNode{
			Val:   preorder[0],
			Left:  left,
			Right: right,
		}

		return cur
	}

	return build(preorder, inorder)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	record := make(map[int]int, len(preorder))
	for k, v := range inorder {
		record[v] = k
	}

	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {

		if preStart > preEnd {
			return nil
		}

		idx := record[preorder[preStart]]

		leftSize := idx - inStart

		left := build(preStart+1, preStart+leftSize, inStart, idx-1)
		right := build(preStart+leftSize+1, preEnd, idx+1, inEnd)

		node := &TreeNode{
			Val:   preorder[preStart],
			Left:  left,
			Right: right,
		}

		return node
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinaryTreeFromPreorderAndInorderTraversal(t *testing.T) {

}
