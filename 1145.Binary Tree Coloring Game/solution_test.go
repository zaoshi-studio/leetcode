package leetcode

import (
	"testing"
)

/**
Two players play a turn based game on a binary tree. We are given the root of
this binary tree, and the number of nodes n in the tree. n is odd, and each node
has a distinct value from 1 to n.

 Initially, the first player names a value x with 1 <= x <= n, and the second
player names a value y with 1 <= y <= n and y != x. The first player colors the
node with value x red, and the second player colors the node with value y blue.

 Then, the players take turns starting with the first player. In each turn,
that player chooses a node of their color (red if player 1, blue if player 2) and
colors an uncolored neighbor of the chosen node (either the left child, right
child, or parent of the chosen node.)

 If (and only if) a player cannot choose such a node in this way, they must
pass their turn. If both players pass their turn, the game ends, and the winner is
the player that colored more nodes.

 You are the second player. If it is possible to choose such a y to ensure you
win the game, return true. If it is not possible, return false.


 Example 1:


Input: root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
Output: true
Explanation: The second player can choose the node with value 2.


 Example 2:


Input: root = [1,2,3], n = 3, x = 1
Output: false



 Constraints:


 The number of nodes in the tree is n.
 1 <= x <= n <= 100
 n is odd.
 1 <= Node.val <= n
 All the values of the tree are unique.


 Related Topics 树 深度优先搜索 二叉树 👍 148 👎 0

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
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {

	var xNode *TreeNode

	var childCnt func(node *TreeNode) int
	childCnt = func(node *TreeNode) int {
		if nil == node {
			return 0
		}

		if node.Val == x {
			xNode = node
		}

		return childCnt(node.Left) + childCnt(node.Right) + 1
	}

	childCnt(root)

	xLeftNodeCnt := childCnt(xNode.Left)
	xRightNodeCnt := childCnt(xNode.Right)

	// 本体关键点是: 以 xNode 为中心，将整个二叉树分为三个区域，xNode 的左右子树和其他节点
	// 玩家只用在这三个区域中选择节点数大于等于一半的区域即可
	// 为什么是 n + 1，题目中明确指出 n 为基数，直接使用 n/2 对导致向下取整
	// 会导致不满一半的情况，导致计数被少算
	if xLeftNodeCnt >= (n+1)/2 || xRightNodeCnt >= (n+1)/2 {
		return true
	}

	return n-xLeftNodeCnt-xRightNodeCnt-1 >= (n+1)/2
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeColoringGame(t *testing.T) {

}
