package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 39 👎 0

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
func rightSideView(root *TreeNode) []int {
	if nil == root {
		return nil
	}

	var ans []int

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

		ans = append(ans, queue[size-1].Val)

		queue = queue[size:]

	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWNC0Lk(t *testing.T) {

}
