package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 树 广度优先搜索 二叉树 👍 260 👎 0

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

	var (
		ans   [][]int
		queue = []*TreeNode{root}
		flag  bool
	)

	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, size)

		for i := 0; i < size; i++ {
			// 弹出一个结点
			node := queue[0]
			queue = queue[1:]

			level[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

		}

		if flag {
			for i, n := 0, len(level); i < n/2; i++ {
				level[i], level[n-1-i] = level[n-1-i], level[i]
			}
		}

		ans = append(ans, level)

		flag = !flag

	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCongShangDaoXiaDaYinErChaShuIiiLcof(t *testing.T) {

}
