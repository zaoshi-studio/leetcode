package leetcode

import (
	"fmt"
	"testing"
)

/**
Given the root of a binary tree, return all root-to-leaf paths in any order.

 A leaf is a node with no children.


 Example 1:


Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]


 Example 2:


Input: root = [1]
Output: ["1"]



 Constraints:


 The number of nodes in the tree is in the range [1, 100].
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 861 👎 0

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
func binaryTreePaths(root *TreeNode) []string {
	var ans []string

	var dfs func(node *TreeNode, str string)
	dfs = func(node *TreeNode, str string) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			str = fmt.Sprintf("%s%d", str, node.Val)
			ans = append(ans, str)
			return
		} else {
			str = fmt.Sprintf("%s%d->", str, node.Val)
			dfs(node.Left, str)
			dfs(node.Right, str)
		}

	}

	dfs(root, "")

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreePaths(t *testing.T) {

}
