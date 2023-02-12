package leetcode

import (
	"math"
	"testing"
)

/**
Given the root of a binary tree, return an array of the largest value in each
row of the tree (0-indexed).


 Example 1:


Input: root = [1,3,2,5,3,null,9]
Output: [1,3,9]


 Example 2:


Input: root = [1,2,3]
Output: [1,3]



 Constraints:


 The number of nodes in the tree will be in the range [0, 10⁴].
 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 298 👎 0

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
	if nil == root {
		return nil
	}

	var (
		ans []int
	)

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		size := len(queue)

		levelMax := math.MinInt32

		for i := 0; i < size; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			levelMax = max(levelMax, node.Val)
		}
		queue = queue[size:]

		ans = append(ans, levelMax)
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

func TestFindLargestValueInEachTreeRow(t *testing.T) {

}
