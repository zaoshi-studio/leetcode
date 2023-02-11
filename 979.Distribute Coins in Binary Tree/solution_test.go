package leetcode

import (
	"testing"
)

/**
You are given the root of a binary tree with n nodes where each node in the
tree has node.val coins. There are n coins in total throughout the whole tree.

 In one move, we may choose two adjacent nodes and move one coin from one node
to another. A move may be from parent to child, or from child to parent.

 Return the minimum number of moves required to make every node have exactly
one coin.


 Example 1:


Input: root = [3,0,0]
Output: 2
Explanation: From the root of the tree, we move one coin to its left child, and
one coin to its right child.


 Example 2:


Input: root = [0,3,0]
Output: 3
Explanation: From the left child of the root, we move two coins to the root [
taking two moves]. Then, we move one coin from the root of the tree to the right
child.



 Constraints:


 The number of nodes in the tree is n.
 1 <= n <= 100
 0 <= Node.val <= n
 The sum of all Node.val is n.


 Related Topics 树 深度优先搜索 二叉树 👍 346 👎 0

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
func distributeCoins(root *TreeNode) int {

	var ans int

	var postorder func(node *TreeNode) int
	postorder = func(node *TreeNode) int {
		if nil == node {
			return 0
		}

		left := postorder(node.Left)
		right := postorder(node.Right)

		// ans 记录需要的步骤
		// 位于树末尾的结点每计算一次就表示到达末尾结点需要在多一步
		// 也就是深度更深一层
		ans += abs(left) + abs(right)

		// 表示从当前的 node 结点，需要给子节点一共分配 left + right 枚硬币
		// 计算 left 和 right 为负数，则表明 node 结点倒欠子节点这么多硬币
		// 可以理解为当前 node 向父节点请求 node.Val - 1 + right + left 枚硬币
		// 如果为正，则可以反馈给父节点
		// 如果为负，则需要这么多
		return node.Val - 1 + left + right
	}

	postorder(root)

	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDistributeCoinsInBinaryTree(t *testing.T) {

}
