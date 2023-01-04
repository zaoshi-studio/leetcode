package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the bottom-up level order traversal of
its nodes' values. (i.e., from left to right, level by level from leaf to root).



 Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[15,7],[9,20],[3]]


 Example 2:


Input: root = [1]
Output: [[1]]


 Example 3:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 2000].
 -1000 <= Node.val <= 1000


 Related Topics 树 广度优先搜索 二叉树 👍 640 👎 0

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
func levelOrderBottom(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}

	var ans [][]int

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)

		level := make([]int, size)

		for i := 0; i < size; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			level[i] = node.Val
		}

		queue = queue[size:]
		ans = append(ans, level)
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeLevelOrderTraversalIi(t *testing.T) {

}
