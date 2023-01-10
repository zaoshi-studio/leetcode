package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the zigzag level order traversal of its
nodes' values. (i.e., from left to right, then right to left for the next level
and alternate between).


 Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]


 Example 2:


Input: root = [1]
Output: [[1]]


 Example 3:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 2000].
 -100 <= Node.val <= 100


 Related Topics 树 广度优先搜索 二叉树 👍 727 👎 0

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}

	var (
		ans  [][]int
		flag = false // 标记是否反向
	)

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		var level []int

		if flag {
			for i := size - 1; i >= 0; i-- {
				level = append(level, queue[i].Val)
			}
		} else {
			for i := 0; i < size; i++ {
				level = append(level, queue[i].Val)
			}
		}

		ans = append(ans, level)

		queue = queue[size:]

		flag = !flag
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeZigzagLevelOrderTraversal(t *testing.T) {

}
