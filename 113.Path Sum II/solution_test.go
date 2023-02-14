package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree and an integer targetSum, return all root-to-
leaf paths where the sum of the node values in the path equals targetSum. Each
path should be returned as a list of the node values, not node references.

 A root-to-leaf path is a path starting from the root and ending at any leaf
node. A leaf is a node with no children.


 Example 1:


Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
Output: [[5,4,11,2],[5,8,4,5]]
Explanation: There are two paths whose sum equals targetSum:
5 + 4 + 11 + 2 = 22
5 + 8 + 4 + 5 = 22


 Example 2:


Input: root = [1,2,3], targetSum = 5
Output: []


 Example 3:


Input: root = [1,2], targetSum = 0
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 5000].
 -1000 <= Node.val <= 1000
 -1000 <= targetSum <= 1000


 Related Topics 树 深度优先搜索 回溯 二叉树 👍 907 👎 0

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
func pathSum(root *TreeNode, targetSum int) [][]int {

	var ans [][]int

	var preorder func(node *TreeNode, preSum int, path []int)
	preorder = func(node *TreeNode, preSum int, path []int) {
		if nil == node {
			return
		}
		preSum += node.Val
		path = append(path, node.Val)

		if preSum == targetSum && node.Left == nil && node.Right == nil {

			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		preorder(node.Left, preSum, path)
		preorder(node.Right, preSum, path)
	}

	preorder(root, 0, []int{})

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestPathSumIi(t *testing.T) {

}
