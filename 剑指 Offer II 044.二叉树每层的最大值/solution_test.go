package leetcode

import (
	"math"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 36 👎 0

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
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ans []int

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)

		m := math.MinInt

		for i := 0; i < size; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			m = max(m, node.Val)
		}

		queue = queue[size:]

		ans = append(ans, m)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)

func TestHPov7L(t *testing.T) {

}
