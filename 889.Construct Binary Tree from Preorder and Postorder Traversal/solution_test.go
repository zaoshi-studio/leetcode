package leetcode

import (
	"testing"
)

/**
Given two integer arrays, preorder and postorder where preorder is the preorder
traversal of a binary tree of distinct values and postorder is the postorder
traversal of the same tree, reconstruct and return the binary tree.

 If there exist multiple answers, you can return any of them.


 Example 1:


Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
Output: [1,2,3,4,5,6,7]


 Example 2:


Input: preorder = [1], postorder = [1]
Output: [1]



 Constraints:


 1 <= preorder.length <= 30
 1 <= preorder[i] <= preorder.length
 All the values of preorder are unique.
 postorder.length == preorder.length
 1 <= postorder[i] <= postorder.length
 All the values of postorder are unique.
 It is guaranteed that preorder and postorder are the preorder traversal and
postorder traversal of the same binary tree.


 Related Topics 树 数组 哈希表 分治 二叉树 👍 290 👎 0

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
func constructFromPrePost1(preorder []int, postorder []int) *TreeNode {

	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}

	var idx int

	for ; postorder[idx] != preorder[1]; idx++ {
	}

	left := constructFromPrePost1(preorder[1:len(postorder[:idx+1])+1], postorder[:idx+1])
	right := constructFromPrePost1(preorder[len(postorder[:idx+1])+1:], postorder[idx+1:len(postorder)-1])

	node := &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}

	return node
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {

	record := make(map[int]int, len(postorder))

	for k, v := range postorder {
		record[v] = k
	}

	var build func(preStart, preEnd, postStart, postEnd int) *TreeNode
	build = func(preStart, preEnd, postStart, postEnd int) *TreeNode {
		if preStart > preEnd || postStart > postEnd {
			return nil
		}

		if preStart == preEnd {
			return &TreeNode{
				Val: preorder[preStart],
			}
		}

		node := &TreeNode{
			Val: postorder[postEnd],
		}

		leftNeighbourVal := preorder[preStart+1]
		leftNeighbourIdx := record[leftNeighbourVal]
		leftSize := leftNeighbourIdx - postStart + 1

		node.Left = build(preStart+1, preStart+leftSize, postStart, leftNeighbourIdx)
		node.Right = build(preStart+leftSize+1, preEnd, leftNeighbourIdx+1, postEnd-1)

		return node

	}

	return build(0, len(preorder)-1, 0, len(postorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinaryTreeFromPreorderAndPostorderTraversal(t *testing.T) {

}
