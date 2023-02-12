package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, find the maximum value v for which there exist
different nodes a and b where v = |a.val - b.val| and a is an ancestor of b.

 A node a is an ancestor of b if either: any child of a is equal to b or any
child of a is an ancestor of b.


 Example 1:


Input: root = [8,3,10,1,6,null,14,null,null,4,7,13]
Output: 7
Explanation: We have various ancestor-node differences, some of which are given
below :
|8 - 3| = 5
|3 - 7| = 4
|8 - 1| = 7
|10 - 13| = 3
Among all possible differences, the maximum value of 7 is obtained by |8 - 1| =
7.

 Example 2:


Input: root = [1,null,2,null,0,3]
Output: 3



 Constraints:


 The number of nodes in the tree is in the range [2, 5000].
 0 <= Node.val <= 10⁵


 Related Topics 树 深度优先搜索 二叉树 👍 135 👎 0

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
func maxAncestorDiff(root *TreeNode) int {

	if nil == root {
		return 0
	}

	var ans int

	// 将 globalMax 和 globalMin 当做一个上下文去维护
	// 注意 0，如果有返回值，0 在 nil 时返回，也可以作为结点值参与计算，比较麻烦
	// 直接忽略 0，关注节点内值的范围，以结点值构成一个离散区间
	var postorder func(node *TreeNode, globalMax, globalMin int)
	postorder = func(node *TreeNode, globalMax, globalMin int) {
		if nil == node {
			return
		}

		globalMax = max(node.Val, globalMax)
		globalMin = min(node.Val, globalMin)

		ans = max(ans, globalMax-globalMin)

		postorder(node.Left, globalMax, globalMin)
		postorder(node.Right, globalMax, globalMin)
	}

	postorder(root, root.Val, root.Val)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumDifferenceBetweenNodeAndAncestor(t *testing.T) {

}
