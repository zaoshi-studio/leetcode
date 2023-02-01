package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the leftmost value in the last row of
the tree.


 Example 1:


Input: root = [2,1,3]
Output: 1


 Example 2:


Input: root = [1,2,3,4,null,5,6,null,null,7]
Output: 7



 Constraints:


 The number of nodes in the tree is in the range [1, 10⁴].
 -2³¹ <= Node.val <= 2³¹ - 1


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 428 👎 0

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

	var (
		ans, maxDepth int
	)

	// 为什么这样可以
	// 采用后序遍历，遍历顺序为 左右中
	// 假设最后一层有多个节点，先左右，左侧结点一定先被记录到 ans，再到右侧结点，depth <= maxDepth
	// 假设最后一层只有一个节点，无论该节点是左还是右结点，都一样

	var postorder func(node *TreeNode, depth int)
	postorder = func(node *TreeNode, depth int) {

		if node.Left != nil {
			postorder(node.Left, depth+1)
		}

		if node.Right != nil {
			postorder(node.Right, depth+1)
		}

		if depth > maxDepth {
			maxDepth = depth
			ans = node.Val
		}

	}

	postorder(root, 1)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindBottomLeftTreeValue(t *testing.T) {

}
