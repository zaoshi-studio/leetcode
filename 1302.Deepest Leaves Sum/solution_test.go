package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the sum of values of its deepest leaves.



 Example 1:


Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15


 Example 2:


Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 19



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 1 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 152 👎 0

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
func deepestLeavesSum(root *TreeNode) int {
	// 使用层序遍历完成

	var ans int

	queue := []*TreeNode{root}

	for len(queue) != 0 {

		size := len(queue)

		var levelSum int

		for i := 0; i < size; i++ {
			node := queue[i]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			levelSum += node.Val
		}

		queue = queue[size:]
		ans = levelSum
	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDeepestLeavesSum(t *testing.T) {

}
