package leetcode

import (
	"testing"
)

/**
Given an array of integers preorder, which represents the preorder traversal of
a BST (i.e., binary search tree), construct the tree and return its root.

 It is guaranteed that there is always possible to find a binary search tree
with the given requirements for the given test cases.

 A binary search tree is a binary tree where for every node, any descendant of
Node.left has a value strictly less than Node.val, and any descendant of Node.
right has a value strictly greater than Node.val.

 A preorder traversal of a binary tree displays the value of the node first,
then traverses Node.left, then traverses Node.right.


 Example 1:


Input: preorder = [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]


 Example 2:


Input: preorder = [1,3]
Output: [1,null,3]



 Constraints:


 1 <= preorder.length <= 100
 1 <= preorder[i] <= 1000
 All the values of preorder are unique.


 Related Topics 栈 树 二叉搜索树 数组 二叉树 单调栈 👍 243 👎 0

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
func bstFromPreorder(preorder []int) *TreeNode {
	return buildTree(preorder, 0, len(preorder)-1)
}

func buildTree(preorder []int, left, right int) *TreeNode {

	if left > right {
		return nil
	}

	node := &TreeNode{
		Val: preorder[left],
	}

	var idx = left + 1

	// 注意这里使用 idx <= right
	// idx 是可以取到 right 的
	for ; idx <= right && preorder[idx] < node.Val; idx++ {
	}

	node.Left = buildTree(preorder, left+1, idx-1)
	node.Right = buildTree(preorder, idx, right)

	return node
}

//leetcode submit region end(Prohibit modification and deletion)

func TestConstructBinarySearchTreeFromPreorderTraversal(t *testing.T) {
	arr := []int{4, 2}
	bstFromPreorder(arr)
}
