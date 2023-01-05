package leetcode

import (
	"testing"
)

/**
Given a root node reference of a BST and a key, delete the node with the given
key in the BST. Return the root node reference (possibly updated) of the BST.

 Basically, the deletion can be divided into two stages:


 Search for a node to remove.
 If the node is found, delete the node.



 Example 1:


Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and
delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also
accepted.



 Example 2:


Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.


 Example 3:


Input: root = [], key = 0
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 10⁴].
 -10⁵ <= Node.val <= 10⁵
 Each node has a unique value.
 root is a valid binary search tree.
 -10⁵ <= key <= 10⁵



 Follow up: Could you solve it with time complexity O(height of tree)?

 Related Topics 树 二叉搜索树 二叉树 👍 1035 👎 0

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
func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}

	if root.Val == key {
		// 1. 如果是叶节点直接删除
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// 2. 如果只有一颗子树
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		// 删除的是中间节点
		rightMin := root.Right
		// 找到右子树最左侧(最小)的结点
		for rightMin.Left != nil {
			rightMin = rightMin.Left
		}

		// 先在右子树上删除最小值结点
		root.Right = deleteNode(root.Right, rightMin.Val)

		rightMin.Left = root.Left
		rightMin.Right = root.Right

		root = rightMin

	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)

	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDeleteNodeInABst(t *testing.T) {

}
