package leetcode

import (
	"testing"
)

/**
Given the root of a binary search tree, return a balanced binary search tree
with the same node values. If there is more than one answer, return any of them.

 A binary search tree is balanced if the depth of the two subtrees of every
node never differs by more than 1.


 Example 1:


Input: root = [1,null,2,null,3,null,4,null,null]
Output: [2,1,3,null,null,null,4]
Explanation: This is not the only correct answer, [3,1,4,null,2] is also
correct.


 Example 2:


Input: root = [2,1,3]
Output: [2,1,3]



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 1 <= Node.val <= 10⁵


 Related Topics 贪心 树 深度优先搜索 二叉搜索树 分治 二叉树 👍 156 👎 0

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
func balanceBST(root *TreeNode) *TreeNode {
	// 由于是二叉搜索树
	// 利用中序遍历将数值存进数组，得到一个有序序列
	// 回到有序序列构造二叉树上
	// 由于每次都是取得中间值，一定平衡
	var arr []int

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if nil == node {
			return
		}

		inorder(node.Left)
		arr = append(arr, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	return buildTree(arr, 0, len(arr)-1)
}

func buildTree(arr []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2

	node := &TreeNode{
		Val: arr[mid],
	}

	node.Left = buildTree(arr, left, mid-1)
	node.Right = buildTree(arr, mid+1, right)
	return node
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBalanceABinarySearchTree(t *testing.T) {

}
