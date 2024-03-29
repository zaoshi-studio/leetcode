package leetcode

import (
	"testing"
)

/**
You are given the root node of a binary search tree (BST) and a value to insert
into the tree. Return the root node of the BST after the insertion. It is
guaranteed that the new value does not exist in the original BST.

 Notice that there may exist multiple valid ways for the insertion, as long as
the tree remains a BST after insertion. You can return any of them.


 Example 1:


Input: root = [4,2,7,1,3], val = 5
Output: [4,2,7,1,3,5]
Explanation: Another accepted tree is:



 Example 2:


Input: root = [40,20,60,10,30,50,70], val = 25
Output: [40,20,60,10,30,50,70,null,null,25]


 Example 3:


Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
Output: [4,2,7,1,3,5]



 Constraints:


 The number of nodes in the tree will be in the range [0, 10⁴].
 -10⁸ <= Node.val <= 10⁸
 All the values Node.val are unique.
 -10⁸ <= val <= 10⁸
 It's guaranteed that val does not exist in the original BST.


 Related Topics 树 二叉搜索树 二叉树 👍 417 👎 0

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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestInsertIntoABinarySearchTree(t *testing.T) {

}
