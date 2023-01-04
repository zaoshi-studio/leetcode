package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the level order traversal of its nodes'
values. (i.e., from left to right, level by level).


 Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]


 Example 2:


Input: root = [1]
Output: [[1]]


 Example 3:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 2000].
 -1000 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 1542 👎 0

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
func levelOrder(root *TreeNode) [][]int {

	if nil == root {
		return nil
	}

	var ans [][]int

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)

		var level []int

		for i := 0; i < size; i++ {
			node := queue[i]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]

		ans = append(ans, level)
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeLevelOrderTraversal(t *testing.T) {

}
