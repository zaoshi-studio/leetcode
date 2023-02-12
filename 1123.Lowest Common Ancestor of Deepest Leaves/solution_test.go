package leetcode

import (
	"testing"
)

/**
Given the root of a binary tree, return the lowest common ancestor of its
deepest leaves.

 Recall that:


 The node of a binary tree is a leaf if and only if it has no children
 The depth of the root of the tree is 0. if the depth of a node is d, the depth
of each of its children is d + 1.
 The lowest common ancestor of a set S of nodes, is the node A with the largest
depth such that every node in S is in the subtree with root A.



 Example 1:


Input: root = [3,5,1,6,2,0,8,null,null,7,4]
Output: [2,7,4]
Explanation: We return the node with value 2, colored in yellow in the diagram.
The nodes coloured in blue are the deepest leaf-nodes of the tree.
Note that nodes 6, 0, and 8 are also leaf nodes, but the depth of them is 2,
but the depth of nodes 7 and 4 is 3.

 Example 2:


Input: root = [1]
Output: [1]
Explanation: The root is the deepest node in the tree, and it's the lca of
itself.


 Example 3:


Input: root = [0,1,3,null,2]
Output: [2]
Explanation: The deepest leaf node in the tree is 2, the lca of one node is
itself.



 Constraints:


 The number of nodes in the tree will be in the range [1, 1000].
 0 <= Node.val <= 1000
 The values of the nodes in the tree are unique.



 Note: This question is the same as 865: https://leetcode.com/problems/smallest-
subtree-with-all-the-deepest-nodes/

 Related Topics 树 深度优先搜索 广度优先搜索 哈希表 二叉树 👍 140 👎 0

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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {

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

func TestLowestCommonAncestorOfDeepestLeaves(t *testing.T) {

}
