package leetcode

import (
	"testing"
)

/**
Given two integer arrays inorder and postorder where inorder is the inorder
traversal of a binary tree and postorder is the postorder traversal of the same
tree, construct and return the binary tree.


 Example 1:


Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]


 Example 2:


Input: inorder = [-1], postorder = [-1]
Output: [-1]



 Constraints:


 1 <= inorder.length <= 3000
 postorder.length == inorder.length
 -3000 <= inorder[i], postorder[i] <= 3000
 inorder and postorder consist of unique values.
 Each value of postorder also appears in inorder.
 inorder is guaranteed to be the inorder traversal of the tree.
 postorder is guaranteed to be the postorder traversal of the tree.


 Related Topics 树 数组 哈希表 分治 二叉树 👍 904 👎 0

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
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	var idx int

	for ; inorder[idx] != postorder[len(postorder)-1]; idx++ {
	}
	//
	//fmt.Println("node: ", postorder[len(postorder)-1], "left: ", inorder[:idx], "right: ", inorder[idx+1:])
	//fmt.Println("node: ", postorder[len(postorder)-1], "left: ", postorder[:len(inorder[:idx])], "right: ", postorder[len(inorder[:idx]):len(postorder)-1])

	left := buildTree(inorder[:idx], postorder[:len(inorder[:idx])])
	// [left,right,node]
	// 中间不能断开
	right := buildTree(inorder[idx+1:], postorder[len(inorder[:idx]):len(postorder)-1])

	node := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  left,
		Right: right,
	}

	return node
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	record := make(map[int]int, len(inorder))
	for k, v := range inorder {
		record[v] = k
	}

	var build func(inStart, inEnd, postStart, postEnd int) *TreeNode
	build = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
		if postStart > postEnd {
			return nil
		}

		node := &TreeNode{
			Val: postorder[postEnd],
		}

		idx := record[node.Val]
		leftSize := idx - inStart

		node.Left = build(inStart, idx-1, postStart, postStart+leftSize-1)
		node.Right = build(idx+1, inEnd, postStart+leftSize, postEnd-1)

		return node
	}
	return build(0, len(inorder)-1, 0, len(postorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinaryTreeFromInorderAndPostorderTraversal(t *testing.T) {

}
