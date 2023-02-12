package leetcode

import (
	"testing"
)

/**
T1 and T2 are two very large binary trees. Create an algorithm to determine if
T2 is a subtree of T1.

 A tree T2 is a subtree of T1 if there exists a node n in T1 such that the
subtree of n is identical to T2. That is, if you cut off the tree at node n, the two
trees would be identical.

 Note: This problem is slightly different from the original problem.

 Example1:


 Input: t1 = [1, 2, 3], t2 = [2]
 Output: true


 Example2:


 Input: t1 = [1, null, 2, 4], t2 = [3, 2]
 Output: false


 Note:


 The node numbers of both tree are in [0, 20000].


 Related Topics 树 深度优先搜索 二叉树 字符串匹配 哈希函数 👍 72 👎 0

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
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {

	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if nil == node {
			return false
		}

		var sub bool

		if node.Val == t2.Val {
			sub = compare(node, t2)
		}

		return sub || dfs(node.Left) || dfs(node.Right)
	}

	return dfs(t1)
}

func compare(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil && node2 != nil || node1 != nil && node2 == nil {
		return false
	}

	if node1.Val == node2.Val {
		return compare(node1.Left, node2.Left) && compare(node1.Right, node2.Right)
	}

	//return compare(node1.Left, node2) || compare(node1.Right, node2)
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCheckSubtreeLcci(t *testing.T) {

}
