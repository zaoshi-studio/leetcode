package leetcode

import (
	"testing"
)

/**
English description is not available for the problem. Please switch to Chinese.
 Related Topics 树 深度优先搜索 回溯 二叉树 👍 385 👎 0

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
func pathSum(root *TreeNode, target int) [][]int {
	var ans [][]int

	var dfs func(node *TreeNode, sum int, buffer []int)
	dfs = func(node *TreeNode, sum int, buffer []int) {
		if nil == node {
			return
		}

		sum += node.Val
		buffer = append(buffer, node.Val)

		if sum == target && node.Left == nil && node.Right == nil {
			temp := make([]int, len(buffer))
			copy(temp, buffer)
			ans = append(ans, temp)
			return
		}


		dfs(node.Left, sum, buffer)
		dfs(node.Right, sum, buffer)
	}

	dfs(root, 0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestErChaShuZhongHeWeiMouYiZhiDeLuJingLcof(t *testing.T) {

}
