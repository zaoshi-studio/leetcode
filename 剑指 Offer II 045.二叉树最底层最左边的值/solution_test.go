package leetcode

import (
	"fmt"
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 35 👎 0

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
func findBottomLeftValue(root *TreeNode) int {
	if nil == root {
		return 0
	}

	var ans int

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		//size := len(queue)
		node := queue[0]
		ans = node.Val

		fmt.Println(node.Val)

		// 从右向左的层序遍历，保证底层最左侧结点在最后一个
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		queue = queue[1:]

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

func TestLwUNpT(t *testing.T) {

}
