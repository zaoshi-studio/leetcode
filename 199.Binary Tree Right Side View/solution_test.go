package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, imagine yourself standing on the right side of
it, return the values of the nodes you can see ordered from top to bottom.


 Example 1:


Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]


 Example 2:


Input: root = [1,null,3]
Output: [1,3]


 Example 3:


Input: root = []
Output: []



 Constraints:


 The number of nodes in the tree is in the range [0, 100].
 -100 <= Node.val <= 100


 Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 810 👎 0

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
func rightSideView(root *TreeNode) []int {

	var (
		ans []int
	)

	// 以右子树优先遍历，每层第一个被访问的一定是最右侧的结点
	// 当前层，最右侧结点被加入后，本层其他节点一概不考虑
	// 通过 len(ans) 来控制下一个要找的层深度
	//
	// 也可以使用层序遍历去完成，每次取每层最末尾的结点
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if nil == node {
			return
		}

		if depth == len(ans) {
			ans = append(ans, node.Val)
		}

		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}

	dfs(root, 0)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeRightSideView(t *testing.T) {

}
