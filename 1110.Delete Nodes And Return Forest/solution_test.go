package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, each node in the tree has a distinct value.

 After deleting all nodes with a value in to_delete, we are left with a forest (
a disjoint union of trees).

 Return the roots of the trees in the remaining forest. You may return the
result in any order.


 Example 1:


Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
Output: [[1,2,null,4],[6],[7]]


 Example 2:


Input: root = [1,2,4,null,3], to_delete = [3]
Output: [[1,2,4]]



 Constraints:


 The number of nodes in the given tree is at most 1000.
 Each node has a distinct value between 1 and 1000.
 to_delete.length <= 1000
 to_delete contains distinct values between 1 and 1000.


 Related Topics 树 深度优先搜索 二叉树 👍 201 👎 0

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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {

	var ans []*TreeNode

	dummy := &TreeNode{Left: root}

	var postorder func(node *TreeNode) *TreeNode
	postorder = func(node *TreeNode) *TreeNode {
		if nil == node {
			return nil
		}

		node.Left = postorder(node.Left)
		node.Right = postorder(node.Right)

		if checkInSlice(node.Val, to_delete) {

			if node.Left != nil {
				ans = append(ans, node.Left)
			}

			if node.Right != nil {
				ans = append(ans, node.Right)
			}
			return nil
		}
		//fmt.Println(node.Val)
		return node
	}

	postorder(dummy)

	if dummy.Left != nil {
		ans = append(ans, dummy.Left)
	}

	return ans
}

func checkInSlice(val int, slice []int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDeleteNodesAndReturnForest(t *testing.T) {

}
