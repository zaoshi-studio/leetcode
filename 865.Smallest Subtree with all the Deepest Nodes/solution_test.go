package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, the depth of each node is the shortest
distance to the root.

 Return the smallest subtree such that it contains all the deepest nodes in the
original tree.

 A node is called the deepest if it has the largest depth possible among any
node in the entire tree.

 The subtree of a node is a tree consisting of that node, plus the set of all
descendants of that node.


 Example 1:


Input: root = [3,5,1,6,2,0,8,null,null,7,4]
Output: [2,7,4]
Explanation: We return the node with value 2, colored in yellow in the diagram.
The nodes coloured in blue are the deepest nodes of the tree.
Notice that nodes 5, 3 and 2 contain the deepest nodes in the tree but node 2
is the smallest subtree among them, so we return it.


 Example 2:


Input: root = [1]
Output: [1]
Explanation: The root is the deepest node in the tree.


 Example 3:


Input: root = [0,1,3,null,2]
Output: [2]
Explanation: The deepest node in the tree is 2, the valid subtrees are the
subtrees of nodes 2, 1 and 0 but the subtree of node 2 is the smallest.



 Constraints:


 The number of nodes in the tree will be in the range [1, 500].
 0 <= Node.val <= 500
 The values of the nodes in the tree are unique.



 Note: This question is the same as 1123: https://leetcode.com/problems/lowest-
common-ancestor-of-deepest-leaves/

 Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 👍 187 👎 0

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
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {

	var (
		ans *TreeNode
	)

	var postorder func(node *TreeNode, depth int) (*TreeNode, int)
	postorder = func(node *TreeNode, depth int) (*TreeNode, int) {
		if nil == node {
			return node, depth
		}

		leftNode, leftDepth := postorder(node.Left, depth+1)
		rightNode, rightDepth := postorder(node.Right, depth+1)

		// 当 leftDepth == rightDepth 时
		// 当前节点就是公共节点，在叶子结点处一定先拥有该状态
		// 也就是叶子结点本身一定是公共节点
		// 仅当左右 depth 平衡被打破时，由于公共节点已经找到(无论是叶子结点本身还是其祖先或者父节点)
		// 只需要吧之前的公共节点向上传递即可
		// 由于要最大深度的叶子结点的公共节点
		// 因此根据左右子树的深度判断，向上传递哪一侧的公共节点
		if leftDepth == rightDepth {
			return node, leftDepth
		}

		if leftDepth > rightDepth {
			return leftNode, leftDepth
		}

		return rightNode, rightDepth
	}

	ans, _ = postorder(root, 1)

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSmallestSubtreeWithAllTheDeepestNodes(t *testing.T) {

}
